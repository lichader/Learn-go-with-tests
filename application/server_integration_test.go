package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/lichader/learn-go-with-test-application"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := poker.CreateTempFile(t, "[]")
	defer cleanDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)
	poker.AssertNoError(t, err)
	server, _ := poker.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostScoreRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertContentType(t, response, "application/json")
		got := getLeagueResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, got, want)
	})
}
