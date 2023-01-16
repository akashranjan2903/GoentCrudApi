package v1routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gocrud/db"
	"github.com/gocrud/src/controller"
)

func authRoutes(router *gin.RouterGroup) {
	sqlClient := db.NewSqlClient()
	// User Routes
	auth := router.Group("/auth")
	authApi := controller.NewAuthController(sqlClient)

	{

		auth.POST("/login", authApi.Login)
		auth.PATCH("/change-password", authApi.ChangePassword)

	}
}
