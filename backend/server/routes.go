package server

import (
	"fmt"
	"net/http"

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

	testApi(router);
	getAllData(router, db);
	addProduct(router, db);
	updateProduct(router, db);

	router.Run(":8080");
}


// Funções de Rota
func testApi(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Api funcionando",
		})
	})
}

func getAllData(router *gin.Engine, db *gorm.DB) {
	router.GET("/products", func(c *gin.Context) {

		alldata, err := database.GetAllProduct(db); if err != nil {
			fmt.Printf("An error occours trying to get the data: %s\n", err);
		}

		c.JSON(http.StatusOK, gin.H{
			"data": alldata,
		})
	})
}

func addProduct(router *gin.Engine, db *gorm.DB) {
	
	router.POST("/product", func(c *gin.Context) {

		product := database.Product{};

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to add data",
			})
		}

		database.AddProduct(db, product);

		c.JSON(http.StatusOK, gin.H{
			"message": "Produto adicionado com sucesso",
		})
	})
}

func updateProduct(router *gin.Engine, db *gorm.DB) {



	router.PUT("/product/:id", func(c *gin.Context) {

		product := database.Product{};

		str_id := c.Param("id");

		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error trying to add data",
			})
		}

		database.UpdateProduct(db, uuid.MustParse(str_id), product);

		c.JSON(http.StatusOK, gin.H{
			"message": "Produto adicionado com sucesso",
		})
	})
}