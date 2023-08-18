import type { User } from '$lib/types'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async () => {

    // load user
    return {
        user: {
            id: 1,
            name: "testuser",
            email: "test@user.loc"
        } as User
    }
}