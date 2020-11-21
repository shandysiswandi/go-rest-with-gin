package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shandysiswandi/go-rest-with-gin/utils"
)

// UpdateUser is
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, utils.AppSuccess{
		Error:   false,
		Message: "welcome to update user",
		Data:    id,
	})
}
