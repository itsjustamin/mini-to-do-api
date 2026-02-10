package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Task struct {
	ID        int       `db:"id"         json:"id"`
	Title     string    `db:"title"      json:"title"`
	Done      bool      `db:"done"       json:"done"`
	StartTime time.Time `db:"start_time" json:"start_time"`
	EndTime   time.Time `db:"end_time"   json:"end_time"`
}

var db *sqlx.DB

func GetTasks(w http.ResponseWriter, r *http.Request) {

	var task []Task
	err := db.Select(&task, "SELECT * FROM tasks ORDER BY id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func PostTask(w http.ResponseWriter, r *http.Request) {

	var task Task
	json.NewDecoder(r.Body).Decode(&task)

	query := `
		INSERT INTO tasks (title, done, start_time, end_time)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := db.QueryRow(
		query,
		task.Title,
		task.Done,
		task.StartTime,
		task.EndTime,
	).Scan(&task.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func PutTask(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var task Task
	json.NewDecoder(r.Body).Decode(&task)

	_, err := db.Exec(`
		UPDATE tasks
		SET title=$1, done=$2, start_time=$3, end_time=$4
		WHERE id=$5
	`,
		task.Title,
		task.Done,
		task.StartTime,
		task.EndTime,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	task.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	_, err := db.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func PatchTask(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var task Task
	err := db.Get(&task, `
		UPDATE tasks
		SET done=true
		WHERE id=$1
		RETURNING *
	`, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}

func main() {

	var err error
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=taskdb sslmode=disable"

	db, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetTasks).Methods("GET")
	router.HandleFunc("/tasks", PostTask).Methods("POST")
	router.HandleFunc("/tasks/{id:[0-9]+}", PutTask).Methods("PUT")
	router.HandleFunc("/tasks/{id:[0-9]+}", DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/done/{id:[0-9]+}", PatchTask).Methods("PATCH")

	println("Server started at :8080")
	http.ListenAndServe(":8080", router)
}
