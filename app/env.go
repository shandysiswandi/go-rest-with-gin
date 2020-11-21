package app

import (
	"log"

	"github.com/joho/godotenv"
)

func env() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
}
