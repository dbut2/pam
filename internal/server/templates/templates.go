package templates

import (
	_ "embed"
	"html/template"
)

//go:embed index.tmpl.html
var index string

//go:embed entry.tmpl.html
var entry string

//go:embed entries.tmpl.html
var entries string

func GetTemplate() (*template.Template, error) {
	t := template.New("Template")

	t, err := t.Parse(index)
	if err != nil {
		return nil, err
	}

	t, err = t.Parse(entry)
	if err != nil {
		return nil, err
	}

	t, err = t.Parse(entries)
	if err != nil {
		return nil, err
	}

	return t, nil
}
