package controllers

import (
	services "gym-api/backend/services/activityServices"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
