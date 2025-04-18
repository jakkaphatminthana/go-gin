package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkaphatminthana/go-gin/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauth2Config *oauth2.Config

func init() {
	googleOauth2Config = &oauth2.Config{
		ClientID:     config.Config.ClientID,
		ClientSecret: config.Config.ClientSecret,
		RedirectURL:  "http://localhost:3000/dashboard/callback/google",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func HandlerGoogleLogin(c *gin.Context) {
	url := googleOauth2Config.AuthCodeURL("champ", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusFound, url)
}
