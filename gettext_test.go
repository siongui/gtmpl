package gtmpl

import (
	"testing"
)

func TestGettext(t *testing.T) {
	setupMessagesDomain("messages", "locale/")

	SetLanguage("zh_TW")
	if Translate("Pāḷi Dictionary") != "巴利字典" {
		t.Error("translate Pāḷi Dictionary (zh_TW) fail")
	}

	SetLanguage("vi_VN")
	if Translate("Pāḷi Dictionary") != "Từ điển Pāḷi" {
		t.Error("translate Pāḷi Dictionary (vi_VN) fail")
	}

	SetLanguage("fr_FR")
	if Translate("Pāḷi Dictionary") != "Dictionnaire Pāḷi" {
		t.Error("translate Pāḷi Dictionary (fr_FR) fail")
	}
}
