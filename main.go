package main

import (
	"fmt"
	"os"
	usercontroller "startrail/controllers"
	"startrail/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
	godotenv.Load()
	_, err := database.GetDB()
	if err != nil {
		fmt.Printf("Error while initializing database: %v", err)
		os.Exit(1)
	}

	router := gin.Default()
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", usercontroller.RegisterUser)
		userGroup.POST("/login", usercontroller.LoginUser)
	}

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
