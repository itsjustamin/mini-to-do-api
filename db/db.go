package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init() {
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=taskdb sslmode=disable"

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
}
