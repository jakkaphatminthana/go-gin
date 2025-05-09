package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jakkaphatminthana/go-gin/config"
	"github.com/jakkaphatminthana/go-gin/pkg/custom"
)

type AuthorizationMiddleware struct {
	conf *config.Config
}

func NewAuthorizationMiddleware(conf *config.Config) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{conf: conf}
}

func (m *AuthorizationMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 1. get token from Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.Error(custom.ErrorUnauthorized("authorization token not found"))
			ctx.Abort()
			return
		}

		// 2. check and strip Bearer
		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			ctx.Error(custom.ErrorUnauthorized("authorization token must begin with Bearer"))
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, prefix)

		// 3. parse and validate token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(m.conf.EVNValue.JWTSaltKey), nil
		})
		if err != nil || !token.Valid {
			ctx.Error(custom.ErrorUnauthorized("invalid or expired token"))
			ctx.Abort()
			return
		}

		// 4. set claims to context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx.Set("user", claims)
		}

		ctx.Next()
	}
}
