package controllers

import (
	services "gym-api/services/enrollmentService"
	"gym-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EnrollmentController struct {
	EnrollmentService services.EnrollmentServiceInterface
}
type EnrollmentControllersInterface interface {
	CreateEnrollment(ctx *gin.Context)
	GetEnrollment(ctx *gin.Context)
	CancelEnrollment(ctx *gin.Context)
	CheckEnrollment(ctx *gin.Context)
	GetActivityCapacity(ctx *gin.Context)
}

func (ec EnrollmentController) CreateEnrollment(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	if !exists {
		// Handle error - el usuario no tiene los claims necesarios para inscribirse
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Type assert to get your CustomClaims
	customClaims, ok := claims.(*utils.CustomClaims)
	if !ok {
		// Handle error - invalid claims type
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
		return
	}
	// Now you can use the user ID
	userID := customClaims.UserID
	type EnrollmentRequest struct {
		ActivityId string `json:"activity_id"`
	}
	var request EnrollmentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	activityId, _ := strconv.Atoi(request.ActivityId)
	err := ec.EnrollmentService.CreateEnrollment(userID, activityId)
	if err != nil {
		if err.Error() == "activity is at full capacity" {
			ctx.JSON(http.StatusConflict, gin.H{"error": "La actividad ha alcanzado su capacidad máxima"})
		} else if err.Error() == "user is already enrolled in this activity" {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Ya estás inscripto a esta actividad"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al inscribirse en la actividad"})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "enrollment created successfully"})
}

func (ec EnrollmentController) CancelEnrollment(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	customClaims, ok := claims.(*utils.CustomClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
		return
	}

	userID := customClaims.UserID
	type CancelEnrollmentRequest struct {
		ActivityId int `json:"activity_id"`
	}
	var request CancelEnrollmentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ec.EnrollmentService.CancelEnrollment(userID, request.ActivityId)
	if err != nil {
		if err.Error() == "user is not enrolled in this activity" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No estás inscripto a esta actividad"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al cancelar la inscripción"})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "enrollment cancelled successfully"})
}

func (ec EnrollmentController) CheckEnrollment(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	customClaims, ok := claims.(*utils.CustomClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
		return
	}

	userID := customClaims.UserID
	activityIdStr := ctx.Query("activity_id")
	if activityIdStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "activity_id is required"})
		return
	}

	activityId, err := strconv.Atoi(activityIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity_id"})
		return
	}

	isEnrolled, err := ec.EnrollmentService.IsEnrolled(userID, activityId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking enrollment"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"is_enrolled": isEnrolled})
}

func (ec EnrollmentController) GetActivityCapacity(ctx *gin.Context) {
	activityIdStr := ctx.Query("activity_id")
	if activityIdStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "activity_id is required"})
		return
	}

	activityId, err := strconv.Atoi(activityIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity_id"})
		return
	}

	availableSpots, err := ec.EnrollmentService.GetAvailableSpots(activityId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting activity capacity"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"available_spots": availableSpots,
	})
}

func (ec EnrollmentController) GetEnrollment(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	customClaims, ok := claims.(*utils.CustomClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
		return
	}

	activities, err := ec.EnrollmentService.GetUserEnrollments(customClaims.UserID)
	if err != nil {
		// Devuelve un array vacío en caso de error
		ctx.JSON(http.StatusOK, []interface{}{})
		return
	}

	ctx.JSON(http.StatusOK, activities)
}
