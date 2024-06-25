package maps

import "testing"

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "this is just a test"}

	t.Run("Key exists", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Key does not exist", func(t *testing.T) {
		_, err := dic.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		if err != ErrNotFound {
			t.Errorf("got %q want %q", err, ErrNotFound)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add a new key", func(t *testing.T) {
		dic := Dictionary{}
		key := "test"
		value := "this is just a test"
		dic.Add(key, value)

		got, _ := dic.Search(key)

		if got != value {
			t.Errorf("got %q want %q", got, value)
		}
	})

	t.Run("Add an existing key", func(t *testing.T) {
		dic := Dictionary{"test": "this is just a test"}
		key := "test"
		value := "this should replace the existing value"
		dic.Add(key, value)

		got, _ := dic.Search(key)

		if got != value {
			t.Errorf("got %q want %q", got, value)
		}
	})
}
