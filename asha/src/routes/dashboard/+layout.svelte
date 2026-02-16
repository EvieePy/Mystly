<script lang="ts">
	import "./layout.css"
	import Side from '$lib/components/side.svelte';

	let { children } = $props();
</script>

<div class="wrapper">
	<div class="layout">
		<Side />
		<main class="rd-2 bg-surface-700">
			{@render children()}
		</main>
	</div>
</div>

