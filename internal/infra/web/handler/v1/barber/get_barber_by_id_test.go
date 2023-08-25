package barber_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Sup3r-Us3r/barber-server/internal/domain/entity"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/repository/memory"
	barberHandler "github.com/Sup3r-Us3r/barber-server/internal/infra/web/handler/v1/barber"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase/barber"
	"github.com/Sup3r-Us3r/barber-server/test/factory"
	"github.com/go-chi/chi/v5"
)

func getBarberByIdController() (*memory.BarberRepositoryMemory, *barberHandler.GetBarberByIdHandler) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewGetBarberByIdUseCase(repositoryContainer)
	handler := barberHandler.NewGetBarberByIdHandler(sut)

	return barberRepositoryMemory, handler
}

func Test_Should_Be_Able_To_Get_Barber_By_Id(t *testing.T) {
	repository, handler := getBarberByIdController()

	barberData := factory.MakeBarber(entity.Barber{})
	repository.Barbers = []entity.Barber{
		barberData,
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/barber/get/"+barberData.ID, nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", barberData.ID)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %v, got %d", http.StatusOK, rr.Code)
	}

	var response barberHandler.GetBarberByIdHandlerResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("unable to parse response data")
	}

	if response.ID != barberData.ID {
		t.Errorf("Expected result %v, got %v", barberData, response)
	}
}

func Test_Should_Not_Be_Able_To_Get_Barber_By_Id_When_Barber_Not_Exists(t *testing.T) {
	repository, handler := getBarberByIdController()

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"
	barberData := factory.MakeBarber(entity.Barber{})
	repository.Barbers = []entity.Barber{
		barberData,
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/barber/get/"+barberId, nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", barberId)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status %v, got %d", http.StatusNotFound, rr.Code)
	}
}

func Test_Should_Not_Be_Able_To_Get_Barber_By_Id_When_Http_Method_Is_Wrong(t *testing.T) {
	_, handler := getBarberByIdController()

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"

	req := httptest.NewRequest(http.MethodPost, "/v1/barber/get/"+barberId, nil)

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %v, got %d", http.StatusMethodNotAllowed, rr.Code)
	}
}
