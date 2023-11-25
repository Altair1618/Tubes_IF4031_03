<script lang="ts">
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { DotsVertical } from 'radix-icons-svelte';
	import Button from './ui/button/button.svelte';
	import { IconLogin2, IconLogout2, IconUser } from '@tabler/icons-svelte';
	import type { User } from '$lib/types/common';
	import { enhance } from '$app/forms';
	import { invalidateAll } from '$app/navigation';
	import toast from 'svelte-french-toast';

	export let user: User | undefined;
</script>

<div class="w-full border-b border-slate-200 fixed top-0 z-20 bg-zinc-50/10 backdrop-blur-md">
	<div class="w-full max-w-2xl mx-auto flex items-center px-4 py-2 justify-between">
		<a class="text-2xl sm:text-4xl font-semibold" href="/">TESSERA</a>
		<DropdownMenu.Root preventScroll={false}>
			<DropdownMenu.Trigger>
				<Button class="p-0" variant="ghost">
					<DotsVertical size={16} />
				</Button>
			</DropdownMenu.Trigger>
			<DropdownMenu.Content>
				<DropdownMenu.Group>
					{#if user}
						<DropdownMenu.Item>
							<a href={`/profile/${user.userId}`} class="flex items-center gap-2">
								<IconUser size={18} />
								Profile
							</a>
						</DropdownMenu.Item>
					{:else}
						<DropdownMenu.Item>
							<a href="/signin" class="flex items-center gap-2">
								<IconLogin2 size={16} />
								Sign In</a
							>
						</DropdownMenu.Item>
					{/if}
				</DropdownMenu.Group>
			</DropdownMenu.Content>
		</DropdownMenu.Root>
	</div>
</div>
