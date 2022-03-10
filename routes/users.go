package routes

import (
	"github.com/alifnuryana/link-galaxy/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesUsers(router fiber.Router) {
	users := router.Group("/users")

	users.Get("/", handlers.GetUsers)
	users.Get("/:id", handlers.GetUser)
	users.Post("/", handlers.AddUser)
	users.Put("/:id", handlers.UpdateUser)
	users.Delete("/:id", handlers.DeleteUser)
}
