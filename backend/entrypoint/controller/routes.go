package server

import (
	"fmt"
	_"os"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"github.com/Kaua-Matheus/fitstore/backend/core/dataprovider"
	"github.com/Kaua-Matheus/fitstore/backend/config/utils"
	"github.com/Kaua-Matheus/fitstore/backend/entrypoint/controller/handler"
)

func Run() {

	// IP
	ip, err := utils.GetLocalIP();

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{
			"http://localhost:3000", 
			fmt.Sprintf("http://%s:3000", ip),
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	db, err := db.NewConnection();
	if err != nil {
		fmt.Printf("Error database connection %s\n", err)
		return
	} else {
		fmt.Println("Conectado.")
	}

	handler.Debug(router);
	handler.Product(router, db);
	handler.Image(router, db);
	handler.SetupFileRoutes(router);
	handler.User(router, db);

	// Setup
	err = godotenv.Load(); if err != nil {
		fmt.Printf("Error setup %s\n", err)
		// return // Adicionar retorno de erro
	}

	fmt.Printf("[\033[32m Info \033[0m] - Server running in \033[32m %s:80 \033[0m\n", ip)
		router.Run(":80")
}