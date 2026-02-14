package main

import (
	"net/http"
	router "to-do-api/api/router"
	"to-do-api/db"
)

func main() {
	db.Init()

	r := router.Init()

	println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
