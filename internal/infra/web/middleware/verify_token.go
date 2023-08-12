package middleware

import (
	"net/http"

	"github.com/Sup3r-Us3r/barber-server/config"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
)

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		tokenData, err := config.VerifyToken(token)
		if err != nil {
			apperr.NewHttpError(w, err)

			return
		}

		repository := repository.RepositoryContainerInstance
		_, err = repository.BarberRepository.GetBarberById(r.Context(), tokenData.ID)
		if err != nil {
			apperr.NewHttpError(w, apperr.NewUnauthorizedError("unauthorized access"))

			return
		}

		next.ServeHTTP(w, r)
	})
}
