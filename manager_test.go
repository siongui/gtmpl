package gotm

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

var htmlZhtw = `<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <title>Test Page</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
</head>
<body>

<div>Hello 巴利字典!</div>

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

	tm := NewTemplateManager("")
	err := tm.ParseDirectory("theme/template")
	if err != nil {
		t.Error(err)
		return
	}

	var b bytes.Buffer
	err = tm.ExecuteTemplate(&b, "index.html", &data)
	if err != nil {
		t.Error(err)
		return
	}
	if b.String() != html {
		t.Error("bad html")
	}
}

func TestI18nTemplateToHtml(t *testing.T) {
	SetupMessagesDomain("locale/")
	data := TemplateData{
		Title: "Test Page",
	}

	tm := NewTemplateManager("")
	err := tm.ParseDirectoryWithGettextFunction("theme/template-i18n")
	if err != nil {
		t.Error(err)
		return
	}

	SetLocale("zh_TW")
	var b bytes.Buffer
	err = tm.ExecuteTemplate(&b, "index.html", &data)
	if err != nil {
		t.Error(err)
		return
	}
	if b.String() != htmlZhtw {
		t.Error("bad zh_TW html")
	}
}
