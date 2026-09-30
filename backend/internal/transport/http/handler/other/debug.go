package other_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Debug(router *gin.Engine) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "The API is working",
		})
	})
}
