import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { register } from "../api/auth";

function Register() {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [error, setError] = useState(null)
    const navigate = useNavigate()

    const handleSubmit = async (e) => {
        e.preventDefault()
        setError(null)

        const res = await register(email, password)

        const data = await res.json()

        if (!res.ok) {
            setError(data.error)
            return
        }

        navigate('/login')
    }

    return (
        <div>
            <h1>Registration</h1>
            {error && <p style={{ color: 'red' }}>{error}</p>}
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
                <button type="submit">Register</button>
            </form>
            <button onClick={() => navigate('/login')}>Already have an account? Login</button>
        </div>
    )
}

export default Register