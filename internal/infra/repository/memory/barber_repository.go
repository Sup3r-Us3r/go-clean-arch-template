package memory

import (
	"context"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
)

type BarberRepositoryMemory struct {
	Barbers []entity.Barber
}

func NewBarberRepositoryMemory() *BarberRepositoryMemory {
	return &BarberRepositoryMemory{
		Barbers: make([]entity.Barber, 0),
	}
}

func (brm *BarberRepositoryMemory) CreateBarber(ctx context.Context, barber *entity.Barber) *apperr.AppErr {
	barberAlreadyExists := false

	for _, currentBarber := range brm.Barbers {
		if currentBarber.Email == barber.Email {

			barberAlreadyExists = true
		}
	}

	if barberAlreadyExists {
		return apperr.ErrBarberAlreadyExists
	}

	brm.Barbers = append(brm.Barbers, *barber)

	return nil
}

func (brm *BarberRepositoryMemory) FetchBarbers(ctx context.Context) []entity.Barber {
	return brm.Barbers
}

func (brm *BarberRepositoryMemory) GetBarberById(ctx context.Context, barberId string) (entity.Barber, *apperr.AppErr) {
	var barber entity.Barber

	for _, currentBarber := range brm.Barbers {
		if currentBarber.ID == barberId {
			barber = currentBarber
			break
		}
	}

	if barber.ID == "" {
		return entity.Barber{}, apperr.ErrBarberNotFound
	}

	return barber, nil
}

func (brm *BarberRepositoryMemory) GetBarberByEmail(ctx context.Context, email string) (entity.Barber, *apperr.AppErr) {
	var barber entity.Barber

	for _, currentBarber := range brm.Barbers {
		if currentBarber.Email == email {
			barber = currentBarber
			break
		}
	}

	if barber.ID == "" {
		return entity.Barber{}, apperr.ErrBarberNotFound
	}

	return barber, nil
}
