<script>
	import Navbar from '$lib/components/Navbar.svelte';
	import '../app.postcss';
	// Supports weights 100-900
	import '@fontsource-variable/lexend';
</script>

<Navbar />
<div class="w-full max-w-2xl mx-auto min-h-screen flex flex-col px-4 py-16">
	<slot />
</div>
