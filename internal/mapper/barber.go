package mapper

import (
	"time"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
)

type BarberMongo struct {
	ID           string    `bson:"_id,omitempty"`
	Name         string    `bson:"name"`
	Email        string    `bson:"email"`
	PasswordHash string    `bson:"password_hash"`
	Phone        string    `bson:"phone"`
	CreatedAt    time.Time `bson:"created_at"`
}

func MapBarberEntityToMongo(barber *entity.Barber) *BarberMongo {
	return &BarberMongo{
		ID:           barber.ID,
		Name:         barber.Name,
		Email:        barber.Email,
		PasswordHash: barber.PasswordHash,
		Phone:        barber.Phone,
		CreatedAt:    barber.CreatedAt,
	}
}

func MapBarberMongoToEntity(barber BarberMongo) *entity.Barber {
	return &entity.Barber{
		ID:           barber.ID,
		Name:         barber.Name,
		Email:        barber.Email,
		PasswordHash: barber.PasswordHash,
		Phone:        barber.Phone,
		CreatedAt:    barber.CreatedAt,
	}
}
