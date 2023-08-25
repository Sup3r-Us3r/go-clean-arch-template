package barber

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/barber"
	"github.com/go-chi/chi/v5"
)

type UpdateBarberHandlerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UpdateBarberHandler struct {
	UpdateBarberUseCase barber.UpdateBarberUseCaseInterface
}

func NewUpdateBarberHandler(UpdateBarberUseCase barber.UpdateBarberUseCaseInterface) *UpdateBarberHandler {
	return &UpdateBarberHandler{
		UpdateBarberUseCase: UpdateBarberUseCase,
	}
}

// UpdateBarber   godoc
// @Summary       Update a barber
// @Description   Update barber data
// @Tags          barber
// @Accept        json
// @Produce       json
// @Param         request body barber.UpdateBarberHandlerRequest true "update barber data"
// @Success       200
// @Failure       400 {object} apperr.AppErr
// @Failure       404 {object} apperr.AppErr
// @Failure       500 {object} apperr.AppErr
// @Router        /v1/barber/update/{id} [put]
func (ubh UpdateBarberHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		apperr.NewHttpError(w, apperr.NewBadRequestError(err.Error()))

		return
	}

	if !json.Valid(body) {
		apperr.NewHttpError(w, apperr.NewBadRequestError("invalid json"))

		return
	}

	var requestInput UpdateBarberHandlerRequest
	err = json.Unmarshal(body, &requestInput)
	if err != nil {
		apperr.NewHttpError(w, apperr.NewBadRequestError(err.Error()))

		return
	}

	barberId := chi.URLParam(r, "id")

	useCaseErr := ubh.UpdateBarberUseCase.Execute(r.Context(), barber.UpdateBarberUseCaseInputDTO{
		ID:    barberId,
		Name:  requestInput.Name,
		Email: requestInput.Email,
		Phone: requestInput.Phone,
	})
	if useCaseErr != nil {
		apperr.NewHttpError(w, useCaseErr)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
