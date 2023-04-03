<script>
	import { DialogTitle } from "@rgossiaux/svelte-headlessui";
	import { Icon, X } from "svelte-hero-icons";
	import Modal from "../components/Modal.svelte";
	import { getUTCOffset } from "../utils/dateUtils";
	export let isOpen;
	export let events;

	let summary = "";
	let location = "";
	let description = "";
	let startTime = "";
	let endTime = "";

	const createTemplate = (e) => {
		e.preventDefault();
		events = [
			...events,
			{
				summary,
				location,
				description,
				startTime: startTime + getUTCOffset(),
				endTime: endTime + getUTCOffset(),
			},
		];
		isOpen = false;
		summary = "";
		location = "";
		description = "";
		startTime = "";
		endTime = "";
	};
</script>

<Modal bind:isOpen>
	<div class="flex justify-between mb-4">
		<DialogTitle class="text-white text-xl font-extrabold mb-2"
			>Create a new template</DialogTitle
		>
		<button
			on:click={() => {
				isOpen = false;
			}}
		>
			<Icon src={X} class="h-6 w-6 text-white" />
		</button>
	</div>

	<form on:submit={createTemplate}>
		<div class="flex flex-col gap-2">
			<div class="flex gap-2">
				<div class="flex flex-col text-white w-1/2">
					<label class="text-lg">Summary</label>
					<input
						class="pl-2 bg-highlightPurple h-[40px] rounded-lg"
						bind:value={summary}
					/>
				</div>
				<div class="flex flex-col text-white w-1/2">
					<label class="text-lg">Location</label>
					<input
						class="pl-2 bg-highlightPurple h-[40px] rounded-lg"
						bind:value={location}
					/>
				</div>
			</div>
			<div class="flex flex-col text-white mb-2">
				<label class="text-lg">Description</label>
				<textarea
					class="pl-2 bg-highlightPurple h-[100px] rounded-lg"
					bind:value={description}
				/>
			</div>
			<div class="flex gap-2">
				<div class="flex flex-col text-white w-1/2">
					<label class="text-lg">Start Time</label>
					<input
						class="pl-2 bg-highlightPurple h-[40px] rounded-lg"
						bind:value={startTime}
					/>
				</div>
				<div class="flex flex-col text-white w-1/2">
					<label class="text-lg">End Time</label>
					<input
						class="pl-2 bg-highlightPurple h-[40px] rounded-lg"
						bind:value={endTime}
					/>
				</div>
			</div>

			<button class="w-full bg-blue-600 text-white py-2 mt-2 rounded-lg" type="submit"
				>Add Template</button
			>
		</div>
	</form>
</Modal>
