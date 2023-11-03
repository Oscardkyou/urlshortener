package main

import (
	"urlshortener/internal/config"
	"urlshortener/internal/logger"
	"urlshortener/internal/server"
)

type App struct {
	server *server.Server
}

func (a *App) Initialize() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	log := logger.Initialize(cfg)

	// Initialize server
	a.server = server.Initialize(cfg, log)
}

func (a *App) Run() {
	a.server.Start()
}

func main() {
	app := &App{}
	app.Initialize()
	app.Run()
	// app.Server()
}
