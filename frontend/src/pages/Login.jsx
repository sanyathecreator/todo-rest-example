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
            setError(data.error || 'Login failed')
            return
        }

        localStorage.setItem('token', data.token)
        navigate('/tasks')
    }

    return (
        <div style={{ maxWidth: '400px', margin: '50px auto', padding: '20px' }}>
            <h1>Login</h1>
            {error && <p style={{ color: 'red' }}>{error}</p>}

            <form onSubmit={handleSubmit}>
                <input
                    type="email"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                    style={{ width: '100%', padding: '10px', marginBottom: '10px', boxSizing: 'border-box' }}
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    required
                    style={{ width: '100%', padding: '10px', marginBottom: '10px', boxSizing: 'border-box' }}
                />
                <button type="submit" style={{ width: '100%', padding: '10px' }}>Login</button>
            </form>
            <p>Don't have an account? <a href="/register">Register</a></p>
        </div>
    )
}

export default Login