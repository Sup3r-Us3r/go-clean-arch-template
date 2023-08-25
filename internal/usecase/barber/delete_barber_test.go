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

func Test_Should_Be_Able_To_Delete_Barber(t *testing.T) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewDeleteBarberUseCase(repositoryContainer)

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"
	barberData := factory.MakeBarber(entity.Barber{ID: barberId})
	barberRepositoryMemory.Barbers = []entity.Barber{barberData}

	inputDTO := barber.DeleteBarberUseCaseInputDTO{
		ID: barberId,
	}

	ctx := context.Background()
	err := sut.Execute(ctx, inputDTO)
	if err != nil {
		t.Errorf("Expected error %v, got %v", nil, err)
	}

	if len(barberRepositoryMemory.Barbers) == 0 {
		t.Errorf("Expected result %v, got %v", 0, len(barberRepositoryMemory.Barbers))
	}
}

func Test_Should_Not_Be_Able_To_Delete_Barber_When_Barber_Not_Exists(t *testing.T) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewDeleteBarberUseCase(repositoryContainer)

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"
	barberData := factory.MakeBarber(entity.Barber{ID: barberId})
	barberRepositoryMemory.Barbers = []entity.Barber{barberData}

	inputDTO := barber.DeleteBarberUseCaseInputDTO{
		ID: "non-existent-id",
	}

	ctx := context.Background()
	err := sut.Execute(ctx, inputDTO)
	if err != apperr.ErrBarberNotFound {
		t.Errorf("Expected error %v, got %v", apperr.ErrBarberNotFound, err)
	}
}
