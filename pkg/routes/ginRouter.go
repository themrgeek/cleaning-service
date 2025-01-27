package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/themrgeek/cleaning-service/pkg/controllers"
)

func InitializeRouter() *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/booking", controllers.CreateBooking)
	}

	return r
}
