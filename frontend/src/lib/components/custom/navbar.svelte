<script lang="ts">
    import { Button, buttonVariants } from "$lib/components/ui/button";
    import * as Avatar from "$lib/components/ui/avatar/index.js";
    import { user } from "$lib/store/user.svelte";
    import Menu from "@lucide/svelte/icons/menu";
    import X from "@lucide/svelte/icons/x";
    import User from "@lucide/svelte/icons/user";
    import LogOut from "@lucide/svelte/icons/log-out";
    import Home from "@lucide/svelte/icons/home";
    import TrendingUp from "@lucide/svelte/icons/trending-up";
    import Bell from "@lucide/svelte/icons/bell";
    import Crown from "@lucide/svelte/icons/crown";
    import MessageCircleMore from "@lucide/svelte/icons/message-circle-more";
    import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
    import * as Sheet from "$lib/components/ui/sheet";
    import { AuthService } from "$lib/api-client";
    import { toast } from "svelte-sonner";
    import { cn, isActiveUrl } from "$lib/utils";
    import { resolve } from "$app/paths";
    import { goto } from "$app/navigation";
    import DarkModeToggler from "./dark-mode-toggler.svelte";

    let mobileMenuOpen = $state(false);
    let navLinksRefs: HTMLAnchorElement[] = [];
    let indicatorLeft = $state(0);
    let indicatorWidth = $state(0);

    const navLinks = [
        { name: "Home", href: "/", icon: Home },
        { name: "Explore", href: "/explore", icon: TrendingUp },
        { name: "Notifications", href: "/notifications", icon: Bell },
        { name: "Rankings", href: "/rankings", icon: Crown },
        { name: "Messaging", href: "/messaging", icon: MessageCircleMore },
    ];

    function updateIndicator() {
        const activeIndex = navLinks.findIndex((link) =>
            isActiveUrl(link.href),
        );
        if (activeIndex !== -1 && navLinksRefs[activeIndex]) {
            const activeLink = navLinksRefs[activeIndex];
            if (activeLink) {
                indicatorLeft = activeLink.offsetLeft;
                indicatorWidth = activeLink.offsetWidth;
            }
        }
    }

    $effect(() => {
        navLinks.forEach((link) => isActiveUrl(link.href));
        updateIndicator();
    });

    async function handleLogout() {
        try {
            await AuthService.postAuthLogout();
            user.current = null;
            toast.success("Logged out successfully");
            window.location.href = "/login";
        } catch {
            toast.error("Failed to logout");
        }
    }
</script>

<!-- eslint-disable svelte/no-navigation-without-resolve -->

<nav
    class="fixed top-0 left-0 right-0 z-50 border-b bg-background/95 backdrop-blur supports-backdrop-filter:bg-background/60"
