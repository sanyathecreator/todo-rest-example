import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { login } from "../api/auth"

function Login() {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [error, setError] = useState(null)
    const navigate = useNavigate()

    const handleSubmit = async (e) => {
        e.preventDefault()
        setError(null)

        const res = await login(email, password)

        const data = await res.json()

        if (!res.ok) {
            setError(data.error)
            return
        }

        localStorage.setItem('token', data.token)
        navigate('/tasks')
    }

    return (
        <div>
            <h1>Login</h1>
            {error && <p style={{color: 'red'}}>{error}</p>}
            <form onSubmit={handleSubmit}>
                <input
                    type="email"
                    placeholder="Email"
                    value={email}
                    onChange={e => setEmail(e.target.value)}
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={e => setPassword(e.target.value)}
                />
                <button type="submit">Login</button>
            </form>
            <button onClick={() => navigate('/register')}>Don't have an account? Register</button>
        </div>
    )
}

export default Login