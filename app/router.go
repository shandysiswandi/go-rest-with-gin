package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shandysiswandi/go-rest-with-gin/controllers/home"
	"github.com/shandysiswandi/go-rest-with-gin/controllers/user"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", home.GetHome)

	v1 := router.Group("api/v1")
	{
		users := v1.Group("users")
		{
			users.GET("", user.GetAllUsers)
			users.GET(":id", user.GetOneUser)
			users.POST("", user.CreateUser)
			users.PUT(":id", user.UpdateUser)
			users.DELETE(":id", user.DeleteUser)
		}
	}

	return router
}

func routes() {
	// check if development or production
	if os.Getenv("APP_MODE") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// initial all router
	r := initRouter()

	// check if any error, if not then run application
	if err := r.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatal(err.Error())
	}
}
