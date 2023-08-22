import { redirect, fail, type Actions } from '@sveltejs/kit'
import { API_URL } from '$env/static/private'

export const actions = {
    default: async ({ request, cookies, fetch }) => {
        const data = await request.formData()
        const email = data.get("email") as string
        const username = data.get("username") as string
        const password = data.get("password") as string
        const confirmPassword = data.get("confirm_password") as string

        // Validate password
        if (password !== confirmPassword) {
            return fail(400, { passwordMismatch: true, msg: "Passwords do not match" })
        }

        if (!email.match("^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$")) {
            return fail(400, { emailMissing: true, msg: "Email is required" })
        }

        const res = await fetch(`${API_URL}/register`, {
            method: 'POST',
            body: JSON.stringify({
                email,
                username,
                password
            })
        })

        if (res.ok) {
            const { access_token, refresh_token } = await res.json()
            cookies.set('access_token', access_token, { path: '/' })
            cookies.set('refresh_token', refresh_token, { path: '/' })
            throw redirect(303, '/dashboard')
        }

        return fail(400, { msg: "Something went wrong" })
    }
} satisfies Actions