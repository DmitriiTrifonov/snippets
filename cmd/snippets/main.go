package main

import (
	"log"
	"net/http"
)

const port = ":8080"

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("This is a root of snippets"))
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	log.Printf("Starting server on %s", port)
	defer log.Printf("Stopped server on %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Println(err)
	}
}
