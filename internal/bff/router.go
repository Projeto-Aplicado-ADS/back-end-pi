package bff

import (
	"fmt"
	"strconv"

	"projeto-api/internal/app"

	"github.com/gofiber/fiber/v2"
)

func (e implBFF) GetAllUsers(c *fiber.Ctx) (err error) {
	out, err := e.user.Get().AllUsers(c.Context())
	if err != nil {
		fmt.Println(err)
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(out)
}

func (e implBFF) GetUserById(c *fiber.Ctx) (err error) {
	id, err := strconv.Unquote(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid parameter:"+c.Params("id"))
	}
	out, err := e.user.Get().GetOneUserById(c.Context(), string(id))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(out)
}

func (e implBFF) CreateUser(c *fiber.Ctx) (err error) {
	var body app.UserIn

	err = c.BodyParser(&body)
	if err != nil {
		fmt.Println(err)
		return fiber.ErrBadRequest
	}

	_, err = e.user.Get().CreateUser(c.Context(), body)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusCreated)
}
