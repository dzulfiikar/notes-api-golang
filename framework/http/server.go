package http

import (
	"fmt"
	"log"
	"os"

	gateway "notes-api-golang/adapter/gateways"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func StartServer() {
	loadEnv()

	r := gin.Default()

	gateway.CreateRoute(&r.RouterGroup)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r.Run(":" + port)
	fmt.Println("Server running on port " + port)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("Error loading .env file")
	}

}
