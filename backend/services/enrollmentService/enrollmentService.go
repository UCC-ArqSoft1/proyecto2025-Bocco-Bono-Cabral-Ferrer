package services

import clients "gym-api/backend/clients/enrollmentClient"

type EnrollmentServiceImpl struct {
	Repo clients.EnrollmentRepositoryInterface
}

type EnrollmentServiceInterface interface {
	CreateEnrollment(userId, activityId int) error
}

//func (es EnrollmentServiceImpl) CreateEnrollment(userId, activityId int) error {

//}
