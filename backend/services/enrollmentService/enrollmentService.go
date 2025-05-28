package services

import clients "gym-api/backend/clients/enrollmentClient"

type EnrollmentServiceImpl struct {
	Repo clients.EnrollmentRepositoryInterface
}
