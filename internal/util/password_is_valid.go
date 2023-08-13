package util

import "regexp"

// PasswordIsValid check if the password has at least 8 digits, 1 lowercase letter,
// 1 uppercase letter, 1 number and 1 special character
func PasswordIsValid(password string) bool {
	if len(password) < 8 {
		return false
	}

	if !regexp.MustCompile("[a-z]").MatchString(password) {
		return false
	}

	if !regexp.MustCompile("[A-Z]").MatchString(password) {
		return false
	}

	if !regexp.MustCompile(`[\W_]`).MatchString(password) {
		return false
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		return false
	}

	return true
}
