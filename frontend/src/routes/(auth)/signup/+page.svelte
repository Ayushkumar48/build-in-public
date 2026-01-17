<script lang="ts">
    import { resolve } from "$app/paths";
    import { AuthService, ApiError } from "$lib/api-client";
    import CompanyLogo from "$lib/assets/company-logo.svelte";
    import GithubLogo from "$lib/assets/github-logo.svelte";
    import GoogleIcon from "$lib/assets/google-icon.svelte";
    import LinkedinLogo from "$lib/assets/linkedin-logo.svelte";
    import MicrosoftIcon from "$lib/assets/microsoft-icon.svelte";
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";
    import { loginWithOAuth } from "$lib/utils";
    import { toast } from "svelte-sonner";

    let signupData = $state({
        firstName: "",
        lastName: "",
        email: "",
        password: "",
    });

    async function signup() {
        try {
            const data = await AuthService.postAuthSignup({
                firstName: signupData.firstName,
                lastName: signupData.lastName,
                email: signupData.email,
                password: signupData.password,
            });

            toast.success(
                `Welcome, ${data.first_name || "User"} ${data.last_name || ""}! Account created successfully.`,
            );
            window.location.href = "/";
        } catch (error) {
            if (error instanceof ApiError) {
                const errorBody = error.body as { error?: string };
                toast.error(errorBody?.error || "Signup failed");
            } else {
                toast.error("Signup failed");
            }
            console.error(error);
        }
    }
</script>

<section
    class="flex min-h-[calc(100vh-64px)] bg-zinc-50 px-4 py-16 md:py-32 dark:bg-transparent"
>
    <form
        onsubmit={signup}
        class="bg-card m-auto h-fit w-full max-w-lg rounded-[calc(var(--radius)+.125rem)] border p-0.5 shadow-md dark:[--color-muted:var(--color-zinc-900)]"
    >
        <div class="p-8 pb-6">
            <div>
                <a href={resolve("/")} aria-label="go home">
                    <CompanyLogo />
                </a>
                <h1 class="text-title mb-1 mt-4 text-xl font-semibold">
                    Create a Tailus UI Account
                </h1>
                <p class="text-sm">Welcome! Create an account to get started</p>
            </div>

            <div class="mt-6 grid grid-cols-2 gap-3">
                <Button
                    type="button"
                    variant="outline"
                    onclick={() => loginWithOAuth("google")}
                >
                    <GoogleIcon />
                    <span>Google</span>
                </Button>
                <Button
                    type="button"
                    variant="outline"
                    onclick={() => loginWithOAuth("microsoft")}
                >
                    <MicrosoftIcon />
                    <span>Microsoft</span>
                </Button>
                <Button
                    type="button"
                    variant="outline"
                    onclick={() => loginWithOAuth("github")}
                >
                    <GithubLogo />
                    <span>GitHub</span>
                </Button>
                <Button
                    type="button"
                    variant="outline"
                    onclick={() => loginWithOAuth("linkedin")}
                >
                    <LinkedinLogo />
                    <span>LinkedIn</span>
                </Button>
            </div>

            <hr class="my-4 border-dashed" />

            <div class="space-y-5">
                <div class="flex gap-x-2">
                    <div class="space-y-2 w-full">
                        <Label for="firstName" class="block text-sm"
                            >First Name</Label
                        >
                        <Input
                            type="text"
                            required
                            name="firstName"
                            id="firstName"
                            bind:value={signupData.firstName}
                        />
                    </div>
                    <div class="space-y-2 w-full">
                        <Label for="lastName" class="block text-sm"
                            >Last Name</Label
                        >
                        <Input
                            type="text"
                            required
                            name="lastName"
                            id="lastName"
                            bind:value={signupData.lastName}
                        />
                    </div>
                </div>

                <div class="space-y-2">
                    <Label for="email" class="block text-sm">Username</Label>
                    <Input
                        type="email"
                        required
                        name="email"
                        id="email"
                        bind:value={signupData.email}
                    />
                </div>

                <div class="space-y-2">
                    <Label for="pwd" class="text-title text-sm">Password</Label>
                    <Input
                        type="password"
                        required
                        name="pwd"
                        id="pwd"
                        class="input sz-md variant-mixed"
                        bind:value={signupData.password}
                    />
                </div>

                <Button class="w-full" type="submit">Create Account</Button>
            </div>
        </div>

        <div class="bg-muted rounded-(--radius) border p-3">
            <p class="text-accent-foreground text-center text-sm">
                Have an account ?
                <Button href="/login" variant="link" class="px-2"
                    >Sign in</Button
                >
            </p>
        </div>
    </form>
</section>
