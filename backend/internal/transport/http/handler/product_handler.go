package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	config "github.com/Kaua-Matheus/fitstore/backend/internal/config"
	domain "github.com/Kaua-Matheus/fitstore/backend/internal/domain"
	"github.com/Kaua-Matheus/fitstore/backend/internal/transport/http/dto/request"
	usecase "github.com/Kaua-Matheus/fitstore/backend/internal/usecase"
)

func Product(router *gin.Engine, db *gorm.DB) {

	ip, err := config.GetLocalIP()
	if err != nil {
		ip = "localhost"
	}

	// GET
	router.GET("/product", func(ctx *gin.Context) {

		products, err := usecase.GetAllProduct(db)
		if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err)
		}

		var resultList []map[string]any

		for _, product := range products {
			image, err := usecase.GetImage(db, product.IdImage)
			if err != nil {
				fmt.Printf("An error occours trying to get the image: %s\n", err)
				return
			} else {
				resultList = append(resultList, map[string]any{
					"product":   product,
					"url_image": fmt.Sprintf("http://%s:80/files/%s/%s%s", ip, image.FilePath, image.FileName, image.ContentType),
				})
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": resultList,
		})
	})

	router.GET("/product/:id", func(ctx *gin.Context) {

		id := ctx.Param("id")

		product, err := usecase.GetProduct(db, uuid.MustParse(id))
		if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err)
			return
		}

		image, err := usecase.GetImage(db, product.IdImage)
		if err != nil {
			fmt.Printf("An error occours trying to get the image: %s\n", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
			"image": gin.H{
				"file_name": image.FileName + image.ContentType,
				"url":       fmt.Sprintf("http://%s:80/files/%s/%s%s", ip, image.FilePath, image.FileName, image.ContentType),
			},
		})
	})

	// POST
	router.POST("/product", func(ctx *gin.Context) {

		productReq := request.CreateProductRequest{}
		if err := ctx.BindJSON(&productReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to add data",
			})
		}

		// User uuid added manually
		// Handle err!
		product_id, _ := uuid.NewUUID()

		product := domain.Product{
			ID:                 product_id,
			ProductName:        productReq.ProductName,
			ProductDescription: productReq.ProductDescription,
			ProductPrice:       productReq.ProductPrice,
			IdImage:            productReq.IdImage,
		}

		usecase.AddProduct(db, product)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Produto adicionado com sucesso",
		})
	})

	// PUT
	router.PUT("/product/:id", func(ctx *gin.Context) {

		product := domain.Product{}

		str_id := ctx.Param("id")

		if err := ctx.BindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to add data",
			})
		}

		usecase.UpdateProduct(db, uuid.MustParse(str_id), product)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Produto adicionado com sucesso",
		})
	})

}
