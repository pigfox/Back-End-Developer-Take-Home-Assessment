package handlers

import (
	"bedtha/db"
	"bedtha/structs"
	"bedtha/utils"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func Create(w http.ResponseWriter, req *http.Request) {
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}
	var task structs.Task
	err = json.Unmarshal(reqBody, &task)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}

	if task.Title == "" || task.Description == "" || task.Status == "" {
		data := structs.Response{Data: "all fields required"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(data)
		return
	}

	if !utils.IsValidDate(task.DueDate) {
		data := structs.Response{Data: "invalid date format, need YYYY-MM-DD"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(data)
		return
	}

	lastInsertID, err := db.Create(task)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	} else {
		data := structs.Response{Data: "Last insert id:" + strconv.Itoa(lastInsertID)}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(data)
		return
	}
}

func Update(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	taskId, err := strconv.Atoi(parts[2])
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}

	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}
	var task structs.Task
	err = json.Unmarshal(reqBody, &task)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}

	if task.Title == "" || task.Description == "" || task.Status == "" {
		data := structs.Response{Data: "all fields required"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(data)
		return
	}

	if !utils.IsValidDate(task.DueDate) {
		data := structs.Response{Data: "invalid date format, YYYY-MM-DD required"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(data)
		return
	}

	affectedRows, err := db.Update(task, taskId)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}

	if affectedRows == 1 {
		data := structs.Response{Data: "Number of affected rows:" + strconv.Itoa(affectedRows)}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(data)
		return
	}
}

func Delete(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	taskId, err := strconv.Atoi(parts[2])
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
		return
	}

	affectedRows, err := db.Delete(taskId)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
		return
	}

	data := structs.Response{Data: "affected rows: " + strconv.Itoa(affectedRows)}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	return
}

func List(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	pageNumber, err := strconv.Atoi(parts[2])
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}

	tasks, err := db.List(pageNumber)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(tasks)
	return
}

func View(w http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	taskId, err := strconv.Atoi(parts[2])
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}

	tasks, err := db.View(taskId)
	if err != nil {
		data := structs.Response{Data: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(data)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if len(tasks) == 1 {
		json.NewEncoder(w).Encode(tasks[0])
		return
	}

}
