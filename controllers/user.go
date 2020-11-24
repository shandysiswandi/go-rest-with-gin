package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shandysiswandi/go-rest-with-gin/services"
)

type userController struct {
	service services.UserService
}

// NewUser is constructor of this file
func NewUser(s services.UserService) Controller {
	return &userController{service: s}
}

func (uc userController) Inject(e *gin.Engine) {
	user := e.Group("api/v1/users")
	{
		user.GET("/", uc.service.GetAll)
		user.GET("/:id", uc.service.GetOne)
		// user.POST("/", uc.service.CreateOne)
		// user.PUT("/:id", uc.service.UpdateOne)
		// user.DELETE("/:id", uc.service.DeleteOne)
	}
}
