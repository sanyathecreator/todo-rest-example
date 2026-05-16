import { useState, useEffect } from 'react'
import '../styles/Modal.css'

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
        <div className="modal-overlay">
            <div className="modal-content">
                <h2 className="modal-header">Edit Task</h2>
                {error && <div className="modal-error">{error}</div>}

                <form className="modal-form" onSubmit={handleSubmit}>
                    <input
                        type="text"
                        placeholder="Task title"
                        value={title}
                        onChange={(e) => setTitle(e.target.value)}
                        required
                    />
                    <textarea
                        placeholder="Task description"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                    />
                    <label>
                        <input
                            type="checkbox"
                            checked={completed}
                            onChange={(e) => setCompleted(e.target.checked)}
                        />
                        Mark as completed
                    </label>

                    <div className="modal-buttons">
                        <button type="button" className="modal-btn-cancel" onClick={onClose}>Cancel</button>
                        <button type="submit" className="modal-btn-primary">Update</button>
                    </div>
                </form>
            </div>
        </div>
    )
}

export default EditTaskModal