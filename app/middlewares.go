package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func middlewares(r *gin.Engine) {
	r.Use(cors.Default())
}
