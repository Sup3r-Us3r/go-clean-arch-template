package util

import (
	"crypto/rand"

	"github.com/Sup3r-Us3r/barber-server/log"
)

const (
	SALT_SIZE int = 16
)

// GenerateRandomSalt generate 16 bytes randomly and securely using the
// Cryptographically secure pseudorandom number generator (CSPRNG)
// in the crypto.rand package
func GenerateRandomSalt(saltSize int) []byte {
	salt := make([]byte, saltSize)

	_, err := rand.Read(salt[:])
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	return salt
}
