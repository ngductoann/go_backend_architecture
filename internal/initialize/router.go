package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/ngductoann/go_backend_architecture/global"
	"github.com/ngductoann/go_backend_architecture/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	// r.Use() // logging
	// r.Use() // cross
	// r.Use() // limiter global

	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/v1/2025")
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		managerRouter.InitUserRouter(MainGroup)
		managerRouter.InitAdminRouter(MainGroup)
	}

	return r
}
