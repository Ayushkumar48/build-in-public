<script lang="ts">
    import * as Card from "$lib/components/ui/card";
    import { Button } from "$lib/components/ui/button";
    import * as Avatar from "$lib/components/ui/avatar";
    import { Separator } from "$lib/components/ui/separator";
    import Heart from "@lucide/svelte/icons/heart";
    import MessageCircle from "@lucide/svelte/icons/message-circle";
    import Share2 from "@lucide/svelte/icons/share-2";
    import TrendingUp from "@lucide/svelte/icons/trending-up";
    import Users from "@lucide/svelte/icons/users";
    import Rocket from "@lucide/svelte/icons/rocket";
    import Quote from "@lucide/svelte/icons/quote";

    // Mock data for posts
    const posts = [
        {
            id: 1,
            author: {
                name: "Sarah Chen",
                username: "@sarahchen",
                avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah",
            },
            content:
                "Day 15 of building my SaaS: Finally shipped the authentication system! Spent 2 days debugging OAuth issues but learned so much about security in the process. The best part? Users can now sign up! 🚀",
            timestamp: "2h ago",
            likes: 47,
            comments: 12,
            tags: ["authentication", "saas", "milestone"],
        },
        {
            id: 2,
            author: {
                name: "Alex Rodriguez",
                username: "@alexdev",
                avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=Alex",
            },
            content:
                "Just hit $1K MRR! 🎉 Started this project 6 months ago as a side hustle. To everyone starting out: consistency beats perfection. Ship fast, learn faster, and don't be afraid to pivot.",
            timestamp: "5h ago",
            likes: 234,
            comments: 45,
            tags: ["milestone", "revenue", "motivation"],
        },
        {
            id: 3,
            author: {
                name: "Maya Patel",
                username: "@mayacodes",
                avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=Maya",
            },
            content:
                "Working on my landing page redesign. Here's what I learned: less is more. Removed 60% of the text and conversion rate went up by 40%. Sometimes we overthink things.",
            timestamp: "1d ago",
            likes: 89,
            comments: 23,
            tags: ["design", "conversion", "learning"],
        },
        {
            id: 4,
            author: {
                name: "Jordan Kim",
                username: "@jordanbuilds",
                avatar: "https://api.dicebear.com/7.x/avataaars/svg?seed=Jordan",
            },
            content:
                "Failed my launch today. Got 3 signups instead of the 100 I was hoping for. But you know what? I learned more from this failure than all my successes combined. Tomorrow, we iterate. 💪",
            timestamp: "2d ago",
            likes: 156,
            comments: 67,
            tags: ["failure", "learning", "resilience"],
        },
    ];

    const stats = [
        { label: "Active Builders", value: "12.5K", icon: Users },
        { label: "Projects Shipped", value: "3.2K", icon: Rocket },
        { label: "Success Stories", value: "890", icon: TrendingUp },
    ];
</script>

