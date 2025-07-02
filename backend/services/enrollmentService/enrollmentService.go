package services

import (
	"errors"
	activityClients "gym-api/clients/activityclient"
	clients "gym-api/clients/enrollmentClient"
	"gym-api/dao"
	"gym-api/domain"
	"time"
)

type EnrollmentService struct {
	Repo         clients.EnrollmentRepositoryInterface
	ActivityRepo activityClients.ActivityRepositoryInterface
}

type EnrollmentServiceInterface interface {
	CreateEnrollment(userId, activityId int) error
	GetUserEnrollments(userId int) ([]domain.Activity, error)
	CancelEnrollment(userId, activityId int) error
	IsEnrolled(userId, activityId int) (bool, error)
	GetAvailableSpots(activityId int) (int, error)
}

func (es EnrollmentService) CreateEnrollment(userId, activityId int) error {
	// Check if user is already enrolled
	isEnrolled, err := es.Repo.IsEnrolled(userId, activityId)
	if err != nil {
		return err
	}
	if isEnrolled {
		return errors.New("user is already enrolled in this activity")
	}

	// Check if there's capacity available
	currentEnrollments, capacity, err := es.Repo.CountEnrollmentsAndCapacity(activityId)
	if err != nil {
		return err
	}
	if currentEnrollments >= capacity {
		return errors.New("activity is at full capacity")
	}

	// Create the enrollment
	return es.Repo.CreateEnrollment(userId, activityId, time.Now())
}

func (es EnrollmentService) CancelEnrollment(userId, activityId int) error {
	// Check if user is enrolled
	isEnrolled, err := es.Repo.IsEnrolled(userId, activityId)
	if err != nil {
		return err
	}
	if !isEnrolled {
		return errors.New("user is not enrolled in this activity")
	}

	// Cancel the enrollment
	return es.Repo.CancelEnrollment(userId, activityId)
}

func (es EnrollmentService) IsEnrolled(userId, activityId int) (bool, error) {
	return es.Repo.IsEnrolled(userId, activityId)
}

func (es EnrollmentService) GetAvailableSpots(activityId int) (int, error) {
	currentEnrollments, totalCapacity, err := es.Repo.CountEnrollmentsAndCapacity(activityId)
	if err != nil {
		return 0, err
	}

	availableSpots := totalCapacity - currentEnrollments

	return availableSpots, nil
}

func (es EnrollmentService) GetUserEnrollments(userId int) ([]domain.Activity, error) {
	enrollments, err := es.Repo.GetUserEnrollments(userId)
	if err != nil {
		return nil, err
	}

	var activities []domain.Activity
	for _, enrollment := range enrollments {
		// Obtener la actividad completa usando el ActivityId
		activityDao, err := es.ActivityRepo.GetActivityByID(enrollment.ActivityId)
		if err != nil {
			continue // Skip this activity if there's an error
		}
		activity := dao.DaoToDto(activityDao)
		activities = append(activities, activity)
	}

	return activities, nil
}
