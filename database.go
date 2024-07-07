package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"path"
)

type DB struct {
	location string
}

func DBInit() DB {
	home_path := os.Getenv("HOME")
	db_path := path.Join(home_path, ".local/share/planner-on-the-go")

	os.MkdirAll(db_path, 0755)

	db_file := path.Join(db_path, "tasks.db")

	db_object := DB{location: db_file}

	return db_object
}

func (d *DB) OpenConn() *sql.DB {
	db, err := sql.Open("sqlite3", d.location)

	if err != nil {
		log.Fatal("Can't open database")
	}

	return db
}

func TableInit(conn *sql.DB) {
	sql := `
CREATE TABLE IF NOT EXISTS tasks (id integer not null primary key, task text, is_done text, day int);
	`
	conn.Exec(sql)
	defer conn.Close()
}

func AddToDB(conn *sql.DB, task Task, day int) {
	tx, err := conn.Begin()
	if err != nil {
		log.Fatal(err)
	}
	sql := `
INSERT INTO tasks (task, is_done, day) values (?, ?, ?)
	`
	stmt, err := tx.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	stmt.Exec(task.task, task.is_done, day)
	tx.Commit()
}

func CurrentStatus(conn *sql.DB, id int) bool {
	sql := `
SELECT is_done FROM tasks WHERE id = ?
	`
	row, err := conn.Query(sql, id)

	if err != nil {
		log.Fatal(err)
	}

	var is_done bool

	for row.Next() {
		row.Scan(&is_done)

	}

	return is_done
}

func UpdateStatus(conn *sql.DB, id int, status bool) {
	sql := `
UPDATE tasks SET is_done = ? WHERE id = ?
	`
	conn.Exec(sql, status, id)
	defer conn.Close()
}

func RemoveFromDB(conn *sql.DB, id int) {
	sql := `
DELETE FROM tasks WHERE id = ?
	`
	conn.Exec(sql, id)
	defer conn.Close()
}

func GetTasksFromDB(conn *sql.DB) []Task {
	sql := `
SELECT id, task, is_done, day FROM tasks WHERE is_done = 0
	`
	rows, err := conn.Query(sql)

	if err != nil {
		log.Fatal(err)
	}

	var tasks []Task

	for rows.Next() {
		var id int
		var task string
		var is_done bool
		var day int

		err = rows.Scan(&id, &task, &is_done, &day)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, Task{id: id, task: task, is_done: is_done, day: day})
	}

	return tasks
}

func GetFinishedFromDB(conn *sql.DB) []Task {
	sql := `
SELECT id, task, is_done, day FROM tasks WHERE is_done = 1
	`
	rows, err := conn.Query(sql)

	if err != nil {
		log.Fatal(err)
	}

	var tasks []Task

	for rows.Next() {
		var id int
		var task string
		var is_done bool
		var day int

		err = rows.Scan(&id, &task, &is_done, &day)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, Task{id: id, task: task, is_done: is_done, day: day})
	}

	return tasks
}
