package repositories

import (
	"github.com/shandysiswandi/go-rest-with-gin/models"
)

// GetAllUsers is
func GetAllUsers() []models.User {
	users := []models.User{
		models.User{
			ID:       "1",
			Name:     "name 1",
			Email:    "email1@email.com",
			Password: "password1",
		},
		models.User{
			ID:       "2",
			Name:     "name 2",
			Email:    "email2@email.com",
			Password: "password2",
		},
	}

	return users
}
