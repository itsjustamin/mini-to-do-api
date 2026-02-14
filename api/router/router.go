package api

import (
	handlers "to-do-api/api/handlers"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handlers.PostTask).Methods("POST")
	r.HandleFunc("/tasks/{id:[0-9]+}", handlers.PutTask).Methods("PUT")
	r.HandleFunc("/tasks/{id:[0-9]+}", handlers.DeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/done/{id:[0-9]+}", handlers.PatchTask).Methods("PATCH")

	return r
}
