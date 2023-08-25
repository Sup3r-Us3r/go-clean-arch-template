package handler

import (
	"github.com/Sup3r-Us3r/barber-server/internal/infra/web/handler/v1/auth"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/web/handler/v1/barber"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/web/middleware"
	"github.com/Sup3r-Us3r/barber-server/internal/infra/web/webserver"
	"github.com/Sup3r-Us3r/barber-server/internal/usecase"
	httpSwagger "github.com/swaggo/http-swagger"
)

func NewHandlerContainer(useCaseContainer usecase.UseCaseContainer, webserver webserver.WebServer) {
	signInHandler := auth.NewSignInHandler(useCaseContainer.SignInUseCase)
	getBarberById := barber.NewGetBarberByIdHandler(useCaseContainer.GetBarberByIdUseCase)
	fetchBarbersHandler := barber.NewFetchBarbersHandler(useCaseContainer.FetchBarbersUseCase)
	createBarberHandler := barber.NewCreateBarberHandler(useCaseContainer.CreateBarberUseCase)
	updateBarberHandler := barber.NewUpdateBarberHandler(useCaseContainer.UpdateBarberUseCase)
	deleteBarberHandler := barber.NewDeleteBarberHandler(useCaseContainer.DeleteBarberUseCase)

	webserver.AddHandler("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json")),
	)

	webserver.AddHandler("/v1/auth/sign-in", signInHandler.Handle)

	webserver.AddHandler("/v1/barber/fetch", fetchBarbersHandler.Handle, middleware.VerifyToken)
	webserver.AddHandler("/v1/barber/get/{id}", getBarberById.Handle)
	webserver.AddHandler("/v1/barber/create", createBarberHandler.Handle)
	webserver.AddHandler("/v1/barber/update/{id}", updateBarberHandler.Handle)
	webserver.AddHandler("/v1/barber/delete/{id}", deleteBarberHandler.Handle)
}
