package app

import (
	"fmt"
	"log"
	"os"

	"github.com/shandysiswandi/go-rest-with-gin/config"
	"github.com/shandysiswandi/go-rest-with-gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// StartApp is
func StartApp() {
	// setting environment
	config.InitEnv()

	// inital db
	var errDB error
	config.DB, errDB = gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if errDB != nil {
		fmt.Println("Status:", errDB)
	}
	config.DB.AutoMigrate(&models.User{})

	// initial routing
	r := routes()

	// run app
	if err := r.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatal(err.Error())
	}
}
