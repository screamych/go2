package models

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Task struct {
	Id   int       `json:"id"`
	Text string    `json:"text"`
	Tags string    `json:"tags"`
	Due  time.Time `json:"due"`
}

func InitDB() error {
	db, err := sql.Open("sqlite3", "./tasks.db?")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER NOT NULL PRIMARY KEY,
			text TEXT,
			tags STRING,
			due TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func CreateTask(text string, tags string, due time.Time) (bool, error) {
	task := Task{
		Text: text,
		Tags: tags,
		Due:  due,
	}

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}
	stmt, err := tx.Prepare("INSERT INTO tasks (text, tags, due) VALUES (?, ?, ?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Text, task.Tags, task.Due)
	if err != nil {
		return false, err
	}
	tx.Commit()

	return true, nil
}

func GetAllTasks() ([]Task, error) {
	rows, err := DB.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	tasks := make([]Task, 0)
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.Id, &task.Text, &task.Tags, &task.Due)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskById(id int) (Task, error) {
	stmt, err := DB.Prepare("SELECT * FROM tasks WHERE id = ?")
	if err != nil {
		return Task{}, err
	}

	task := Task{}
	sqlErr := stmt.QueryRow(id).Scan(&task.Id, &task.Text, &task.Tags, &task.Due)
	if sqlErr != nil {
		return Task{}, sqlErr
	}

	return task, nil
}

func DeleteAllTasks() (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare("DELETE FROM tasks")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func DeleteTaskById(id int) (bool, error) {
	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
