package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

type LoggerContextKeyType string

var loggerContextKey LoggerContextKeyType = "logger"

func LoggerContextMiddleware(baseLogger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqLogger := baseLogger.With("path", r.URL.Path, "request_id", middleware.GetReqID(r.Context()))
			ctx := context.WithValue(r.Context(), loggerContextKey, reqLogger)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerContextKey).(*slog.Logger); ok {
		return logger
	}

	// provide default fallback
	return slog.Default()
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := LoggerFromContext(r.Context())
		logger.Info("Incoming request")
		next.ServeHTTP(w, r)
	})
}
