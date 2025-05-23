package services

import (
	"gym-api/backend/dao"
	"gym-api/backend/db"
	"gym-api/backend/domain"
)

func GetActivities() ([]domain.Activity, error) {
	return dao.GetActivities(db.DB)
}
func GetactivityByID(id int) (domain.Activity, error) {
	return dao.GetActivityByID(db.DB, id)
}
