package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dbconfig "github.com/sidz111/jwt-second-project/dbConfig"
	"github.com/sidz111/jwt-second-project/models"
	"github.com/sidz111/jwt-second-project/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
}

func (ac *AuthController) login(c *gin.Context) {
	var user models.User
	var foundUser models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	err = dbconfig.DB.Where("id = ?", user.ID).First(&foundUser).Error
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}
	dbconfig.DB.Where("id = ?", user.ID).First(&foundUser)
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(foundUser.Password),
	); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid Password",
		})
		return
	}

	token, err := utils.GenerateJWT(user.Name, int(user.ID))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token})

}
