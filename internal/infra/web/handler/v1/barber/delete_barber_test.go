package barber_test

import (
	"context"
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

func deleteBarberController() (*memory.BarberRepositoryMemory, *barberHandler.DeleteBarberHandler) {
	barberRepositoryMemory := memory.NewBarberRepositoryMemory()
	repositoryContainer := repository.RepositoryContainer{
		BarberRepository: barberRepositoryMemory,
	}
	sut := barber.NewDeleteBarberUseCase(repositoryContainer)
	handler := barberHandler.NewDeleteBarberHandler(sut)

	return barberRepositoryMemory, handler
}

func Test_Should_Be_Able_To_Delete_Barber(t *testing.T) {
	repository, handler := deleteBarberController()

	barberData := factory.MakeBarber(entity.Barber{})
	repository.Barbers = []entity.Barber{
		barberData,
	}

	req := httptest.NewRequest(http.MethodDelete, "/v1/barber/delete/"+barberData.ID, nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", barberData.ID)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("Expected status %v, got %d", http.StatusNoContent, rr.Code)
	}
}

func Test_Should_Not_Be_Able_To_Delete_Barber_When_Barber_Not_Exists(t *testing.T) {
	repository, handler := deleteBarberController()

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"
	barberData := factory.MakeBarber(entity.Barber{})
	repository.Barbers = []entity.Barber{
		barberData,
	}

	req := httptest.NewRequest(http.MethodDelete, "/v1/barber/delete/"+barberId, nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", barberId)

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status %v, got %d", http.StatusNotFound, rr.Code)
	}
}

func Test_Should_Not_Be_Able_To_Delete_Barber_When_Http_Method_Is_Wrong(t *testing.T) {
	_, handler := deleteBarberController()

	barberId := "bd09785b-f68a-46bc-800e-51a676804203"

	req := httptest.NewRequest(http.MethodPost, "/v1/barber/delete/"+barberId, nil)

	rr := httptest.NewRecorder()
	handler.Handle(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status %v, got %d", http.StatusMethodNotAllowed, rr.Code)
	}
}
