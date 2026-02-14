<script lang="ts">
	import { onMount } from 'svelte';
	import { error as SError } from '@sveltejs/kit';
	import OverviewCard from '$lib/components/overview-card.svelte';
	import { apiBase } from '../../../stores';
	import Icon from '@iconify/svelte';

	const URL: string = `${$apiBase}/overviewstats`;

	interface SysUsage {
		cpu_per: number;
		mem_total: number;
		mem_used: number;
		mem_free: number;
		mem_used_per: number;
		physical_cores: number;
		logical_cores: number;
		disk_used: number;
		disk_total: number;
		disk_free: number;
		disk_per: number;
	}
	let stats: SysUsage | undefined;

	onMount(async () => {
		stats = await fetchStats();
	});

	async function fetchStats(): Promise<SysUsage | undefined> {
		let resp: Response;

		try {
			resp = await fetch(URL);
		} catch (error: any) {
			throw SError(500, error);
		}

		if (!resp.ok) {
			if (resp.status === 401) {
				// Redirect...
				return;
			}
			throw SError(500, `Unable to fetch data from API. Status: ${resp.status}`);
		}

		let data: SysUsage = await resp.json();

		return data;
	}
</script>

<div>
	<h3 class="text-tertiary">Dashboard Overview</h3>
	<hr />
</div>


<div class="flex-r g-1 jc-center">
	<OverviewCard title="Services">
		<div class="flex-c g-1 text-3.6 jc-space-evenly text-surface-100">
			<div class="flex-r g-1 ai-center">
				<Icon icon="mingcute:game-2-fill" class="text-primary"/>
				<span><b>3</b> <span class="text-3.2 text-surface-200"> Installed</span></span>
			</div>
			<div class="flex-r g-1 ai-center">
				<Icon icon="mdi:progress-tick" class="text-success"/>
				<span><b>2</b> <span class="text-3.2 text-surface-200"> Running</span></span>
			</div>
			<div class="flex-r g-1 ai-center">
				<Icon icon="heroicons-solid:status-offline" class="text-error"/>
				<span><b>1</b> <span class="text-3.2 text-surface-200"> Offline</span></span>
			</div>
		</div>
	</OverviewCard>

	<OverviewCard title="System Memory">
		{#if stats}
			<div class="flex-c g-1 text-3.6 jc-space-evenly text-surface-100">
				<div class="flex-r g-1 ai-center">
					<Icon icon="ph:memory-fill"/>
					<span>{Math.round(stats.mem_used / Math.pow(1024, 3))} GiB</span>
					<span class="text-surface-200 text-3.2"> / {Math.round(stats.mem_total / Math.pow(1024, 3))} GiB</span>
				</div>
				<div class="flex-r g-1 ai-center">
					<Icon icon="ph:percent-fill"/>
					<div class="flex-r flexn-1 ai-center jc-flex-end g-1">
						<span>{Math.round(stats.mem_used_per)} %</span>
						<span class="text-surface-200 text-3.2"> / 100 %</span>
					</div>
				</div>
			</div>
		{/if}
	</OverviewCard>

	<OverviewCard title="System CPU" >
		{#if stats}
			<div class="flex-c g-1 text-3.6 jc-space-evenly text-surface-100">
				<div class="flex-r g-1 ai-center">
					<Icon icon="ph:cpu-fill"/>
					<span><b>{stats.physical_cores}</b> <span class="text-3.2 text-surface-200">Physical</span></span>
					<span><b>{stats.logical_cores}</b> <span class="text-3.2 text-surface-200">Logical</span></span>
				</div>
				<div class="flex-r g-1 ai-center">
					<Icon icon="ph:percent-fill"/>
					<div class="flex-r flexn-1 ai-center jc-flex-end g-1">
						<span>{Math.round(stats.cpu_per)} %</span>
						<span class="text-surface-200 text-3.2"> / 100 %</span>
					</div>
				</div>
			</div>
		{/if}
	</OverviewCard>

	<OverviewCard title="Disk Usage" >
		{#if stats}
			<div class="flex-c g-1 text-3.6 jc-space-evenly text-surface-100">
				<div class="flex-r g-1 ai-center">
					<Icon icon="icon-park-solid:hard-disk"/>
					<span>{Math.round(stats.disk_used / Math.pow(1024, 3))} GiB</span>
					<span class="text-surface-200 text-3.2"> / {Math.round(stats.disk_total / Math.pow(1024, 3))} GiB</span>
				</div>
				<div class="flex-r g-1 ai-center">
					<Icon icon="ph:percent-fill"/>
					<div class="flex-r flexn-1 ai-center jc-flex-end g-1">
						<span>{Math.round(stats.disk_per)} %</span>
						<span class="text-surface-200 text-3.2"> / 100 %</span>
					</div>
				</div>
			</div>
		{/if}
	</OverviewCard>

	<OverviewCard title="Network Usage">
		{#if stats}
			<div class="flex-c g-1 text-3.6 jc-space-evenly text-surface-100">
				<div class="flex-r g-1 ai-center">
					<Icon icon="ph:memory-fill"/>
					<span>{Math.round(stats.mem_used / Math.pow(1024, 3))} GiB</span>
					<span class="text-surface-200 text-3.2"> / {Math.round(stats.mem_total / Math.pow(1024, 3))} GiB</span>
				</div>
				<div class="flex-r g-1 ai-center">
					<span>{stats.mem_used_per} %</span>
					<span class="text-surface-200 text-3.2"> / 100 %</span>
				</div>
			</div>
		{/if}
	</OverviewCard>

</div>


<div class="flex-c g-1 bg-surface-600 rd-2 p-1">
	<h4 class="text-primary-400">Services Overview</h4>
</div>