>
    <div class="container mx-auto px-4">
        <div class="flex h-14 items-center justify-between">
            <a href={resolve("/")} class="flex items-center gap-2">
                <div
                    class="flex h-8 w-8 items-center justify-center rounded-lg bg-primary text-primary-foreground"
                >
                    <span class="text-lg font-bold">B</span>
                </div>
                <span class="hidden text-lg font-semibold sm:inline-block"
                    >Build In Public</span
                >
            </a>
            <div class="relative hidden items-center gap-1 md:flex">
                {#each navLinks as link, i (link.href)}
                    {@const Icon = link.icon}
                    <a
                        bind:this={navLinksRefs[i]}
                        href={link.href}
                        class={cn(
                            "relative flex flex-col items-center gap-1 rounded-md px-3 py-1 text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground",
                            isActiveUrl(link.href)
                                ? "text-foreground"
                                : "text-muted-foreground",
                        )}
                    >
                        <Icon class="h-5 w-5" strokeWidth={1.5} />
                        <span class="text-xs">{link.name}</span>
                    </a>
                {/each}
                {#if indicatorWidth > 0}
                    <span
                        class="absolute bottom-0 h-0.5 bg-primary transition-all duration-300 ease-out"
                        style="left: {indicatorLeft}px; width: {indicatorWidth}px; transform: translateZ(0);"
                    ></span>
                {/if}
            </div>

            <div class="flex items-center gap-3">
                <DarkModeToggler />
                {#if user.isLoading}
                    <div
                        class="h-8 w-8 animate-pulse rounded-full bg-muted"
                    ></div>
                {:else if user.isAuthenticated}
                    <DropdownMenu.Root>
                        <DropdownMenu.Trigger>
                            <Avatar.Root>
                                <Avatar.Fallback>
                                    {user.current?.name
                                        ?.charAt(0)
                                        .toUpperCase() || "U"}
                                </Avatar.Fallback>
                            </Avatar.Root>
                        </DropdownMenu.Trigger>
                        <DropdownMenu.Content
                            class="w-56"
                            align="end"
                            sideOffset={8}
                        >
                            <DropdownMenu.Label class="font-normal">
                                <div class="flex flex-col space-y-1">
                                    <p class="text-sm font-medium leading-none">
                                        {user.current?.name || "User"}
                                    </p>
                                    <p
                                        class="text-xs leading-none text-muted-foreground"
                                    >
                                        {user.current?.email || ""}
                                    </p>
                                </div>
                            </DropdownMenu.Label>
                            <DropdownMenu.Separator />
                            <DropdownMenu.Item
                                onclick={() => goto(resolve("/profile"))}
                            >
                                <User class="mr-2 h-4 w-4" />
                                <span>Profile</span>
                            </DropdownMenu.Item>
                            <DropdownMenu.Separator />
                            <DropdownMenu.Item
                                onclick={handleLogout}
                                class="text-red-600 focus:text-red-600"
                            >
                                <LogOut class="mr-2 h-4 w-4" />
                                <span>Log out</span>
                            </DropdownMenu.Item>
                        </DropdownMenu.Content>
                    </DropdownMenu.Root>
                {:else}
                    <div class="hidden items-center gap-2 sm:flex">
                        <Button href="/login" variant="ghost" size="sm">
                            Log in
                        </Button>
                        <Button href="/signup" size="sm">Sign up</Button>
                    </div>
                {/if}

                <Sheet.Root bind:open={mobileMenuOpen}>
                    <Sheet.Trigger
                        class={cn(
                            "md:hidden",
                            buttonVariants({ variant: "ghost", size: "icon" }),
                        )}
                    >
                        {#if mobileMenuOpen}
                            <X class="h-5 w-5" />
                        {:else}
                            <Menu class="h-5 w-5" />
                        {/if}
                        <span class="sr-only">Toggle menu</span>
                    </Sheet.Trigger>
                    <Sheet.Content side="right" class="w-75 sm:w-100">
                        <Sheet.Header>
                            <Sheet.Title>Menu</Sheet.Title>
                        </Sheet.Header>
                        <div class="mt-6 flex flex-col gap-4">
                            {#each navLinks as link (link.href)}
                                {@const Icon = link.icon}
                                <a
                                    href={link.href}
                                    onclick={() => (mobileMenuOpen = false)}
                                    class={cn(
                                        "flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors hover:bg-accent",
                                        isActiveUrl(link.href)
                                            ? "bg-accent text-accent-foreground"
                                            : "text-muted-foreground",
                                    )}
                                >
                                    <Icon class="h-5 w-5" />
                                    {link.name}
                                </a>
                            {/each}

                            {#if !user.isAuthenticated && !user.isLoading}
                                <div
                                    class="mt-4 flex flex-col gap-2 border-t pt-4"
                                >
                                    <Button
                                        href="/login"
                                        variant="outline"
                                        class="w-full"
                                    >
                                        Log in
                                    </Button>
                                    <Button href="/signup" class="w-full">
                                        Sign up
                                    </Button>
                                </div>
                            {/if}
                        </div>
                    </Sheet.Content>
                </Sheet.Root>
            </div>
        </div>
    </div>
</nav>

<div class="h-16"></div>
