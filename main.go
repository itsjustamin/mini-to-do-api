package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

var tasks []Task
var nextID = 1

func GetTasks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func PostTask(w http.ResponseWriter, r *http.Request) {

	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func PutTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)


	for i, t := range tasks {
		if t.ID == id {
			var updated Task
			json.NewDecoder(r.Body).Decode(&updated)
			tasks[i].Title = updated.Title
			tasks[i].Done = updated.Done
			tasks[i].StartTime = updated.StartTime
			tasks[i].EndTime = updated.EndTime

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)


	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func PatchTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)


	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(tasks[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks", PostTask).Methods("POST")
	router.HandleFunc("/tasks/{{id:[0-9]+}", PutTask).Methods("PUT")
	router.HandleFunc("/tasks/{id:[0-9]+}", DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/done/{id:[0-9]+}", PatchTask).Methods("PATCH")

	println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
