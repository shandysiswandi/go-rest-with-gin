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
	GetOne(c *gin.Context)
	// CreateOne(c *gin.Context)
	// UpdateOne(c *gin.Context)
	// DeleteOne(c *gin.Context)
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

func (us *userService) GetOne(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := repositories.GetOneUser(&user, id)

	if err != nil {
		utils.Success(c, http.StatusNotFound, "user not found", user)
		return
	}

	utils.Success(c, http.StatusOK, "get user by id", user)
}
