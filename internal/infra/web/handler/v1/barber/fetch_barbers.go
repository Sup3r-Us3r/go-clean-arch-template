package barber

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Sup3r-Us3r/barber-server/internal/usecase/barber"
)

type FetchBarbersHandlerResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"createdAt"`
}

type FetchBarbersHandler struct {
	FetchBarbersUseCase barber.FetchBarbersUseCaseInterface
}

func NewFetchBarbersHandler(fetchBarbersUseCase barber.FetchBarbersUseCaseInterface) *FetchBarbersHandler {
	return &FetchBarbersHandler{
		FetchBarbersUseCase: fetchBarbersUseCase,
	}
}

// FetchBarbers   godoc
// @Summary       Fetch barbers
// @Description   Get list of all barbers
// @Tags          barber
// @Accept        json
// @Produce       json
// @Success       200 {array} barber.FetchBarbersHandlerResponse
// @Failure       500 {object} apperr.AppErr
// @Router        /v1/barber/fetch [get]
// @Security      BearerAuth
func (fbh FetchBarbersHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	useCaseResponse := fbh.FetchBarbersUseCase.Execute(r.Context())

	var result []FetchBarbersHandlerResponse
	for _, barber := range useCaseResponse {
		result = append(result, FetchBarbersHandlerResponse{
			ID:        barber.ID,
			Name:      barber.Name,
			Email:     barber.Email,
			Phone:     barber.Phone,
			CreatedAt: barber.CreatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
