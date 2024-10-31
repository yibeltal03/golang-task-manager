document.addEventListener("DOMContentLoaded", () => {
    const taskForm = document.getElementById("taskForm");
    const taskList = document.getElementById("taskList");

    // Fetch and display tasks
    async function fetchTasks() {
        const response = await fetch("/api/tasks");
        const tasks = await response.json();
        taskList.innerHTML = "";
        tasks.forEach(task => {
            const li = document.createElement("li");
            li.innerHTML = `
                ${task.name}
                <button class="delete-button" data-id="${task.id}">Delete</button>
            `;
            taskList.appendChild(li);
        });
    }

    // Add a new task
    taskForm.addEventListener("submit", async (e) => {
        e.preventDefault();
        const taskName = document.getElementById("taskName").value;

        await fetch("/api/tasks", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ name: taskName, done: false })
        });

        document.getElementById("taskName").value = "";
        fetchTasks();
    });

    // Delete a task
    taskList.addEventListener("click", async (e) => {
        if (e.target.classList.contains("delete-button")) {
            const taskId = e.target.getAttribute("data-id");
            await fetch(`/api/tasks/${taskId}`, {
                method: "DELETE"
            });
            fetchTasks();
        }
    });

    fetchTasks();
});
