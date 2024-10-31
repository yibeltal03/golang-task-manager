# golang-task-manager

# Instructions
go mod init task-manager
go get -u github.com/gorilla/mux

# Running app Method 1
go run main.go handlers.go models.go middleware.go

# Running app Method 2
go build
./task-manager


# Test endpoints
Create a Task:
`curl -X POST -H "Content-Type: application/json" -d '{"name": "Learn Go"} http://localhost:8081/api/tasks`
Get All Tasks:
`curl http://localhost:8081/api/tasks`
Delete a Task:
`curl -X DELETE http://localhost:8081/api/tasks/{id}`



