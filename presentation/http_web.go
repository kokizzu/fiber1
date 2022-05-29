package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/S"
)

func RenderHtml(c *fiber.Ctx, path string, bind interface{}) error {
	return c.Render(`business/`+S.Replace(path[1:], `/`, `_`)+`.html`, bind)
}

func HandleGet(app *fiber.App, path string) {
	app.Get(path, func(c *fiber.Ctx) error {
		return RenderHtml(c, path, nil)
	})
}
