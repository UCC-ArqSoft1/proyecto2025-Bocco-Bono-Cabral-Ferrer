package utils

import (
	"proyecto2025/backend/controllers"
	"proyecto2025/backend/dao"
	"proyecto2025/backend/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Instancias necesarias
	activityRepo := dao.NewActivityDAO() // si existe
	activityService := services.NewActivityService(activityRepo)
	activityController := controllers.NewActivityController(activityService)

	// Endpoints
	r.GET("/activities", activityController.GetAll)

	return r
}
