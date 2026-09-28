package middleware

import (
	"net/http"
	"strings"

	cookies "github.com/Kaua-Matheus/fitstore/backend/pkg/cookies"
	jwt "github.com/Kaua-Matheus/fitstore/backend/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string

		cookieToken, err := cookies.GetAuthCookie(ctx)
		if err == nil && cookieToken != "" {
			token = cookieToken
		} else {
			authHeader := ctx.GetHeader("Authorization")
			if authHeader == "" {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "Auth header is empty",
				})
				ctx.Abort()
				return
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "Token format is invalid",
				})
				ctx.Abort()
				return
			}

			token = tokenParts[1]
		}

		err = jwt.VerifyToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
