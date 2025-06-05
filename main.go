package main

import (
	activityClients "gym-api/backend/clients/activityclient"
	userClients "gym-api/backend/clients/userClients"
	userControllers "gym-api/backend/controllers/user"

	enrollmentClients "gym-api/backend/clients/enrollmentClient"
	enrollmentController "gym-api/backend/controllers/enrollment"
	"gym-api/backend/db"
	activityServices "gym-api/backend/services/activityServices"
	enrollmentServices "gym-api/backend/services/enrollmentService"
	userServices "gym-api/backend/services/userServices"

	activityControllers "gym-api/backend/controllers/activity"
	"gym-api/backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	dbInstance := db.MySQLDB{}
	dbInstance.Connect()
	dbInstance.Migrate()
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5178"} // Add your frontend URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

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

	enrollmentRepo := enrollmentClients.EnrollmentRepository{DB: dbInstance.DB}
	enrollmentService := enrollmentServices.EnrollmentService{Repo: enrollmentRepo}
	enrollmentController := enrollmentController.EnrollmentController{EnrollmentService: enrollmentService}

	router.POST("/users/login", userController.Login)
	router.POST("/users/register", userController.Register)

	router.GET("/activities", activityController.GetActivities)
	router.GET("/activities/:id", activityController.GetActivityByID)
	router.POST("/activities", activityController.CreateActivity)
	router.DELETE("/activities/:id", activityController.DeleteActivity)
	router.PUT("/activities/:id", activityController.UpdateActivity)

	authorized := router.Group("/")

	authorized.Use(middleware.AuthMiddleware())

	authorized.POST("/enrollment", enrollmentController.CreateEnrollment)
	//router.GET("/enrollment", enrollmentController.GetEnrollment)

	log.Info("Starting server")
	router.Run(":8080")
}
