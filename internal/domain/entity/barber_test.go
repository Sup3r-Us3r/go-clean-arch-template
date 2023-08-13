package entity_test

import (
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
