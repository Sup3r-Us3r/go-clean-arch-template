package factory

import (
	"fmt"
	"strings"
	"time"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/jaswdr/faker"
)

func MakeBarber(override entity.Barber) entity.Barber {
	fake := faker.New()

	generateSalt := entity.GenerateRandomSalt(16)
	password := fmt.Sprintf(
		"%v%v%v%v%v",
		fake.Internet().Password(),
		strings.ToLower(fake.RandomLetter()),
		strings.ToUpper(fake.RandomLetter()),
		fake.RandomDigit(),
		fake.RandomStringElement([]string{"!", "@", "#", "$", "&", "*", "(", ")"}),
	)
	hashedPassword := entity.HashPassword(password, generateSalt)

	barber := &entity.Barber{
		ID:           uuid.New().String(),
		Name:         fake.Person().Name(),
		Email:        fake.Person().Contact().Email,
		Password:     password,
		PasswordHash: hashedPassword,
		Phone:        fake.Person().Contact().Phone,
		CreatedAt:    time.Now(),
	}

	if override.ID != "" {
		barber.ID = override.ID
	}

	if override.Name != "" {
		barber.Name = override.Name
	}

	if override.Email != "" {
		barber.Email = override.Email
	}

	if override.Password != "" {
		barber.Password = override.Password
	}

	if override.PasswordHash != "" {
		barber.PasswordHash = override.PasswordHash
	}

	if override.Phone != "" {
		barber.Phone = override.Phone
	}

	if !override.CreatedAt.IsZero() {
		barber.CreatedAt = override.CreatedAt
	}

	return *barber
}
