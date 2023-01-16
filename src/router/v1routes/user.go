package v1routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gocrud/db"
	"github.com/gocrud/src/controller"
)

func userRoutes(router *gin.RouterGroup) {
	sqlClient := db.NewSqlClient()
	// User Routes
	user := router.Group("/user")
	userApi := controller.NewUserController(sqlClient)

	{
		user.POST("", userApi.Create)
		user.DELETE("/:id", userApi.Delete)
		user.GET("/:id", userApi.Get)
		user.GET("/", userApi.GetAll)
		user.PATCH("/:id", userApi.Update)
	}
}
