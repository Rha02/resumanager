<script lang="ts">
	import type { ActionData } from "./$types";

    export let form: ActionData
</script>

<form method="post">
    {#if form?.emailMissing}
        <p>
            Must be a valid email address
        </p>
    {/if}
    <label for="email">Email</label>
    <input type="email" name="email">
    <br>
    {#if form?.passwordMismatch}
        <p>
            Passwords must match
        </p>
    {/if}
    <label for="password">Password</label>
    <input type="password" name="password">
    <br>
    <label for="confirm-password">Confirm Password</label>
    <input type="password" name="confirm_password">
    <br>
    <button type="submit">Sign Up</button>
</form>