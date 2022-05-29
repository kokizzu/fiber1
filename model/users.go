package model

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kokizzu/gotro/S"
)

type User struct {
	db       *sqlx.DB `json:"-"`
	Id       int64    `json:"id"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	Name     string   `json:"name"`
}

func NewUser(db *sqlx.DB) *User {
	return &User{db: db}
}

func (u *User) Migrate() (err error) {
	_, err = u.db.Exec(`
CREATE TABLE IF NOT EXISTS users (
  id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  email VARCHAR(64) NOT NULL DEFAULT '',
  password VARCHAR(64) NOT NULL DEFAULT ''
)`)
	if err != nil {
		log.Fatalf("error creating table: %v", err)
	}

	_, err = u.db.Exec(`ALTER TABLE users ADD column name VARCHAR(64) NOT NULL DEFAULT ''`)
	if err != nil {
		log.Printf("error adding column: %v", err)
	}

	_, err = u.db.Exec(`ALTER TABLE users ADD constraint email_unique UNIQUE(email)`)
	if err != nil {
		log.Printf("error creating unique email: %v", err)
	}
	return
}

func (u *User) FindByEmail(email string) (found bool) {
	err := u.db.Get(u, `
SELECT id, email, password, name
FROM users 
WHERE email = ?`, email)
	if err != nil {
		log.Printf(`User.FindByEmail: %v`, err)
	}
	return err == nil && u.Id > 0
}

func (u *User) CheckPassword(password string) bool {
	err := S.CheckPassword(u.Password, password)
	if err != nil {
		log.Printf(`User.CheckPassword: %v`, err)
		return false
	}
	return true
}

func (u *User) SetPassword(password string) {
	u.Password = S.EncryptPassword(password)
}

func (u *User) Insert() bool {
	ra, err := u.db.Exec(`
INSERT INTO users (email, password, name)
VALUES (?, ?, ?)
`, u.Email, u.Password, u.Name)
	if err != nil {
		log.Printf(`User.Insert: %v`, err)
		return false
	}
	id, err := ra.LastInsertId()
	if err != nil {
		log.Printf(`User.LastInsertId: %v`, err)
		return false
	}
	u.Id = id
	return true
}

func (u *User) Clean() *User {
	u.db = nil
	return u
}
