package restapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/semicolon-ina/semicolon-url-shortener/repo/domain/url"
	"github.com/semicolon-ina/semicolon-url-shortener/repo/infra/inmem"
	"github.com/semicolon-ina/semicolon-url-shortener/url-shortener/config"
	"github.com/semicolon-ina/semicolon-url-shortener/url-shortener/handlers"
)

// RegisterRoutes cuma tugasnya nempelin handler ke path
func RegisterRoutes(app *fiber.App, cfg config.Config) {
	// 2. Init Layers
	rdb := inmem.Init(cfg.Redis)

	svc := url.New(cfg.URL, rdb)
	h := handlers.NewHTTPHandler(svc)

	// Root level redirect
	app.Get("/:code", h.Redirect)

	// API Group
	api := app.Group("/api/v1")
	api.Get("/get/:code", h.GetData)
	api.Post("/shorten", h.ShortenURL)
}
