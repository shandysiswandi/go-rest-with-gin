package config

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
	var errDB error
	DB, errDB = gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{PrepareStmt: true})

	if errDB != nil {
		fmt.Println("Status:", errDB)
	}

	DB.AutoMigrate(&models.User{})
}
