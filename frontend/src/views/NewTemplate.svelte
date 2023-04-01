<script lang="ts">
	import { toast } from "@zerodevx/svelte-toast";
	import { _axios } from "../utils/_axios";
	import { Icon, LocationMarker, ViewList, Calendar, Plus, XCircle } from "svelte-hero-icons";
	import CreateTemplateModal from "./CreateEventModal.svelte";
	import { userAuthToken } from "../utils/userUtils";
	import { router } from "tinro";
	let tagInput = "";
	let isOpen = false;
	let name = "";
	let tags = [];
	let events = [];
	const removeTags = (index: number) => {
		tags = [...tags.splice(0, index), ...tags.splice(index + 1)];
	};

	const openModal = (e) => {
		e.preventDefault();
		isOpen = true;
	};

	const createTemplate = async (e) => {
		e.preventDefault();
		console.log({ events, name, tags });
		try {
			const res = await _axios.post(
				"/template/new",
				{
					name,
					tags: tags.map((tag: string) => ({ name: tag })),
					events,
					timeZone: new window.Intl.DateTimeFormat().resolvedOptions().timeZone,
				},
				{
					headers: {
						Authorization: `Bearer ${userAuthToken()}`,
					},
				}
			);
			if (res.status === 200) {
				toast.push("Template creation succesful");
				router.goto("/");
			}
		} catch (e) {
			toast.push(e.message);
		}
	};

	const IsTagInputValid = (tags) => tags.length > 3;

	const createAndPushTags = (e) => {
		if (e.key === " ") {
			tags = [...tags, tagInput];
			tagInput = "";
		}
	};

	const onSubmit = () => {};
</script>

<div class="col-span-12  pl-4 pt-8">
	<h1 class="text-3xl mb-8">Create a new template</h1>
	<form class="flex flex-col" on:submit={createTemplate}>
		<div class="flex gap-6">
			<div class="flex flex-col">
				<label>Name</label>
				<input
					type="text"
					class="pl-2 bg-highlightPurple w-[320px] min-h-[40px] rounded-lg"
					bind:value={name}
				/>
			</div>
			<div class="flex flex-col">
				<label>Tags</label>
				<div
					class="flex pl-2 bg-highlightPurple w-[400px] overflow-x-scroll min-h-[40px] rounded-lg focus-within:ring-1 focus-within:ring-slate-100\40 hide-scroll"
				>
					<ul class="inline-flex gap-2 items-center">
						{#each tags as tag, index}
							<li class="bg-blue-600 px-2 rounded-lg inline-flex gap-1 items-center">
								{tag}<button on:click={() => removeTags(index)}
									><Icon src={XCircle} class="h-4 w-4" /></button
								>
							</li>
						{/each}
					</ul>
					<input
						type="text"
						class="pl-2 bg-highlightPurple w-[320px] min-h-[40px] rounded-lg focus:outline-none"
						on:keydown={createAndPushTags}
						bind:value={tagInput}
						disabled={IsTagInputValid(tags)}
					/>
				</div>
			</div>
		</div>
		<div class="flex text-[24px] mt-2 mb-4 gap-4 items-center">
			<h1>Events</h1>
			<button class="bg-supabaseGreen px-2 h-6 rounded-md" on:click={openModal}
				><Icon src={Plus} class="h-4 w-4" /></button
			>
			<CreateTemplateModal bind:isOpen bind:events />
		</div>
		<div class="flex gap-6 flex-wrap">
			{#if events.length === 0}
				<div>Wow so empty 😶</div>
			{:else}
				{#each events as event}
					<div
						class="pt-4 bg-highlightPurple h-[220px] w-[400px]  backdrop-saturate-[180%] backdrop-blur-sm text-md rounded-xl"
					>
						<h1 class="ml-4 text-[22px] font-extrabold">{event.summary}</h1>
						<hr class="border-white/10" />
						<div class="pt-4 pl-4">
							<div class="flex gap-2">
								<Icon src={LocationMarker} class="h-6 w-6 stroke-green-400" />
								<div class="flex items-center">
									<p class="text-[16px]">{event.location}</p>
								</div>
							</div>
						</div>
						<div class="pt-4 pl-4">
							<div class="flex gap-2">
								<Icon src={ViewList} class="h-6 w-6 stroke-fuchsia-400" />
								<div class="flex items-center">
									<p class="text-[16px]">
										{event.description}
									</p>
								</div>
							</div>
						</div>
						<div class="pt-4 pl-4">
							<div class="flex gap-2">
								<Icon src={Calendar} class="h-5 w-5 stroke-red-400" />
								<div class="flex items-center">
									<p class="text-[16px]">{event.startTime} - {event.endTime}</p>
								</div>
							</div>
						</div>
					</div>
				{/each}
			{/if}
		</div>
		<button
			type="submit"
			class="py-2 bg-supabaseGreen mt-4 max-w-[200px] rounded-lg font-extrabold"
			>Create Template</button
		>
	</form>
</div>

<style>
	/* Hide scrollbar for Chrome, Safari and Opera */
	.hide-scroll::-webkit-scrollbar {
		display: none;
	}

	/* Hide scrollbar for IE, Edge and Firefox */
	.hide-scroll {
		-ms-overflow-style: none; /* IE and Edge */
		scrollbar-width: none; /* Firefox */
	}
</style>
