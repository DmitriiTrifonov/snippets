package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"snippets/internal/handlers"
	"syscall"
)

const (
	port = ":8080"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Root)
	mux.HandleFunc("/snippet", handlers.GetSnippet)
	mux.HandleFunc("/snippet/add", handlers.AddSnippet)

	log.Printf("Starting server on %s", port)
	go func() {
		if err := http.ListenAndServe(port, mux); err != nil {
			log.Println(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	sig := <-stop
	log.Printf("Got a signal: %s", sig)
	log.Printf("Stopped server on %s", port)
}
