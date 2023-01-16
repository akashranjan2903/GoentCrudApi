package v1routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gocrud/db"
	"github.com/gocrud/src/controller"
)

func carRoutes(router *gin.RouterGroup) {
	sqlClient := db.NewSqlClient()
	// User Routes
	car := router.Group("/car")
	carApi := controller.NewcarController(sqlClient)

	{
		car.POST("", carApi.Create)
		car.GET("/:id", carApi.GetCarsbyOwner)
		car.DELETE("/:id", carApi.Delete)
	}
}
