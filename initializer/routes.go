package initializer

import (
	"go-brain/internal/handler"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置路由
func SetupRoutes(router *gin.Engine, userHandler *handler.UserHandler, thingHandler *handler.ThingHandler) {
	// Ping Route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"query":   c.Request.URL.Query(),
		})
	})
	// Time Route
	router.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Data":      time.Now().Format("2006/01/02"),
			"Time":      time.Now().Format("15:04:05"),
			"TimeStamp": time.Now().Unix(),
		})
	})

	// User Routes
	router.POST("/users", userHandler.Create)
	router.GET("/users/:id", userHandler.GetByID)
	router.PUT("/users/:id", userHandler.Update)
	router.DELETE("/users/:id", userHandler.Delete)

	// Thing Routes
	router.POST("/things", thingHandler.Create)
	router.GET("/things/:id", thingHandler.GetByID)
	router.PUT("/things/:id", thingHandler.Update)
	router.DELETE("/things/:id", thingHandler.Delete)
}
