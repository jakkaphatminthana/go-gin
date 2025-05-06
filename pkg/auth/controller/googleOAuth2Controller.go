package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jakkaphatminthana/go-gin/config"
	"github.com/jakkaphatminthana/go-gin/entities"
	"github.com/jakkaphatminthana/go-gin/pkg/custom"
	_userService "github.com/jakkaphatminthana/go-gin/pkg/user/service"
	"github.com/jakkaphatminthana/go-gin/types"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type googleOAuth2Controller struct {
	conf        *config.Config
	userService _userService.UserService
}

func NewGoogleOAuth2(conf *config.Config, userService _userService.UserService) OAuth2Controller {
	return &googleOAuth2Controller{conf, userService}
}

var oauthStateString = "champ" //easy

// implement
func (g *googleOAuth2Controller) GoogleLogin(c *gin.Context) {
	oauthConfig := g.getOAuth2Config()
	url := oauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

// implement
func (g *googleOAuth2Controller) GoogleCallback(c *gin.Context) {
	// 1.validate the state
	if c.Query("state") != oauthStateString {
		custom.ErrorBadRequest("Invalid OAuth state")
		return
	}

	// 2.validate the code and get token
	code := c.Query("code")
	conf := g.getOAuth2Config()
	token, err := conf.Exchange(c, code)
	if err != nil {
		custom.ErrorInternalServerError("Failed to exchange token")
		return
	}

	// 3.get user-info using token
	client := conf.Client(c, token)
	resp, err := client.Get(g.conf.OAuth2.UserInfoUrl)
	if err != nil {
		custom.ErrorInternalServerError("Failed to get user info")
		return
	}
	defer resp.Body.Close()

	// 4. Decode response to userInfo
	var userInfo struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Picture string `json:"picture"`
		ID      string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		custom.ErrorInternalServerError("Failed to decode user info")
		return
	}

	// 5. Check and Create User
	//check user
	user, err := g.userService.FindByEmail(userInfo.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = nil
		} else {
			custom.ErrorInternalServerError("Find user error: " + err.Error())
			return
		}
	}

	if user == nil {
		user = &entities.User{
			Name:    userInfo.Name,
			Email:   userInfo.Email,
			Picture: userInfo.Picture,
		}

		provider := &entities.Provider{
			Provider:   string(types.ProviderEnumGoogle),
			ProviderID: userInfo.ID,
		}

		_, err := g.userService.Create(user, provider)
		if err != nil {
			custom.ErrorInternalServerError("Create user error: " + err.Error())
			return
		}
	}

	// 6. Generate jwt
	jwtToken, err := g.generateJWT(user.Email, user.Name)
	if err != nil {
		custom.ErrorInternalServerError("Failed to generate token")
		return
	}

	log.Printf("🔑 jwtToken = %s", jwtToken)

	custom.Success(c, http.StatusOK, gin.H{
		"token":   jwtToken,
		"email":   userInfo.Email,
		"name":    userInfo.Name,
		"picture": userInfo.Picture,
	})
}

func (g *googleOAuth2Controller) getOAuth2Config() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     g.conf.OAuth2.GoogleClientID,
		ClientSecret: g.conf.OAuth2.GoogleClientSecret,
		RedirectURL:  g.conf.OAuth2.GoogleRedirectURL,
		Scopes:       g.conf.OAuth2.Scopes,
		Endpoint:     google.Endpoint,
	}
}

func (g *googleOAuth2Controller) generateJWT(email, name string) (string, error) {
	clamis := jwt.MapClaims{
		"email": email,
		"name":  name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clamis)
	return token.SignedString([]byte(g.conf.EVNValue.JWTSaltKey))
}
