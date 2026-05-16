import { useState } from "react";
import '../styles/Modal.css'

function NewTaskModal({ isOpen, onClose, onTaskCreated }) {
    const [title, setTitle] = useState('')
    const [description, setDescription] = useState('')
    const [error, setError] = useState(null)

    const handleSubmit = async (e) => {
        e.preventDefault()
        setError(null)

        try {
            await onTaskCreated(title, description)
            setTitle('')
            setDescription('')
            onClose()
        } catch (err) {
            setError(err.message)
        }
    }

    if (!isOpen) {
        return null
    }

    return (
        <div className="modal-overlay">
            <div className="modal-content">
                <h2 className="modal-header">Create New Task</h2>
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
                        placeholder="Task description (optional)"
                        value={description}
                        onChange={(e) => setDescription(e.target.value)}
                    />
                    <div className="modal-buttons">
                        <button type="button" className="modal-btn-cancel" onClick={onClose}>Cancel</button>
                        <button type="submit" className="modal-btn-primary">Create</button>
                    </div>
                </form>
            </div>
        </div>
    )
}

export default NewTaskModal