package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Kaua-Matheus/fitstore/backend/database"
	"github.com/Kaua-Matheus/fitstore/backend/server/handler"
)

func Run() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	db, err := database.NewConnection()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Conectado.")
	}

	handler.Debug(router)
	handler.Product(router, db)
	handler.Image(router, db)
	handler.SetupFileRoutes(router)

	router.Run(":8080")
}