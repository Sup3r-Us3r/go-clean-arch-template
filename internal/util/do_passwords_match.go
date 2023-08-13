package util

import "encoding/hex"

// DoPasswordsMatch check if two passwords match
func DoPasswordsMatch(hashedPassword string, password string) bool {
	// Decode the hashed password from hex to bytes
	hashedPasswordBytes, err := hex.DecodeString(hashedPassword)
	if err != nil {
		return false
	}

	// Extract the salt from the hashed password
	salt := hashedPasswordBytes[:SALT_SIZE]

	// Hash the input password using the extracted salt
	currentPasswordHash := HashPassword(password, salt)

	return currentPasswordHash == hashedPassword
}
