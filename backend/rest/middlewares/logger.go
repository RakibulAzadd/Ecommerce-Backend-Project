package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("Ami loger  age print hobo")
		next.ServeHTTP(w, r)
		log.Println("Ami loger pore print hobo")
		log.Println(r.Method, r.URL.Path, time.Since(start))
	})
}
