package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shandysiswandi/go-rest-with-gin/models"
	"github.com/shandysiswandi/go-rest-with-gin/repositories"
	"github.com/shandysiswandi/go-rest-with-gin/utils"
)

// UserService is .
type UserService interface {
	GetAll(c *gin.Context)
}

// userService is private struct
type userService struct {
	users []models.User
}

// NewUser is contstructor
func NewUser() UserService {
	return &userService{}
}

func (us *userService) GetAll(c *gin.Context) {
	var users []models.User
	err := repositories.GetAllUsers(&users)

	if err != nil {
		utils.Success(c, http.StatusNotFound, "users empty", users)
		return
	}

	utils.Success(c, http.StatusOK, "welcome to all users", users)
}
