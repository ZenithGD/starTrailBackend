package usercontroller

import (
	"fmt"
	"net/http"
	"os"
	database "startrail/database"
	"startrail/models"
	crypt "startrail/utils"

	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
	Descr    string `json:"description"`
}

type LoginForm struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	// Get database instance
	db, err := database.GetDB()
	if err != nil {
		fmt.Printf("Error while getting database: %v", err)
		os.Exit(1)
	}

	// Bind JSON body to variable
	var userForm RegisterForm
	err = c.BindJSON(&userForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encryptedPass, err := crypt.Encrypt(userForm.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	userData := models.User{
		Nickname: userForm.Nickname,
		Email:    userForm.Email,
		Password: string(encryptedPass),
		Descr:    userForm.Descr,
	}
	// Create record
	err = db.Create(&userData).Error
	if err != nil {
		// Handle error
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}

func LoginUser(c *gin.Context) {
	// Get database instance
	db, err := database.GetDB()
	if err != nil {
		fmt.Printf("Error while getting database: %v", err)
		os.Exit(1)
	}

	var userForm LoginForm
	err = c.BindJSON(&userForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dest models.User
	db.First(&dest, "nickname = ?", userForm.Nickname)

}
