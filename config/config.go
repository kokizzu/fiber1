package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectMysql(connStr string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", connStr)
	if err != nil {
		log.Printf("error connecting to mysql: %v", err)
		return nil
	}
	return db
}
