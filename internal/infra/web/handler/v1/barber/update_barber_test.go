package barber_test

import (
	"bytes"
	"context"
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
	"github.com/go-chi/chi/v5"
)

func updateBarberController() (*memory.BarberRepositoryMemory, *barberHandler.UpdateBarberHandler) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewUpdateBarberUseCase(repositoryContainer)
	handler := barberHandler.NewUpdateBarberHandler(sut)

	return barberRepositoryMemory, handler
}

func Test_Should_Be_Able_To_Update_Barber(t *testing.T) {
	repository, handler := updateBarberController()

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"
	barberData := factory.MakeBarber(entity.Barber{ID: barberId})
	repository.Barbers = []entity.Barber{
		barberData,
	}
	payload := barberHandler.UpdateBarberHandlerRequest{
		Name:  "Barber1",
		Email: "barber1@mail.com",
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPut, "/v1/barber/update/"+barberId, bytes.NewBuffer(body))

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", barberId)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %v, got %v", http.StatusOK, rr.Code)
	}

	if repository.Barbers[0].Name != payload.Name {
		t.Errorf("Expected result %v, got %v", payload.Name, repository.Barbers[0].Name)
	}

	if repository.Barbers[0].Email != payload.Email {
		t.Errorf("Expected result %v, got %v", payload.Email, repository.Barbers[0].Email)
	}
}

func Test_Should_Not_Be_Able_To_Update_Barber_When_Barber_Not_Exists(t *testing.T) {
	repository, handler := updateBarberController()

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"
	barberData := factory.MakeBarber(entity.Barber{})
	repository.Barbers = []entity.Barber{
		barberData,
	}
	payload := barberHandler.UpdateBarberHandlerRequest{
		Name:  "Barber1",
		Email: "barber1@mail.com",
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPut, "/v1/barber/update/"+barberId, bytes.NewBuffer(body))

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", barberId)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

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

func Test_Should_Not_Be_Able_To_Update_Barber_When_Http_Method_Is_Wrong(t *testing.T) {
	_, handler := updateBarberController()

	barberData := factory.MakeBarber(entity.Barber{})

	req := httptest.NewRequest(http.MethodGet, "/v1/barber/update/"+barberData.ID, nil)

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %v, got %d", http.StatusMethodNotAllowed, rr.Code)
	}
}

type updateBarberErrorReader struct{}

func (uper *updateBarberErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated error while reading request body")
}

func Test_Should_Not_Be_Able_To_Update_Barber_When_Unable_To_Read_Body_Content(t *testing.T) {
	_, handler := updateBarberController()

	barberData := factory.MakeBarber(entity.Barber{})

	req := httptest.NewRequest(http.MethodPut, "/v1/barber/update/"+barberData.ID, &updateBarberErrorReader{})

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %d", http.StatusBadRequest, rr.Code)
	}
}

func Test_Should_Not_Be_Able_To_Update_Barber_When_Body_Is_Not_A_JSON(t *testing.T) {
	_, handler := updateBarberController()

	barberData := factory.MakeBarber(entity.Barber{})
	body := []byte("invalid json")

	req := httptest.NewRequest(http.MethodPut, "/v1/barber/update/"+barberData.ID, bytes.NewBuffer(body))

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

func Test_Should_Not_Be_Able_To_Update_Barber_When_Payload_Does_Not_Satisfy_Input_DTO(t *testing.T) {
	_, handler := updateBarberController()

	barberData := factory.MakeBarber(entity.Barber{})
	payload := `{"name": 1, "email": "barber1@mail.com", "phone": 12934567890}`
	body := []byte(payload)

	req := httptest.NewRequest(http.MethodPut, "/v1/barber/update/"+barberData.ID, bytes.NewBuffer(body))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status %v, got %d", http.StatusBadRequest, rr.Code)
	}
}
