package barber

import (
	"context"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/gateway"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
)

type UpdateBarberUseCaseInputDTO struct {
	ID    string
	Name  string
	Email string
	Phone string
}

type UpdateBarberUseCaseInterface interface {
	Execute(ctx context.Context, input UpdateBarberUseCaseInputDTO) *apperr.AppErr
}

type UpdateBarberUseCase struct {
	BarberRepository gateway.BarberGatewayInterface
}

func NewUpdateBarberUseCase(repositoryContainer repository.RepositoryContainer) *UpdateBarberUseCase {
	return &UpdateBarberUseCase{
		BarberRepository: repositoryContainer.BarberRepository,
	}
}

func (ubuc *UpdateBarberUseCase) Execute(ctx context.Context, input UpdateBarberUseCaseInputDTO) *apperr.AppErr {
	err := ubuc.BarberRepository.UpdateBarber(ctx, input.ID, &entity.Barber{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	})
	if err != nil {
		return err
	}

	return nil
}
