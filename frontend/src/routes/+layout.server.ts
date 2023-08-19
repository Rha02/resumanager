import type { User } from '$lib/types'
import type { LayoutServerLoad } from './$types'

export const load: LayoutServerLoad = async () => {
    // load user
    return {
        user: null
    }
}