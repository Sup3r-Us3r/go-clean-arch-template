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
	createBarberHandler := barber.NewCreateBarberHandler(useCaseContainer.CreateBarberUseCase)

	webserver.AddHandler("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/doc.json")),
	)

	webserver.AddHandler("/v1/auth/sign-in", signInHandler.Handle)

	webserver.AddHandler("/v1/barber/create", createBarberHandler.Handle, middleware.VerifyToken)
}
