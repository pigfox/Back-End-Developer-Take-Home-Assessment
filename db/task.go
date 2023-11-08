package db

import (
	"bedtha/config"
	"bedtha/structs"
	"fmt"
	"strconv"
)

func Create(task structs.Task) (int, error) {
	db, err := conn()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	statement, err := db.Prepare("INSERT INTO tasks (title, description, due_date, status) VALUES (?,?,?,?)")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	res, err := statement.Exec(task.Title, task.Description, task.DueDate, task.Status)
	if err != nil {
		return 0, err
	}

	LastInsertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(LastInsertId), nil
}

func Update(task structs.Task, taskID int) (int, error) {
	db, err := conn()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	statement, err := db.Prepare("UPDATE tasks SET title = ?, description = ?, due_date = ?, status = ? WHERE id = ?")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	res, err := statement.Exec(task.Title, task.Description, task.DueDate, task.Status, taskID)
	if err != nil {
		return 0, err
	}

	RowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(RowsAffected), nil
}

func View(taskID int) ([]structs.Task, error) {
	db, err := conn()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, title, description, due_date, status FROM tasks WHERE id=?", taskID)
	if err != nil {
		return nil, err
	}
	var task structs.Task
	var tasks []structs.Task
	for rows.Next() {
		rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func List(page int) ([]structs.Task, error) {
	db, err := conn()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT id, title, description, due_date, status FROM tasks LIMIT ?," + strconv.Itoa(config.RowsPerPage*page)
	rows, err := db.Query(query, page-1)
	if err != nil {
		return nil, err
	}
	var task structs.Task
	var tasks []structs.Task
	for rows.Next() {
		rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.Status)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func Delete(taskID int) (int, error) {
	db, err := conn()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	statement, err := db.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	res, err := statement.Exec(taskID)
	if err != nil {
		return 0, err
	}

	RowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(RowsAffected), nil
}
