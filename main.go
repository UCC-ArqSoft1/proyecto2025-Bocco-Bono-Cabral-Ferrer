package main

import (
	activityClients "gym-api/backend/clients/activityclient"
	userClients "gym-api/backend/clients/userClients"
	userControllers "gym-api/backend/controllers/user"

	enrollmentController "gym-api/backend/controllers/enrollment"
	"gym-api/backend/db"
	activityServices "gym-api/backend/services/activityServices"
	userServices "gym-api/backend/services/userServices"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"

	activityControllers "gym-api/backend/controllers/activity"
)

func main() {

	dbInstance := db.MySQLDB{}
	dbInstance.Connect()
	dbInstance.Migrate()
	router := gin.Default()

	// Inyección manual de dependencias
	activityRepo := activityClients.ActivityRepository{DB: dbInstance.DB}
	activityService := activityServices.ActivityServiceImpl{Repo: activityRepo}
	activityController := activityControllers.ActivityController{ActivityService: activityService}
	//userRepo seria como el empleado que puede obtener la información de la base de datos
	//userService seria el jefe del empleado que le dice como trabajar
	//userController seria el que pasa la informacion entre el client y el userService
	userRepo := userClients.UserRepository{DB: dbInstance.DB}
	userService := userServices.UserService{Repo: userRepo}
	userController := userControllers.UserController{UserService: userService}

	router.POST("/users/login", userController.Login)
	router.POST("/users/register", userController.Register)

	router.GET("/activities", activityController.GetActivities)
	router.GET("/activities/:id", activityController.GetActivityByID)
	router.POST("/activities", activityController.CreateActivity)

	router.POST("/enrollment", enrollmentController.CreateEnrollment)
	router.GET("/enrollment", enrollmentController.GetEnrollment)

	log.Info("Starting server")
	router.Run(":8080")
}
