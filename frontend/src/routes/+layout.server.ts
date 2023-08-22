import type { LayoutServerLoad } from './$types'
import { API_URL } from '$env/static/private'
import { redirect } from '@sveltejs/kit'

export const load: LayoutServerLoad = async ({ cookies, url }) => {
    // check if there is access token
    const accessToken = cookies.get('access_token')
    const refreshToken = cookies.get('refresh_token')
    if (!accessToken) return {
        user: {
            id: 0,
            name: "",
            email: "",
        }
    }

    const res = await fetch(`${API_URL}/checkauth`, {
        headers: {
            Authorization: `Bearer ${accessToken}`
        }
    })

    if (res.ok) {
        const data = await res.json()

        return {
            user: {
                id: data.id,
                name: data.username,
                email: "user@test.loc"
            }
        }
    }

    console.log(res.status)

    if (res.status === 401 && refreshToken) {
        const res = await fetch(`${API_URL}/refresh`, {
            method: 'POST',
            headers: {
                Authorization: `Bearer ${refreshToken}`
            }
        })

        if (res.status === 200) {
            const { access_token } = await res.json()
            cookies.set('access_token', access_token, { path: '/' })
            throw redirect(307, url.pathname)
        }
    }

    return {
        user: {
            id: 0,
            name: "",
            email: "",
        }
    }
}