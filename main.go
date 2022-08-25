package main

import (
	"errors"
	"log"

	"fiber1/business"
	"fiber1/config"
	"fiber1/model"
	"fiber1/presentation"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db := config.ConnectMysql(`root:password@(127.0.0.1:3306)/test1`)
	if db == nil {
		panic(errors.New(`cannot connect to database`))
	}

	guest := business.Guest{
		Db: db,
		// UploadToS3: s3.UploadToS3,
	}
	_ = guest

	user := model.NewUser(db)
	_ = user.Migrate()

	app := fiber.New()

	presentation.HandlePost[business.Guest_LoginIn](app, business.Guest_LoginPath, guest.Login)
	presentation.HandleGet(app, business.Guest_LoginPath)

	app.Get(business.Guest_LoginPath, func(c *fiber.Ctx) error {
		return presentation.RenderHtml(c, business.Guest_LoginPath, nil)
	})

	presentation.HandlePost[business.Guest_RegisterIn](app, business.Guest_RegisterPath, guest.Register)
	presentation.HandleGet(app, business.Guest_RegisterPath)

	log.Fatal(app.Listen(":3000"))
}
