package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndResult(t *testing.T) {
	store := NewInmemoryPlayerStore()
	server := PlayerServer{store}
	player := "stevie"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))

	assertResponseStatusCode(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
