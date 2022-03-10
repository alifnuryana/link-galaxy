package handlers

import (
	"fmt"
	"github.com/alifnuryana/link-galaxy/models"
	"github.com/alifnuryana/link-galaxy/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetUsers(ctx *fiber.Ctx) error {
	users, err := services.GetUsers()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
			"data":    nil,
		})
	}
	if len(users) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "success",
			"message": "tidak ada data users",
			"data":    users,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   users,
	})
}
func GetUser(ctx *fiber.Ctx) error {
	params := ctx.Params("id")

	id, err := strconv.Atoi(params)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "error convert param string to integer",
			"data":    nil,
		})
	}

	user, err := services.GetUser(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}
func AddUser(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var userInput models.User
	err := ctx.BodyParser(&userInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
			"data":    nil,
		})
	}

	user, err := services.AddUser(userInput)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
			"data":    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}
func UpdateUser(ctx *fiber.Ctx) error {
	params := ctx.Params("id")

	id, err := strconv.Atoi(params)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
			"data":    nil,
		})
	}

	var userInput models.User
	err = ctx.BodyParser(&userInput)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
			"data":    nil,
		})
	}
	user, err := services.UpdateUser(userInput, uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
			"data":    nil,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}
func DeleteUser(ctx *fiber.Ctx) error {
	params := ctx.Params("id")

	id, err := strconv.Atoi(params)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
			"data":    nil,
		})
	}

	err = services.DeleteUser(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("id with %d deleted", id),
	})
}
