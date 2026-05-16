import Task from "./Task";

function TaskList({ tasks }) {
    return (
        <ul>
            {tasks?.map((task) => (
                <li key={task.id}>
                    <Task task={task} />
                </li>
            ))}
        </ul>
    )
}

export default TaskList