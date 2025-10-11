package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil { //nolint:staticcheck
		// TODO: handle returned error here.
	}
}

func (s *Server) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(LoggerContextMiddleware(slog.Default()))
	r.Use(LoggingMiddleware)

	r.Method(http.MethodGet, "/v1/health", Handler(s.GetHealthHandler))
	return r
}

func NewServer() *Server {
	s := &Server{
		httpServer: &http.Server{
			Addr: ":53430",

			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       120 * time.Second,

			MaxHeaderBytes: 1 << 20, // 1MB
		},
	}

	s.httpServer.Handler = s.Routes()
	return s
}
