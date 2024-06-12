package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

var port = 8080

type Server struct {
	port int
	db   *sql.DB
}

func NewServer() *http.Server {
	NewServer := &Server{
		port: port,
		db:   nil,
	}

	mux := chi.NewRouter()

	// mux := http.NewServeMux()
	NewServer.RegisterRoutes(mux)

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", 8080),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
