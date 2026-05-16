import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import TaskList from "../components/TaskList";
import { getTasks, createTask, deleteTask, updateTask } from "../api/tasks";
import EditTaskModal from "../components/EditTaskModal";
import NewTaskModal from "../components/NewTaskModal";
import '../styles/Tasks.css'

function Tasks() {
    const [tasks, setTasks] = useState(null)
    const [error, setError] = useState(null)
    const [showNewModal, setShowNewModal] = useState(false)
    const [editingTask, setEditingTask] = useState(null)
    const navigate = useNavigate()

    useEffect(() => {
        const token = localStorage.getItem('token')

        if (!token) {
            navigate('/login')
            return
        }

        loadTasks()
    }, [navigate])

    const loadTasks = async () => {
        try {
            const res = await getTasks()
            if (res.status === 401) {
                localStorage.removeItem('token')
                navigate('/login')
                return
            }
            const data = await res.json()
            setTasks(data || [])
        } catch (err) {
            setError('Failed to load tasks')
        }
    }

    const handleCreateTask = async (title, description) => {
        const res = await createTask(title, description)
        const data = await res.json()

        if (!res.ok) {
            throw new Error(data.error || 'Failed to create task')
        }

        setTasks([data, ...tasks])
    }

    const handleUpdateTask = async (taskId, updates) => {
        try {
            const res = await updateTask(taskId, updates)
            const data = await res.json()

            if (!res.ok) {
                throw new Error(data.error || 'Failed to update task')
            }

            setTasks(tasks.map(t => t.id === taskId ? data : t))
        } catch (err) {
            setError(err.message)
        }
    }

    const handleDeleteTask = async (taskId) => {
        try {
            const res = await deleteTask(taskId)

            if (!res.ok) {
                const data = await res.json()
                throw new Error(data.error || 'Failed to delete task')
            }

            setTasks(tasks.filter(t => t.id !== taskId))
        } catch (err) {
            setError(err.message)
        }
    }

    const handleToggleComplete = async (taskId) => {
        const task = tasks.find(t => t.id === taskId)
        await handleUpdateTask(taskId, { completed: !task.completed })
    }

    const handleLogout = () => {
        localStorage.removeItem('token')
        navigate('/login')
    }

    return (
        <div className="tasks-page">
            <div className="tasks-container">
                <div className="tasks-header">
                    <h1 className="tasks-title">My Tasks</h1>
                    <button className="tasks-btn-logout" onClick={handleLogout}>Logout</button>
                </div>
                {error && <div className="tasks-error">{error}</div>}

                <button className="tasks-btn-create" onClick={() => setShowNewModal(true)}>
                    + Create New Task
                </button>
                <TaskList
                    tasks={tasks}
                    onToggleComplete={handleToggleComplete}
                    onEdit={setEditingTask}
                    onDelete={handleDeleteTask}
                />
                <NewTaskModal
                    isOpen={showNewModal}
                    onClose={() => setShowNewModal(false)}
                    onTaskCreated={handleCreateTask}
                />
                <EditTaskModal
                    task={editingTask}
                    isOpen={editingTask !== null}
                    onClose={() => setEditingTask(null)}
                    onTaskUpdated={handleUpdateTask}
                />
            </div>
        </div>
    )
}

export default Tasks