package gotm

import (
	"testing"
)

func TestGettext(t *testing.T) {
	SetupMessagesDomain("locale/")

	SetLocale("zh_TW")
	if Translate("Pāḷi Dictionary") != "巴利字典" {
		t.Error("translate Pāḷi Dictionary (zh_TW) fail")
	}

	SetLocale("vi_VN")
	if Translate("Pāḷi Dictionary") != "Từ điển Pāḷi" {
		t.Error("translate Pāḷi Dictionary (vi_VN) fail")
	}

	SetLocale("fr_FR")
	if Translate("Pāḷi Dictionary") != "Dictionnaire Pāḷi" {
		t.Error("translate Pāḷi Dictionary (fr_FR) fail")
	}
}
