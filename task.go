package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	tasks   = []Task{}
	nextID  = 1
	tasksMu sync.Mutex
)

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasksMu.Lock()
	defer tasksMu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasksMu.Lock()
	defer tasksMu.Unlock()

	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.ID = nextID
	nextID++
	tasks = append(tasks, task)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func removeTaskHandler(w http.ResponseWriter, r *http.Request) {
	tasksMu.Lock()
	defer tasksMu.Unlock()

	var id struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, task := range tasks {
		if task.ID == id.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
}
