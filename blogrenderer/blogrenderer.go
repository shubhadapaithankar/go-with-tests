package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"
)

// Post represents a blog post with metadata and content.
type Post struct {
	Title, Description, Body string
	Tags                     []string
}

//go:embed "templates/*"
var postTemplates embed.FS

// Render renders a single post to the provided writer.
func Render(w io.Writer, p Post) error {
	tmpl, err := template.ParseFS(postTemplates, "templates/blog.gohtml")
	if err != nil {
		return err
	}
	return tmpl.Execute(w, p)
}

// SanitizeTitle replaces spaces with hyphens and converts the title to lowercase.
func (p Post) SanitizeTitle() string {
	return strings.Replace(strings.ToLower(p.Title), " ", "-", -1)
}
