package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jakkaphatminthana/go-gin/config"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// read Authorization rawToken from request header
		rawToken := ctx.GetHeader("Authorization")
		if rawToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
			ctx.Abort()
			return
		}

		// remove Bearer from token
		if len(rawToken) > 7 && rawToken[:7] == "Bearer " {
			rawToken = rawToken[7:]
		}

		// validate token
		token, err := jwt.Parse(rawToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Config.JWTSaltKey), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		// set user info to context
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			ctx.Set("user", claims)
		}

		ctx.Next() //pass
	}
}
