package main

import (
	activityClients "gym-api/clients/activityclient"
	userClients "gym-api/clients/userClients"
	userControllers "gym-api/controllers/user"
	"log"

	enrollmentClients "gym-api/clients/enrollmentClient"
	enrollmentController "gym-api/controllers/enrollment"
	activityServices "gym-api/services/activityServices"
	enrollmentServices "gym-api/services/enrollmentService"
	userServices "gym-api/services/userServices"

	activityControllers "gym-api/controllers/activity"
	db "gym-api/database"
	"gym-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dbInstance := db.MySQLDB{}
	dbInstance.Connect()
	dbInstance.Migrate()
	router := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Add your frontend URL
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
	enrollmentService := enrollmentServices.EnrollmentService{Repo: enrollmentRepo, ActivityRepo: activityRepo}
	enrollmentController := enrollmentController.EnrollmentController{EnrollmentService: enrollmentService}

	router.POST("/users/login", userController.Login)
	router.POST("/users/register", userController.Register)
	router.GET("/activities/search", activityController.GetActivitiesByFilters)
	router.GET("/activities/:id", activityController.GetActivityByID)
	router.GET("/activities", activityController.GetActivities)
	authorized := router.Group("/")

	authorized.Use(middleware.AuthMiddleware())

	authorized.POST("/enrollment", enrollmentController.CreateEnrollment)
	authorized.DELETE("/enrollment", enrollmentController.CancelEnrollment)
	authorized.GET("/enrollment/check", enrollmentController.CheckEnrollment)
	authorized.GET("/enrollment/capacity", enrollmentController.GetActivityCapacity)
	authorized.GET("/enrollments", enrollmentController.GetEnrollment)

	// Admin routes - require both authentication and admin privileges
	admin := router.Group("/")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())

	admin.POST("/activities", activityController.CreateActivity)
	admin.DELETE("/activities/:id", activityController.DeleteActivity)
	admin.PUT("/activities/:id", activityController.UpdateActivity)

	log.Println("Starting server")
	router.Run(":8080")
}
