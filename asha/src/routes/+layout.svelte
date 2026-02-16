<script lang="ts">
	import 'virtual:uno.css';
	import "./layout.css"
	import "./fonts.css";
	import favicon from '$lib/assets/favicon.svg';


	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children()}


