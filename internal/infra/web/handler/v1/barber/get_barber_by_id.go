package barber

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/barber"
	"github.com/go-chi/chi/v5"
)

type GetBarberByIdHandlerResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetBarberByIdHandler struct {
	GetBarberByIdUseCase barber.GetBarberByIdUseCaseInterface
}

func NewGetBarberByIdHandler(getBarberByIdUseCase barber.GetBarberByIdUseCaseInterface) *GetBarberByIdHandler {
	return &GetBarberByIdHandler{
		GetBarberByIdUseCase: getBarberByIdUseCase,
	}
}

// GetBarberById   godoc
// @Summary        Get barber
// @Description    Get barber by ID
// @Tags           barber
// @Accept         json
// @Produce        json
// @Param          id path string true "barber id" Format(uuid)
// @Success        200 {object} barber.GetBarberByIdHandlerResponse
// @Failure        404 {object} apperr.AppErr
// @Failure        500 {object} apperr.AppErr
// @Router         /v1/barber/get/{id} [get]
// @Security       BearerAuth
func (gbbih GetBarberByIdHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	barberId := chi.URLParam(r, "id")

	useCaseResponse, useCaseErr := gbbih.GetBarberByIdUseCase.Execute(
		r.Context(),
		barber.GetBarberByIdUseCaseInputDTO{ID: barberId},
	)
	if useCaseErr != nil {
		apperr.NewHttpError(w, useCaseErr)

		return
	}

	result := GetBarberByIdHandlerResponse{
		ID:        useCaseResponse.ID,
		Name:      useCaseResponse.Name,
		Email:     useCaseResponse.Email,
		Phone:     useCaseResponse.Phone,
		CreatedAt: useCaseResponse.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
