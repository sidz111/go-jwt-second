package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/jwt-second-project/controller"
)

func SetupRoutes(r *gin.Engine, authController controller.AuthController, userController *controller.UserController) *gin.Engine {
	router := gin.Default()
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
	}
	usersGroup := router.Group("/users")
	{
		usersGroup.POST("/", userController.CreateUser)
		usersGroup.GET("/:id", userController.GetUserByID)
		usersGroup.PUT("/:id", userController.UpdateUser)
		usersGroup.DELETE("/:id", userController.DeleteUser)
		usersGroup.GET("/", userController.GetAllUsers)
	}
	return router
}
