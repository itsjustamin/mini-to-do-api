package main

import (
	"fmt"
	"log"
	"net/http"
	router "to-do-api/api/router"
	"to-do-api/config"
	"to-do-api/db"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env файл не найден, используются только системные переменные")
	}
	cfg := config.LoadConfig()
	db.Init()

	r := router.Init()
	log.Printf("Server started at :%d", cfg.App.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port), r))

}
