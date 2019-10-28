package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", Healthz)
	mux.HandleFunc("/greet", Greet)

	log.Println("Starting server :8080")
	s := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("."))
}

func Greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	fmt.Fprintf(w, `{"response": "Hello world!"}`)
}
