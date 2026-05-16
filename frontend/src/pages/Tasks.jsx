import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { apiCall } from "../api/client";

function Tasks() {
    const [tasks, setTasks] = useState(null)
    const navigate = useNavigate()

    useEffect(() => {
        const token = localStorage.getItem('token')

        if (!token) {
            navigate('/login')
            return
        }

        apiCall()
            .then(res => {
                if (res.status === 401) {
                    localStorage.removeItem('token')
                    navigate('/login')
                    return null
                }
                return res.json()
            })
            .then(json => {
                if (json) setTasks(json)
            })
    }, [navigate])

    const handleLogout = () => {
        localStorage.removeItem('token')
        navigate('/login')
    }

    return (
        <div>
            <button onClick={handleLogout}>Logout</button>
            <ul>
                {tasks?.map((task) => (
                    <li key={task.id}>
                        <p>{task.title}</p>
                        <p>{task.completed ? 'completed' : 'uncompleted'}</p>
                    </li>
                ))}
            </ul>
        </div>
    )
}

export default Tasks