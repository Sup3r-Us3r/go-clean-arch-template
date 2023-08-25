package memory_test

import (
	"context"
	"testing"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository/memory"
	"github.com/Sup3r-Us3r/barber-server/test/factory"
)

func TestBarberRepositoryMemory_GetBarberById(t *testing.T) {
	type testCase struct {
		test           string
		barberId       string
		expectedResult entity.Barber
		expectedErr    *apperr.AppErr
	}

	barberData := factory.MakeBarber(entity.Barber{})

	testCases := []testCase{
		{
			test:           "should be able get barber by id",
			barberId:       barberData.ID,
			expectedResult: barberData,
			expectedErr:    nil,
		},
		{
			test:           "should not be able get barber by id when barber does not exists",
			barberId:       "",
			expectedResult: entity.Barber{},
			expectedErr:    apperr.ErrBarberNotFound,
		},
	}

	ctx := context.Background()
	barberRepository := memory.NewBarberRepositoryMemory()

	barberRepository.CreateBarber(ctx, &barberData)

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			result, err := barberRepository.GetBarberById(ctx, testCase.barberId)

			if result.ID != testCase.expectedResult.ID {
				t.Errorf("Expected result %v, got %v", testCase.expectedResult, result)
			}

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}
}

func TestBarberRepositoryMemory_GetBarberByEmail(t *testing.T) {
	type testCase struct {
		test           string
		barberEmail    string
		expectedResult entity.Barber
		expectedErr    *apperr.AppErr
	}

	barberData := factory.MakeBarber(entity.Barber{})

	testCases := []testCase{
		{
			test:           "should be able get barber by email",
			barberEmail:    barberData.Email,
			expectedResult: barberData,
			expectedErr:    nil,
		},
		{
			test:           "should not be able get barber by email when barber does not exists",
			barberEmail:    "",
			expectedResult: entity.Barber{},
			expectedErr:    apperr.ErrBarberNotFound,
		},
	}

	ctx := context.Background()
	barberRepository := memory.NewBarberRepositoryMemory()

	barberRepository.CreateBarber(ctx, &barberData)

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			result, err := barberRepository.GetBarberByEmail(ctx, testCase.barberEmail)

			if result.ID != testCase.expectedResult.ID {
				t.Errorf("Expected result %v, got %v", testCase.expectedResult, result)
			}

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}
}

func TestBarberRepositoryMemory_FetchBarbers(t *testing.T) {
	type testCase struct {
		test           string
		expectedResult []entity.Barber
	}

	barberData := factory.MakeBarber(entity.Barber{})

	testCases := []testCase{
		{
			test: "should be able fetch barber list",
			expectedResult: []entity.Barber{
				barberData,
			},
		},
	}

	barberRepository := memory.NewBarberRepositoryMemory()

	ctx := context.Background()
	barberRepository.CreateBarber(ctx, &barberData)

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			result := barberRepository.FetchBarbers(ctx)
			resultsMatch := []entity.Barber{}

			for _, expectedResult := range testCase.expectedResult {
				for _, currentResult := range result {
					if expectedResult.ID == currentResult.ID {
						resultsMatch = append(resultsMatch, currentResult)
					}
				}
			}

			if len(resultsMatch) == 0 {
				t.Errorf("Expected result %v, got %v", testCase.expectedResult, resultsMatch)
			}
		})
	}
}

func TestBarberRepositoryMemory_CreateBarber(t *testing.T) {
	type testCase struct {
		test        string
		barber      entity.Barber
		expectedErr *apperr.AppErr
	}

	barberData := factory.MakeBarber(entity.Barber{})

	testCases := []testCase{
		{
			test:        "should be able to create a new barber",
			barber:      barberData,
			expectedErr: nil,
		},
		{
			test:        "should not be able create a new barber when barber already exists",
			barber:      barberData,
			expectedErr: apperr.ErrBarberAlreadyExists,
		},
	}

	ctx := context.Background()
	barberRepository := memory.NewBarberRepositoryMemory()

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			err := barberRepository.CreateBarber(ctx, &testCase.barber)

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}
}

func TestBarberRepositoryMemory_UpdateBarber(t *testing.T) {
	type testCase struct {
		test        string
		barberId    string
		updateData  entity.Barber
		expectedErr *apperr.AppErr
	}

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"
	barberData := factory.MakeBarber(entity.Barber{ID: barberId})
	updateData := barberData
	updateData.Name = "Barber1"
	updateData.Email = "barber1@mail.com"

	testCases := []testCase{
		{
			test:        "should be able update barber",
			barberId:    barberId,
			updateData:  updateData,
			expectedErr: nil,
		},
		{
			test:        "should not be able update barber when barber not exists",
			barberId:    "non-existent-id",
			updateData:  updateData,
			expectedErr: apperr.ErrBarberNotFound,
		},
	}

	ctx := context.Background()
	barberRepository := memory.NewBarberRepositoryMemory()

	barberRepository.Barbers = []entity.Barber{
		barberData,
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			err := barberRepository.UpdateBarber(ctx, testCase.barberId, &updateData)

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}
}

func TestBarberRepositoryMemory_DeleteBarber(t *testing.T) {
	type testCase struct {
		test        string
		barberId    string
		expectedErr *apperr.AppErr
	}

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"
	barberData := factory.MakeBarber(entity.Barber{ID: barberId})

	testCases := []testCase{
		{
			test:        "should be able delete barber",
			barberId:    barberId,
			expectedErr: nil,
		},
		{
			test:        "should not be able delete barber when barber not exists",
			barberId:    "non-existent-id",
			expectedErr: apperr.ErrBarberNotFound,
		},
	}

	ctx := context.Background()
	barberRepository := memory.NewBarberRepositoryMemory()

	barberRepository.Barbers = []entity.Barber{
		barberData,
	}

	for _, testCase := range testCases {
		t.Run(testCase.test, func(t *testing.T) {
			err := barberRepository.DeleteBarber(ctx, testCase.barberId)

			if err != testCase.expectedErr {
				t.Errorf("Expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}
}
