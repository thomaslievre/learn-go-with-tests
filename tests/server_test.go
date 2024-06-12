package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/thomaslievre/learn-go-with-tests/internal/server"
)

func TestGETPlayers(t *testing.T) {
	server := api.NewServer()
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := executeRequest(server, request)

		var v string

		err := json.NewDecoder(response.Body).Decode(&v)

		if err != nil {
			t.Errorf("error decode")
		}

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, v, "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := executeRequest(server, request)

		var v string

		err := json.NewDecoder(response.Body).Decode(&v)

		if err != nil {
			t.Errorf("error decode")
		}

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, v, "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func executeRequest(server *http.Server, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	server.Handler.ServeHTTP(rr, req)

	return rr
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
