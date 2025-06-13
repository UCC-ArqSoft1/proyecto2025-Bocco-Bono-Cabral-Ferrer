package services

import (
	"errors"
	"fmt"
	clients "gym-api/backend/clients/userClients"
	"gym-api/backend/utils"
)

type UserService struct {
	Repo clients.UserRepositoryInterface
}

type UserServiceInterface interface {
	Login(email string, password string) (int, string, int, error)
	Register(name string, lastName string, email string, password string, birth_date string, sex string) (int, int, error)
}

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidPassword = errors.New("invalid password")

func (services UserService) Login(email string, password string) (int, string, int, error) {
	daoUser, err := services.Repo.GetUserByEmail(email)
	if err != nil {
		return 0, "", 0, ErrUserNotFound
	}
	if utils.HashPassword(password) != daoUser.PasswordHash {
		return 0, "", 0, ErrInvalidPassword
	}
	token, err := utils.GenerateJWT(daoUser.Id, daoUser.UserTypeId)
	if err != nil {
		return 0, "", 0, fmt.Errorf("error generating token: %w", err)
	}
	return daoUser.Id, token, daoUser.UserTypeId, nil
}

var ErrEmailAlreadyExists = errors.New("email already in use")

func (services UserService) Register(name string, lastName string, email string, password string, birth_date string, sex string) (int, int, error) {
	err := services.Repo.EmailAlreadyExists(email)
	if err != nil {
		return 0, 0, ErrEmailAlreadyExists
	}
	hashedPassword := utils.HashPassword(password)
	user, err := services.Repo.InsertUser(name, lastName, email, hashedPassword, birth_date, sex)
	return user.Id, user.UserTypeId, err
}
