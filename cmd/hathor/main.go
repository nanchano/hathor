package main

import (
	"os"
	"os/signal"

	"github.com/nanchano/hathor/internal/config"
	"github.com/nanchano/hathor/internal/core"
	"github.com/nanchano/hathor/internal/repository"
	"github.com/nanchano/hathor/internal/server"
	"golang.org/x/exp/slog"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout)
	logger := slog.New(handler)

	logger.Info("Reading config")
	config, err := config.Load(".env")
	if err != nil {
		logger.Error("Failed parsing config", "error", err)
		panic(err)
	}

	repo, err := repository.New(config.Database.URL())
	if err != nil {
		logger.Error("Failed connecting to the database", "error", err)
		panic(err)
	}

	service := core.NewService(logger, repo)

	server := server.New(logger, service)
	server.Start(config.Server.Host, config.Server.Port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
