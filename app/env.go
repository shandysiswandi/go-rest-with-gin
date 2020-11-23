package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func env() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	// check if development or production
	if os.Getenv("APP_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
}
