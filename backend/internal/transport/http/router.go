package router

import (
	"fmt"
	_ "os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	config "github.com/Kaua-Matheus/fitstore/backend/internal/config"
	db "github.com/Kaua-Matheus/fitstore/backend/internal/repository/postgres"

	handler "github.com/Kaua-Matheus/fitstore/backend/internal/transport/http/handler"
	other_handler "github.com/Kaua-Matheus/fitstore/backend/internal/transport/http/handler/other"
)

func Run() {

	// IP
	ip, err := config.GetLocalIP()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			fmt.Sprintf("http://%s:3000", ip),
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	db, err := db.NewConnection()
	if err != nil {
		fmt.Printf("Error database connection %s\n", err)
		return
	} else {
		fmt.Println("Conectado.")
	}

	handler.User(router, db)
	handler.Product(router, db)

	other_handler.Debug(router)
	other_handler.Image(router, db)
	other_handler.SetupFileRoutes(router)

	// Setup
	err = godotenv.Load()
	if err != nil {
		fmt.Printf("Error setup %s\n", err)
		// return // Adicionar retorno de erro
	}

	fmt.Printf("[\033[32m Info \033[0m] - Server running in \033[32m %s:80 \033[0m\n", ip)
	router.Run(":80")
}
