package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"pahan/config"
	"pahan/internal/controller/http/v1"
	"pahan/internal/usecase"
	"pahan/internal/usecase/repo"
	"pahan/pkg/httpserver"
	"pahan/pkg/postgres"
	"syscall"
)

func Run(cfg *config.Config) {

	pg, err := postgres.New(cfg)
	if err != nil {
		log.Fatal("Error in creating Postgres Instance")
	}

	engineUseCase := usecase.NewEngineUseCase(repo.NewEngineRepo(pg))
	suspensionUseCase := usecase.NewSuspensionUseCase(repo.NewSuspensionRepo(pg))
	// http Server
	handler := gin.New()
	v1.NewRouter(handler, engineUseCase, suspensionUseCase)

	serv := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
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
