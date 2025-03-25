package middleware

import (
	"context"
	"net/http"

	customI18n "github.com/akifkadioglu/vocapedia/pkg/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"golang.org/x/time/rate"
	"github.com/sirupsen/logrus"
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
		ctx = context.WithValue(ctx, customI18n.CONTEXT, localizer)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

var limiter = rate.NewLimiter(1, 3)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		next.ServeHTTP(w, r)
	})
}

var log = logrus.New()

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(logrus.Fields{
			"method": r.Method,
			"path":   r.URL.Path,
			"ip":     r.RemoteAddr,
		}).Info("Request received")
		next.ServeHTTP(w, r)
	})
}
