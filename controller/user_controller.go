package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sidz111/jwt-second-project/models"
	"github.com/sidz111/jwt-second-project/service"
)

type UserController struct {
	serv service.UserService
}

func NewUserController(serv service.UserService) *UserController {
	return &UserController{serv: serv}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id, err := c.serv.CreateUser(ctx.Request.Context(), &user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"id": id})

}
func (c *UserController) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid user id"})
		return
	}
	user, err := c.serv.GetUserByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(200, user)
}
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.serv.GetAllUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, users)
}
func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid user id"})
		return
	}
	if err := c.serv.DeleteUser(ctx.Request.Context(), id); err != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(204, gin.H{
		"message": "user deleted successfully",
	})
}
func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.serv.UpdateUser(ctx.Request.Context(), &user); err != nil {
		ctx.JSON(404, gin.H{"error": "user not found"})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "user updated successfully",
	})
}
