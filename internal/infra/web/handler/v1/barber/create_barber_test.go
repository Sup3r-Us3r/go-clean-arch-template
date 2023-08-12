package barber_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/apperr"
	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository/memory"
	barberHandler "github.com/Sup3r-Us3r/barber-server/internal/infra/web/handler/v1/barber"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/barber"
	"github.com/Sup3r-Us3r/barber-server/test/factory"
)

func controller() (*memory.BarberRepositoryMemory, *barberHandler.CreateBarberHandler) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewCreateBarberUseCase(repositoryContainer)
	handler := barberHandler.NewCreateBarberHandler(sut)

	return barberRepositoryMemory, handler
}

func Test_Should_Be_Able_To_Create_A_New_Barber(t *testing.T) {
	_, handler := controller()

	barberData := factory.MakeBarber(entity.Barber{})
	payload := barberHandler.CreateBarberHandlerRequest{
		Name:     barberData.Name,
		Email:    barberData.Email,
		Password: barberData.Password,
		Phone:    barberData.Phone,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/v1/barber/create", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status %v, got %d", http.StatusCreated, rr.Code)
	}
}

func Test_Should_Not_Be_Able_To_Create_A_New_Barber_When_Barber_Already_Exists(t *testing.T) {
	repository, handler := controller()

	barberData := factory.MakeBarber(entity.Barber{})
	repository.Barbers = []entity.Barber{
		barberData,
	}
	payload := barberHandler.CreateBarberHandlerRequest{
		Name:     barberData.Name,
		Email:    barberData.Email,
		Password: barberData.Password,
		Phone:    barberData.Phone,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/v1/barber/create", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	var response apperr.AppErr
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("unable to parse response data")
	}

	if response != *apperr.ErrBarberAlreadyExists {
		t.Errorf("Expected error %v, got %v", apperr.ErrBarberAlreadyExists.Error(), response.Error())
	}
}

func Test_Should_Not_Be_Able_To_Create_A_New_Barber_When_Http_Method_Is_Wrong(t *testing.T) {
	_, handler := controller()

	req := httptest.NewRequest(http.MethodGet, "/v1/barber/create", nil)

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %v, got %d", http.StatusMethodNotAllowed, rr.Code)
	}
}

type errorReader struct{}

func (er *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated error while reading request body")
}

func Test_Should_Not_Be_Able_To_Create_A_New_Barber_When_Unable_To_Read_Body_Content(t *testing.T) {
	_, handler := controller()

	req := httptest.NewRequest(http.MethodPost, "/v1/barber/create", &errorReader{})

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %d", http.StatusBadRequest, rr.Code)
	}
}

func Test_Should_Not_Be_Able_To_Create_A_New_Barber_When_Body_Is_Not_A_JSON(t *testing.T) {
	_, handler := controller()

	body := []byte("invalid json")

	req := httptest.NewRequest(http.MethodPost, "/v1/barber/create", bytes.NewBuffer(body))

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

func Test_Should_Not_Be_Able_To_Create_A_New_Barber_When_Payload_Does_Not_Satisfy_Input_DTO(t *testing.T) {
	_, handler := controller()

	payload := `{"name": 1, "email": "barber1@mail.com", "phone": 12934567890}`
	body := []byte(payload)

	req := httptest.NewRequest(http.MethodPost, "/v1/barber/create", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %d", http.StatusBadRequest, rr.Code)
	}
}

func Test_Should_Not_Be_Able_To_Create_A_New_Barber_When_Payload_Is_Empty(t *testing.T) {
	_, handler := controller()

	payload := `{}`
	body := []byte(payload)

	req := httptest.NewRequest(http.MethodPost, "/v1/barber/create", bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %d", http.StatusBadRequest, rr.Code)
	}
}
