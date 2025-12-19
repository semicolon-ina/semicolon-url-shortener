package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// GET /:code
func (h *HTTPHandler) Redirect(c *fiber.Ctx) error {
	code := c.Params("code") // Echo: c.Param, Fiber: c.Params

	data, err := h.uSvc.GetOriginalURL(c.UserContext(), code)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "url not found"})
	}

	return c.Redirect(data.OriginalURL, fiber.StatusMovedPermanently)
}
