import Task from "./Task";

function TaskList({ tasks, onToggleComplete, onEdit, onDelete }) {
    if (!tasks) {
        return <p>Loading tasks...</p>
    }

    if (tasks.length === 0) {
        return <p>No tasks yet. Create one to get started!</p>
    }

    return (
        <div>
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