package controllers

import "github.com/gofiber/fiber/v3"

func Header(c fiber.Ctx) error {
	return c.Render("includes/header", fiber.Map{}, "includes/header")
}
