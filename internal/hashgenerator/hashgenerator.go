package hashGenerator

import (
	"smart-recommendation/internal/errorhandler"

	"golang.org/x/crypto/bcrypt"
)

// Generate generates hash of data
func Generate(data string) (string, errorhandler.ErrorData) {
	pwd := []byte(data)
	pwd, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", errorhandler.ErrorHash(err)
	}
	response := string(pwd[:])
	return response, errorhandler.ErrorData{}
}

// Verify verifies password with hashedPassword
func Verify(hashedPassword string, password string) (bool, errorhandler.ErrorData) {
	hshPwd := []byte(hashedPassword)
	pwd := []byte(password)
	err := bcrypt.CompareHashAndPassword(hshPwd, pwd)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, errorhandler.ErrorData{}
		}
		return false, errorhandler.ErrorHash(err)
	}
	return true, errorhandler.ErrorData{}
}
