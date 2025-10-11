package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestServer() *Server {
	s := &Server{}
	return s
}

func performRequest(t *testing.T, r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	t.Helper()
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, body)
	r.ServeHTTP(rr, req)
	return rr
}

func parseJSONResponse[T any](t *testing.T, rr *httptest.ResponseRecorder) T {
	t.Helper()
	var resp T
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	return resp
}
