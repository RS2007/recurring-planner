<script lang="ts">
	import { Dialog, DialogOverlay } from "@rgossiaux/svelte-headlessui";
	import { fade } from "svelte/transition";
	export let isOpen: boolean;
</script>

{#if isOpen}
	<Dialog
		open={isOpen}
		on:close={() => (isOpen = false)}
		class="fixed inset-0 p-4 pt-[20vh] flex h-fit justify-center"
	>
		<DialogOverlay class="fixed top-0 left-0 w-full h-full bg-white bg-opacity-10" />
		<div
			class="bg-backgroundDashboard w-[500px] min-w-[300px] overflow-y-scroll h-fit rounded-xl modal flex-col z-10 p-8"
			transition:fade
		>
			<slot />
		</div>
	</Dialog>
{/if}

<style>
	/* Hide scrollbar for Chrome, Safari and Opera */
	.modal::-webkit-scrollbar {
		display: none;
	}

	/* Hide scrollbar for IE, Edge and Firefox */
	.modal {
		-ms-overflow-style: none; /* IE and Edge */
		scrollbar-width: none; /* Firefox */
	}
</style>
