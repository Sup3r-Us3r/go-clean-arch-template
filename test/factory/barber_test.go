package factory_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/test/factory"
)

func TestFactory_MakeBarber(t *testing.T) {
	type testCase struct {
		test           string
		barber         entity.Barber
		expectedResult entity.Barber
	}

	barberData := entity.Barber{}
	modifiedBarberData := entity.Barber{
		ID:           "barber-id",
		Name:         "Barber 1",
		Email:        "barber1@mail.com",
		Password:     "!Aa12345678",
		PasswordHash: "a45ed3ef2af41f0f091148764e3b1876f34b78334538ddf484b63e41c380823e794e55714af9a3befcf6be34ad1b5a2714e74409d569489e63f74e82f58efd74c594a6d649d43e3a70a66824bbda4cb9",
		Phone:        "12934567890",
		CreatedAt:    time.Now(),
	}

	testCases := []testCase{
		{
			test:   "should be able to create an barber using factory",
			barber: barberData,
		},
		{
			test:           "should be able to create an barber using factory with override data",
			barber:         modifiedBarberData,
			expectedResult: modifiedBarberData,
		},
	}

	for _, testCase := range testCases {
		if testCase.test == "should be able to create an barber using factory" {
			t.Run(testCase.test, func(t *testing.T) {
				currentBarberData := factory.MakeBarber(testCase.barber)

				if currentBarberData.ID == "" {
					t.Errorf("Expected result %v, got %v", "non-empty ID", "''")
				}

				if currentBarberData.Name == "" {
					t.Errorf("Expected result %v, got %v", "non-empty Name", "''")
				}

				if currentBarberData.Email == "" {
					t.Errorf("Expected result %v, got %v", "non-empty Email", "''")
				}

				if currentBarberData.Password == "" {
					t.Errorf("Expected result %v, got %v", "non-empty Password", "''")
				}

				if currentBarberData.PasswordHash == "" {
					t.Errorf("Expected result %v, got %v", "non-empty Password Hash", "''")
				}

				if currentBarberData.Phone == "" {
					t.Errorf("Expected result %v, got %v", "non-empty Phone", "''")
				}

				if currentBarberData.CreatedAt.IsZero() {
					t.Errorf("Expected result %v, got %v", "non-empty CreatedAt", "''")
				}
			})
		}

		if testCase.test == "should be able to create an barber using factory with override data" {
			t.Run(testCase.test, func(t *testing.T) {
				currentBarberData := factory.MakeBarber(testCase.barber)

				if currentBarberData.ID != testCase.expectedResult.ID {
					t.Errorf("Expected result %v, got %v", testCase.expectedResult, currentBarberData)
				}

				if currentBarberData.Name != testCase.expectedResult.Name {
					t.Errorf("Expected result %v, got %v", testCase.expectedResult, currentBarberData)
				}

				if currentBarberData.Email != testCase.expectedResult.Email {
					t.Errorf("Expected result %v, got %v", testCase.expectedResult, currentBarberData)
				}

				if currentBarberData.Password != testCase.expectedResult.Password {
					t.Errorf("Expected result %v, got %v", testCase.expectedResult, currentBarberData)
				}

				if currentBarberData.PasswordHash != testCase.expectedResult.PasswordHash {
					t.Errorf("Expected result %v, got %v", testCase.expectedResult, currentBarberData)
				}

				if currentBarberData.Phone != testCase.expectedResult.Phone {
					t.Errorf("Expected result %v, got %v", testCase.expectedResult, currentBarberData)
				}

				fmt.Println(currentBarberData.CreatedAt, testCase.barber.CreatedAt, !currentBarberData.CreatedAt.Equal(testCase.barber.CreatedAt))

				if !currentBarberData.CreatedAt.Equal(testCase.barber.CreatedAt) {
					t.Errorf("Expected result %v, got %v", testCase.expectedResult, testCase.barber)
				}
			})
		}
	}
}
