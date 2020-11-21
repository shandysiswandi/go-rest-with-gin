package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shandysiswandi/go-rest-with-gin/utils"
)

// CreateUser is
func CreateUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, utils.AppSuccess{
		Error:   false,
		Message: "welcome to create user",
		Data:    id,
	})
}
