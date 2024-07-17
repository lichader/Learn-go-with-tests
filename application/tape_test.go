package poker_test

import (
	"io"
	"testing"

	poker "github.com/lichader/learn-go-with-test-application"
)

func TestTape_Write(t *testing.T) {
	file, clean := poker.CreateTempFile(t, "12345")
	defer clean()

	tape := &poker.Tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
