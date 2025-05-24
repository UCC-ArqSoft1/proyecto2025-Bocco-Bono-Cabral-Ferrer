package controllers

import (
	"errors"
	services "gym-api/backend/services/activityServices"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetActivities(ctx *gin.Context) {
	activities, err := services.GetActivities()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, activities)
}

func GetActivityByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	activity, err := services.GetactivityByID(idInt)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, activity)
}

func GetActivitiesByFilters(ctx *gin.Context) {
	filters := map[string]string{
		"keyword":  ctx.Query("keyword"),
		"category": ctx.Query("category"),
		"hour":     ctx.Query("hour"),
	}
	activities, err := services.GetActivitiesByFilters(filters)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron actividades"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, activities)
}
