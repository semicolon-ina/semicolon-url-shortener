package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/semicolon-ina/semicolon-url-shortener/repo/common/config"
)

// NewFiberServer ngebalikin *fiber.App yang udah "Batteries Included"
func NewFiberServer(cfg config.DefaultConfig) *fiber.App {
	// 1. Init Engine
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
		// DisableStartupMessage: true, // Uncomment kalau mau terminal bersih
	})

	// 2. Pasang Standard Middleware
	app.Use(recover.New()) // Biar gak crash kalo panic
	app.Use(logger.New())  // Access Log
	app.Use(cors.New())    // CORS Allow All (bisa dituning)

	// 3. Default Health Check (Standardize buat semua service)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": cfg.AppName,
		})
	})

	return app
}
