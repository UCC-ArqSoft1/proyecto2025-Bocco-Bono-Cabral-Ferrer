package controllers

import (
	"net/http"
	"proyecto2025/backend/services"

	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	activityService services.ActivityService
}

// crea una nueva instancia del controlador con el servicio inyectado.
func NewActivityController(s services.ActivityService) *ActivityController {
	return &ActivityController{activityService: s}
}

// maneja el endpoint GET /activities y devuelve todas las actividades.
func (ac *ActivityController) GetAll(c *gin.Context) {
	activities, err := ac.activityService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener actividades"})
		return
	}
	c.JSON(http.StatusOK, activities)
}
