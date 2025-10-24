package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/kamil-koziol/issuefinder/api/internal/config"
	"github.com/kamil-koziol/issuefinder/api/internal/store"
)

type Server struct {
	config     config.Config
	httpServer *http.Server
	db         *pgx.Conn
	querier    store.Querier
}

func (s *Server) Run() error {
	db, err := pgx.Connect(context.Background(), s.config.PostgreSQLURL.String())
	if err != nil {
		return fmt.Errorf("unable to connect to db: %w", err)
	}
	defer db.Close(context.Background()) //nolint:errcheck

	s.db = db
	s.querier = store.New(s.db)

	s.httpServer.Handler = s.Routes()

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
	r.Use(middleware.RealIP)
	r.Use(LoggerContextMiddleware(slog.Default()))
	r.Use(LoggingMiddleware)

	r.Method(http.MethodGet, "/v1/health", Handler(s.GetHealthHandler))
	return r
}

func NewServer(cfg config.Config) (*Server, error) {
	s := &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf(":%d", cfg.Port),

			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       120 * time.Second,

			MaxHeaderBytes: 1 << 20, // 1MB
		},
		config: cfg,
	}

	return s, nil
}
