package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // postgres driver
)

type PostgresConfig struct {
	DbConnectionStr string
}

func NewPostgres(config PostgresConfig) *sql.DB {
	db, err := sql.Open("postgres", config.DbConnectionStr)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
