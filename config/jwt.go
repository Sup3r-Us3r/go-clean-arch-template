package config

import (
	"strings"
	"time"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/golang-jwt/jwt"
)

type TokenData struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GenerateToken(tokenData TokenData) (string, *apperr.AppErr) {
	jwtSecretKey := GetEnvAsString("JWT_SECRET_KEY")
	jwtExpirationInMinutes := time.Duration(GetEnvAsInt("JWT_EXPIRATION_IN_MINUTES"))

	claims := jwt.MapClaims{
		"id":    tokenData.ID,
		"name":  tokenData.Name,
		"email": tokenData.Email,
		"exp":   time.Now().Add(time.Minute * jwtExpirationInMinutes).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", apperr.NewInternalServerError("error trying to generate jwt token")
	}

	return tokenString, nil
}

func VerifyToken(tokenValue string) (TokenData, *apperr.AppErr) {
	jwtSecretKey := GetEnvAsString("JWT_SECRET_KEY")

	token, err := jwt.Parse(removeBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(jwtSecretKey), nil
		}

		return nil, apperr.NewBadRequestError("invalid token")
	})
	if err != nil {
		return TokenData{}, apperr.NewUnauthorizedError(
			strings.ToLower(err.Error()),
		)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return TokenData{}, apperr.NewUnauthorizedError("invalid token")
	}

	return TokenData{
		ID:    claims["id"].(string),
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
	}, nil
}

func removeBearerPrefix(token string) string {
	token = strings.TrimPrefix(token, "Bearer ")

	return token
}
