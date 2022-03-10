package router

import (
	"github.com/alifnuryana/link-galaxy/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	routes.SetupRoutesUsers(v1)
}

func ServeStatic(app *fiber.App) {
	app.Static("/", "view/dist")

	app.Get("/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("view/dist/index.html")
	})
}
