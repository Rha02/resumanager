<script lang="ts">
	import type { ActionData } from "./$types";

    export let form: ActionData
</script>

<form method="post">
    {#if form?.emailMissing}
        <p>
            {form.msg}
        </p>
    {/if}
    <label for="email">Email</label>
    <input type="email" name="email" required>
    <br>
    <label for="username">Username</label>
    <input type="text" name="username" required>
    <br>
    {#if form?.passwordMismatch}
        <p>
            {form.msg}
        </p>
    {/if}
    <label for="password">Password</label>
    <input type="password" name="password" required>
    <br>
    <label for="confirm-password">Confirm Password</label>
    <input type="password" name="confirm_password" required>
    <br>
    <button type="submit">Sign Up</button>
</form>