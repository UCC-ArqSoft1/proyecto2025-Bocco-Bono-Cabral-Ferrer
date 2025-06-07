package controllers

import (
	services "gym-api/backend/services/enrollmentService"
	"gym-api/backend/utils"
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
}

func (ec EnrollmentController) CreateEnrollment(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	if !exists {
		// Handle error - no claims found
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "enrollment created successfully"})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, activities)
}
