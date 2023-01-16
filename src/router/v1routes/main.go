package v1routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func InitRouter() *gin.Engine {
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	log := logrus.New()
	router.Use(cors.Default())
	// Setup CORS for the API
	// router.Use(middleware.Cors())
	router.Use(ginlogrus.Logger(log), gin.Recovery())

	// Set the API Version
	api := router.Group("/api")
	userRoutes(api)
	carRoutes(api)
	authRoutes(api)

	return router
}
