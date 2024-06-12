package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Print("Executing middlewareOne again")
	})
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello tout le monde titi\n")
}

func (s *Server) RegisterRoutes(
	mux *chi.Mux,
	// logger *logging.Logger,
	// config Config,
	// tenantsStore        *TenantsStore,
	// commentsStore       *CommentsStore,
	// conversationService *ConversationService,
	// chatGPTService      *ChatGPTService,
	// authProxy           *authProxy
) {
	// mux.Handle("/api/v1/", handleTenantsGet(logger, tenantsStore))
	// mux.Handle("/oauth2/", handleOAuth2Proxy(logger, authProxy))
	// mux.HandleFunc("/healthz", handleHealthzPlease(logger))
	// mux.Handle("/", http.NotFoundHandler())
	// mux.HandleFunc("/hello", hello)

	mux.Get("/players/{name}", GetPlayerScore)

	// log.Fatal(http.ListenAndServe(":5000", handler))

}
