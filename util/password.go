package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword return a bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hashed password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword check if the provided password is correct
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
