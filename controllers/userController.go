package usercontroller

import (
	"fmt"
	"net/http"
	"os"
	"startrail/database"
	"startrail/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	// Get database instance
	db, err := database.GetDB()
	if err != nil {
		fmt.Printf("Error while getting database: %v", err)
		os.Exit(1)
	}

	// Bind JSON body to variable
	var user models.User
	err = c.BindJSON(&user)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	fmt.Println(user)

	// Create record
	err = db.Create(&user).Error
	if err != nil {
		// Handle error
		c.Status(http.StatusConflict)
		return
	}

	c.Status(http.StatusOK)
}
