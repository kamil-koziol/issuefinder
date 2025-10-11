package server

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetHealthHandler(t *testing.T) {
	s := newTestServer()
	router := s.Routes()

	rr := performRequest(t, router, http.MethodGet, "/v1/health", nil)
	require.Equal(t, http.StatusOK, rr.Code)

	resp := parseJSONResponse[GetHealthResponse](t, rr)
	require.Equal(t, "healthy", resp.Status)
}
