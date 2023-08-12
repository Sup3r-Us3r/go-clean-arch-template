package barber_test

import (
	"context"
	"testing"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository/memory"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/barber"
	"github.com/Sup3r-Us3r/barber-server/test/factory"
)

func Test_Should_Be_Able_To_Create_A_New_Barber(t *testing.T) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewCreateBarberUseCase(repositoryContainer)

	barberData := factory.MakeBarber(entity.Barber{})
	inputDTO := barber.CreateBarberUseCaseInputDTO{
		Name:     barberData.Name,
		Email:    barberData.Email,
		Password: barberData.Password,
		Phone:    barberData.Phone,
	}

	ctx := context.Background()
	err := sut.Execute(ctx, inputDTO)
	if err != nil {
		t.Errorf("Expected error %v, got %v", nil, err)
	}

	if len(barberRepositoryMemory.Barbers) == 0 {
		t.Errorf("Expected result %v, got %v", 1, len(barberRepositoryMemory.Barbers))
	}
}

func Test_Should_Not_Be_Able_To_Create_A_New_Barber_When_Input_Is_Invalid(t *testing.T) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewCreateBarberUseCase(repositoryContainer)

	barberData := factory.MakeBarber(entity.Barber{})
	inputDTO := barber.CreateBarberUseCaseInputDTO{
		Name:  barberData.Name,
		Phone: barberData.Phone,
	}

	ctx := context.Background()
	err := sut.Execute(ctx, inputDTO)
	if err != apperr.ErrBarberFieldEmailIsRequired {
		t.Errorf("Expected error %v, got %v", apperr.ErrBarberFieldEmailIsRequired, err)
	}

	if len(barberRepositoryMemory.Barbers) != 0 {
		t.Errorf("Expected result %v, got %v", 0, len(barberRepositoryMemory.Barbers))
	}
}

func Test_Should_Not_Be_Able_To_Create_A_New_Barber_When_Barber_Already_Exists(t *testing.T) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewCreateBarberUseCase(repositoryContainer)

	barberData := factory.MakeBarber(entity.Barber{})
	barberRepositoryMemory.Barbers = []entity.Barber{
		barberData,
	}
	inputDTO := barber.CreateBarberUseCaseInputDTO{
		Name:     barberData.Name,
		Email:    barberData.Email,
		Password: barberData.Password,
		Phone:    barberData.Phone,
	}

	ctx := context.Background()
	err := sut.Execute(ctx, inputDTO)
	if err != apperr.ErrBarberAlreadyExists {
		t.Errorf("Expected error %v, got %v", apperr.ErrBarberAlreadyExists, err)
	}
}
