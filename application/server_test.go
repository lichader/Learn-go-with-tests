package poker_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	poker "github.com/lichader/learn-go-with-test-application"
)

func TestGETPlayers(t *testing.T) {
	store := poker.StubPlayerStore{
		Scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := mustMakePlayerServer(t, &store)

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
	server := mustMakePlayerServer(t, &store)

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
		server, _ := poker.NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertContentType(t, response, "application/json")

		got := getLeagueResponse(t, response.Body)
		poker.AssertLeague(t, got, wantedLegua)
	})
}

func TestGame(t *testing.T) {
	t.Run("Get /game returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, &poker.StubPlayerStore{})
		request, _ := http.NewRequest(http.MethodGet, "/game", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("when we get a message over a websocket it is a winner of a game", func(t *testing.T) {
		store := &poker.StubPlayerStore{}
		winner := "Ruth"

		pokerServer := mustMakePlayerServer(t, store)
		server := httptest.NewServer(pokerServer)
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
		if err != nil {
			t.Fatalf("could not open a ws connection on %s %v", wsURL, err)
		}
		defer ws.Close()

		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		poker.AssertPlayerWin(t, store, winner)
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

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore) *poker.PlayerServer {
	t.Helper()
	server, err := poker.NewPlayerServer(store)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}
