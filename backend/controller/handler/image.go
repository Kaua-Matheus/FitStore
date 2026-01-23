package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/model/entitie"
	"github.com/Kaua-Matheus/fitstore/backend/model/handler"
)

func Image(router *gin.Engine, db *gorm.DB) {

	// GET
	// Adquire informações de uma entidade Image
	router.GET("/image/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")
		image, err := handler.GetImage(db, uuid.MustParse(id))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Erro ao tentar adquirir a imagem",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": image,
		})

	})

	// POST
	// Cria um registro de uma entidade Image
	router.POST("/image", func(ctx *gin.Context) {
		image := entitie.Image{}

		if err := ctx.BindJSON(&image); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Erro ao tentar adicionar a imagem",
			})
			return
		}

		handler.AddImage(db, image)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Imagem adicionada com sucesso",
		})

	})
}