package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectMysql(connStr string, silent ...bool) *sqlx.DB {
	db, err := sqlx.Connect("mysql", connStr)
	if err != nil {
		if len(silent) == 0 || !silent[0] {
			log.Printf("error connecting to mysql: %v", err)
		}
		return nil
	}
	return db
}
