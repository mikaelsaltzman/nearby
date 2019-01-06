package server

import (
	"encoding/json"
	"local/nearby/places"
	"log"
	"net/http"
)

// Server is the Server struct.
type Server struct {
	Logger       *log.Logger
	PlacesClient *places.Client
}

// NewServer creates a new Server struct.
func NewServer(logger *log.Logger, c *places.Client) *Server {
	return &Server{
		Logger:       logger,
		PlacesClient: c,
	}
}

// Nearby is the nearby endpoint handler method.
func (s *Server) Nearby(w http.ResponseWriter, r *http.Request) {
	var k, l, p string
	keys := r.URL.Query()
	if k = keys.Get("key"); k == "" {
		http.Error(w, "URL parameter 'key' missing", http.StatusUnauthorized)
	}
	if l = keys.Get("location"); l == "" {
		l = "59.326165362,18.058666432"
	}
	if p = keys.Get("type"); p == "" {
		p = "bicycle_store"
	}

	params := places.Params{
		Key:       k,
		Location:  l,
		PlaceType: p,
	}

	res, err := s.PlacesClient.Nearby(params)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
	}

	results, err := json.Marshal(res.Results)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(results)
}
