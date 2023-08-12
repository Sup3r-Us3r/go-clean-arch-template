package barber

import (
	"context"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/gateway"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
)

type CreateBarberUseCaseInputDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type CreateBarberUseCaseInterface interface {
	Execute(ctx context.Context, input CreateBarberUseCaseInputDTO) *apperr.AppErr
}

type CreateBarberUseCase struct {
	BarberRepository gateway.BarberGatewayInterface
}

func NewCreateBarberUseCase(repositoryContainer repository.RepositoryContainer) *CreateBarberUseCase {
	return &CreateBarberUseCase{
		BarberRepository: repositoryContainer.BarberRepository,
	}
}

func (cbuc *CreateBarberUseCase) Execute(ctx context.Context, input CreateBarberUseCaseInputDTO) *apperr.AppErr {
	barber, err := entity.NewBarber(input.Name, input.Email, input.Phone, input.Password)
	if err != nil {
		return err
	}

	err = cbuc.BarberRepository.CreateBarber(ctx, barber)
	if err != nil {
		return err
	}

	return nil
}
