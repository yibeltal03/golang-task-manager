package main

// Task represents a task model
type Task struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Done bool   `json:"done"`
}

// Initialize an empty slice of Task to hold the tasks
var tasks = []Task{}

// Next ID to assign to new tasks
var nextID = 1
