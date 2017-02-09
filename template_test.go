package gossg

import (
	"bytes"
	"testing"
)

var html = `<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <title>Test Page</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>

<div>Hello World!</div>

<footer>Powered by
  <a href="https://golang.org/">Go</a>
</footer>
</body>
</html>
`

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

	var b bytes.Buffer
	err = tmpl.ExecuteTemplate(&b, "index.html", &data)
	if err != nil {
		t.Error(err)
	}
	if b.String() != html {
		t.Error("bad html")
	}
}
