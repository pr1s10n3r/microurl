package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pr1s10n3r/microurl/cmd/servd/routes/handlers"
	"github.com/pr1s10n3r/microurl/pkg/url"
)

type Router struct {
	UrlRepo url.Repository
}

func (r Router) Start() error {
	app := fiber.New()
	app.Get("/ping", handlers.Ping)

	urlHandler := handlers.NewURLHandler(r.UrlRepo)
	app.Get("/:code", urlHandler.GoTo)
	app.Get("/new", urlHandler.NewURL)
	app.Post("/new", urlHandler.StoreURL)

	return app.Listen(":3000")
}
