package controllers

import (
	"errors"
	"gym-api/domain"
	services "gym-api/services/userServices"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserServiceInterface
}
type UserControllerInterface interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

func (uc UserController) Login(ctx *gin.Context) {
	var request domain.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, token, userTypeId, err := uc.UserService.Login(request.Email, request.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidPassword) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña Invalida"})
			return
		}
		if errors.Is(err, services.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"user_id":      userID,
		"token":        token,
		"user_type_id": userTypeId,
	})

}

func (uc UserController) Register(ctx *gin.Context) {
	var request domain.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, userTypeId, err := uc.UserService.Register(request.FirstName, request.LastName, request.Email, request.Password, request.Birth_date, request.Sex)
	if err != nil {
		if errors.Is(err, services.ErrEmailAlreadyExists) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Email ya en uso"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"user_id":      userID,
		"user_type_id": userTypeId,
		"message":      "User registered successfully",
	})
}
