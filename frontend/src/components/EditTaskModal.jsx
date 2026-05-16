import { useState, useEffect } from 'react'

function EditTaskModal({ task, isOpen, onClose, onTaskUpdated }) {
    const [title, setTitle] = useState('')
    const [description, setDescription] = useState('')
    const [completed, setCompleted] = useState(false)
    const [error, setError] = useState(null)

    useEffect(() => {
        if (task) {
            setTitle(task.title)
            setDescription(task.description)
            setCompleted(task.completed)
        }
    }, [task, isOpen])

    const handleSubmit = async (e) => {
        e.preventDefault()
        setError(null)

        try {
            await onTaskUpdated(task.id, { title, description, completed })
            onClose()
        } catch (err) {
            setError(err.message)
        }
    }

    if (!isOpen || !task) {
        return null
    }

    return (
        <div style={{
            position: 'fixed',
            inset: 0,
            backgroundColor: 'rgba(0, 0, 0, 0.5)',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            zIndex: 1000
        }}>
            <div style={{
                backgroundColor: 'white',
                padding: '20px',
                borderRadius: '8px',
                width: '90%',
                maxWidth: '400px'
            }}>
                <h2>Edit Task</h2>
                {error && <p style={{ color: 'red' }}>{error}</p>}

                <form onSubmit={handleSubmit}>
                    <input
                        type="text"
                        placeholder="Task title"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        required
                        style={{ width: '100%', padding: '8px', marginBottom: '10px', boxSizing: 'border-box' }}
                    />
                    <textarea
                        placeholder="Task description"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                        style={{ width: '100%', padding: '8px', marginBottom: '10px', boxSizing: 'border-box' }}
                    />
                    <label>
                        <input
                            type="checkbox"
                            checked={completed}
                            onChange={(e) => setCompleted(e.target.checked)}
                        />
                        Mark as completed
                    </label>

                    <div style={{ display: 'flex', gap: '10px', justifyContent: 'flex-end', marginTop: '15px' }}>
                        <button type="button" onClick={onClose}>Cancel</button>
                        <button type="submit">Update</button>
                    </div>
                </form>
            </div>
        </div>
    )
}

export default EditTaskModal