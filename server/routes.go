package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/youthcongress/nepal/scripts/auth"
)

func Routes(app *fiber.App) {
	pages := map[string]string{
		"/": "index",
		"/login": "login",
		"/register": "register",
		"/profile": "profile",
	}

	for route, page := range pages {
		app.Get(route, func(p string) fiber.Handler {
			return func(c *fiber.Ctx) error {
				return c.Render(p, nil)
			}
		}(page))
	}

	// Register authentication routes from auth package
	auth.Register(app)
}