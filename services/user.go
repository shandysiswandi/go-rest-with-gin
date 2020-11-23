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
	c.JSON(http.StatusOK, utils.AppSuccess{
		Error:   false,
		Message: "welcome to all users",
		Data:    repositories.GetAllUsers(),
	})
}
