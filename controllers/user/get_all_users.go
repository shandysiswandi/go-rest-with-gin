package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shandysiswandi/go-rest-with-gin/utils"
)

// GetAllUsers is
func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, utils.AppSuccess{
		Error:   false,
		Message: "welcome to all users",
	})
}
