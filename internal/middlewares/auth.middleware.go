package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ngductoann/go_backend_architecture/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(ctx, response.ErrInvalidToken, "Unauthorized access")
			ctx.Abort() // Stop the request processing
			return
		}
		ctx.Next() // Continue to the next handler if the token is valid
	}
}
