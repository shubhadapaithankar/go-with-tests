package blogrenderer_test

import (
	"bytes"
	"testing"

	"github.com/shubhadapaithankar/go-with-tests/Templating/blogrenderer"
)

func TestRender(t *testing.T) {
	renderer, err := blogrenderer.NewRenderer()
	if err != nil {
		t.Fatal(err)
	}

	post := blogrenderer.Post{
		Title:       "Hello World",
		Description: "This is a post",
		Body:        "# Markdown Title\nThis is the body in Markdown.",
		Tags:        []string{"test", "blog"},
	}

	var buf bytes.Buffer
	err = renderer.Render(&buf, post)
	if err != nil {
		t.Fatal(err)
	}

	want := `<h1>Hello World</h1><p>This is a description</p><div>Markdown Title<p>This is the body in Markdown.</p></div><ul><li>test</li><li>blog</li></ul>`
	got := buf.String()
	if got != want {
		t.Errorf("Rendered HTML was incorrect, got: %s, want: %s", got, want)
	}
}
