package app

import (
	"log"
	"os"
)

// StartApp is
func StartApp() {
	// setting environment
	env()

	// routing
	r := routes()

	// run app
	if err := r.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatal(err.Error())
	}
}
