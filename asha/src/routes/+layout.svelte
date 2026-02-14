
<script lang="ts">
	import 'virtual:uno.css';
	import "./layout.css";
	import "./fonts.css";
	import favicon from '$lib/assets/favicon.svg';
	import Side from '$lib/components/side.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="wrapper">
	<div class="layout">
		<Side />
		<main class="rd-2 bg-surface-700">
			{@render children()}
		</main>
	</div>
</div>

