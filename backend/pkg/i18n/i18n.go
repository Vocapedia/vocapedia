package i18n

import (
	"log"
	"net/http"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

const (
	LOCALIZER = "localizer"
)

func InitI18n() {
	bundle = i18n.NewBundle(language.Turkish)

	bundle.LoadMessageFile("pkg/i18n/tr.json")
	bundle.LoadMessageFile("pkg/i18n/en.json")
	log.Println("i18n is ready")
}

func Bundle() *i18n.Bundle {
	return bundle
}

func Localizer(r *http.Request, key string) string {
	localizer := r.Context().Value(LOCALIZER).(*i18n.Localizer)
	val, err := localizer.Localize(&i18n.LocalizeConfig{MessageID: key})
	if err != nil {
		localizer = i18n.NewLocalizer(bundle, language.Turkish.String())
		return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: key})
	} else {
		return val
	}
}
