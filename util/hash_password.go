package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	cost := bcrypt.DefaultCost

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
