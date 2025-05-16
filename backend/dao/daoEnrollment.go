package dao

import "time"

type Enrollment struct {
	Id             int       `json:"id"`
	UserId         int       `json:"user_id"`
	User           User      `json:"user"`
	ActivityId     int       `json:"activity_id"`
	Activity       Activity  `json:"activity"`
	EnrollmentDate time.Time `json:"enrollment_date"`
}

type Enrollments []Enrollment
