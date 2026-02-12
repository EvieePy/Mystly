<script lang="ts">
	import "./layout.css";
	import "./fonts.css";
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="layout">
	<aside></aside>
	<main>
		{@render children()}
	</main>
</div>
