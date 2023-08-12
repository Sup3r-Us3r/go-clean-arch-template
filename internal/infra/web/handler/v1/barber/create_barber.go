package barber

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/barber"
)

type CreateBarberHandlerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type CreateBarberHandler struct {
	CreateBarberUseCase barber.CreateBarberUseCaseInterface
}

func NewCreateBarberHandler(createBarberUseCase barber.CreateBarberUseCaseInterface) *CreateBarberHandler {
	return &CreateBarberHandler{
		CreateBarberUseCase: createBarberUseCase,
	}
}

// CreateBarber   godoc
// @Summary       Create a barber
// @Description   Create a new barber
// @Tags          barber
// @Accept        json
// @Produce       json
// @Param         request body barber.CreateBarberHandlerRequest true "barber data"
// @Success       201
// @Failure       400 {object} apperr.AppErr
// @Failure       404 {object} apperr.AppErr
// @Failure       500 {object} apperr.AppErr
// @Router        /v1/barber/create [post]
func (cbh CreateBarberHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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

	var requestInput CreateBarberHandlerRequest
	err = json.Unmarshal(body, &requestInput)
	if err != nil {
		apperr.NewHttpError(w, apperr.NewBadRequestError(err.Error()))

		return
	}

	useCaseErr := cbh.CreateBarberUseCase.Execute(r.Context(), barber.CreateBarberUseCaseInputDTO{
		Name:     requestInput.Name,
		Email:    requestInput.Email,
		Password: requestInput.Password,
		Phone:    requestInput.Phone,
	})
	if useCaseErr != nil {
		apperr.NewHttpError(w, useCaseErr)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
