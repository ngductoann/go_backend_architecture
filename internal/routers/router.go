package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ngductoann/go_backend_architecture/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// set router with grouping
	v1 := r.Group(("/v1/2025"))
	{
		// This is a route group for version 1 of the API, with a base path of "/v1/2025"
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.GET("/user/1", controller.NewUserController().GetUserByID) // Register the route for GET requests to "/v1/2025/ping"
		// v1.POST("/ping", Pong)   // Register the route for POST requests to "/v1/2025/ping"
		// v1.PUT("/ping", Pong)    // Register the route for PUT requests to "/v1/2025/ping"
		// v1.DELETE("/ping", Pong) // Register the route for DELETE requests to "/v1/2025/ping"
	}

	/*
		r.GET("/ping", Pong)    // Register the route for GET requests to "/ping"
		r.POST("/ping", Pong)   // Register the route for POST requests to "/ping"
		r.PUT("/ping", Pong)    // Register the route for PUT requests to "/ping"
		r.DELETE("/ping", Pong) // Register the route for DELETE requests to "/ping"
		// You can add more routes as needed
	*/

	return r
}
