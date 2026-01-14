package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Debug(router *gin.Engine) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Api funcionando",
		})
	})
}