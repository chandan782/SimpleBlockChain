package routes

import (
	"fmt"
	"net/http"

	"github.com/chandan782/SimpleBlockChain/api/resources/health"
	"github.com/chandan782/SimpleBlockChain/api/resources/version"
	"github.com/rohanraj7316/middleware/libs/response"

	"github.com/gofiber/fiber/v2"
)

type Router func(fiber.Router)

type Route struct {
	path   string
	router Router
}

func New(app *fiber.App) {
	// list down all the routes and their handlers
	routes := []Route{
		{
			path:   "/health",
			router: health.Router,
		},
		{
			path:   "/version",
			router: version.Router,
		},
	}

	for i := 0; i < len(routes); i++ {
		routes[i].router(app.Group(routes[i].path))
	}

	app.Use("*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Cannot %s %s", c.Method(), c.Path()) // Cannot GET /healths
		return response.NewBody(c, http.StatusInternalServerError, msg, nil, nil)
	})
}
