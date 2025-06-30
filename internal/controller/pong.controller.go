package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "ngductoann") // Get the "name" query parameter, defaulting to "ngductoann" if not provided
	uid := c.Query("uid")

	// gin.Context is used to handle HTTP requests and responses
	c.JSON(http.StatusOK, gin.H{
		// gin.H is represent a map[string]interface{} (key-value pairs)
		"message": "pong.hhhh..ping" + name,
		"uuid":    uid,
		"users":   []string{"cr7", "messi", "ngductoann"},
	})
}
