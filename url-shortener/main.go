package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/semicolon-ina/semicolon-url-shortener/repo/common/logger"
	server "github.com/semicolon-ina/semicolon-url-shortener/repo/common/server" // Import Common Server
	"github.com/semicolon-ina/semicolon-url-shortener/url-shortener/config"

	"github.com/semicolon-ina/semicolon-url-shortener/url-shortener/routers/restapi"
	// Jangan lupa import redis driver
)

func main() {
	// setup logger
	logger.SetupZeroLog()
	// 1. Init Config & Infra (Sama kayak sebelumnya)
	config.LoadConfig()
	cfg := config.Get()

	fmt.Printf("config: %+v\n", cfg)

	// 3. Init Fiber App (Pake Common!)
	// Liat ini, kita cuma manggil satu baris dari common
	app := server.NewFiberServer(cfg.DefaultConfig)

	// 4. Register Routes (Inject Handler)
	restapi.RegisterRoutes(app, cfg)

	// 5. Start Server
	port := cfg.AppPort
	if port == "" {
		port = "8080"
	}

	// Graceful Shutdown Channel
	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // Block sampe ada sinyal kill
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()
	log.Println("Running Cleanup Tasks...")
}
