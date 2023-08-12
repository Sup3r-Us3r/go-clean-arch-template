package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/auth"
)

type SignInHandlerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInHandlerResponse struct {
	Token string `json:"token"`
}

type SignInHandler struct {
	SignInUseCase auth.SignInUseCaseInterface
}

func NewSignInHandler(signInUseCase auth.SignInUseCaseInterface) *SignInHandler {
	return &SignInHandler{
		SignInUseCase: signInUseCase,
	}
}

// SignIn         godoc
// @Summary       Authentication
// @Description   Authentication with email and password
// @Tags          auth
// @Accept        json
// @Produce       json
// @Param         request body auth.SignInHandlerRequest true "authentication data"
// @Success       200 {object} auth.SignInHandlerResponse
// @Failure       400 {object} apperr.AppErr
// @Failure       404 {object} apperr.AppErr
// @Failure       500 {object} apperr.AppErr
// @Router        /v1/auth/sign-in [post]
func (sih *SignInHandler) Handle(w http.ResponseWriter, r *http.Request) {
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

	var requestInput SignInHandlerRequest
	err = json.Unmarshal(body, &requestInput)
	if err != nil {
		apperr.NewHttpError(w, apperr.NewBadRequestError(err.Error()))

		return
	}

	useCaseResponse, useCaseErr := sih.SignInUseCase.Execute(r.Context(), auth.SignInUseCaseInputDTO{
		Email:    requestInput.Email,
		Password: requestInput.Password,
	})
	if useCaseErr != nil {
		apperr.NewHttpError(w, useCaseErr)

		return
	}

	result := SignInHandlerResponse{
		Token: useCaseResponse.Token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+result.Token)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
