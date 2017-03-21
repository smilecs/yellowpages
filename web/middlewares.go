package web

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/context"
	"github.com/smilecs/yellowpages/messages"
)

// Middlewares

func RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				messages.WriteError(w, messages.ErrInternalServer)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func AcceptHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept") != "application/json" {
			messages.WriteError(w, messages.ErrNotAcceptable)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func ContentTypeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			messages.WriteError(w, messages.ErrUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func BodyHandler(v interface{}) func(http.Handler) http.Handler {
	t := reflect.TypeOf(v)

	m := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			val := reflect.New(t).Interface()
			err := json.NewDecoder(r.Body).Decode(val)

			if err != nil {
				log.Println(err)
				messages.WriteError(w, messages.ErrBadRequest)
				return
			}

			if next != nil {
				context.Set(r, "body", val)
				next.ServeHTTP(w, r)
			}
		}

		return http.HandlerFunc(fn)
	}

	return m
}
