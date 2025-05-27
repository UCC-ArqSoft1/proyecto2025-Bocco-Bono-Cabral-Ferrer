package controllers

import (
	"gym-api/backend/domain"
	services "gym-api/backend/services/userServices"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserServices
}
type ControllerMethods interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

func (uc UserController) Login(ctx *gin.Context) {
	var request domain.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, token, err := uc.UserService.Login(request.Email, request.Password)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"token":   token,
	})

}

func (uc UserController) Register(ctx *gin.Context) {
	var request domain.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := uc.UserService.Register(request.FirstName, request.LastName, request.Email, request.Password, request.Birth_date, request.Sex)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"message": "User registered successfully",
	})
}
