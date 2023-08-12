package usecase

import (
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/auth"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/barber"
)

type UseCaseContainer struct {
	SignInUseCase       auth.SignInUseCaseInterface
	CreateBarberUseCase barber.CreateBarberUseCaseInterface
}

func GetUseCases(repositoryContainer repository.RepositoryContainer) *UseCaseContainer {
	return &UseCaseContainer{
		SignInUseCase:       auth.NewSignInUseCase(repositoryContainer),
		CreateBarberUseCase: barber.NewCreateBarberUseCase(repositoryContainer),
	}
}
