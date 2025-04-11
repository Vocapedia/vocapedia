package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/akifkadioglu/vocapedia/pkg/cache"
	"github.com/akifkadioglu/vocapedia/pkg/database"
	"github.com/akifkadioglu/vocapedia/pkg/entities"
	customI18n "github.com/akifkadioglu/vocapedia/pkg/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
	"golang.org/x/time/rate"
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

var limiter = rate.NewLimiter(1, 6)

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

func HandleToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := database.Manager()
		token := strings.Split(r.Header.Get("Authorization"), " ")[1]
		var tokenEntity entities.Token
		err := db.Where("token = ?", token).First(&tokenEntity).Error
		if err != nil {
			log.Println("err", err)
			http.Error(w, "Token is not valid", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})

}

func HandleVocatoken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rds := cache.Redis()
		db := database.Manager()
		token := r.Header.Get("Vocatoken")
		if token == "" {
			http.Error(w, "Missing vocatoken", http.StatusUnauthorized)
			return
		}

		var userEntity entities.User
		err := db.Where("vocatoken = ?", token).First(&userEntity).Error
		if err != nil {
			log.Println("err:", err)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		email := userEntity.Email
		redisKey := fmt.Sprintf("voca:usage:%s", email)

		val, err := rds.Get(context.Background(), redisKey).Int()
		if err == redis.Nil {
			err := rds.Set(context.Background(), redisKey, 29, 24*time.Hour).Err()
			if err != nil {
				log.Println("Redis set error:", err)
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}
		} else if err != nil {
			log.Println("Redis get error:", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		} else {
			if val <= 0 {
				http.Error(w, customI18n.Localizer(r, "err.something_went_wrong"), http.StatusTooManyRequests)
				return
			}

			if err := rds.Decr(context.Background(), redisKey).Err(); err != nil {
				log.Println("Redis decr error:", err)
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
