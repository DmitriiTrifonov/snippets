package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
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

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("ui/static/")})

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

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

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err = nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
