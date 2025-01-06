package controllers

import "github.com/gofiber/fiber/v3"

// Serving the Index page:
func Index(c fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
