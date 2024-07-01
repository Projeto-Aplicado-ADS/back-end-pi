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
		HTTPOnly: false,
	})

	return c.SendStatus(fiber.StatusCreated)
}

func (e implBFF) getMe(c *fiber.Ctx) (err error) {
	email, ok := c.Locals("email").(string)

	if !ok || email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Email not found")
	}

	out, err := e.user.Get().ListUserByEmail(c.Context(), email)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(out)
}

func (e implBFF) UpdateUsersEmail (c *fiber.Ctx) (err error) {
	var body app.UpdateUsersEmailIn
	if err = c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}
	_, err = e.user.Get().UpdateUserByEmail(c.Context(), app.UpdateUsersEmailIn{
		ID: c.Params("id"),
		Email: body.Email,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (e implBFF) UpdateUsersPhone(c *fiber.Ctx) (err error) {
	var body app.UpdateUsersPhoneIn
	if err = c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}
	_, err = e.user.Get().UpdateUserPhone(c.Context(), app.UpdateUsersPhoneIn{
		ID: c.Params("id"),
		Phone: body.Phone,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusNoContent)
}


func (e implBFF) Logout(c *fiber.Ctx) (err error) {
	c.ClearCookie("token")
	return c.SendStatus(fiber.StatusOK)
}

func (e implBFF) GetHospedes(c *fiber.Ctx) (err error) {
	out, err := e.user.Get().GetHospedes(c.Context())
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(out)
}

func (e implBFF) GetHospedeById (c *fiber.Ctx) (err error) {
	out, err := e.user.Get().GetOneHospedeById(c.Context(), c.Params("id"))
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(out)
}

func (e implBFF) CreateHospede (c *fiber.Ctx) (err error) {
	var body app.HospedesIn

	err = c.BodyParser(&body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	_, err = e.user.Get().CreateHospede(c.Context(), body)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusCreated)
}


func (e implBFF) DeleteHospede(c *fiber.Ctx) (err error) {
	var body app.RemoveHospedesIn
	if err = c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	_, err = e.user.Get().RemoveHospede(c.Context(), app.RemoveHospedesIn{
		ID: c.Params("id"),
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}	

	return c.SendStatus(fiber.StatusNoContent)
}

func (e implBFF) UpdateHospede(c *fiber.Ctx) (err error) {
	var body app.UpdateHospedesIn
	if err = c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}
	_, err = e.user.Get().UpdateHospedesIn(c.Context(), app.UpdateHospedesIn{
		ID: c.Params("id"),
		Nome: body.Nome,
		Email: body.Email,
		Telefone: body.Telefone,
		Cpf: body.Cpf,
		DataNascimento: body.DataNascimento,
		Sexo: body.Sexo,
	})
	if err != nil {
		fmt.Println(err)
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusNoContent)
}


func (e implBFF) CreateQuarto(c *fiber.Ctx) (err error) {
	var body app.QuartosIn

	err = c.BodyParser(&body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	_, err = e.user.Get().CreateQuarto(c.Context(), body)
	if err != nil {
		return fiber.ErrInternalServerError
	}


	return c.SendStatus(fiber.StatusCreated)
}

func (e implBFF) GetQuartos(c *fiber.Ctx) (err error) {
	out, err := e.user.Get().GetQuartos(c.Context())
	if err != nil {
		return err
	}
	
	return c.Status(fiber.StatusOK).JSON(out)
}


func (e implBFF) UpdateStatusQuarto(c *fiber.Ctx) (err error) {
	var body app.UpdateQuartosStatusIn
	if err = c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	_, err = e.user.Get().UpdateStatusQuarto(c.Context(), app.UpdateQuartosStatusIn{
		ID: c.Params("id"),
		StatusQuarto: body.StatusQuarto,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}


func (e implBFF) GetReservas (c *fiber.Ctx) (err error) {
	out, err := e.user.Get().GetReservas(c.Context())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(out)
}

func (e implBFF) CreateReserva (c *fiber.Ctx) (err error) {
	var body app.ReservasIn

	err = c.BodyParser(&body)
	if err != nil {
		return fiber.ErrBadRequest
	}

	body.QuartoId = c.Params("quartoId")
	body.HospedeId = c.Params("hospedeId")

	_, err = e.user.Get().CreateReserva(c.Context(), body)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusCreated)
}