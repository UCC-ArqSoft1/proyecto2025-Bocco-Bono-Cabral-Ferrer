package main

import (
	"gym-api/backend/db"

	"gym-api/backend/controllers"
	"gym-api/backend/dao"
	"gym-api/backend/services"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/log"

	userController "gym-api/backend/controllers/user"
)

func main() {
	db.InitDatabase()
	router := gin.Default()
	activityDAO := &dao.ActivityDAO{}
	activityService := services.NewActivityService(activityDAO)
	activityController := controllers.NuevoActivityController(activityService)

	router.POST("/users/login", userController.Login)
	router.POST("/users/register", userController.Register)
	router.GET("/activity", activityController.GetActivity)
	router.GET("/activity/:id", activityController.GetActivityByID)
	router.POST("/activity", activityController.CreateActivity)
	router.PUT("/activity/:id", activityController.UpdateActivity)
	router.DELETE("/activity/:id", activityController.DeleteActivity)
	//mapUrls()

	log.Info("Starting server")
	router.Run(":8080")

	//app.StartRoute()
	// router := gin.Default
	// router.Use(controllers.AllowCORS)
	// router.POST("/login", userControllers.Login)
}
