package controllers

import "github.com/gofiber/fiber/v3"

func Index(c fiber.Ctx) error {
	return c.Render("task/index", fiber.Map{})
}

func Store(c fiber.Ctx) error {
	return c.SendString("aaaa")
}
