package mocks

import (
	"gym-api/backend/domain"
)

type MockActivityService struct{}

func (m MockActivityService) GetActivities() ([]domain.Activity, error) {
	return []domain.Activity{
		{
			Id:          1,
			Name:        "Natación",
			Description: "Clases para adultos",
			Category:    "Acuático",
			Profesor:    "Juan",
			Capacity:    20,
		},
		{
			Id:          2,
			Name:        "Boxeo",
			Description: "Clases para niños principiantes",
			Category:    "Deporte de Contacto",
			Profesor:    "Myke Tyson",
			Capacity:    15,
		},
	}, nil
}

func (m MockActivityService) GetactivityByID(id int) (domain.Activity, error) {
	return domain.Activity{Id: id, Name: "Test"}, nil
}

func (m MockActivityService) GetActivitiesByFilters(keyword string) ([]domain.Activity, error) {
	return []domain.Activity{
		{Name: "Filtrado", Description: "Con keyword"},
	}, nil
}
