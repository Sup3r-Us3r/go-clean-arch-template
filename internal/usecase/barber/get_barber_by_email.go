package barber

import (
	"context"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/gateway"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
)

type GetBarberByEmailUseCaseInputDTO struct {
	Email string
}

type GetBarberByEmailUseCaseOutputDTO = entity.Barber

type GetBarberByEmailUseCaseInterface interface {
	Execute(ctx context.Context, input GetBarberByEmailUseCaseInputDTO) (GetBarberByEmailUseCaseOutputDTO, *apperr.AppErr)
}

type GetBarberByEmailUseCase struct {
	BarberRepository gateway.BarberGatewayInterface
}

func NewGetBarberByEmailUseCase(repositoryContainer repository.RepositoryContainer) *GetBarberByEmailUseCase {
	return &GetBarberByEmailUseCase{
		BarberRepository: repositoryContainer.BarberRepository,
	}
}

func (gbbeuc GetBarberByEmailUseCase) Execute(ctx context.Context, input GetBarberByEmailUseCaseInputDTO) (GetBarberByEmailUseCaseOutputDTO, *apperr.AppErr) {
	barber, err := gbbeuc.BarberRepository.GetBarberByEmail(ctx, input.Email)
	if err != nil {
		return GetBarberByEmailUseCaseOutputDTO{}, err
	}

	return barber, nil
}
