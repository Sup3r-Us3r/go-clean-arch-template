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

func Test_Should_Be_Able_To_Get_Barber_By_Email(t *testing.T) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewGetBarberByEmailUseCase(repositoryContainer)

	barberEmail := "barber1@mail.com"
	barberData := factory.MakeBarber(entity.Barber{Email: barberEmail})
	barberRepositoryMemory.Barbers = []entity.Barber{barberData}

	inputDTO := barber.GetBarberByEmailUseCaseInputDTO{
		Email: barberEmail,
	}

	ctx := context.Background()
	barber, err := sut.Execute(ctx, inputDTO)
	if err != nil {
		t.Errorf("Expected error %v, got %v", nil, err)
	}

	if barber.Email != barberEmail {
		t.Errorf("Expected result %v, got %v", barberEmail, barber.ID)
	}
}

func Test_Should_Not_Be_Able_To_Get_Barber_By_Email_When_Barber_Not_Exists(t *testing.T) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewGetBarberByEmailUseCase(repositoryContainer)

	barberEmail := "barber1@mail.com"
	barberData := factory.MakeBarber(entity.Barber{Email: barberEmail})
	barberRepositoryMemory.Barbers = []entity.Barber{barberData}

	inputDTO := barber.GetBarberByEmailUseCaseInputDTO{
		Email: "non_existent_email@mail.com",
	}

	ctx := context.Background()
	_, err := sut.Execute(ctx, inputDTO)
	if err != apperr.ErrBarberNotFound {
		t.Errorf("Expected error %v, got %v", apperr.ErrBarberNotFound, err)
	}
}
