package blogposts_test

import (
	"errors"
	"hello/blogposts"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFs(fs)
	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}
	assertPost(t, got, want)
}

func TestNewBlogPosts_WhenReadingDirFails_ReturnsError(t *testing.T) {
	fs := StubFailingFS{}

	_, err := blogposts.NewPostsFromFs(fs)
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
