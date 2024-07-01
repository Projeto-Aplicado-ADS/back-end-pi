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
  auth := middleware.Authorization()

	g := router.Group("/users")
	g.Get("/", e.GetAllUsers)
	g.Get("/me", auth, e.getMe)
	g.Get("/:id", e.GetUserById)
	g.Post("/login", e.Login)
	g.Post("/", e.CreateUser)
	g.Put("/email/:id", e.UpdateUsersEmail)
	g.Put("/phone/:id", e.UpdateUsersPhone)

	/* Hospedes */
  gH := router.Group("/hospedes")
	gH.Get("/", auth, e.GetHospedes)
	gH.Get("/:id", auth, e.GetHospedeById)
	gH.Post("/", auth, e.CreateHospede)
	gH.Put("/delete/:id", auth, e.DeleteHospede)
	gH.Put("/update/:id", auth, e.UpdateHospede)

	/* Quartos */

	gQ := router.Group("/quartos")
	gQ.Get("/", auth, e.GetQuartos)
	gQ.Post("/", auth, e.CreateQuarto)
	gQ.Put("/:id", auth, e.UpdateStatusQuarto)

	/* Reservas */

	gR := router.Group("/reservas")
	gR.Get("/", auth, e.GetReservas)
	gR.Post("/:quartoId/:hospedeId", auth, e.CreateReserva)
	gR.Put("/remove/:id", auth, e.RemoveReserva)
	gR.Put("/:id", auth, e.UpdateReservas)

}

func Server(ctx context.Context, e *implBFF) error {
	fmt.Printf("BFF listener available on %v\n", e.bff.String())

	f := fiber.New()

	e.createRouter(ctx, f)

	return f.Listener(e.bff)
}
