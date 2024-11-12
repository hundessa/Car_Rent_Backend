package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HandleError(err error) {
    if err != nil {
        log.Printf("Error: %v\n", err) // Log the error instead of panicking
        // Optionally, you could return the error or handle it in a way that fits your application's needs
    }
}

// HashAndSalt hashes and salts the password
func HashAndSalt(pass []byte) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
    if err != nil {
        return "", err // Return the error to be handled by the caller
    }
    return string(hashed), nil
}

// CheckPasswordHash checks if the provided password matches the hashed password
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
