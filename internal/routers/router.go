package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ngductoann/go_backend_architecture/internal/controller"
	"github.com/ngductoann/go_backend_architecture/internal/middlewares"
)

// Middleware AA
func AA() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before --> AA")
		ctx.Next() // Call the next handler in the chain
		fmt.Println("After --> AA")
	}
}

// Middleware BB
func BB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Before --> BB")
		ctx.Next() // Call the next handler in the chain
		fmt.Println("After --> BB")
	}
}

// Middleware CC
func CC(ctx *gin.Context) {
	// receiver gin.Context not gin.HandlerFunc (lastware)
	fmt.Println("Before --> CC")
	ctx.Next() // Call the next handler in the chain
	fmt.Println("After --> CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	// use the middleware
	r.Use(middlewares.AuthMiddleware(), AA(), BB(), CC) // A --> B --> C

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
