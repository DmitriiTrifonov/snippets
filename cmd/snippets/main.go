package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"snippets/internal/application"
	"snippets/internal/fs"
	"snippets/internal/handlers"
	"syscall"

	"go.uber.org/zap"
)

const (
	port = ":8080"
)

// Config collects an initial config data
type Config struct {
	Address       string
	StaticDirPath string
}

func main() {
	cfg := Config{}

	flag.StringVar(&cfg.Address, "addr", port, "HTTP-server address")
	flag.StringVar(&cfg.StaticDirPath, "static-dir-path", "ui/static/", "Path to static dir")
	flag.Parse()

	app := application.NewContainer()

	logger, err := app.GetLogger()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", &handlers.Root{})
	mux.Handle("/snippet", &handlers.SnippetGetter{})
	mux.Handle("/snippet/add", &handlers.SnippetAdder{})

	fileServer := http.FileServer(fs.NewNeuteredFS(http.Dir(cfg.StaticDirPath), logger))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	startingMessage := "starting server on " + cfg.Address
	logger.Info(startingMessage)
	go func() {
		if errServe := http.ListenAndServe(cfg.Address, mux); errServe != nil {
			logger.Error("cannot init http-server", zap.Error(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	sig := <-stop
	signalMessage := "got a signal: " + sig.String()
	logger.Info(signalMessage)
	stopMessage := "stopped server on " + cfg.Address
	logger.Info(stopMessage)
}
