package controllers_test

import (
	"fmt"
	controllers "gym-api/backend/controllers/activity"
	"gym-api/backend/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetActivities(t *testing.T) {
	// Configurar el contexto de Gin en modo test
	gin.SetMode(gin.TestMode)

	// Crear el mock y el controller
	mockService := mocks.MockActivityService{}
	controller := controllers.ActivityController{ActivityService: mockService}

	// Simular request
	req, _ := http.NewRequest(http.MethodGet, "/activities", nil)
	resp := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(resp)
	ctx.Request = req

	// Ejecutar el handler
	controller.GetActivities(ctx)

	// Verificaciones
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Natación")
	fmt.Println("Response Body:", resp.Body.String())
}

func TestGetActivititesByFilters(t *testing.T) {
	// Configurar el contexto de Gin en modo test
	gin.SetMode(gin.TestMode)

	// Crear el mock y el controller
	mockService := mocks.MockActivityService{}
	controller := controllers.ActivityController{ActivityService: mockService}

	// Simular request con keyword
	req, _ := http.NewRequest(http.MethodGet, "/activities?keyword=Filtrado", nil)
	resp := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(resp)
	ctx.Request = req

	// Ejecutar el handler
	controller.GetActivities(ctx)

	// Verificaciones
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Filtrado")
	fmt.Println("Response Body:", resp.Body.String())
}
