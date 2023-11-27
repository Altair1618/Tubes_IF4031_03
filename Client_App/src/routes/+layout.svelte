<script>
	import Navbar from '$lib/components/Navbar.svelte';
	import '../app.postcss';
	import '@fontsource-variable/lexend';
	import { Toaster } from 'svelte-french-toast';

	export let data;
</script>

<Toaster />
<Navbar user={data.user} />
<div class="w-full max-w-4xl mx-auto min-h-screen flex flex-col px-4 py-16">
	<slot />
</div>
