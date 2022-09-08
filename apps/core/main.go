package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/server"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appMode := os.Getenv("APP_MODE")
	if appMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	server := server.NewServer()
	server.Start()
	// server.GracefulShutdown()
}
