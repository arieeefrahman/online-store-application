package util

import (
	"errors"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	
	if secretKey == "" {
		return "", errors.New("JWT_SECRET_KEY is empty in .env file")
	  }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return webtoken, nil
}
