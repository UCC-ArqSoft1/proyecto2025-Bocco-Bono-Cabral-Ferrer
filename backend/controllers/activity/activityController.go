package controllers

import (
	"gym-api/domain"
	services "gym-api/services/activityServices"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	ActivityService services.ActivityServiceInterface
}
type ActivityControllersInterface interface {
	GetActivities(ctx *gin.Context)
	GetActivitiesByFilters(ctx *gin.Context)
	GetActivityByID(ctx *gin.Context)
	CreateActivity(ctx *gin.Context)
	UpdateActivity(ctx *gin.Context)
}

func (ac ActivityController) GetActivities(ctx *gin.Context) {
	if _, ok := ctx.GetQuery("keyword"); !ok {
		dtoActivities, err := ac.ActivityService.GetActivities()
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, dtoActivities)
		return
	}
}
func (ac ActivityController) GetActivitiesByFilters(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	dtoActivities, err := ac.ActivityService.GetActivitiesByFilters(keyword)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, dtoActivities)
}
func (ac ActivityController) GetActivityByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	dtoActivity, err := ac.ActivityService.GetactivityByID(idInt)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, dtoActivity)
}
func (ac ActivityController) CreateActivity(ctx *gin.Context) {
	var activityRequest domain.Activity
	if err := ctx.BindJSON(&activityRequest); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := ac.ActivityService.CreateActivity(activityRequest.Name,
		activityRequest.Description,
		activityRequest.Capacity,
		activityRequest.Category,
		activityRequest.Profesor,
		activityRequest.ImageUrl,
		activityRequest.Schedules)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Activity created successfully"})
}
func (ac ActivityController) DeleteActivity(ctx *gin.Context) {
	// Obtener el id desde el path
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	// Llamar al servicio para borrar la actividad
	err = ac.ActivityService.DeleteActivity(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
}
func (ac ActivityController) UpdateActivity(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid activity ID"})
		return
	}

	var activityRequest domain.Activity
	if err := ctx.BindJSON(&activityRequest); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = ac.ActivityService.UpdateActivity(
		id,
		activityRequest.Name,
		activityRequest.Description,
		activityRequest.Capacity,
		activityRequest.Category,
		activityRequest.Profesor,
		activityRequest.ImageUrl,
		activityRequest.Schedules,
	)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Activity updated successfully"})
}
