package blogrenderer_test

import (
	"bytes"
	"hello/blogrenderer"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	aPost := blogrenderer.Post{
		Title:       "Hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post intl HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello world"}, {Title: "Hello world 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	aPost := blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	postRenderer, _ := blogrenderer.NewPostRenderer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
