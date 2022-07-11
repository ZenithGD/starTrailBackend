package usercontroller

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	database "startrail/database"
	"startrail/models"
	crypt "startrail/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

// Register a new user into the system
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

// Login user if it exists and the provided credentials are correct.
func LoginUser(c *gin.Context) {
	// Get database instance
	db, err := database.GetDB()
	if err != nil {
		fmt.Printf("Error while getting database: %v", err)
		os.Exit(1)
	}

	// Bind JSON body to variable
	var userForm LoginForm
	err = c.BindJSON(&userForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user
	var dest models.User
	result := db.First(&dest, "nickname = ?", userForm.Nickname)
	if result != nil {
		err = bcrypt.CompareHashAndPassword([]byte(dest.Password), []byte(userForm.Password))
	}

	// Error if user doesn't exist or passwords don't match
	if err != nil || errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect username and/or password"})
		return
	}

	// Sign token on credential match
	token, err := crypt.SignToken(userForm.Nickname)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token sign error"})
		return
	}

	// Send response back
	c.Header("Authorization", fmt.Sprintf("Bearer %s", token))
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}
