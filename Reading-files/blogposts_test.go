package blogposts_test

import (
	"testing"
	"testing/fstest"

	"github.com/yourusername/yourproject/blogposts"
)

func TestNewPostsFromFS(t *testing.T) {
	fs := fstest.MapFS{
		"hello_world.md": &fstest.MapFile{
			Data: []byte(`
Title: Hello, TDD world!
Description: First post on our wonderful blog
Tags: tdd, go
---
Hello world!`),
		},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != 1 {
		t.Errorf("got %d posts, wanted %d posts", len(posts), 1)
	}

	got := posts[0]
	want := blogposts.Post{
		Title:       "Hello, TDD world!",
		Description: "First post on our wonderful blog",
		Tags:        []string{"tdd", "go"},
		Body:        "Hello world!",
	}

	if got.Title != want.Title || got.Description != want.Description || !equal(got.Tags, want.Tags) || got.Body != want.Body {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
