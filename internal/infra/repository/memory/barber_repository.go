package memory

import (
	"context"
	"slices"

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

func (brm BarberRepositoryMemory) GetBarberById(ctx context.Context, barberId string) (entity.Barber, *apperr.AppErr) {
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

func (brm BarberRepositoryMemory) GetBarberByEmail(ctx context.Context, email string) (entity.Barber, *apperr.AppErr) {
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

func (brm BarberRepositoryMemory) FetchBarbers(ctx context.Context) []entity.Barber {
	return brm.Barbers
}

func (brm *BarberRepositoryMemory) CreateBarber(ctx context.Context, barber *entity.Barber) *apperr.AppErr {
	barberAlreadyExists := false

	for _, currentBarber := range brm.Barbers {
		if currentBarber.Email == barber.Email {
			barberAlreadyExists = true
			break
		}
	}

	if barberAlreadyExists {
		return apperr.ErrBarberAlreadyExists
	}

	brm.Barbers = append(brm.Barbers, *barber)

	return nil
}

func (brm *BarberRepositoryMemory) UpdateBarber(ctx context.Context, id string, updateData *entity.Barber) *apperr.AppErr {
	barberIndex := slices.IndexFunc(brm.Barbers, func(barber entity.Barber) bool {
		return barber.ID == id
	})

	if barberIndex == -1 {
		return apperr.ErrBarberNotFound
	}

	if updateData.Name != "" {
		brm.Barbers[barberIndex].Name = updateData.Name
	}

	if updateData.Email != "" {
		brm.Barbers[barberIndex].Email = updateData.Email
	}

	if updateData.Phone != "" {
		brm.Barbers[barberIndex].Phone = updateData.Phone
	}

	return nil
}

func (brm *BarberRepositoryMemory) DeleteBarber(ctx context.Context, id string) *apperr.AppErr {
	barberAlreadyExists := slices.IndexFunc(brm.Barbers, func(barber entity.Barber) bool {
		return barber.ID == id
	})

	if barberAlreadyExists == -1 {
		return apperr.ErrBarberNotFound
	}

	slices.DeleteFunc(brm.Barbers, func(barber entity.Barber) bool {
		return barber.ID == id
	})

	return nil
}
