import type { dto_UserResponse } from "$lib/api-client";

/**
 * Global user store using Svelte 5 $state
 *
 * States:
 * - undefined: Not yet loaded (initial state)
 * - null: Loaded but user is not authenticated
 * - dto_UserResponse: User is authenticated
 *
 * Usage in components:
 *
 * ```svelte
 * <script lang="ts">
 *   import { user } from "$lib/store/user.svelte";
 *
 *   // Access user data reactively
 *   $effect(() => {
 *     console.log(user.current?.name);
 *   });
 * </script>
 *
 * {#if user.isLoading}
 *   <p>Loading...</p>
 * {:else if user.isAuthenticated}
 *   <p>Welcome, {user.current?.name}!</p>
 * {:else}
 *   <p>Please log in</p>
 * {/if}
 * ```
 */
export const user: {
  current: dto_UserResponse | null | undefined;
  isLoading: boolean;
  isAuthenticated: boolean;
} = $state({
  current: undefined,
  isLoading: true,
  get isAuthenticated() {
    return this.current !== null && this.current !== undefined;
  },
});
