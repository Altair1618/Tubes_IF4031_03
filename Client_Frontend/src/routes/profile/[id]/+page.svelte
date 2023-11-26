<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import { IconLogout2, IconTrash } from '@tabler/icons-svelte';
	import toast from 'svelte-french-toast';
	import { superForm } from 'sveltekit-superforms/client';
	import { invalidateAll } from '$app/navigation';
	import { enhance as enhance2 } from '$app/forms';

	export let data;

	const { form, enhance, errors, constraints } = superForm(data.form, {
		onResult: ({ result }) => {
			if (result.type === 'failure') {
				toast.error(result.data?.form.message);
			} else if (result.type === 'success') {
				toast.success(result.data?.message);
			}
		}
	});
</script>

<div class="flex flex-1 flex-col gap-4 w-full py-4">
	<h2 class="text-2xl">Profile</h2>
	<div class="border-b" />
	<div class="flex justify-center w-full">
		<img alt={data.data.name} class="bg-cover w-20 h-20 rounded-full" src={data.data.picture} />
	</div>
	<form action="?/updateProfile" class="flex flex-col gap-4" method="post" use:enhance>
		<div class="flex flex-col gap-1">
			<span>Email</span>
			<div class="p-2 border rounded-md bg-zinc-100 text-zinc-500">
				<p>{data.data.email}</p>
			</div>
		</div>
		<div class="flex flex-col gap-1">
			<label for="name">Name</label>
			<input
				{...$constraints.name}
				aria-invalid={$errors.name ? 'true' : undefined}
				bind:value={$form.name}
				class={`p-2 rounded-md ${$errors.name ? 'border border-destructive' : 'border'}`}
				id="name"
				name="name"
				type="text"
			/>
			{#if $errors.name}
				<span class="text-destructive">{$errors.name}</span>
			{/if}
		</div>
		<Button type="submit" variant="ghost" disabled={data.form.data.name === $form.name}
			>Update Profile</Button
		>
	</form>
	<div class="border-b" />
	<div class="flex flex-col gap-4 my-4">
		<form
			action="?/signOut"
			method="post"
			use:enhance2={() => {
				return async ({ result }) => {
					if (result.type === 'error') {
						toast.error(result.error.message);
					} else if (result.type === 'success') {
						toast.success('Sign Out success');
						await invalidateAll();
					}
				};
			}}
		>
			<Button class="flex items-center gap-2 w-full" type="submit">
				<IconLogout2 size={18} />
				Sign Out
			</Button>
		</form>
		<form
			action="?/deleteProfile"
			method="post"
			use:enhance2={() => {
				return async ({ result }) => {
					if (result.type === 'error') {
						toast.error(result.error.message);
					} else if (result.type === 'success') {
						toast.success('Account successfully deleted');
						await invalidateAll();
					}
				};
			}}
		>
			<Button
				class="flex w-full items-center gap-2 hover:border-destructive hover:text-destructive"
				variant="outline"
				type="submit"
			>
				<IconTrash size={16} />
				Delete Account
			</Button>
		</form>
	</div>
</div>
