package util_test

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"testing"

	"github.com/Sup3r-Us3r/barber-server/internal/util"
)

type mockRand struct{}

func (mr *mockRand) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated error while reading salt")
}

func TestUtil_GenerateRandomSalt(t *testing.T) {
	t.Run("should be able to generate random salt", func(t *testing.T) {
		generatedSalt := util.GenerateRandomSalt(util.SALT_SIZE)
		saltHex := hex.EncodeToString(generatedSalt)

		if len(saltHex) <= 0 {
			t.Errorf("Expected size greater than 0, got %v", len(saltHex))
		}
	})

	t.Run("should not be able to generate salt when panic is happens", func(t *testing.T) {
		rand.Reader = &mockRand{}

		func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic with an error")
				}
			}()

			util.GenerateRandomSalt(util.SALT_SIZE)
		}()
	})
}
