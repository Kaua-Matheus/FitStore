package server

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Kaua-Matheus/fitstore/backend/model"
	"github.com/Kaua-Matheus/fitstore/backend/controller/utils"
	"github.com/Kaua-Matheus/fitstore/backend/controller/handler"
)

func Run() {

	ip, err := utils.GetLocalIP();

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{
			"http://localhost:5173", 
			"http://localhost:5174",
			fmt.Sprintf("http://%s:5173", ip),
		},
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

	if err == nil {
		fmt.Printf("[\033[32m Info \033[0m] - Server running in \033[32m %s:8080 \033[0m\n", ip)
		router.Run(ip + ":8080")
	} else {
		fmt.Printf("[\033[31m Error \033[0m] - Couldn't run the server\n")
	}
}