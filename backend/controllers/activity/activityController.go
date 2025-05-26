package controllers

import (
	services "gym-api/backend/services/activityServices"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	ActivityService services.ActivityService
}
type ControllerMethods interface {
	GetActivities(ctx *gin.Context)
	GetActivityByID(ctx *gin.Context)
}

func (ac ActivityController) GetActivities(ctx *gin.Context) {
	if keyword, ok := ctx.GetQuery("keyword"); !ok {
		Activities, err := ac.ActivityService.GetActivities()
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, Activities)
		return
	} else {
		Activities, err := ac.ActivityService.GetActivitiesByFilters(keyword)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, Activities)
	}

}

func (ac ActivityController) GetActivityByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)
	activity, err := ac.ActivityService.GetactivityByID(idInt)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, activity)
}
