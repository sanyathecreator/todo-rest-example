import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { login } from "../api/auth"
import '../styles/Auth.css'

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
        <div className="auth-page">
            <div className="auth-container">
                <h1 className="auth-title">Login</h1>
                {error && <div className="auth-error">{error}</div>}

                <form className="auth-form" onSubmit={handleSubmit}>
                    <input
                        type="email"
                        placeholder="Email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                    <input
                        type="password"
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                    <button type="submit" className="auth-btn-submit">Login</button>
                </form>

                <div className="auth-footer">
                    Don't have an account? <a href="/register">Register</a>
                </div>
            </div>
        </div>
    )
}

export default Login