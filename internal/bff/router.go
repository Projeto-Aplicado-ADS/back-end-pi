package bff

import (
	"fmt"
	"time"

	"projeto-api/internal/app"
	"projeto-api/pkg/token"

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
	userId := c.Params("id")

	if userId == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid parameter:"+c.Params("id"))
	}

	out, err := e.user.Get().GetOneUserById(c.Context(), string(userId))
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

func (e implBFF) Login(c *fiber.Ctx) (err error) {
	var body app.UserGetByEmailAndPassword

	err = c.BodyParser(&body)
	if err != nil {
		fmt.Println(err)
		return fiber.ErrBadRequest
	}

	err = e.user.Get().GetUserByEmailAndPassword(c.Context(), body.Email, body.Password)
	if err != nil {
		return err
	}

	tokenString, err := token.New().CreateNewToken(body.Email, body.IsAdmin)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	

	return c.SendStatus(fiber.StatusCreated)
}

func (e implBFF) ListUserByEmail(c *fiber.Ctx) (err error) {

	var tokenFromCookie = c.Cookies("token")
	if tokenFromCookie == "" {
		return fiber.ErrUnauthorized
	}

	err = token.New().ValidateToken(tokenFromCookie)
	if err != nil {
		return err
	}

	email := c.Params("email")

	if email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "invalid parameter:"+c.Params("email"))
	}

	out, err := e.user.Get().ListUserByEmail(c.Context(), string(email))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(fiber.StatusOK).JSON(out)
}