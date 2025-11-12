package httpserver

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndResult(t *testing.T) {
	database, cleanDB := createTempFile(t, `[]`)
	defer cleanDB()

	store, err := NewFileSystemPlayerStore(database)
	if err != nil {
		log.Fatalf("problem creating system player store %v", err)
	}
	server := NewPlayerServer(store)

	player := "stevie"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		AssertResponseStatusCode(t, response.Code, http.StatusOK)
		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		AssertResponseStatusCode(t, response.Code, http.StatusOK)
		got := getLeagueResponse(t, response.Body)
		want := []Player{
			{"stevie", 3},
		}
		AssertLeague(t, got, want)
	})
}
