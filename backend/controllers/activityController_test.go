package controllers_test

import (
	"fmt"
	controllers "gym-api/controllers/activity"
	"gym-api/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
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
	req, _ := http.NewRequest(http.MethodGet, "/activities?keyword=Yoga", nil)
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
func TestCreateActivity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := mocks.MockActivityService{}
	controller := controllers.ActivityController{ActivityService: mockService}

	jsonData := `{
		"name": "Yoga Test",
		"description": "Clase de yoga de prueba",
		"capacity": 10,
		"category": "Bienestar",
		"profesor": "María",
		"schedules": [
			{"day": "Lunes",
			 "start_time": "08:00",
			  "end_time": "09:00"
			}
		]
	}`

	req, _ := http.NewRequest(http.MethodPost, "/activities", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(resp)
	ctx.Request = req

	controller.CreateActivity(ctx)

	assert.Equal(t, http.StatusCreated, resp.Code)
	assert.Contains(t, resp.Body.String(), "Activity created successfully")
	fmt.Println("Response Body:", resp.Body.String())
}

func TestUpdateActivity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := mocks.MockActivityService{}
	controller := controllers.ActivityController{ActivityService: mockService}

	jsonData := `{
		"name": "Yoga Editado",
		"description": "Clase editada",
		"capacity": 12,
		"category": "Bienestar",
		"profesor": "María",
		"schedules": [
			{"day": "Martes", "start_time": "09:00", "end_time": "10:00"}
		]
	}`

	req, _ := http.NewRequest(http.MethodPut, "/activities/1", strings.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(resp)
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	ctx.Request = req

	controller.UpdateActivity(ctx)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Activity updated successfully")
	fmt.Println("Response Body:", resp.Body.String())
}

func TestDeleteActivity(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := mocks.MockActivityService{}
	controller := controllers.ActivityController{ActivityService: mockService}

	req, _ := http.NewRequest(http.MethodDelete, "/activities/1", nil)
	resp := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(resp)
	ctx.Params = []gin.Param{{Key: "id", Value: "1"}}
	ctx.Request = req

	controller.DeleteActivity(ctx)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Activity deleted successfully")
	fmt.Println("Response Body:", resp.Body.String())
}
