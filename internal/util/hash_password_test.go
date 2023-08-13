package util_test

import (
	"testing"

	"github.com/Sup3r-Us3r/barber-server/internal/util"
)

func TestUtil_HashPassword(t *testing.T) {
	t.Run("should be able generate password hash", func(t *testing.T) {
		result := util.HashPassword("!Aa12345678", []byte{})

		if result == "" {
			t.Errorf("Expected hash password, got ''")
		}
	})
}
