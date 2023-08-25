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

func Test_Should_Be_Able_To_Update_Barber(t *testing.T) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewUpdateBarberUseCase(repositoryContainer)

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"
	barberData := factory.MakeBarber(entity.Barber{ID: barberId})
	barberRepositoryMemory.Barbers = []entity.Barber{barberData}
	updateName := "Barber1"
	updateEmail := "barber1@mail.com"

	inputDTO := barber.UpdateBarberUseCaseInputDTO{
		ID:    barberId,
		Name:  updateName,
		Email: updateEmail,
	}

	ctx := context.Background()
	err := sut.Execute(ctx, inputDTO)
	if err != nil {
		t.Errorf("Expected error %v, got %v", nil, err)
	}

	if barberRepositoryMemory.Barbers[0].Name != updateName {
		t.Errorf("Expected result %v, got %v", updateName, barberRepositoryMemory.Barbers[0].Name)
	}

	if barberRepositoryMemory.Barbers[0].Email != updateEmail {
		t.Errorf("Expected result %v, got %v", updateEmail, barberRepositoryMemory.Barbers[0].Email)
	}
}

func Test_Should_Not_Be_Able_To_Update_Barber_When_Barber_Not_Exists(t *testing.T) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewUpdateBarberUseCase(repositoryContainer)

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"
	barberData := factory.MakeBarber(entity.Barber{ID: barberId})
	barberRepositoryMemory.Barbers = []entity.Barber{barberData}
	updateName := "Barber1"
	updateEmail := "barber1@mail.com"

	inputDTO := barber.UpdateBarberUseCaseInputDTO{
		ID:    "non-existing-id",
		Name:  updateName,
		Email: updateEmail,
	}

	ctx := context.Background()
	err := sut.Execute(ctx, inputDTO)
	if err != apperr.ErrBarberNotFound {
		t.Errorf("Expected error %v, got %v", apperr.ErrBarberNotFound, err)
	}
}
