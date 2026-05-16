import '../styles/Task.css'

function Task({ task, onToggleComplete, onEdit, onDelete }) {
    return (
        <div className="task">
            <input
                type="checkbox"
                checked={task.completed}
                onChange={() => onToggleComplete(task.id)}
            />
            <div className="task-info">
                <div className="task-title">{task.title}</div>
                <div className="task-description">{task.description}</div>
                <div className="task-status">{task.completed ? '✓ Completed' : 'Pending'}</div>
            </div>
            <div className="task-buttons">
                <button className="task-btn-edit" onClick={() => onEdit(task)}>Edit</button>
                <button className="task-btn-delete" onClick={() => onDelete(task.id)}>Delete</button>
            </div>
        </div>
    )
}

export default Task