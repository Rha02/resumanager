<script lang="ts">
    import {Navbar, Footer} from "$lib/components"
	import { user } from "../store";
	import type { LayoutData } from "./$types";

    export let data: LayoutData;
</script>

<Navbar />
<main>
    <slot />
</main>
<Footer />