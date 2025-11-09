package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndResult(t *testing.T) {
	store := NewInmemoryPlayerStore()
	server := NewPlayerServer(store)

	player := "stevie"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertResponseStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		assertResponseStatusCode(t, response.Code, http.StatusOK)
		got := getLeagueResponse(t, response.Body)
		want := []Player{
			{"stevie", 3},
		}
		assertLeague(t, got, want)
	})
}
