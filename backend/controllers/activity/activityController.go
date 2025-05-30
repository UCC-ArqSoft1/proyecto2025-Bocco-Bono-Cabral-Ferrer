package controllers

import (
	"gym-api/backend/domain"
	services "gym-api/backend/services/activityServices"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	ActivityService services.ActivityServiceInterface
}
type ActivityControllersInterface interface {
	GetActivities(ctx *gin.Context)
	GetActivityByID(ctx *gin.Context)
	CreateActivity(ctx *gin.Context)
}

func (ac ActivityController) GetActivities(ctx *gin.Context) {
	if keyword, ok := ctx.GetQuery("keyword"); !ok {
		dtoActivities, err := ac.ActivityService.GetActivities()
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, dtoActivities)
		return
	} else {
		dtoActivities, err := ac.ActivityService.GetActivitiesByFilters(keyword)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, dtoActivities)
	}

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
	validSchedules := []domain.ActivitySchedule{}
	for _, s := range activityRequest.Schedules {
		if s.Day != "" && s.StartTime != "" && s.EndTime != "" {
			validSchedules = append(validSchedules, s)
		}
	}
	err := ac.ActivityService.CreateActivity(activityRequest.Name, activityRequest.Description, activityRequest.Capacity, activityRequest.Category, activityRequest.Profesor, validSchedules)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Activity created successfully"})
}
