package util_test

import (
	"testing"

	"github.com/Sup3r-Us3r/barber-server/internal/util"
)

func TestUtil_EmailIsValid(t *testing.T) {
	t.Run("should be able returns true when email is valid", func(t *testing.T) {
		result := util.EmailIsValid("barber1@mail.com")

		if !result {
			t.Errorf("Expected result %v, got %v", true, result)
		}
	})

	t.Run("should be able returns false when email is invalid", func(t *testing.T) {
		result := util.EmailIsValid("barber1@mailcom")

		if result {
			t.Errorf("Expected result %v, got %v", false, result)
		}
	})
}
