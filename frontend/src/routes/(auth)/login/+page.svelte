<script lang="ts">
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input";
    import { AuthService, ApiError } from "$lib/api-client";
    import { toast } from "svelte-sonner";
    import { Label } from "$lib/components/ui/label";
    import GoogleIcon from "$lib/assets/google-icon.svelte";
    import MicrosoftIcon from "$lib/assets/microsoft-icon.svelte";
    import CompanyLogo from "$lib/assets/company-logo.svelte";
    import { resolve } from "$app/paths";

    let loginData = $state({
        email: "",
        password: "",
    });

    async function login() {
        try {
            const data = await AuthService.postAuthLogin({
                email: loginData.email,
                password: loginData.password,
            });

            toast.success(data.success || "Logged In");
            window.location.href = "/";
        } catch (error) {
            if (error instanceof ApiError) {
                const errorBody = error.body as { error?: string };
                toast.error(errorBody?.error || "Login failed");
            } else {
                toast.error("Login failed");
            }
            console.error(error);
        }
    }
</script>

<section
    class="flex min-h-screen bg-zinc-50 px-4 py-16 md:py-32 dark:bg-transparent"
>
    <form
        action=""
        class="bg-card m-auto h-fit w-full max-w-sm rounded-[calc(var(--radius)+.125rem)] border p-0.5 shadow-md dark:[--color-muted:var(--color-zinc-900)]"
    >
        <div class="p-8 pb-6">
            <div>
                <a href={resolve("/")} aria-label="go home">
                    <CompanyLogo />
                </a>
                <h1 class="mb-1 mt-4 text-xl font-semibold">
                    Sign In to Tailus UI
                </h1>
                <p class="text-sm">Welcome back! Sign in to continue</p>
            </div>

            <div class="mt-6 grid grid-cols-2 gap-3">
                <Button type="button" variant="outline">
                    <GoogleIcon />
                    <span>Google</span>
                </Button>
                <Button type="button" variant="outline">
                    <MicrosoftIcon />
                    <span>Microsoft</span>
                </Button>
            </div>

            <hr class="my-4 border-dashed" />

            <div class="space-y-6">
                <div class="space-y-2">
                    <Label for="email" class="block text-sm">Email</Label>
                    <Input
                        type="email"
                        required
                        name="email"
                        id="email"
                        bind:value={loginData.email}
                    />
                </div>

                <div class="space-y-0.5">
                    <div class="flex items-center justify-between">
                        <Label for="pwd" class="text-title text-sm"
                            >Password</Label
                        >
                        <Button
                            variant="link"
                            href="#"
                            size="sm"
                            class="link intent-info variant-ghost text-sm"
                        >
                            Forgot your Password ?
                        </Button>
                    </div>
                    <Input
                        type="password"
                        required
                        name="pwd"
                        id="pwd"
                        class="input sz-md variant-mixed"
                        bind:value={loginData.password}
                    />
                </div>

                <Button class="w-full" onclick={login}>Sign In</Button>
            </div>
        </div>

        <div class="bg-muted rounded-(--radius) border p-3">
            <p class="text-accent-foreground text-center text-sm">
                Don't have an account ?
                <Button href="/signup" variant="link" class="px-2"
                    >Create account</Button
                >
            </p>
        </div>
    </form>
</section>
