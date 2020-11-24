package app

import (
	"log"
	"os"

	"github.com/shandysiswandi/go-rest-with-gin/config"
)

// StartApp is
func StartApp() {
	// setting environment
	config.InitEnv()

	// inital db
	config.InitDB()

	// initial routing
	r := routes()

	// run app
	if err := r.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatal(err.Error())
	}
}