<div class="min-h-screen bg-linear-to-b from-background to-muted/20">
    <!-- Hero Section -->
    <section class="relative overflow-hidden border-b bg-background">
        <div
            class="absolute inset-0 bg-[linear-gradient(to_right,hsl(var(--border))_1px,transparent_1px),linear-gradient(to_bottom,hsl(var(--border))_1px,transparent_1px)] bg-size-[4rem_4rem] opacity-5"
        ></div>
        <div
            class="container relative mx-auto max-w-6xl px-4 py-8 sm:px-6 lg:px-8"
        >
            <div class="flex flex-col items-center text-center">
                <!-- Main Heading -->
                <div
                    class="mb-4 inline-flex items-center gap-2 rounded-full border bg-muted/50 px-4 py-2 text-sm font-medium"
                >
                    <Rocket class="h-4 w-4 text-primary" />
                    <span>Build. Learn. Share. Grow.</span>
                </div>

                <h1
                    class="mb-4 text-3xl font-bold tracking-tight sm:text-4xl lg:text-5xl"
                >
                    Build Your Dreams
                    <span
                        class="block bg-linear-to-r from-primary to-chart-1 bg-clip-text text-transparent"
                    >
                        In Public
                    </span>
                </h1>

                <p
                    class="mb-6 max-w-2xl text-base text-muted-foreground sm:text-lg"
                >
                    Join a community of makers, builders, and entrepreneurs
                    sharing their journey. Learn from failures, celebrate wins,
                    and grow together.
                </p>

                <!-- Quote Card -->
                <Card.Root
                    class="relative mb-6 max-w-xl overflow-hidden border-2"
                >
                    <div class="absolute right-3 top-3 opacity-10">
                        <Quote class="h-16 w-16" />
                    </div>
                    <Card.Content class="relative p-4">
                        <blockquote class="text-center">
                            <p
                                class="mb-2 text-base font-medium italic sm:text-lg"
                            >
                                "I hate to lose, more than I enjoy winning."
                            </p>
                            <footer class="text-xs text-muted-foreground">
                                — <span class="font-semibold"
                                    >Michael Phelps</span
                                >
                            </footer>
                        </blockquote>
                    </Card.Content>
                </Card.Root>

                <!-- Stats -->
                <div class="mb-6 grid w-full max-w-2xl grid-cols-3 gap-3">
                    {#each stats as stat (stat.label)}
                        <div
                            class="flex flex-col items-center rounded-lg border bg-card p-4 transition-all hover:shadow-md"
                        >
                            <svelte:component
                                this={stat.icon}
                                class="mb-1 h-5 w-5 text-primary"
                            />
                            <div class="text-2xl font-bold">{stat.value}</div>
                            <div class="text-xs text-muted-foreground">
                                {stat.label}
                            </div>
                        </div>
                    {/each}
                </div>

                <!-- CTA Buttons -->
                <div class="flex flex-wrap items-center justify-center gap-3">
                    <Button class="gap-2">
                        <Rocket class="h-4 w-4" />
                        Start Building
                    </Button>
                    <Button variant="outline">Explore Projects</Button>
                </div>
            </div>
        </div>
    </section>

    <!-- Posts Feed Section -->
    <section class="container mx-auto max-w-4xl px-4 py-12 sm:px-6 lg:px-8">
        <div class="mb-8 flex items-center justify-between">
            <div>
                <h2 class="text-3xl font-bold tracking-tight">
                    Latest Updates
                </h2>
                <p class="mt-2 text-muted-foreground">
                    See what builders are working on right now
                </p>
            </div>
        </div>

        <!-- Posts -->
        <div class="space-y-6">
            {#each posts as post (post.id)}
                <Card.Root class="transition-all hover:shadow-lg">
                    <Card.Header>
                        <div class="flex items-start gap-4">
                            <Avatar.Root class="h-12 w-12">
                                <Avatar.Image
                                    src={post.author.avatar}
                                    alt={post.author.name}
                                />
                                <Avatar.Fallback
                                    >{post.author.name
                                        .slice(0, 2)
                                        .toUpperCase()}</Avatar.Fallback
                                >
                            </Avatar.Root>
                            <div class="flex-1">
                                <div class="flex items-center gap-2">
                                    <h3 class="font-semibold">
                                        {post.author.name}
                                    </h3>
                                    <span class="text-sm text-muted-foreground"
                                        >{post.author.username}</span
                                    >
                                </div>
                                <p class="text-sm text-muted-foreground">
                                    {post.timestamp}
                                </p>
                            </div>
                        </div>
                    </Card.Header>

                    <Card.Content>
                        <p
                            class="mb-4 whitespace-pre-line text-base leading-relaxed"
                        >
                            {post.content}
                        </p>

                        <!-- Tags -->
                        <div class="flex flex-wrap gap-2">
                            {#each post.tags as tag (tag)}
                                <span
                                    class="rounded-full bg-primary/10 px-3 py-1 text-xs font-medium text-primary"
                                >
                                    #{tag}
                                </span>
                            {/each}
                        </div>
                    </Card.Content>

                    <Separator />

                    <Card.Footer class="gap-6 pt-4">
                        <Button
                            variant="ghost"
                            size="sm"
                            class="gap-2 text-muted-foreground hover:text-foreground"
                        >
                            <Heart class="h-4 w-4" />
                            <span class="text-sm font-medium">{post.likes}</span
                            >
                        </Button>
                        <Button
                            variant="ghost"
                            size="sm"
                            class="gap-2 text-muted-foreground hover:text-foreground"
                        >
                            <MessageCircle class="h-4 w-4" />
                            <span class="text-sm font-medium"
                                >{post.comments}</span
                            >
                        </Button>
                        <Button
                            variant="ghost"
                            size="sm"
                            class="gap-2 text-muted-foreground hover:text-foreground"
                        >
                            <Share2 class="h-4 w-4" />
                            <span class="text-sm font-medium">Share</span>
                        </Button>
                    </Card.Footer>
                </Card.Root>
            {/each}
        </div>

        <!-- Load More -->
        <div class="mt-8 flex justify-center">
            <Button variant="outline" size="lg">Load More Updates</Button>
        </div>
    </section>
</div>
