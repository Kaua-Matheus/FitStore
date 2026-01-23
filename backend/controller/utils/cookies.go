package utils

import (
	"time"

	"github.com/gin-gonic/gin"
)

const CookieName = "auth_token"

// Define o cookie JWT
func SetAuthCookie(ctx *gin.Context, token string) {
	ctx.SetCookie(
		CookieName,
		token,
		int(24*time.Hour.Seconds()), // Tempo de duração do token
		"/",                         // Path - Disponível para toda a aplicação
		"",                          // Domain - vazio = apenas domínio atual
		false,                       // Secure - true apenas em produção com https (adicionar no env com variável)
		true,                        // HttpOnly - impede acesso via Javascript
	)
}

// Recupera o cookie JWT
func GetAuthCookie(ctx *gin.Context) (string, error) {
	return ctx.Cookie(CookieName)
}

// Remove o cookie JWT (LogOut)
func ClearAuthCookie(ctx *gin.Context) {
	ctx.SetCookie(
		CookieName,
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
}
