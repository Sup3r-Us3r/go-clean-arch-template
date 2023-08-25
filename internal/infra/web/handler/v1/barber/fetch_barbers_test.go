package barber_test

import (
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
)

func fetchBarbersController() (*memory.BarberRepositoryMemory, *barberHandler.FetchBarbersHandler) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewFetchBarbersUseCase(repositoryContainer)
	handler := barberHandler.NewFetchBarbersHandler(sut)

	return barberRepositoryMemory, handler
}

func Test_Should_Be_Able_To_Fetch_Barbers(t *testing.T) {
	repository, handler := fetchBarbersController()

	barberData := factory.MakeBarber(entity.Barber{})
	repository.Barbers = []entity.Barber{
		barberData,
	}

	req := httptest.NewRequest(http.MethodGet, "/v1/barber/fetch", nil)

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %v, got %d", http.StatusOK, rr.Code)
	}

	var response []entity.Barber
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal("unable to parse response data")
	}

	if len(response) != 1 {
		t.Errorf("Expected result %v, got %v", 1, len(response))
	}
}

func Test_Should_Not_Be_Able_To_Fetch_Barbers_When_Http_Method_Is_Wrong(t *testing.T) {
	_, handler := fetchBarbersController()

	req := httptest.NewRequest(http.MethodPost, "/v1/barber/fetch", nil)

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %v, got %d", http.StatusMethodNotAllowed, rr.Code)
	}
}
