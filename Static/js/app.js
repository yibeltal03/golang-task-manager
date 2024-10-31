// Fetch and display tasks on load
document.addEventListener("DOMContentLoaded", () => {
    loadTasks();
});

async function loadTasks() {
    const response = await fetch("/api/tasks");
    const tasks = await response.json();
    const taskList = document.getElementById("taskList");

    taskList.innerHTML = ""; // Clear list
    tasks.forEach(task => {
        const listItem = document.createElement("li");
        listItem.className = "task-item";
        listItem.innerHTML = `
            <span>${task.name}</span>
            <button onclick="deleteTask(${task.id})">Delete</button>
        `;
        taskList.appendChild(listItem);
    });
}

async function addTask() {
    const taskNameInput = document.getElementById("taskName");
    const taskName = taskNameInput.value.trim();
    if (!taskName) return;

    await fetch("/api/tasks", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ name: taskName, done: false })
    });

    taskNameInput.value = "";
    loadTasks(); // Refresh task list
}

async function deleteTask(id) {
    await fetch(`/api/tasks/${id}`, {
        method: "DELETE"
    });
    loadTasks(); // Refresh task list
}
