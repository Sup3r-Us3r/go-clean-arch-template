package barber_test

import (
	"context"
	"testing"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository/memory"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/barber"
	"github.com/Sup3r-Us3r/barber-server/test/factory"
)

func Test_Should_Be_Able_To_Fetch_Barbers(t *testing.T) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewFetchBarbersUseCase(repositoryContainer)

	barberData1 := factory.MakeBarber(entity.Barber{})
	barberData2 := factory.MakeBarber(entity.Barber{})
	barberRepositoryMemory.Barbers = []entity.Barber{barberData1, barberData2}

	ctx := context.Background()
	barbers := sut.Execute(ctx)

	if len(barbers) != 2 {
		t.Errorf("Expected result %v, got %v", 2, len(barbers))
	}
}
