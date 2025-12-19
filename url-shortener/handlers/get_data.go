package handlers

import "github.com/gofiber/fiber/v2"

func (h *HTTPHandler) GetData(c *fiber.Ctx) error {
	code := c.Params("code") // Echo: c.Param, Fiber: c.Params

	originalURL, err := h.uSvc.GetOriginalURL(c.UserContext(), code)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "url not found"})
	}

	return c.Status(fiber.StatusOK).JSON(originalURL)
}
