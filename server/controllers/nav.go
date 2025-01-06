package controllers

import "github.com/gofiber/fiber/v3"

func Nav(c fiber.Ctx) error {
	return c.Render("includes/nav", fiber.Map{}, "includes/nav")
}
