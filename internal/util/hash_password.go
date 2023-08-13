package util

import (
	"crypto/sha512"
	"encoding/hex"
)

// HashPassword combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password as a hex string
func HashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	passwordBytes := []byte(password)

	// Create SHA-512 hasher
	sha512Hasher := sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	hashedPasswordBytes := sha512Hasher.Sum(nil)

	// Combine salt and hashed password
	combined := append(salt, hashedPasswordBytes...)

	// Convert the combined value to a hex string
	hashedPasswordHex := hex.EncodeToString(combined)

	return hashedPasswordHex
}
