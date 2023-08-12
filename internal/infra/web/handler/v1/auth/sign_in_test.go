package auth_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sup3r-Us3r/barber-server/config"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository/memory"
	authHandler "github.com/Sup3r-Us3r/barber-server/internal/infra/web/handler/v1/auth"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/auth"
	"github.com/Sup3r-Us3r/barber-server/test/factory"
)

func controller() (*memory.BarberRepositoryMemory, *authHandler.SignInHandler) {
	config.LoadConfig("../../../../../..")

	barberRepositoryInMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryInMemory,
	}
	sut := auth.NewSignInUseCase(repositoryContainer)
	handler := authHandler.NewSignInHandler(sut)

	return barberRepositoryInMemory, handler
}

func Test_Should_Be_Able_To_Authenticate(t *testing.T) {
	repository, handler := controller()

	barberData := factory.MakeBarber(entity.Barber{
		Email:        "barber1@mail.com",
		Password:     "!Aa12345678",
		PasswordHash: "a45ed3ef2af41f0f091148764e3b1876f34b78334538ddf484b63e41c380823e794e55714af9a3befcf6be34ad1b5a2714e74409d569489e63f74e82f58efd74c594a6d649d43e3a70a66824bbda4cb9",
	})
	repository.Barbers = []entity.Barber{
		barberData,
	}
	payload := authHandler.SignInHandlerRequest{
		Email:    barberData.Email,
		Password: barberData.Password,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, rr.Code)
	}
}

func Test_Should_Not_Be_Able_To_Authenticate_When_Barber_Does_Not_Exists(t *testing.T) {
	repository, handler := controller()

	barberData := factory.MakeBarber(entity.Barber{
		Email:        "barber1@mail.com",
		Password:     "!Aa12345678",
		PasswordHash: "a45ed3ef2af41f0f091148764e3b1876f34b78334538ddf484b63e41c380823e794e55714af9a3befcf6be34ad1b5a2714e74409d569489e63f74e82f58efd74c594a6d649d43e3a70a66824bbda4cb9",
	})

	repository.Barbers = []entity.Barber{
		barberData,
	}

	payload := authHandler.SignInHandlerRequest{
		Email:    "johndoe@mail.com",
		Password: barberData.Password,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/v1/auth/sign-in", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status %v, got %v", http.StatusNotFound, rr.Code)
	}

	var response apperr.AppErr
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("unable to parse response data")
	}

	if response != *apperr.ErrBarberNotFound {
		t.Errorf("Expected error %v, got %v", apperr.ErrBarberNotFound.Error(), response.Error())
	}
}

func Test_Should_Not_Be_Able_To_Authenticate_When_Http_Method_Is_Wrong(t *testing.T) {
	_, handler := controller()

	req := httptest.NewRequest(http.MethodGet, "/v1/barber/auth", nil)

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %v, got %v", http.StatusMethodNotAllowed, rr.Code)
	}
}

type errorReader struct{}

func (er *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated error while reading request body")
}

func Test_Should_Not_Be_Able_To_Authenticate_When_Unable_To_Read_Body_Content(t *testing.T) {
	_, handler := controller()

	req := httptest.NewRequest(http.MethodPost, "/v1/barber/auth", &errorReader{})

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %d", http.StatusBadRequest, rr.Code)
	}
}

func Test_Should_Not_Be_Able_To_Authenticate_When_Body_Is_Not_A_JSON(t *testing.T) {
	_, handler := controller()

	body := []byte("invalid json")

	req := httptest.NewRequest(http.MethodPost, "/v1/barber/auth", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %d", http.StatusBadRequest, rr.Code)
	}

	var response apperr.AppErr
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("unable to parse response data")
	}

	expectedErr := apperr.NewBadRequestError("invalid json")

	if response != *expectedErr {
		t.Errorf("Expected error %v, got %v", expectedErr.Error(), response.Error())
	}
}

func Test_Should_Not_Be_Able_To_Authenticate_When_Payload_Does_Not_Satisfy_Input_DTO(t *testing.T) {
	_, handler := controller()

	payload := `{"email": "barber1@mail.com", "password": 12934567890}`
	body := []byte(payload)

	req := httptest.NewRequest(http.MethodPost, "/v1/barber/auth", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %d", http.StatusBadRequest, rr.Code)
	}
}
