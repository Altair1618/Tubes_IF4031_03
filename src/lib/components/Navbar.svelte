<script lang="ts">
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { DotsVertical, Enter } from 'radix-icons-svelte';
	import Button from './ui/button/button.svelte';
	import { IconDoorEnter, IconLogin2, IconUser } from '@tabler/icons-svelte';
</script>

<div class="w-full border-b border-slate-200 fixed top-0 z-20 bg-zinc-50/10 backdrop-blur-md">
	<div class="w-full max-w-2xl mx-auto flex items-center px-4 py-2 justify-between">
		<a href="/" class="text-2xl sm:text-4xl font-semibold">TESSERA</a>
		<DropdownMenu.Root preventScroll={false}>
			<DropdownMenu.Trigger>
				<Button variant="ghost" class="p-0">
					<DotsVertical size={16} />
				</Button>
			</DropdownMenu.Trigger>
			<DropdownMenu.Content>
				<DropdownMenu.Group>
					<DropdownMenu.Item>
						<a href="/profile/a" class="flex items-center gap-2">
							<IconUser size={18} /> Profile
						</a>
					</DropdownMenu.Item>
					<DropdownMenu.Item>
						<a href="/signin" class="flex items-center gap-2"><IconLogin2 size={16} />Sign In</a>
					</DropdownMenu.Item>
				</DropdownMenu.Group>
			</DropdownMenu.Content>
		</DropdownMenu.Root>
	</div>
</div>
