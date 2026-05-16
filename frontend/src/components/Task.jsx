function Task({ task, onToggleComplete, onEdit, onDelete }) {
    return (
        <>
            <input 
                type="checkbox" 
                checked={task.completed}
                onChange={() => onToggleComplete(task.id)}
            />
            <div>
                <p>{task.title}</p>
                <p>{task.description}</p>
            </div>
            <button type="button" onClick={() => onEdit(task)}>Edit</button>
            <button type="button" onClick={() => onDelete(task.id)}>Delete</button>
        </>
    )
}

export default Task