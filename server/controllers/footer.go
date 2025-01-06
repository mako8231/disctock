package controllers

import "github.com/gofiber/fiber/v3"

func Footer(c fiber.Ctx) error {
	return c.Render("includes/footer", fiber.Map{}, "includes/footer")
}
