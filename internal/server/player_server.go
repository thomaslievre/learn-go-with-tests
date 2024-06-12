package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Player struct {
	name  string
	score int
}

type PlayerService interface {
	// Register(user Player) (insertedID string, err error)
	// GetPlayerScore(player Player) (score string, err error)
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	service PlayerService
}

// func NewPlayerServer(service PlayerService) *PlayerServer {
// 	return &PlayerServer{service: service}
// }

// func (s *PlayerServer) GetPlayerScore(w http.ResponseWriter, req *http.Request) {
// 	defer req.Body.Close()
// 	fmt.Fprintf(w, "20")

// }

// func GetPlayerScore(name string) string {
// 	if name == "Pepper" {
// 		return "20"
// 	}

// 	if name == "Floyd" {
// 		return "10"
// 	}
// 	return ""
// }

func Encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func Decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

func GetPlayerScore(w http.ResponseWriter, req *http.Request) {
	// defer req.Body.Close()
	// param := chi.URLParam(req, "name")
	param := strings.TrimPrefix(req.URL.Path, "/players/")
	var score string
	if param == "Pepper" {
		score = "20"
	}

	if param == "Floyd" {
		score = "10"
	}
	// fmt.Fprintf(w, score)
	Encode(w, req, 200, score)
	fmt.Fprintln(w)
}

// func (pS *PlayerServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	player := strings.TrimPrefix(req.URL.Path, "/players/")
// 	fmt.Fprint(w, GetPlayerScore(player))
// }
