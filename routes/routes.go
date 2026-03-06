package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/jwt-second-project/controller"
)

func SetupRoutes(authController controller.AuthController, userController *controller.UserController) *gin.Engine {
	router := gin.Default()
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
	}
	return router
}
