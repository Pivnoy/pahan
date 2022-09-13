package app

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pahan/config"
	"pahan/pkg/httpserver"
	"pahan/pkg/postgres"
	"syscall"
)

func Run(cfg *config.Config) {

	_, err := postgres.New(cfg)
	if err != nil {
		log.Fatal("Error in creating Postgres Instance")
	}

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pashok delaet kyrsach"))
	})

	serv := httpserver.New(r, httpserver.Port(cfg.HTTP.Port))

	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err = <-serv.Notify():
		log.Printf("Notify from http server")
	}

	err = serv.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}
