<script lang="ts">
	import {
		Listbox,
		ListboxButton,
		ListboxOptions,
		ListboxOption,
	} from "@rgossiaux/svelte-headlessui";
	import { ChevronDown, Icon } from "svelte-hero-icons";
	import { fade } from "svelte/transition";
	import Empty from "../components/Empty.svelte";
	import { userStore } from "../store/userStore";
	import { getCurrentMonth, getCurrentYear, numDaysInCurrentMonth } from "../utils/dateUtils";
	import { days, monthNames, views } from "../utils/dayConstants";
	import SelectTemplateModal from "./SelectTemplateModal.svelte";
	import type { userCtx } from "../utils/types";
	import { getRange } from "../utils/misc";

	let user: userCtx;
	userStore.subscribe((value) => (user = value));

	let isSelectTemplateModalOpen = false;
	let selectedDay: number;

	const firstDayOfMonth = new Date(new Date().getFullYear(), new Date().getMonth(), 0).getDay();
	const numDays = [
		...getRange(0, firstDayOfMonth - 1), // skip till first day of month
		...[...Array(numDaysInCurrentMonth()).keys()].map((val) => val + firstDayOfMonth), // create the month
	];

	let selectedView = "Month View";
	const changeViewState = (e: CustomEvent) => {
		selectedView = e.detail;
	};
</script>

<div class="col-span-11 px-4 flex flex-col">
	<div
		class="flex justify-between w-[100%] border-[1px] border-white/10 bg-highlightPurple mt-8 p-4 rounded-lg overflow-y-visible h-16"
	>
		<div>{monthNames[getCurrentMonth()]} {getCurrentYear()}</div>
		<Listbox value={selectedView} on:change={changeViewState}>
			<ListboxButton class={`bg-[#3fcf8e] px-3 pt-2 pb-1 rounded-lg flex gap-2 `}
				><span class="font-extrabold">{selectedView}</span><Icon
					src={ChevronDown}
					class="h-6 w-6"
				/></ListboxButton
			>
			<div transition:fade>
				<ListboxOptions
					class={`bg-highlightPurple delay-100  rounded-lg border-[1px] border-white/10 backdrop-saturate-[180%] backdrop-blur-sm`}
				>
					{#each views as view}
						<ListboxOption
							value={view}
							class={({ active }) => `p-2 ${active ? "bg-blue-500 text-white" : ""}`}
							>{view}</ListboxOption
						>{/each}
				</ListboxOptions>
			</div>
		</Listbox>
	</div>
	{#if selectedView == "Month View"}
		<div class="grid grid-cols-7 grid-rows-[repeat(22,1fr)]">
			{#each days as { abbrev }}
				<div
					class="col-span-1 flex justify-center border-[1px] border-white/10 row-span-1 py-1"
				>
					{abbrev}
				</div>
			{/each}

			{#each numDays as numDay}
				<div
					class="col-span-1 pl-6 pt-2  border-[1px] border-white/10 row-span-4"
					on:click={() => {
						isSelectTemplateModalOpen = true;
						selectedDay = numDay;
					}}
					on:keydown={() => {
						isSelectTemplateModalOpen = true;
						selectedDay = numDay;
					}}
					on:keypress={() => {}}
				>
					{#if numDay >= firstDayOfMonth}
						{numDay - firstDayOfMonth + 1}
					{:else}
						{""}
					{/if}
				</div>
			{/each}
		</div>
		{#if isSelectTemplateModalOpen}
			<SelectTemplateModal bind:isOpen={isSelectTemplateModalOpen} bind:selectedDay
				><p class="text-white">Hey guys</p></SelectTemplateModal
			>
		{/if}
	{:else}
		<Empty />
	{/if}
</div>
