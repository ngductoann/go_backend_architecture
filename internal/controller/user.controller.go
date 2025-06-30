package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ngductoann/go_backend_architecture/internal/service"
	"github.com/ngductoann/go_backend_architecture/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc user controller
// us user service
// controller -> service -> repo -> model -> dbs
func (uc *UserController) GetUserByID(c *gin.Context) {
	// name := c.DefaultQuery("name", "ngductoann") // Get the "name" query parameter, defaulting to "ngductoann" if not provided
	// uid := c.Query("uid")

	// gin.Context is used to handle HTTP requests and responses
	// c.JSON(http.StatusOK, gin.H{
	// 	// gin.H is represent a map[string]interface{} (key-value pairs)
	// 	"message": uc.userService.GetInfoUser(),
	// 	"uuid":    uid,
	// 	"users":   []string{"cr7", "messi", "ngductoann"},
	// })

	response.SuccessResponse(c, 20001, []string{"ngductoann", "m10", "cr7"})
}
