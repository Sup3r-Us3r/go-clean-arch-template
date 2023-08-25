package barber

import (
	"net/http"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/barber"
	"github.com/go-chi/chi/v5"
)

type DeleteBarberHandler struct {
	DeleteBarberUseCase barber.DeleteBarberUseCaseInterface
}

func NewDeleteBarberHandler(DeleteBarberUseCase barber.DeleteBarberUseCaseInterface) *DeleteBarberHandler {
	return &DeleteBarberHandler{
		DeleteBarberUseCase: DeleteBarberUseCase,
	}
}

// DeleteBarber   godoc
// @Summary       Delete Barber
// @Description   Delete barber by ID
// @Tags          barber
// @Accept        json
// @Produce       json
// @Param         id path string true "barber id" Format(uuid)
// @Success       204
// @Failure       404 {object} apperr.AppErr
// @Failure       500 {object} apperr.AppErr
// @Router        /v1/barber/delete/{id} [delete]
// @Security      BearerAuth
func (dbh *DeleteBarberHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	barberId := chi.URLParam(r, "id")

	useCaseErr := dbh.DeleteBarberUseCase.Execute(
		r.Context(),
		barber.DeleteBarberUseCaseInputDTO{ID: barberId},
	)
	if useCaseErr != nil {
		apperr.NewHttpError(w, useCaseErr)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
