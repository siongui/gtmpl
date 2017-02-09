package gossg

import (
	"os"
	"testing"
)

type TemplateData struct {
	Title string
}

func TestTemplateToHtml(t *testing.T) {
	data := TemplateData{
		Title: "Test Page",
	}

	tmpl, err := ParseDirectory("theme/template")
	if err != nil {
		t.Error(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "index.html", &data)
	if err != nil {
		t.Error(err)
	}
}
