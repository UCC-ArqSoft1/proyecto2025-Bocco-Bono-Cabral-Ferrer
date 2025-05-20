package services

import (
	"fmt"

	"gym-api/backend/dao"
	"gym-api/backend/db"
	"gym-api/backend/utils"
)

func Login(email string, password string) (int, string, error) {
	userDAO, err := dao.GetUserByEmail(db.DB, email)
	if err != nil {
		return 0, "", fmt.Errorf("error getting user: %w", err)
	}
	if utils.HashPassword(password) != userDAO.PasswordHash {
		return 0, "", fmt.Errorf("invalid password")
	}
	token, err := utils.GenerateJWT(userDAO.Id)
	if err != nil {
		return 0, "", fmt.Errorf("error generating token: %w", err)
	}
	return userDAO.Id, token, nil
}
func Register(name string, lastName string, email string, password string, birth_date string, sex string) (int, error) {
	hashedPassword := utils.HashPassword(password)
	dao.InsertUser(db.DB, name, lastName, email, hashedPassword, birth_date, sex)
	userDAO, err := dao.GetUserByEmail(db.DB, email)
	return userDAO.Id, err
}
