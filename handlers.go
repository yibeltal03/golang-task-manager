package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

// Sample Task struct and in-memory task storage
// type Task struct {
//     ID   int    `json:"id"`
//     Name string `json:"name"`
//     Done bool   `json:"done"`
// }

// var tasks = []Task{}
// var nextID = 1

// getTasks retrieves all tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

// createTask adds a new task
func createTask(w http.ResponseWriter, r *http.Request) {
    var task Task
    json.NewDecoder(r.Body).Decode(&task)
    task.ID = nextID
    nextID++
    tasks = append(tasks, task)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

// deleteTask removes a task by ID
func deleteTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    for i, task := range tasks {
        if task.ID == id {
            tasks = append(tasks[:i], tasks[i+1:]...)
            break
        }
    }
    w.WriteHeader(http.StatusNoContent)
}
