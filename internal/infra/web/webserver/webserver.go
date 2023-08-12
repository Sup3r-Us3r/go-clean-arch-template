package webserver

import (
	"net/http"

	"github.com/Sup3r-Us3r/barber-server/internal/infra/web/middleware"
	"github.com/Sup3r-Us3r/barber-server/log"
	"github.com/go-chi/chi/v5"

	"github.com/go-chi/cors"
)

type Handler struct {
	Handler     http.HandlerFunc
	Middlewares []func(http.Handler) http.Handler
}

type WebServer struct {
	Router        chi.Router
	WebServerPort string
	Handlers      map[string]Handler
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		WebServerPort: webServerPort,
		Handlers:      make(map[string]Handler),
	}
}

func (ws *WebServer) AddHandler(
	path string,
	handler http.HandlerFunc,
	middlewares ...func(http.Handler) http.Handler,
) {
	ws.Handlers[path] = Handler{
		Handler:     handler,
		Middlewares: middlewares,
	}
}

func (ws WebServer) GlobalMiddlewares() {
	ws.Router.Use(middleware.LoggerMiddleware)
	ws.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}

func (ws WebServer) StartServer() {
	ws.GlobalMiddlewares()

	for path, handler := range ws.Handlers {
		if len(handler.Middlewares) > 0 {
			ws.Router.With(handler.Middlewares...).Handle(path, handler.Handler)
		} else {
			ws.Router.Handle(path, handler.Handler)
		}
	}

	if err := http.ListenAndServe(ws.WebServerPort, ws.Router); err != nil {
		log.Error(err.Error())
		panic(err)
	}
}
