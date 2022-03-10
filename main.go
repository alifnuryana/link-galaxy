package main

import (
	"log"

	"github.com/alifnuryana/link-galaxy/database"
	"github.com/alifnuryana/link-galaxy/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupRouter(app)
	router.ServeStatic(app)

	database.Setup()

	log.Fatal(app.Listen(":8080"))
}
