<script lang="ts">
    import {Navbar, Footer} from "$lib/components"
	import { user } from "../store";
	import type { LayoutData } from "./$types";

    export let data: LayoutData;

    $: user.set({
        id: data.user.id,
        name: data.user.name,
        email: data.user.email,
    });
</script>

<Navbar />
<main>
    <slot />
</main>
<Footer />