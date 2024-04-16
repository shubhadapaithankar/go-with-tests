package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

var postTemplates embed.FS

func Render(w io.Writer, p Post) error {
	tmpl, err := template.ParseFS(postTemplates, "templates/blog.gohtml")
	if err != nil {
		return err
	}
	return tmpl.Execute(w, p)
}

func (p Post) SanitizeTitle() string {
	return strings.Replace(strings.ToLower(p.Title), " ", "-", -1)
}
