package util

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	// cost for
	cost = 10
)

// HashPassword will generate hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// CheckPasswordHash will
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

// HashOTP will generate hashed password
func HashOTP(otp string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(otp), cost)
	return string(bytes), err
}

// CheckOTPHash will
func CheckOTPHash(otp, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(otp))

	return err
}

// HashPIN will generate hashed password
func HashPIN(pin string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pin), cost)
	return string(bytes), err
}

// CheckPINHash will
func CheckPINHash(pin, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pin))

	return err
}
