package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/isd-sgcu/oph66-backend/apperror"
	"github.com/isd-sgcu/oph66-backend/cfgldr"
)

func GenTokenStaff(department string, cfg *cfgldr.Config) (string, *apperror.AppError) {
	secretKey := cfg.JWTConfig.SecretKey

	claims := jwt.MapClaims{
		"role":       "staff",
		"department": department,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", apperror.InternalError
	}

	return tokenString, nil
}

func main() {
	config, err := cfgldr.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load configuration:", err)
		return
	}

	department := "IT"
	token, appErr := GenTokenStaff(department, config)
	if appErr != nil {
		fmt.Println("Failed to generate JWT token:", appErr)
		return
	}

	fmt.Println("Generated JWT token for staff in department", department, ":", token)
}
