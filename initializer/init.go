// init.go
package initializer

import (
	"fmt"
	"go-brain/database"
	"go-brain/internal/handler"
	"go-brain/internal/repo"
	"go-brain/internal/service"

	"github.com/gin-gonic/gin"
)

// Init 初始化应用
func Init() (*gin.Engine, func(), error) {
	gin.SetMode(gin.TestMode)
	db, err := database.InitDB()
	if err != nil {
		return nil, nil, err
	}

	// 初始化 Repositories
	userRepo := repo.NewUserRepo(db)
	thingRepo := repo.NewThingRepo(db)

	// 初始化 Services
	userService := service.NewUserService(userRepo)
	thingService := service.NewThingService(thingRepo)

	// 初始化 Handlers
	userHandler := handler.NewUserHandler(userService)
	thingHandler := handler.NewThingHandler(thingService)

	router := gin.Default()
	SetupRoutes(router, userHandler, thingHandler)
	fmt.Println("Initialization complete")
	return router, func() { db.Close() }, nil
}