package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/model/entitie"
	"github.com/Kaua-Matheus/fitstore/backend/model/handler"
)

func Product(router *gin.Engine, db *gorm.DB) {

	// GET
	router.GET("/product", func(ctx *gin.Context) {

		products, err := handler.GetAllProduct(db)
		if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err)
		}

		var resultList []map[string]any

		for _, product := range products {
			image, err := handler.GetImage(db, product.IdImage)
			if err != nil {
				fmt.Printf("An error occours trying to get the image: %s\n", err)
				return
			} else {
				resultList = append(resultList, map[string]any{
					"product":   product,
					"url_image": fmt.Sprintf("http://localhost:8080/files/%s/%s%s", image.FilePath, image.FileName, image.ContentType),
				})
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": resultList,
		})
	})

	router.GET("/product/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")

		product, err := handler.GetProduct(db, uuid.MustParse(id))
		if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err)
			return
		}

		image, err := handler.GetImage(db, product.IdImage)
		if err != nil {
			fmt.Printf("An error occours trying to get the image: %s\n", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
			"image": gin.H{
				"file_name": image.FileName + image.ContentType,
				"url":       fmt.Sprintf("http://localhost:8080/files/%s/%s%s", image.FilePath, image.FileName, image.ContentType),
			},
		})
	})

	// POST
	router.POST("/product", func(ctx *gin.Context) {

		product := entitie.Product{}

		if err := ctx.BindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to add data",
			})
		}

		handler.AddProduct(db, product)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Produto adicionado com sucesso",
		})
	})

	// PUT
	router.PUT("/product/:id", func(ctx *gin.Context) {

		product := entitie.Product{}

		str_id := ctx.Param("id")

		if err := ctx.BindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to add data",
			})
		}

		handler.UpdateProduct(db, uuid.MustParse(str_id), product)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Produto adicionado com sucesso",
		})
	})

}