package gotm

import (
	"github.com/chai2010/gettext-go"
)

// SetLanguage sets the locale to be used to render the template.
//
// see https://godoc.org/github.com/chai2010/gettext-go#SetLanguage for more
// details.
func SetLanguage(locale string) string {
	return gettext.SetLanguage(locale)
}

// Translate returns gettext.PGettext("", msgid)
//
// see https://godoc.org/github.com/chai2010/gettext-go#PGettext for more
// details.
func Translate(msgid string) string {
	return gettext.PGettext("", msgid)
}

func setupMessagesDomain(domain, dir string) {
	gettext.BindLocale(gettext.New(domain, dir))
	gettext.SetDomain(domain)
}
