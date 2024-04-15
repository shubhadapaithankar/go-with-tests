package blogrenderer

import (
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// Embed the templates
//
//go:embed "templates/*"
var postTemplates embed.FS

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

// Renderer holds the templates and can render posts
type Renderer struct {
	templ *template.Template
}

// NewRenderer creates a new Renderer
func NewRenderer() (*Renderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &Renderer{templ: templ}, nil
}

// Render writes the post rendered as HTML to the given writer
func (r *Renderer) Render(w io.Writer, p Post) error {
	mdParser := parser.NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs)
	html := markdown.ToHTML([]byte(p.Body), mdParser, nil)
	data := struct {
		Title, Description string
		Tags               []string
		HTMLBody           template.HTML
	}{
		Title:       p.Title,
		Description: p.Description,
		Tags:        p.Tags,
		HTMLBody:    template.HTML(html),
	}
	return r.templ.ExecuteTemplate(w, "blog.gohtml", data)
}
