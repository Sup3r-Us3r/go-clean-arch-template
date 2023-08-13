package entity

import (
	"strings"
	"time"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/util"
	"github.com/google/uuid"
)

type Barber struct {
	ID           string
	Name         string
	Email        string
	Password     string
	PasswordHash string
	Phone        string
	CreatedAt    time.Time
}

// NewBarber create a new instance of Barber
func NewBarber(name string, email string, phone string, password string) (*Barber, *apperr.AppErr) {
	generateSalt := util.GenerateRandomSalt(util.SALT_SIZE)

	barber := &Barber{
		ID:           uuid.New().String(),
		Name:         name,
		Email:        email,
		Password:     password,
		PasswordHash: util.HashPassword(password, generateSalt),
		Phone:        phone,
		CreatedAt:    time.Now(),
	}

	if err := barber.Validate(); err != nil {
		return nil, err
	}

	return barber, nil
}

// Validate checks that all attributes that need validation are really
// valid upon entity creation
func (b Barber) Validate() *apperr.AppErr {
	if strings.TrimSpace(b.Name) == "" {
		return apperr.ErrBarberFieldNameIsRequired
	}

	email := strings.TrimSpace(b.Email)

	if email == "" {
		return apperr.ErrBarberFieldEmailIsRequired
	}

	if !util.EmailIsValid(email) {
		return apperr.ErrBarberEmailIsInvalid
	}

	password := strings.TrimSpace(b.Password)

	if password == "" {
		return apperr.ErrBarberFieldPasswordIsRequired
	}

	if !util.PasswordIsValid(password) {
		return apperr.ErrBarberPasswordIsInvalid
	}

	if strings.TrimSpace(b.Phone) == "" {
		return apperr.ErrBarberFieldPhoneIsRequired
	}

	return nil
}
