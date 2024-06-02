package bff

import (
	"context"
	"fmt"

	"projeto-api/internal/app"

	"github.com/ServiceWeaver/weaver"
	"github.com/gofiber/fiber/v2"
)

type implBFF struct {
	weaver.Implements[weaver.Main]

	user weaver.Ref[app.Component]

	bff weaver.Listener `weaver:"bff"`
}

func (e *implBFF) createRouter(ctx context.Context, f *fiber.App) {

	middleware := NewMiddleware(NewMiddlewareIn{
		Router: f,
	})

	router := middleware.BaseRouter()

	g := router.Group("/users")
	g.Get("/", e.GetAllUsers)
	g.Get("/:id", e.GetUserById)
	g.Post("/", e.CreateUser)
}

func Serve(ctx context.Context, e *implBFF) error {
	fmt.Printf("BFF listener available on %v\n", e.bff.String())

	f := fiber.New()

	e.createRouter(ctx, f)

	return f.Listener(e.bff)
}
