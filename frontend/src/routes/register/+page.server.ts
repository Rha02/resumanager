import { redirect, fail, type Actions } from '@sveltejs/kit'

export const actions = {
    default: async ({request, cookies}) => {
        const data = await request.formData()
        const email = data.get("email") as string
        const password = data.get("password") as string
        const confirmPassword = data.get("confirm_password") as string

        // Validate password
        if (password !== confirmPassword) {
            return fail(400, { passwordMismatch: true })
        }

        // TODO: Validate email
        if (!email) {
            return fail(400, { emailMissing: true })
        }

        // TODO: Register user

        cookies.set("access_token", "test")
        cookies.set("refresh_token", "test")
        throw redirect(303, "/dashboard")
    }
} satisfies Actions