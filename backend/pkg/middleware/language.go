package middleware

import (
	"context"
	"net/http"

	customI18n "github.com/akifkadioglu/vocapedia/pkg/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func Language(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := r.FormValue("lang")
		if lang == "" {
			lang = r.Header.Get("Accept-Language")
		}
		if lang == "" {
			lang = language.Turkish.String()
		}

		localizer := i18n.NewLocalizer(customI18n.Bundle(), lang)

		ctx := r.Context()
		ctx = context.WithValue(ctx, customI18n.LOCALIZER, localizer)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
