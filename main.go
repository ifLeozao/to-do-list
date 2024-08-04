package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", getTasksHandler)
	http.HandleFunc("/tasks/add", addTaskHandler)
	http.HandleFunc("/tasks/remove", removeTaskHandler)

	log.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
