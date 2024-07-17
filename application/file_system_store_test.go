package poker_test

import (
	"testing"

	poker "github.com/lichader/learn-go-with-test-application"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		got := store.GetLeague()
		want := []poker.Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		poker.AssertLeague(t, got, want)

		got = store.GetLeague()
		poker.AssertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}
        ]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}
        ]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)
		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		poker.AssertScoreEqual(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}
        ]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)
		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		poker.AssertScoreEqual(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := poker.CreateTempFile(t, "")
		defer cleanDatabase()

		_, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)
	})
}
