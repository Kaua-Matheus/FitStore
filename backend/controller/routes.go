package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Kaua-Matheus/fitstore/backend/model"
	"github.com/Kaua-Matheus/fitstore/backend/controller/handler"
)

func Run() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5174"}, // Adicionadas duas origens
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	db, err := model.NewConnection()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Conectado.")
	}

	handler.Debug(router);
	handler.Product(router, db);
	handler.Image(router, db);
	handler.SetupFileRoutes(router);
	handler.User(router, db);

	router.Run(":8080")
}