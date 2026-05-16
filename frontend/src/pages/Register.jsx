import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { register } from "../api/auth";
import '../styles/Auth.css'

function Register() {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [confirmPassword, setConfirmPassword] = useState('')
    const [error, setError] = useState(null)
    const navigate = useNavigate()

    const handleSubmit = async (e) => {
        e.preventDefault()
        setError(null)

        if (password !== confirmPassword) {
            setError('Passwords do not match')
            return
        }

        const res = await register(email, password)
        const data = await res.json()

        if (!res.ok) {
            setError(data.error || 'Registration failed')
            return
        }

        navigate('/login')
    }

    return (
        <div className="auth-page">
            <div className="auth-container">
                <h1 className="auth-title">Register</h1>
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
                    <input
                        type="password"
                        placeholder="Confirm Password"
                        value={confirmPassword}
                        onChange={(e) => setConfirmPassword(e.target.value)}
                        required
                    />
                    <button type="submit" className="auth-btn-submit">Register</button>
                </form>

                <div className="auth-footer">
                    Already have an account? <a href="/login">Login</a>
                </div>
            </div>
        </div>
    )
}

export default Register