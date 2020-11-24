package repositories

import (
	"github.com/shandysiswandi/go-rest-with-gin/config"
	"github.com/shandysiswandi/go-rest-with-gin/models"
)

// GetAllUsers is
func GetAllUsers(mu *[]models.User) error {
	if err := config.DB.Find(mu).Error; err != nil {
		return err
	}

	return nil
}
