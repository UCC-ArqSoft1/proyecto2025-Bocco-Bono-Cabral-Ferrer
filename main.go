package main

import (
	clients "gym-api/backend/clients/activityclient"
	"gym-api/backend/db"
	services "gym-api/backend/services/activityServices"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"

	controllers "gym-api/backend/controllers/activity"
	//userController "gym-api/backend/controllers/user"
)

func main() {
	db.InitDatabase()
	router := gin.Default()

	// Inyección manual de dependencias
	mysql := clients.MySQL{DB: db.DB}
	activityService := services.ActivityServices{ActivityClients: mysql}
	activityController := controllers.ActivityController{ActivityService: activityService}

	/*// Rutas
	router.POST("/users/login", userController.Login)
	router.POST("/users/register", userController.Register)
	*/
	router.GET("/activities", activityController.GetActivities)
	router.GET("/activities/:id", activityController.GetActivityByID)

	log.Info("Starting server")
	router.Run(":8080")
}
