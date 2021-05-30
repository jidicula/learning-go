// Run `curl http://0.0.0.0:8080/hello` to see unauthenticated request, or run
// `curl -H "X-Secret-Password: GOPHER" http://0.0.0.0:8080/hello` to see
// authed request.
package main

import (
	"log"
	"net/http"
	"time"
)

// RequestTimer times a request.
func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
	})
}

var securityMsg = []byte("You didn't give the secret password\n")

// TerribleSecurityProvider accepts a password as config info and function
// returns middleware function that uses the config information.
// Passing values through layers of middleware is done via the context.
func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write(securityMsg)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func main() {
	terribleSecurity := TerribleSecurityProvider("GOPHER")
	mux := http.NewServeMux()
	mux.Handle("/hello", terribleSecurity(RequestTimer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello!\n"))
	}))))

	err := http.ListenAndServe("0.0.0.0:8080", mux)
	if err != nil {
		panic(err)
	}
}
