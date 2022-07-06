package usercontroller

import (
	"fmt"
	"net/http"
	"os"
	database "startrail/database"
	"startrail/models"

	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email"    binding:"required"`
	Password string `json:"password" binding:"required"`
	Descr    string `json:"description"`
}

func RegisterUser(c *gin.Context) {
	// Get database instance
	db, err := database.GetDB()
	if err != nil {
		fmt.Printf("Error while getting database: %v", err)
		os.Exit(1)
	}

	// Bind JSON body to variable
	var user RegisterForm
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData := models.User{
		Nickname: user.Nickname,
		Email:    user.Email,
		Password: user.Password,
		Descr:    user.Descr,
	}
	// Create record
	err = db.Create(&userData).Error
	if err != nil {
		// Handle error
		c.Status(http.StatusConflict)
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}
