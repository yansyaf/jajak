package utils

import (
	"log"
	"net/http"
	"time"
)

func ThrowPanic(e error) {
	if e != nil {
		log.Panicf("throwing panic %v", e)
		panic(e)
	}
}

func CommonPanicHandler() {
	if r := recover(); r != nil {
		log.Fatalf("Unhandled Panic: %v", r)
	}
}

func LoggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(w, r)
		timeEnd := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), timeEnd.Sub(timeStart))
	}

	return http.HandlerFunc(fn)
}
