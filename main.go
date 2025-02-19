package main

import (

	"go-brain/database"
	"go-brain/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.TestMode)
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// 初始化 Repositories
	userRepo := internal.NewUserRepo(db)
	thingRepo := internal.NewThingRepo(db)

	// 初始化 Services
	userService := internal.NewUserService(userRepo)
	thingService := internal.NewThingService(thingRepo)

	// 初始化 Handlers
	userHandler := internal.NewUserHandler(userService)
	thingHandler := internal.NewThingHandler(thingService)

	router := gin.Default()

	router.GET("/ping",func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
				"query":  c.Request.URL.Query(),
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

	router.Run(":8080")
}
