<script lang="ts">
	import { Icon, ChevronUp, ChevronDown } from "svelte-hero-icons";
	import { fade } from "svelte/transition";
	import { userStore } from "../store/userStore";
	import type { schedule, templateType, userCtx } from "../utils/types";
	let user: userCtx;
	userStore.subscribe((value) => (user = value));

	export let template: templateType;

	const collapseTemplate = () => {
		template.expanded = !template.expanded;
	};
</script>

<div
	class={`mb-4 bg-highlightPurple mr-4 backdrop-saturate-[180%] backdrop-blur-sm z-50 w-[90%] rounded-lg`}
>
	<button on:click={collapseTemplate} class="w-full">
		<div class="flex justify-between pr-4 pl-4 py-3">
			<p class="text-lg">{template.name}</p>
			<Icon src={template.expanded ? ChevronUp : ChevronDown} class="h-6 w-6" />
		</div>
	</button>
	{#if template.expanded}
		<div class={` mr-4 rounded-lg`} transition:fade>
			{#each template.events as event}
				<div class="pl-4 py-2 flex flex-col mb-2">
					<div class=" flex justify-between">
						<p class="text-md">{event.summary}</p>
						<p class="text-md text-gray-400">
							{`${new Date(event.startTime).getHours()}:${new Date(
								event.startTime
							).getMinutes()}`} - {`${new Date(event.endTime).getHours()} :${new Date(
								event.endTime
							).getMinutes()}`}
						</p>
					</div>
					<div class={`${user.accentColor.class}} rounded-lg px-2 py-1 w-fit text-base`}>
						{event.location}
					</div>
				</div>
			{/each}
		</div>
	{/if}
</div>
