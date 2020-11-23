package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shandysiswandi/go-rest-with-gin/controllers"
	"github.com/shandysiswandi/go-rest-with-gin/services"
)

func routes() *gin.Engine {
	// initial router engine
	r := gin.Default()

	// route not-found
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	// route home
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to home"})
	})

	// route user
	controllers.NewUser(services.NewUser()).Inject(r)

	return r
}
