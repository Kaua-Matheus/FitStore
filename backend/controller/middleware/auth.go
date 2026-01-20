package middleware

import (
	"net/http"
	_"strings"

	"github.com/Kaua-Matheus/fitstore/backend/controller/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization");
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Auth header is empty",
			});
			ctx.Abort();
			return 
		}

		// Essa parte é importante para tokens envidos pelo frontend, pois os mesmos virão em um formato correto
		// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
		// tokenParts := strings.Split(authHeader, " ");
		// if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{
		// 		"error": "Token format is invalid",
		// 	});
		// 	ctx.Abort();
		// 	return 
		// }

		// token := tokenParts[1]
		err := utils.VerifyToken(authHeader); // Alterar authHeader para token
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			});
			ctx.Abort();
			return 
		}

		ctx.Next();
	}
}