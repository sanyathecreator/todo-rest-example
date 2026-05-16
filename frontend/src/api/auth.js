import { apiCall } from "./client";

export async function login(email, password) {
    const res = await apiCall('/login', {
        method: 'POST',
        body: JSON.stringify({email, password})
    })
    return res
}

export async function register(email, password) {
    const res = await apiCall('/register', {
        method: 'POST',
        body: JSON.stringify({email, password})
    })
    return res
}