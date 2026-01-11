<script lang="ts">
    import "./layout.css";
    import favicon from "$lib/assets/favicon.svg";
    import { Toaster } from "$lib/components/ui/sonner";
    import { AuthService } from "$lib/api-client";
    import { user } from "$lib/store/user.svelte";
    import { onMount } from "svelte";
    import Navbar from "$lib/components/custom/navbar.svelte";
    import { ModeWatcher } from "mode-watcher";
    import { fly, fade } from "svelte/transition";
    import { cubicOut } from "svelte/easing";
    import { page } from "$app/state";

    let { children } = $props();

    onMount(async () => {
        user.isLoading = true;

        try {
            user.current = await AuthService.getUsersMe();
        } catch {
            user.current = null;
        } finally {
            user.isLoading = false;
        }
    });
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<Navbar />
<Toaster />
<ModeWatcher />

<div class="relative min-h-full overflow-x-hidden">
    {#key page.url.pathname}
        <div
            in:fly={{ x: 20, duration: 300, delay: 150, easing: cubicOut }}
            out:fade={{ duration: 150 }}
            class="w-full"
        >
            {@render children()}
        </div>
    {/key}
</div>
