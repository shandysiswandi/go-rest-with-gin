package app

import (
	"fmt"
	"os"

	"github.com/shandysiswandi/go-rest-with-gin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB as
var DB *gorm.DB

// InitDB is
func InitDB() {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		fmt.Println("Status:", err)
	}

	DB = db

	DB.AutoMigrate(&models.User{})
}
