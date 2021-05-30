package main

import (
	"fmt"
	"net/http"
	"os"
)

type HelloHandler struct {
}

// ServeHTTP handles.
func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello!\n"))
}

func main() {
	// s := http.Server{
	// 	Handler:      HelloHandler{},
	// 	Addr:         ":8080",
	// 	ReadTimeout:  30 * time.Second,
	// 	WriteTimeout: 90 * time.Second,
	// 	IdleTimeout:  120 * time.Second,
	// }
	// err := s.ListenAndServe()
	// if err != nil {
	// 	if err != http.ErrServerClosed {
	// 		panic(err)
	// 	}
	// }

	// mux := http.NewServeMux()
	// mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello!\n"))
	// })

	person := http.NewServeMux()
	person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		nBytesWritten, err := w.Write([]byte("greetings!\n"))
		if err != nil || nBytesWritten == 0 {
			fmt.Printf("Error!\n")
			os.Exit(1)
		}
	})
	dog := http.NewServeMux()
	dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		nBytesWritten, err := w.Write([]byte("good puppy!\n"))
		if err != nil || nBytesWritten == 0 {
			fmt.Printf("Error!\n")
			os.Exit(1)
		}
	})

	mux := http.NewServeMux()
	mux.Handle("/person/", http.StripPrefix("/person", person))
	mux.Handle("/dog/", http.StripPrefix("/dog", dog))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
