package entity

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"regexp"
	"strings"
	"time"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/log"
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

const (
	SALT_SIZE int = 16
)

// NewBarber create a new instance of Barber
func NewBarber(name string, email string, phone string, password string) (*Barber, *apperr.AppErr) {
	generateSalt := GenerateRandomSalt(SALT_SIZE)

	barber := &Barber{
		ID:           uuid.New().String(),
		Name:         name,
		Email:        email,
		Password:     password,
		PasswordHash: HashPassword(password, generateSalt),
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

	if !emailIsValid(email) {
		return apperr.ErrBarberEmailIsInvalid
	}

	password := strings.TrimSpace(b.Password)

	if password == "" {
		return apperr.ErrBarberFieldPasswordIsRequired
	}

	if !passwordIsValid(password) {
		return apperr.ErrBarberPasswordIsInvalid
	}

	if strings.TrimSpace(b.Phone) == "" {
		return apperr.ErrBarberFieldPhoneIsRequired
	}

	return nil
}

// emailIsValid check if the email is valid
func emailIsValid(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return regex.MatchString(email)
}

// passwordIsValid check if the password has at least 8 digits, 1 lowercase letter,
// 1 uppercase letter, 1 number and 1 special character
func passwordIsValid(password string) bool {
	if len(password) < 8 {
		return false
	}

	if !regexp.MustCompile("[a-z]").MatchString(password) {
		return false
	}

	if !regexp.MustCompile("[A-Z]").MatchString(password) {
		return false
	}

	if !regexp.MustCompile(`[\W_]`).MatchString(password) {
		return false
	}

	if !regexp.MustCompile(`\d`).MatchString(password) {
		return false
	}

	return true
}

// GenerateRandomSalt generate 16 bytes randomly and securely using the
// Cryptographically secure pseudorandom number generator (CSPRNG)
// in the crypto.rand package
func GenerateRandomSalt(saltSize int) []byte {
	salt := make([]byte, saltSize)

	_, err := rand.Read(salt[:])
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	return salt
}

// HashPassword combine password and salt then hash them using the SHA-512
// hashing algorithm and then return the hashed password as a hex string
func HashPassword(password string, salt []byte) string {
	// Convert password string to byte slice
	passwordBytes := []byte(password)

	// Create SHA-512 hasher
	sha512Hasher := sha512.New()

	// Append salt to password
	passwordBytes = append(passwordBytes, salt...)

	// Write password bytes to the hasher
	sha512Hasher.Write(passwordBytes)

	// Get the SHA-512 hashed password
	hashedPasswordBytes := sha512Hasher.Sum(nil)

	// Combine salt and hashed password
	combined := append(salt, hashedPasswordBytes...)

	// Convert the combined value to a hex string
	hashedPasswordHex := hex.EncodeToString(combined)

	return hashedPasswordHex
}

// DoPasswordsMatch check if two passwords match
func DoPasswordsMatch(hashedPassword string, password string) bool {
	// Decode the hashed password from hex to bytes
	hashedPasswordBytes, err := hex.DecodeString(hashedPassword)
	if err != nil {
		return false
	}

	// Extract the salt from the hashed password
	salt := hashedPasswordBytes[:SALT_SIZE]

	// Hash the input password using the extracted salt
	currentPasswordHash := HashPassword(password, salt)

	return currentPasswordHash == hashedPassword
}
