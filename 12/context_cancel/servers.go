package main

import (
	"net/http"
	"net/http/httptest"
	"time"
)

// slowServer is a slow server.
func slowServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		_, _ = w.Write([]byte("Slow response"))
	}))
	return s
}

// fastServer is a fast server.
func fastServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("error") == "true" {
			_, _ = w.Write([]byte("error"))
			return
		}
		_, _ = w.Write([]byte("ok"))
	}))
	return s
}
