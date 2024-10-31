package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    // API routes
    router.HandleFunc("/api/tasks", withLock(getTasks)).Methods("GET")
    router.HandleFunc("/api/tasks", withLock(createTask)).Methods("POST")
    router.HandleFunc("/api/tasks/{id}", withLock(deleteTask)).Methods("DELETE")

    // Serve static files
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

    log.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8081", router))
}
