package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sidz111/jwt-second-project/controller"
	dbconfig "github.com/sidz111/jwt-second-project/dbConfig"
	"github.com/sidz111/jwt-second-project/models"
	"github.com/sidz111/jwt-second-project/repository"
	"github.com/sidz111/jwt-second-project/routes"
	"github.com/sidz111/jwt-second-project/service"
)

func main() {
	if err := dbconfig.ConnectDB(); err != nil {
		panic("Failed to connect to database")
	}
	r := gin.Default()

	dbconfig.DB.AutoMigrate(&models.User{})
	repos := repository.NewUserRepository(dbconfig.DB)
	serv := service.NewUserService(repos)
	authController := controller.AuthController{}
	userController := controller.NewUserController(serv)

	routes := routes.SetupRoutes(r, authController, userController)
	routes.Run(":8080")
}
