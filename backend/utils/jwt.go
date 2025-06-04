package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtDuration = time.Hour * 1
	jwtSecret   = "jwtSecret"
)

// CustomClaims extends the standard JWT claims with our custom fields
type CustomClaims struct {
	UserID     int `json:"user_id"`
	UserTypeID int `json:"user_type_id"`
	jwt.RegisteredClaims
}

// GetJWTSecret returns the JWT secret key
func GetJWTSecret() string {
	return jwtSecret
}

func GenerateJWT(userID int, userTypeID int) (string, error) {
	// Setear expiracion
	expirationTime := time.Now().Add(jwtDuration)

	// Construir los claims
	claims := CustomClaims{
		UserID:     userID,
		UserTypeID: userTypeID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "backend",
			Subject:   "auth",
		},
	}

	// Crear el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token
	fmt.Println(jwtSecret)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenString, nil
}
