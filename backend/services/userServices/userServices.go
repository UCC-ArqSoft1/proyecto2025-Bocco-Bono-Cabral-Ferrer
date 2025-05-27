package services

import (
	"fmt"
	clients "gym-api/backend/clients/userClients"
	"gym-api/backend/utils"
)

type UserServices struct {
	UserClient clients.MySQLUserRepository
}

type UserService interface {
	Login(email string, password string) (int, string, error)
	Register(name string, lastName string, email string, password string, birth_date string, sex string) (int, error)
}

func (services UserServices) Login(email string, password string) (int, string, error) {
	daoUser, err := services.UserClient.GetUserByEmail(email)
	if err != nil {
		return 0, "", fmt.Errorf("error getting user: %w", err)
	}
	if utils.HashPassword(password) != daoUser.PasswordHash {
		return 0, "", fmt.Errorf("invalid password")
	}
	token, err := utils.GenerateJWT(daoUser.Id)
	if err != nil {
		return 0, "", fmt.Errorf("error generating token: %w", err)
	}
	return daoUser.Id, token, nil
}
func (services UserServices) Register(name string, lastName string, email string, password string, birth_date string, sex string) (int, error) {
	hashedPassword := utils.HashPassword(password)
	services.UserClient.InsertUser(name, lastName, email, hashedPassword, birth_date, sex)
	userDAO, err := services.UserClient.GetUserByEmail(email)
	return userDAO.Id, err
}
