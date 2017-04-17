package gotm

import (
	"github.com/chai2010/gettext-go/gettext"
)

// https://godoc.org/github.com/chai2010/gettext-go/gettext#BindTextdomain
func BindTextdomain(domain, path string, zipData []byte) (domains, paths []string) {
	return gettext.BindTextdomain(domain, path, zipData)
}

// https://godoc.org/github.com/chai2010/gettext-go/gettext#SetLocale
func SetLocale(locale string) string {
	return gettext.SetLocale(locale)
}

// https://godoc.org/github.com/chai2010/gettext-go/gettext#Textdomain
func Textdomain(domain string) string {
	return gettext.Textdomain(domain)
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
	gettext.BindTextdomain("messages", dirpath, nil)
	Textdomain("messages")
}
