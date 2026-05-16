import { apiCall } from "./client"

export async function getTasks(completed = null) {
    const endpoint = completed !== null ? `/tasks?completed=${completed}` : "/tasks"
    const res = await apiCall(endpoint)
    return res
}

export async function createTask(title, description) {
    const res = await apiCall('/tasks/create', {
        method: 'POST',
        body: JSON.stringify({ title, description })
    })
    return res
}

export async function updateTask(id, updates) {
    const res = await apiCall(`/tasks/${id}`, {
        method: 'PATCH',
        body: JSON.stringify(updates)
    })
    return res
}

export async function deleteTask(id) {
    const res = await apiCall(`/tasks/${id}`, {
        method: 'DELETE'
    })
    return res
}