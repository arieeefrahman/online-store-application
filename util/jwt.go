package util

import (
	"errors"
	"fmt"
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

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	tkn, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return tkn, nil
}

func DecodeToken(tokenStr string) (jwt.MapClaims, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.ErrInvalidKeyType
	}
}
