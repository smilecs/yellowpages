package web

import (
	"log"
	"net/http"
	"time"

<<<<<<< HEAD
	"github.com/tonyalaribe/yellowpages/messages"
=======
	"github.com/gorilla/context"
	"github.com/smilecs/yellowpages/messages"
>>>>>>> 48d8113334091894b410b4ed4222cc4868c19898
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
