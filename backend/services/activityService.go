package services

import (
	"gym-api/backend/dao"
	"gym-api/backend/domain"
)

type ActivityService struct {
	activityDAO *dao.ActivityDAO
}

// Constructor para inyectar el DAO
func NewActivityService(d *dao.ActivityDAO) *ActivityService {
	return &ActivityService{activityDAO: d}
}

// Obtener todas las actividades
func (s *ActivityService) GetActivity() ([]domain.Activity, error) {
	return s.activityDAO.GetAll()
}

// Obtener actividad por ID
func (s *ActivityService) GetActivityByID(id int) (domain.Activity, error) {
	return s.activityDAO.GetByID(id)
}

// Crear actividad
func (s *ActivityService) Create(activity domain.Activity) error {
	return s.activityDAO.Create(activity)
}

// Actualizar actividad
func (s *ActivityService) Update(id int, activity domain.Activity) error {
	return s.activityDAO.Update(id, activity)
}

// Eliminar actividad
func (s *ActivityService) Delete(id int) error {
	return s.activityDAO.Delete(id)
}
