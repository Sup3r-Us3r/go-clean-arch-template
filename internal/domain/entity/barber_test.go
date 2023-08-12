package entity_test

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"testing"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
)

func TestEntity_NewBarber(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		email       string
		password    string
		phone       string
		expectedErr *apperr.AppErr
	}

	testCases := []testCase{
		{
			test:        "should be able to create a new instance of barber",
			name:        "Barber 1",
			email:       "barber1@mail.com",
			password:    "!Aa12345678",
			phone:       "12934567890",
			expectedErr: nil,
		},
		{
			test:        "should not be able create a new instance of barber when field name is not provided",
			email:       "barber1@mail.com",
			password:    "!Aa12345678",
			phone:       "12934567890",
			expectedErr: apperr.ErrBarberFieldNameIsRequired,
		},
		{
			test:        "should not be able create a new instance of barber when field email is not provided",
			name:        "Barber 1",
			phone:       "12934567890",
			expectedErr: apperr.ErrBarberFieldEmailIsRequired,
		},
		{
			test:        "should not be able create a new instance of barber when email is invalid",
			name:        "Barber 1",
			email:       "barber1.mail.com",
			password:    "!Aa12345678",
			phone:       "12934567890",
			expectedErr: apperr.ErrBarberEmailIsInvalid,
		},
		{
			test:        "should not be able create a new instance of barber when field phone is not provided",
			name:        "Barber 1",
			email:       "barber1@mail.com",
			password:    "!Aa12345678",
			expectedErr: apperr.ErrBarberFieldPhoneIsRequired,
		},
		{
			test:        "should not be able create a new instance of barber when field password is not provided",
			name:        "Barber 1",
			email:       "barber1@mail.com",
			phone:       "12934567890",
			expectedErr: apperr.ErrBarberFieldPasswordIsRequired,
		},
		{
			test:        "should not be able create a new instance of barber when password is invalid for having less than 8 characters",
			name:        "Barber 1",
			email:       "barber1@mail.com",
			password:    "!Aa123",
			phone:       "12934567890",
			expectedErr: apperr.ErrBarberPasswordIsInvalid,
		},
		{
			test:        "should not be able create a new instance of barber when password is invalid for not contain lower case",
			name:        "Barber 1",
			email:       "barber1@mail.com",
			password:    "!A12345678",
			phone:       "12934567890",
			expectedErr: apperr.ErrBarberPasswordIsInvalid,
		},
		{
			test:        "should not be able create a new instance of barber when password is invalid for not contain upper case",
			name:        "Barber 1",
			email:       "barber1@mail.com",
			password:    "!a12345678",
			phone:       "12934567890",
			expectedErr: apperr.ErrBarberPasswordIsInvalid,
		},
		{
			test:        "should not be able create a new instance of barber when password is invalid for not contain special character",
			name:        "Barber 1",
			email:       "barber1@mail.com",
			password:    "Aa12345678",
			phone:       "12934567890",
			expectedErr: apperr.ErrBarberPasswordIsInvalid,
		},
		{
			test:        "should not be able create a new instance of barber when password is invalid for not contain number",
			name:        "Barber 1",
			email:       "barber1@mail.com",
			password:    "!aABCDEFGH",
			phone:       "12934567890",
			expectedErr: apperr.ErrBarberPasswordIsInvalid,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			_, err := entity.NewBarber(testCase.name, testCase.email, testCase.phone, testCase.password)

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}
}

type mockRand struct{}

func (mr *mockRand) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated error while reading salt")
}

func TestEntity_GenerateRandomSalt(t *testing.T) {
	t.Run("should be able to generate random salt", func(t *testing.T) {
		generatedSalt := entity.GenerateRandomSalt(entity.SALT_SIZE)
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

			entity.GenerateRandomSalt(entity.SALT_SIZE)
		}()
	})
}

func TestEntity_DoPasswordMatch(t *testing.T) {
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
			passwordHasMatch := entity.DoPasswordsMatch(testCase.hashedPassword, testCase.password)

			if testCase.expectedResult != passwordHasMatch {
				t.Errorf("Expected result %v, got %v", testCase.expectedResult, passwordHasMatch)
			}
		})
	}
}
