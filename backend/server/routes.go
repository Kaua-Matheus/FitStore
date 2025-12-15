package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/Kaua-Matheus/fitstore/backend/database"
)

func Run() {
	// Criação do router
	router := gin.Default();
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// Colocar o db aqui!
	db, err := database.NewConnection(); if err != nil {
		fmt.Println(err);
		return;
	} else {
		fmt.Println("Conectado.");
	}

	Debug(router);
	Product(router, db);
	ProductImage(router, db);

	// Files
	setupFileRoutes(router);

	router.Run(":8080");
}


// Funções Debug
func Debug(router *gin.Engine) {

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Api funcionando",
		})
	})
}

// Funções de Produto
func Product(router *gin.Engine, db *gorm.DB) {
	
	// GET
	router.GET("/product", func(ctx *gin.Context) {

		products, err := database.GetAllProduct(db); if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err);
		}

		var resultList []map[string]any;

		for _, product := range products {
			image, err := database.GetImage(db, product.IdImage); if err != nil {
			fmt.Printf("An error occours trying to get the image: %s\n", err);
			return;
		} else {
			resultList = append(resultList, map[string]any{
				"product": product,
				"url_image": fmt.Sprintf("http://localhost:8080/files/%s/%s%s", image.FilePath, image.FileName, image.ContentType),
			})
		}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": resultList,
		})
	})

	router.GET("/product/:id", func(ctx *gin.Context) {

		id := ctx.Param("id");

		product, err := database.GetProduct(db, uuid.MustParse(id)); if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err);
			return;
		}

		image, err := database.GetImage(db, product.IdImage); if err != nil {
			fmt.Printf("An error occours trying to get the image: %s\n", err);
			return;
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
			"image": gin.H{
				"file_name": image.FileName + image.ContentType,
				"url": fmt.Sprintf("http://localhost:8080/files/%s/%s%s", image.FilePath, image.FileName, image.ContentType),
			},
		})
	})

	// POST
	router.POST("/product", func(ctx *gin.Context) {

		product := database.Product{};

		if err := ctx.BindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to add data",
			})
		}

		database.AddProduct(db, product);

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Produto adicionado com sucesso",
		})
	})

	// PUT
	router.PUT("/product/:id", func(ctx *gin.Context) {

		product := database.Product{};

		str_id := ctx.Param("id");

		if err := ctx.BindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to add data",
			})
		}

		database.UpdateProduct(db, uuid.MustParse(str_id), product);

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Produto adicionado com sucesso",
		})
	})

}

func ProductImage(router *gin.Engine, db *gorm.DB) {

	router.GET("/image/:id", func(ctx *gin.Context) {

		id := ctx.Param("id");
		image, err := database.GetImage(db, uuid.MustParse(id)); if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Erro ao tentar adquirir a imagem",
			})
			return;
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": image,
		})

	})

	router.POST("/image", func(ctx *gin.Context) {
		image := database.ProductImage{};

		if err := ctx.BindJSON(&image); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Erro ao tentar adicionar a imagem",
			});
			return;
		}

		database.AddProductImage(db, image);

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Imagem adicionada com sucesso",
		})

	})
}

// Funções de Arquivo
func setupFileRoutes(router *gin.Engine) {

	// GET
	// Adquire todas as imagens de dentro da pasta Produtos
	router.GET("/productfiles", func(ctx *gin.Context) {
		path := filepath.Join("..", "uploads", "Produtos");

		files, err := os.ReadDir(path); if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Error trying to read the Produtos path",
			})
			return;
		}

		var fileList []map[string]string;
		
		for _, file := range(files) {
			if !file.IsDir() {
				fileInfo := map[string]string{
					"filename": file.Name(),
					"url": 		fmt.Sprintf("http://localhost:8080/files/Produtos/%s", file.Name()),
				}
				fileList = append(fileList, fileInfo);
			}
		}

		ctx.JSON(http.StatusOK, gin.H{
			"files": fileList,
		})
	})

	router.GET("/files/:categoria/:filename", func (ctx *gin.Context)  {
		categoria := ctx.Param("categoria");
		filename := ctx.Param("filename");

		filePath := filepath.Join("..", "uploads", categoria, filename);

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("File not found: %s", filePath),
			})
			return;
		}

		ctx.File(filePath)
	})
}