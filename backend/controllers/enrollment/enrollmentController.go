package controllers

import (
	services "gym-api/backend/services/enrollmentService"

	"github.com/gin-gonic/gin"
)

type EnrollmentController struct {
	EnrollmentService services.EnrollmentServiceInterface
}
type EnrollmentControllersInterface interface {
	CreateEnrollment(ctx *gin.Context)
	GetEnrollment(ctx *gin.Context)
}

func CreateEnrollment(ctx *gin.Context) {

}

func GetEnrollment(ctx *gin.Context) {
}
