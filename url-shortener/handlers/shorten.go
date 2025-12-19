package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// POST /api/v1/shorten
func (h *HTTPHandler) ShortenURL(c *fiber.Ctx) error {
	var req ShortenRequest

	// Bedanya Echo vs Fiber: c.Bind() vs c.BodyParser()
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	if req.URL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "url is required"})
	}

	// PENTING: Fiber pake c.UserContext() atau c.Context() buat pass ke layer bawah
	code, err := h.uSvc.ShortenURL(c.UserContext(), req.URL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Construct Full URL
	host := c.Protocol() + "://" + c.Hostname()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":         code,
		"short_url":    host + code,
		"original_url": req.URL,
	})
}
