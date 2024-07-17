package poker_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/lichader/learn-go-with-test-application"
)

func TestGETPlayers(t *testing.T) {
	store := poker.StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := poker.NewPlayerServer(&store)

	t.Run("return Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertResponseBody(t, response.Body.String(), "20")
		poker.AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertResponseBody(t, response.Body.String(), "10")
		poker.AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStorePlayers(t *testing.T) {
	store := poker.StubPlayerStore{
		Scores: map[string]int{},
	}
	server := poker.NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostScoreRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusAccepted)

		if len(store.WinCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.WinCalls), 1)
		}

		if store.WinCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.WinCalls[0], player)
		}
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns 200 on /league", func(t *testing.T) {
		wantedLegua := []poker.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := poker.StubPlayerStore{nil, nil, wantedLegua}
		server := poker.NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertContentType(t, response, "application/json")

		got := getLeagueResponse(t, response.Body)
		poker.AssertLeague(t, got, wantedLegua)
	})
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func getLeagueResponse(t testing.TB, body io.Reader) (league []poker.Player) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf(
			"Unable to parse response from server %q into slice of player, '%v'",
			body,
			err,
		)
	}

	return
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}
