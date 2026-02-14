package repository

import (
	"to-do-api/db"
	"to-do-api/models"
)

func GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := db.DB.Select(&tasks, "SELECT * FROM tasks ORDER BY id")
	return tasks, err
}

func CreateTask(task *models.Task) error {
	query := `
	INSERT INTO tasks (title, done, start_time, end_time)
	VALUES ($1,$2,$3,$4)
	RETURNING id`

	return db.DB.QueryRow(query,
		task.Title,
		task.Done,
		task.StartTime,
		task.EndTime,
	).Scan(&task.ID)
}

func UpdateTask(id int, task *models.Task) error {
	return db.DB.Get(task, `
	UPDATE tasks
	SET title=$1, done=$2, start_time=$3, end_time=$4
	WHERE id=$5
	RETURNING *`,
		task.Title, task.Done, task.StartTime, task.EndTime, id)
}

func DeleteTask(id int) error {
	_, err := db.DB.Exec("DELETE FROM tasks WHERE id=$1", id)
	return err
}

func DoneTask(id int, task *models.Task) error {
	return db.DB.Get(task, `
	UPDATE tasks SET done=true WHERE id=$1 RETURNING *`, id)
}
