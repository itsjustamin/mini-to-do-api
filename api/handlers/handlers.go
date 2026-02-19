package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"to-do-api/models"
	"to-do-api/repository"

	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	err := repository.CreateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func PutTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	err := repository.UpdateTask(id, &task)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	err := repository.DeleteTask(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func PatchTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var task models.Task
	err := repository.DoneTask(id, &task)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	json.NewEncoder(w).Encode(task)
}
