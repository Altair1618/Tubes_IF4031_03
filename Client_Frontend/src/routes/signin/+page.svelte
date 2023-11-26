<script>
	import Button from '$lib/components/ui/button/button.svelte';
	import { IconBrandGoogleFilled } from '@tabler/icons-svelte';

	export let data;
</script>

<div class="flex flex-1 flex-col items-center justify-center">
	<a href={`${data.clientURL}/auth/google`}>
		<Button class="flex items-center gap-2" variant="outline">
			<IconBrandGoogleFilled />
			Continue with Google
		</Button>
	</a>
</div>
