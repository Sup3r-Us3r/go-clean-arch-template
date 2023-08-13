package util_test

import (
	"testing"

	"github.com/Sup3r-Us3r/barber-server/internal/util"
)

func TestUtil_DoPasswordMatch(t *testing.T) {
	type testCase struct {
		test           string
		password       string
		hashedPassword string
		expectedResult bool
	}

	testCases := []testCase{
		{
			test:           "should be able check if two passwords match",
			password:       "!Aa12345678",
			hashedPassword: "a45ed3ef2af41f0f091148764e3b1876f34b78334538ddf484b63e41c380823e794e55714af9a3befcf6be34ad1b5a2714e74409d569489e63f74e82f58efd74c594a6d649d43e3a70a66824bbda4cb9",
			expectedResult: true,
		},
		{
			test:           "should not be able check if two passwords match when password is wrong",
			password:       "!Aa1234567890",
			hashedPassword: "a45ed3ef2af41f0f091148764e3b1876f34b78334538ddf484b63e41c380823e794e55714af9a3befcf6be34ad1b5a2714e74409d569489e63f74e82f58efd74c594a6d649d43e3a70a66824bbda4cb9",
			expectedResult: false,
		},
		{
			test:           "should not be able check if two passwords match when password hash is invalid",
			password:       "!Aa1234567890",
			hashedPassword: "invalid-password-hash",
			expectedResult: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			passwordHasMatch := util.DoPasswordsMatch(testCase.hashedPassword, testCase.password)

			if testCase.expectedResult != passwordHasMatch {
				t.Errorf("Expected result %v, got %v", testCase.expectedResult, passwordHasMatch)
			}
		})
	}
}
