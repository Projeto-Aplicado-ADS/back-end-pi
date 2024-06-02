package bff

import (
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
		AllowOrigins:     "*",
		AllowHeaders:     "Accept, Authorization, Content-Type, X-CSRF-Token, X-API-Key",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		ExposeHeaders:    "ETag",
		AllowCredentials: false,
		MaxAge:           300,
	})
}
