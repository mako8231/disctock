package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mako8231/disctock/server/controllers"
)

func WebRoutes(app *fiber.App) {
	app.Get("/", controllers.Index)
	app.Get("/header", controllers.Header)
	app.Get("/nav", controllers.Nav)
}
