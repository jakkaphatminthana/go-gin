package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jakkaphatminthana/go-gin/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauth2Config *oauth2.Config
var oauthStateString = "champ"

func init() {
	googleOauth2Config = &oauth2.Config{
		ClientID:     config.Config.GoogleClientID,
		ClientSecret: config.Config.GoogleClientSecret,
		RedirectURL:  config.Config.GoogleRedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func HandlerGoogleLogin(c *gin.Context) {
	url := googleOauth2Config.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}

func HandlerGoogleCallback(c *gin.Context) {
	// 1.validate the state
	state := c.Query("state")
	if state != oauthStateString {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OAuth state"})
		return
	}

	// 2.validate the code and get token
	code := c.Query("code")
	token, err := googleOauth2Config.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
		return
	}

	// 3.get user-info using token
	client := googleOauth2Config.Client(c, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}
	defer resp.Body.Close()

	// 4.send user-info to Frontend
	var userInfo struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Picture string `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user info"})
		return
	}

	jwtToken, err := generateJWT(userInfo.Email, userInfo.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   jwtToken,
		"email":   userInfo.Email,
		"name":    userInfo.Name,
		"picture": userInfo.Picture,
	})

}

func generateJWT(email, name string) (string, error) {
	tokenInfo := jwt.MapClaims{
		"email": email,
		"name":  name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenInfo)
	return token.SignedString([]byte(config.Config.JWTSaltKey))
}
