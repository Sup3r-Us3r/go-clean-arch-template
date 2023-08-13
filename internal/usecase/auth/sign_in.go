package auth

import (
	"context"

	"github.com/Sup3r-Us3r/barber-server/config"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/gateway"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
	"github.com/Sup3r-Us3r/barber-server/internal/util"
)

type SignInUseCaseInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInUseCaseOutputDTO struct {
	Token string `json:"token"`
}

type SignInUseCaseInterface interface {
	Execute(ctx context.Context, input SignInUseCaseInputDTO) (SignInUseCaseOutputDTO, *apperr.AppErr)
}

type SignInUseCase struct {
	BarberRepository gateway.BarberGatewayInterface
}

func NewSignInUseCase(repositoryContainer repository.RepositoryContainer) *SignInUseCase {
	return &SignInUseCase{
		BarberRepository: repositoryContainer.BarberRepository,
	}
}

func (siuc *SignInUseCase) Execute(ctx context.Context, input SignInUseCaseInputDTO) (SignInUseCaseOutputDTO, *apperr.AppErr) {
	barber, err := siuc.BarberRepository.GetBarberByEmail(ctx, input.Email)
	if err != nil {
		return SignInUseCaseOutputDTO{}, err
	}

	passwordMatch := util.DoPasswordsMatch(barber.PasswordHash, input.Password)

	if !passwordMatch {
		return SignInUseCaseOutputDTO{}, apperr.ErrBarberUnableToAuthenticate
	}

	token, err := config.GenerateToken(
		config.TokenData{
			ID:    barber.ID,
			Name:  barber.Name,
			Email: barber.Email,
		},
	)
	if err != nil {
		return SignInUseCaseOutputDTO{}, err
	}

	return SignInUseCaseOutputDTO{Token: token}, nil
}
