package usecase

import (
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/auth"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/barber"
)

type UseCaseContainer struct {
	SignInUseCase           auth.SignInUseCaseInterface
	GetBarberByIdUseCase    barber.GetBarberByIdUseCaseInterface
	GetBarberByEmailUseCase barber.GetBarberByEmailUseCaseInterface
	FetchBarbersUseCase     barber.FetchBarbersUseCaseInterface
	CreateBarberUseCase     barber.CreateBarberUseCaseInterface
	UpdateBarberUseCase     barber.UpdateBarberUseCaseInterface
	DeleteBarberUseCase     barber.DeleteBarberUseCaseInterface
}

func GetUseCases(repositoryContainer repository.RepositoryContainer) *UseCaseContainer {
	return &UseCaseContainer{
		SignInUseCase:           auth.NewSignInUseCase(repositoryContainer),
		GetBarberByIdUseCase:    barber.NewGetBarberByIdUseCase(repositoryContainer),
		GetBarberByEmailUseCase: barber.NewGetBarberByEmailUseCase(repositoryContainer),
		FetchBarbersUseCase:     barber.NewFetchBarbersUseCase(repositoryContainer),
		CreateBarberUseCase:     barber.NewCreateBarberUseCase(repositoryContainer),
		UpdateBarberUseCase:     barber.NewUpdateBarberUseCase(repositoryContainer),
		DeleteBarberUseCase:     barber.NewDeleteBarberUseCase(repositoryContainer),
	}
}
