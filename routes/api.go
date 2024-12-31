package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rudSarkar/reminder/controllers"
)

func RegisterAPIRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controllers.Health)
	}
}
