package barber

import (
	"context"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/gateway"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
)

type GetBarberByIdUseCaseInputDTO struct {
	ID string
}

type GetBarberByIdUseCaseOutputDTO = entity.Barber

type GetBarberByIdUseCaseInterface interface {
	Execute(ctx context.Context, input GetBarberByIdUseCaseInputDTO) (GetBarberByIdUseCaseOutputDTO, *apperr.AppErr)
}

type GetBarberByIdUseCase struct {
	BarberRepository gateway.BarberGatewayInterface
}

func NewGetBarberByIdUseCase(repositoryContainer repository.RepositoryContainer) *GetBarberByIdUseCase {
	return &GetBarberByIdUseCase{
		BarberRepository: repositoryContainer.BarberRepository,
	}
}

func (gbbiuc GetBarberByIdUseCase) Execute(ctx context.Context, input GetBarberByIdUseCaseInputDTO) (GetBarberByIdUseCaseOutputDTO, *apperr.AppErr) {
	barber, err := gbbiuc.BarberRepository.GetBarberById(ctx, input.ID)
	if err != nil {
		return GetBarberByIdUseCaseOutputDTO{}, err
	}

	return barber, nil
}
