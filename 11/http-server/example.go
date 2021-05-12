package main

import (
	"net/http"
	"time"
)

type HelloHandler struct {
}

// ServeHTTP handles.
func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello!\n"))
}

func main() {
	s := http.Server{
		Handler:      HelloHandler{},
		Addr:         "0.0.0.0/:8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
