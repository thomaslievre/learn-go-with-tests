package api

type MockPlayerService struct {
}

// func TestGetPlayers(m *testing.M) {
// 	var err error

// 	service := &MockPlayerService{
// 		// RegisterFunc: func(user User) (string, error) {
// 		// 	return expectedInsertedID, nil
// 		// },
// 	}

// 	server := api.NewPlayerServer(service)

// 	req := httptest.NewRequest(http.MethodGet, "/players/", userToJSON(user))
// 	res := httptest.NewRecorder()

// 	// request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
// 	// response := httptest.NewRecorder()

// 	// PlayerServer(response, request)

// 	// got := response.Body.String()
// 	// want := "20"

// 	// if got != want {
// 	// 	t.Errorf("got %q, want %q", got, want)
// 	// }

// 	os.Exit((m.Run()))
// }
