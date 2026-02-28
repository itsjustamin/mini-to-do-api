package db

import (
	"log"

	"to-do-api/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init() {
	cfg := config.LoadConfig()

	dsn := cfg.DBUrl()

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database connected successfully")
}
