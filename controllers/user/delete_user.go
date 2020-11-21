package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shandysiswandi/go-rest-with-gin/utils"
)

// DeleteUser is
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, utils.AppSuccess{
		Error:   false,
		Message: "welcome to delete user",
		Data:    id,
	})
}
