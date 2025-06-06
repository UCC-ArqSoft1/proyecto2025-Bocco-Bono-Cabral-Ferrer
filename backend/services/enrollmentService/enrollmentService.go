package services

import (
	"errors"
	clients "gym-api/backend/clients/enrollmentClient"
	"time"
	"gym-api/backend/domain"
	"gym-api/backend/dao"
)

type EnrollmentService struct {
	Repo clients.EnrollmentRepositoryInterface
}

type EnrollmentServiceInterface interface {
	CreateEnrollment(userId, activityId int) error
	GetUserEnrollments(userId int) ([]domain.Activity, error)
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

func (es EnrollmentService) GetUserEnrollments(userId int) ([]domain.Activity, error) {
	enrollments, err := es.Repo.GetUserEnrollments(userId)
	if err != nil {
		return nil, err
	}

	var activities []domain.Activity
	for _, enrollment := range enrollments {
		activities = append(activities, dao.DaoToDto(enrollment.Activity))
	}

	return activities, nil
}
