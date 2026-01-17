<script lang="ts">
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import { resolve } from "$app/paths";
    import Loader from "@lucide/svelte/icons/loader";

    onMount(() => {
        goto(resolve("/"), { replaceState: true });
    });
</script>

<div class="flex min-h-screen items-center justify-center">
    <div class="flex flex-col items-center text-center">
        <Loader class="h-8 w-8 animate-spin mb-4" />
        <p class="text-muted-foreground">Completing sign in...</p>
    </div>
</div>
