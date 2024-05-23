package racer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run(
		"compares speeds of servers, returnning the url of the fastest one",
		func(t *testing.T) {
			slowServer := mockDelayedServer(20 * time.Millisecond)
			fastServer := mockDelayedServer(0 * time.Millisecond)

			defer slowServer.Close()
			defer fastServer.Close()

			slowURL := slowServer.URL
			fastURL := fastServer.URL

			want := fastURL
			got, _ := CongiruableRacer(slowURL, fastURL, 1*time.Second)

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}

			slowServer.Close()
			fastServer.Close()
		},
	)

	t.Run("returns an error if a server doesn't respond within 10 seconds", func(t *testing.T) {
		serverA := mockDelayedServer(2 * time.Second)
		serverB := mockDelayedServer(3 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := CongiruableRacer(serverA.URL, serverB.URL, 1*time.Second)

		if err == nil {
			t.Error("Expected an error but didn't get one")
		}
	})
}

func mockDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calling it from mock server")
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
