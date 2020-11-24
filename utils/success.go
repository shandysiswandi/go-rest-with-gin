package utils

import (
	"github.com/gin-gonic/gin"
)

type appSuccess struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success is
func Success(c *gin.Context, s int, m string, d interface{}) {
	c.JSON(s, appSuccess{
		Error:   false,
		Message: m,
		Data:    d,
	})
}
