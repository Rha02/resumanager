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

        // TODO: Validate email
        if (!email) {
            return fail(400, { emailMissing: true, msg: "Email is required" })
        }

        // TODO: Register user
        const response = await fetch(`${API_URL}/register`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email,
                username,
                password
            })
        })
        
        const resBody = await response.json()

        if (!response.ok) {
            return fail(response.status, resBody)
        }



        cookies.set("access_token", "test")
        cookies.set("refresh_token", "test")
        throw redirect(303, "/dashboard")
    }
} satisfies Actions