import Task from "./Task";
import '../styles/TaskList.css'

function TaskList({ tasks, onToggleComplete, onEdit, onDelete }) {
    if (!tasks) {
        return <div className="task-list-loading">Loading tasks...</div>
    }

    if (tasks.length === 0) {
        return <div className="task-list-empty">No tasks yet. Create one to get started!</div>
    }

    return (
        <div className="task-list">
            {tasks.map((task) => (
                <Task
                    key={task.id}
                    task={task}
                    onToggleComplete={onToggleComplete}
                    onEdit={onEdit}
                    onDelete={onDelete}
                />
            ))}
        </div>
    )
}

export default TaskList