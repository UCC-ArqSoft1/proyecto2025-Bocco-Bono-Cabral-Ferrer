package main

import (
	"gym-api/backend/db"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"

	activitiesController "gym-api/backend/controllers/activity"
	userController "gym-api/backend/controllers/user"
)

func main() {
	db.InitDatabase()
	router := gin.Default()
	router.POST("/users/login", userController.Login)
	router.POST("/users/register", userController.Register)
	router.GET("/activities", activitiesController.GetActivities)
	router.GET("/activities/:id", activitiesController.GetActivityByID)

	//mapUrls()

	log.Info("Starting server")
	router.Run(":8080")

	//app.StartRoute()
	// router := gin.Default
	// router.Use(controllers.AllowCORS)
	// router.POST("/login", userControllers.Login)
}
