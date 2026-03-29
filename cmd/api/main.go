package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rodrigorahman/cuidapet_api_ace/internal/config"
	"github.com/rodrigorahman/cuidapet_api_ace/internal/handler"
)

func main() {
	dbConfig := config.NewDatabaseConfig()
	db := dbConfig.Connect()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	defer sqlDB.Close()

	r := gin.Default()

	healthHandler := handler.NewHealthHandler()
	healthHandler.Register(r)

	log.Fatal(r.Run(":8080"))
}
