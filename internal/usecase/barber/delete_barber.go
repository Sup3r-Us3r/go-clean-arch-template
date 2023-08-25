package barber

import (
	"context"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/gateway"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
)

type DeleteBarberUseCaseInputDTO struct {
	ID string
}

type DeleteBarberUseCaseInterface interface {
	Execute(ctx context.Context, input DeleteBarberUseCaseInputDTO) *apperr.AppErr
}

type DeleteBarberUseCase struct {
	BarberRepository gateway.BarberGatewayInterface
}

func NewDeleteBarberUseCase(repositoryContainer repository.RepositoryContainer) *DeleteBarberUseCase {
	return &DeleteBarberUseCase{
		BarberRepository: repositoryContainer.BarberRepository,
	}
}

func (dbuc *DeleteBarberUseCase) Execute(ctx context.Context, input DeleteBarberUseCaseInputDTO) *apperr.AppErr {
	err := dbuc.BarberRepository.DeleteBarber(ctx, input.ID)
	if err != nil {
		return err
	}

	return nil
}
