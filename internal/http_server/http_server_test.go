package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}

	server := NewPlayerServer(&store)
	t.Run("return Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		AssertResponseBody(t, response.Body.String(), "20")
		AssertResponseStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		AssertResponseBody(t, response.Body.String(), "10")
		AssertResponseStatusCode(t, response.Code, http.StatusOK)
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Stevie")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		AssertResponseStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}

	server := NewPlayerServer(&store)
	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "stevie"
		request := newPostScoreRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		AssertResponseStatusCode(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("got %d want %d", len(store.scores), 1)
		}
		if store.winCalls[0] != player {
			t.Errorf("got %s want %s", store.winCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns 200 on /league", func(t *testing.T) {
		wantedLeague := []Player{
			{"amberly", 30},
			{"grace", 24},
			{"stevie", 23},
		}
		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := getLeagueResponse(t, response.Body)
		AssertResponseStatusCode(t, response.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
		AssertContentType(t, response, jsonContentType)
	})
}

func newPostScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func getLeagueResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, %v", body, err)
	}
	return league
}
