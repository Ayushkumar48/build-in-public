<script lang="ts">
    import { goto } from "$app/navigation";
    import { resolve } from "$app/paths";
    import { AuthService } from "$lib/api-client";
    import { onMount } from "svelte";

    onMount(async () => {
        try {
            const user = await AuthService.getUsersMe();
            if (user) {
                goto(resolve("/"));
            }
        } catch (error) {
            console.error(error);
        }
    });
    let { children } = $props();
</script>

{@render children()}
