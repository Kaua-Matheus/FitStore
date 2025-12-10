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

	getAllImage(router, db);

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
	router.GET("/products", func(ctx *gin.Context) {

		alldata, err := database.GetAllProduct(db); if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err);
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": alldata,
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

// Funções de Arquivo
func setupFileRoutes(router *gin.Engine) {

	// GET
	router.GET("/files", func(ctx *gin.Context) {
		uploadsPath := filepath.Join("..", "uploads");

		files, err := os.ReadDir(uploadsPath); if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read the uploads path",
			})
			return;
		}

		var fileList []map[string]string;

		for _, file := range files {
			if !file.IsDir() {
				fileInfo := map[string]string{
					"filename": file.Name(),
					"url": 		fmt.Sprintf("http://localhost:8080/files/%s", file.Name()),
				}
				fileList = append(fileList, fileInfo);
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"files": fileList,
		})
	})

	router.GET("/files/:filename", func (ctx *gin.Context)  {
		filename := ctx.Param("filename");

		filePath := filepath.Join("..", "uploads", filename);

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "File not found",
			})
			return;
		}

		ctx.File(filePath)
	})
}

// Funções de Imagem
func getAllImage(router *gin.Engine, db *gorm.DB) {
	router.GET("/images", func(ctx *gin.Context) {

		alldata, err := database.GetAllImage(db); if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err);
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": alldata,
		})
	})
}