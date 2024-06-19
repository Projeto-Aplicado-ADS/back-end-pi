package bff

import (
	"projeto-api/pkg/token"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type (
	NewMiddlewareIn struct {
		Router *fiber.App
	}
	Middleware interface {
		BaseRouter() fiber.Router
		Logger(...logger.Config) func(*fiber.Ctx) error
		Cors() func(*fiber.Ctx) error
		Authorization() func(*fiber.Ctx) error
	}
	implMiddleware struct {
		router fiber.Router
	}
)

func NewMiddleware(in NewMiddlewareIn) Middleware {
	return &implMiddleware{
		router: in.Router,
	}
}

func (e implMiddleware) BaseRouter() fiber.Router {
	e.router.Use(e.Cors())
	return e.router
}

func (e implMiddleware) Logger(log ...logger.Config) func(*fiber.Ctx) error {
	return logger.New(log...)
}

func (e implMiddleware) Cors() func(*fiber.Ctx) error {
	return cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization, Access-Control-Allow-Headers",
		AllowOrigins:     "http://localhost:3000", 
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	})
}

func (e implMiddleware) Authorization() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization")  

		if tokenString == "" {
			return fiber.ErrUnauthorized
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		tokenParsed, err := token.New().ParseToken(tokenString)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		email := tokenParsed.Email
		if email == "" {
				return fiber.NewError(fiber.StatusBadRequest, "email claim not found")
		}

		ctx.Locals("email", email)

		return ctx.Next()
	}
}