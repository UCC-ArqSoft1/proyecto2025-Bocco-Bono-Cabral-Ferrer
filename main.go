package main

import (
	activityClients "gym-api/backend/clients/activityclient"
	userClients "gym-api/backend/clients/userClients"
	userControllers "gym-api/backend/controllers/user"
	"gym-api/backend/db"
	activityServices "gym-api/backend/services/activityServices"
	userServices "gym-api/backend/services/userServices"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"

	activityControllers "gym-api/backend/controllers/activity"
	//userController "gym-api/backend/controllers/user"
)

func main() {
	dbInstance := db.MySQLDB{}
	dbInstance.Connect()
	dbInstance.Migrate()
	router := gin.Default()

	// Inyección manual de dependencias
	mysql := activityClients.MySQLActivityRepository{DB: dbInstance.DB}
	activityService := activityServices.ActivityServices{ActivityClients: mysql}
	activityController := activityControllers.ActivityController{ActivityService: activityService}

	userRepo := userClients.MySQLUserRepository{DB: dbInstance.DB}
	userService := userServices.UserServices{UserClient: userRepo}
	userController := userControllers.UserController{UserService: userService}

	router.POST("/users/login", userController.Login)
	router.POST("/users/register", userController.Register)

	router.GET("/activities", activityController.GetActivities)
	router.GET("/activities/:id", activityController.GetActivityByID)

	log.Info("Starting server")
	router.Run(":8080")
}
