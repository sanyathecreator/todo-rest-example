function Task({ task, onToggleComplete, onEdit, onDelete }) {
    return (
        <div style={{ display: 'flex', gap: '10px', padding: '10px', border: '1px solid #ddd', marginBottom: '10px' }}>
            <input
                type="checkbox"
                checked={task.completed}
                onChange={() => onToggleComplete(task.id)}
            />
            <div style={{ flex: 1 }}>
                <p style={{ fontWeight: 'bold', margin: '0' }}>{task.title}</p>
                <p style={{ margin: '0', color: '#666' }}>{task.description}</p>
                <small>{task.completed ? '✓ Completed' : 'Pending'}</small>
            </div>
            <button type="button" onClick={() => onEdit(task)}>Edit</button>
            <button type="button" onClick={() => onDelete(task.id)}>Delete</button>
        </div>
    )
}

export default Task