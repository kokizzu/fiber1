package presentation

import (
	"fiber1/business"
	"github.com/gofiber/fiber/v2"
)

func RenderRestApi(c *fiber.Ctx, out interface{}, co business.CommonOut) error {
	if co.StatusCode > 0 {
		c.Status(co.StatusCode)
	}
	if co.SetCookie != `` {
		const sessionToken = `sessionToken`
		if co.SetCookie == `CLEAR` {
			c.ClearCookie(sessionToken)
		}
		c.Cookie(&fiber.Cookie{
			Name:  sessionToken,
			Value: co.SetCookie,
		})
	}
	return c.JSON(out)
}
