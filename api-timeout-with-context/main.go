package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func SlowerHandler(w http.ResponseWriter, r *http.Request) {
	// slow external api call
	time.Sleep(2 * time.Second)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "slower handler")
}

func FasterHandler(w http.ResponseWriter, r *http.Request) {
	// faster external api call
	time.Sleep(200 * time.Millisecond)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "faster handler")
}

// WithTimeout acts as a middlware that interrupts the execution once the given handler takes longer than the defined timeout to execute
func WithTimeout(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()

		// let post created contexts shutdown properly
		r = r.WithContext(ctx)

		ch := make(chan struct{})
		go func() {
			next(w, r)
			close(ch)
		}()

		select {
		case <-ch:
		case <-ctx.Done():
			w.WriteHeader(http.StatusGatewayTimeout)
			fmt.Fprintln(w, ctx.Err().Error())
		}
	}
}

func main() {
	mux := http.NewServeMux()

	// wraps the handlers with the WithTimeout middleware. It can be expanded to including as many middlewares as want
	mux.HandleFunc("GET /slow", WithTimeout(SlowerHandler))
	mux.HandleFunc("GET /fast", WithTimeout(FasterHandler))

	http.ListenAndServe(":8080", mux)
}
