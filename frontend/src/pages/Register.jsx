import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { register } from "../api/auth";

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
        <div style={{ maxWidth: '400px', margin: '50px auto', padding: '20px' }}>
            <h1>Register</h1>
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
                <input
                    type="password"
                    placeholder="Confirm Password"
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                    required
                    style={{ width: '100%', padding: '10px', marginBottom: '10px', boxSizing: 'border-box' }}
                />
                <button type="submit" style={{ width: '100%', padding: '10px' }}>Register</button>
            </form>
            <p>Already have an account? <a href="/login">Login</a></p>
        </div>
    )
}

export default Register