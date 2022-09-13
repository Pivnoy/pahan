package app

import (
	"log"
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

	serv := httpserver.New(nil, httpserver.Port(cfg.HTTP.Port))

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
