package auth_test

import (
	"context"
	"testing"

	"github.com/Sup3r-Us3r/barber-server/config"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository/memory"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/auth"
	"github.com/Sup3r-Us3r/barber-server/test/factory"
)

func Test_Should_Be_Able_To_Authenticate(t *testing.T) {
	config.LoadConfig("../../..")

	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := auth.NewSignInUseCase(repositoryContainer)

	barberData := factory.MakeBarber(entity.Barber{})
	barberRepositoryMemory.Barbers = []entity.Barber{
		barberData,
	}
	inputDTO := auth.SignInUseCaseInputDTO{
		Email:    barberData.Email,
		Password: barberData.Password,
	}

	ctx := context.Background()
	result, err := sut.Execute(ctx, inputDTO)
	if err != nil {
		t.Errorf("Expected error %v, got %v", nil, err)
	}

	if result.Token == "" {
		t.Errorf("Expected a non-empty token, got %v", "''")
	}
}

func Test_Should_Not_Be_Able_To_Authenticate_When_Email_Does_Not_Exists(t *testing.T) {
	config.LoadConfig("../../..")

	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := auth.NewSignInUseCase(repositoryContainer)

	barberData := factory.MakeBarber(entity.Barber{
		Email: "barber1@mail.com",
	})
	barberRepositoryMemory.Barbers = []entity.Barber{
		barberData,
	}
	inputDTO := auth.SignInUseCaseInputDTO{
		Email:    "johndoe@mail.com",
		Password: barberData.Password,
	}

	ctx := context.Background()
	result, err := sut.Execute(ctx, inputDTO)
	if err != apperr.ErrBarberNotFound {
		t.Errorf("Expected error %v, got %v", apperr.ErrBarberNotFound, err)
	}

	if result.Token != "" {
		t.Errorf("Expected result %v, got %v", "''", result.Token)
	}
}

func Test_Should_Not_Be_Able_To_Authenticate_When_Password_Is_Wrong(t *testing.T) {
	config.LoadConfig("../../..")

	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := auth.NewSignInUseCase(repositoryContainer)

	barberData := factory.MakeBarber(entity.Barber{
		Password: "!Aa12345678",
	})
	barberRepositoryMemory.Barbers = []entity.Barber{
		barberData,
	}
	inputDTO := auth.SignInUseCaseInputDTO{
		Email:    barberData.Email,
		Password: "!Aa12345678901234567890",
	}

	ctx := context.Background()
	result, err := sut.Execute(ctx, inputDTO)
	if err != apperr.ErrBarberUnableToAuthenticate {
		t.Errorf("Expected error %v, got %v", apperr.ErrBarberUnableToAuthenticate, err)
	}

	if result.Token != "" {
		t.Errorf("Expected result %v, got %v", "''", result.Token)
	}
}
