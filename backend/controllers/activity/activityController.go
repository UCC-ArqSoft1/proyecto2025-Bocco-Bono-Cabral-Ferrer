package controllers

import (
	"gym-api/backend/domain"
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
func CreateActivity(ctx *gin.Context) {
	var activity domain.Activity
	if err := ctx.BindJSON(&activity); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdActivity, err := services.CreateActivity(activity)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, createdActivity)
}
