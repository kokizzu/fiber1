package business

import (
	"fiber1/model"
	"github.com/jmoiron/sqlx"
	"github.com/kokizzu/gotro/S"
)

type Guest struct {
	Db *sqlx.DB
	//UploadToS3 func(filename string, content string) error
}

type Guest_LoginIn struct {
	CommonIn
	Email    string
	Password string
}

type Guest_LoginOut struct {
	CommonOut
	User *model.User
}

const Guest_LoginPath = `/guest/login`

func (g *Guest) Login(in *Guest_LoginIn) (out Guest_LoginOut) {
	if len(in.Email) < 3 {
		out.SetError(400, `invalid email`)
		return
	}
	if len(in.Password) < 3 {
		out.SetError(400, `invalid password`)
		return
	}

	user := model.NewUser(g.Db)

	if !user.FindByEmail(in.Email) {
		out.SetError(400, `user not found`)
		return
	}

	if !user.CheckPassword(in.Password) {
		out.SetError(400, `user or password does not match`)
		return
	}

	out.SetCookie = S.RandomCB63(32)
	// TODO: insert to session table (redis)

	out.User = user.Clean()
	return
}

type Guest_RegisterIn struct {
	CommonIn
	Email    string
	Password string
	Name     string
}

type Guest_RegisterOut struct {
	CommonOut
	User *model.User `json:"user,omitempty"`
}

const Guest_RegisterPath = `/guest/register`

func (g *Guest) Register(in *Guest_RegisterIn) (out Guest_RegisterOut) {
	if len(in.Email) < 3 {
		out.SetError(400, `invalid email`)
		return
	}
	if len(in.Password) < 3 {
		out.SetError(400, `invalid password`)
		return
	}
	if len(in.Name) < 3 {
		out.SetError(400, `invalid name`)
		return
	}

	user := model.NewUser(g.Db)
	if user.FindByEmail(in.Email) {
		out.SetError(400, `email already used`)
		return
	}

	user.Email = in.Email
	user.Name = in.Name
	user.SetPassword(in.Password)

	if !user.Insert() {
		out.SetError(500, `failed to insert user to database`)
		return
	}

	out.User = user.Clean()
	return
}
