package barber

import (
	"context"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/gateway"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
)

type FetchBarbersUseCaseOutputDTO = []entity.Barber

type FetchBarbersUseCaseInterface interface {
	Execute(ctx context.Context) FetchBarbersUseCaseOutputDTO
}

type FetchBarbersUseCase struct {
	BarberRepository gateway.BarberGatewayInterface
}

func NewFetchBarbersUseCase(repositoryContainer repository.RepositoryContainer) *FetchBarbersUseCase {
	return &FetchBarbersUseCase{
		BarberRepository: repositoryContainer.BarberRepository,
	}
}

func (fbuc FetchBarbersUseCase) Execute(ctx context.Context) FetchBarbersUseCaseOutputDTO {
	barbers := fbuc.BarberRepository.FetchBarbers(ctx)

	return barbers
}
