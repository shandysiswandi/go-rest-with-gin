package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shandysiswandi/go-rest-with-gin/utils"
)

// GetHome is
func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, utils.AppSuccess{
		Error:   false,
		Message: "welcome to home",
	})
}
