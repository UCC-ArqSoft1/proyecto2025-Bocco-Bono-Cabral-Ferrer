package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtDuration = time.Minute * 15
)

// CustomClaims extends the standard JWT claims with our custom fields
type CustomClaims struct {
	UserID     int `json:"user_id"`
	UserTypeID int `json:"user_type_id"`
	jwt.RegisteredClaims
}

// GetJWTSecret returns the JWT secret key
func GetJWTSecret() string {
	return os.Getenv("JWT_PASSWORD")
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
	fmt.Println(os.Getenv("JWT_PASSWORD"))
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_PASSWORD")))
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenString, nil
}
