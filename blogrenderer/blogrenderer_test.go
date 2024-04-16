package blogrenderer_test

import (
    "bytes"
    "testing"
    "github.com/shubhadapaithankar/go-with-tests/blogrenderer" 

func TestRenderBlogPost(t *testing.T) {
    // Example blog post data
    post := blogrenderer.Post{
        Title:       "Test Post",
        Description: "This is a test post.",
        Body:        "<p>This is the content of the test post.</p>",
        Tags:        []string{"test", "blog"},
    }

    // Create a buffer to write our output to
    var buf bytes.Buffer

    // Attempt to render the post
    err := blogrenderer.Render(&buf, post)
    if err != nil {
        t.Errorf("Render failed: %v", err)
    }

   
    if !bytes.Contains(buf.Bytes(), []byte(post.Title)) {
        t.Errorf("Rendered content does not contain the post title '%s'", post.Title)
    }
}
