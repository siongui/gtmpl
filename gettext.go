package gotm

import (
	"github.com/chai2010/gettext-go"
)

// https://pkg.go.dev/github.com/chai2010/gettext-go?tab=doc#BindLocale
func BindTextdomain(domain, path string, zipData []byte) {
	gettext.BindLocale(gettext.New(domain, path))
}

// https://pkg.go.dev/github.com/chai2010/gettext-go?tab=doc#SetLanguage
func SetLocale(locale string) string {
	return gettext.SetLanguage(locale)
}

// https://pkg.go.dev/github.com/chai2010/gettext-go?tab=doc#SetDomain
func Textdomain(domain string) string {
	return gettext.SetDomain(domain)
}

// return gettext.PGettext("", msgid)
func Translate(msgid string) string {
	return gettext.PGettext("", msgid)
}

/*

Shorthand for

	BindTextdomain("messages", dirpath, nil)
	Textdomain("messages")

*/
func SetupMessagesDomain(dirpath string) {
	gettext.BindLocale(gettext.New("messages", dirpath))
	Textdomain("messages")
}
