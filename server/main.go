package main

import (
	"log"
	"server/config"
	"server/controller"
	"server/repo"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer config.CloseDB()

	serverRepository := repo.NewUserRepository(db)
	serverController := controller.NewServerController(serverRepository)
	
	config.ListenAndServeGrpc(serverController)
}