package server

import (
	"encoding/json"
	"net/http"
)

type GetHealthResponse struct {
	Status string `json:"status"`
}

func (s *Server) GetHealthHandler(w http.ResponseWriter, r *http.Request) error {
	resp := GetHealthResponse{
		Status: "healthy",
	}
	return json.NewEncoder(w).Encode(resp)
}
