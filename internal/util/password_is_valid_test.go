package util_test

import (
	"testing"

	"github.com/Sup3r-Us3r/barber-server/internal/util"
)

func TestUtil_PasswordIsValid(t *testing.T) {
	type testCase struct {
		test           string
		password       string
		expectedResult bool
	}

	testCases := []testCase{
		{
			test:           "should be able returns true when password is valid",
			password:       "!Aa12345678",
			expectedResult: true,
		},
		{
			test:           "should be able returns false when password is not provided",
			expectedResult: false,
		},
		{
			test:           "should be able returns false when password is invalid for having less than 8 characters",
			password:       "!Aa123",
			expectedResult: false,
		},
		{
			test:           "should be able returns false when password is invalid for not contain lower case",
			password:       "!A12345678",
			expectedResult: false,
		},
		{
			test:           "should be able returns false when password is invalid for not contain upper case",
			password:       "!a12345678",
			expectedResult: false,
		},
		{
			test:           "should be able returns false when password is invalid for not contain special character",
			password:       "Aa12345678",
			expectedResult: false,
		},
		{
			test:           "should be able returns false when password is invalid for not contain number",
			password:       "!aABCDEFGH",
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			result := util.PasswordIsValid(testCase.password)

			if result != testCase.expectedResult {
				t.Errorf("Expected result %v, got %v", testCase.expectedResult, result)
			}
		})
	}
}
