<script lang="ts">
	import { Icon, FolderAdd } from "svelte-hero-icons";
	import TemplateCard from "./TemplateCard.svelte";
	import { _axios } from "../utils/_axios";
	import { templatesStore } from "../store/templateStore";
	import { userAuthToken } from "../utils/userUtils";
	import type { templateType } from "../utils/types";

	let templates: Array<templateType>;
	let isLoading = false;

	templatesStore.subscribe((value) => (templates = value));

	const fetchTemplates = async () => {
		isLoading = true;
		const templatesFromServer = (
			await _axios.get("/template/all", {
				headers: {
					Authorization: `Bearer ${userAuthToken()}`,
				},
			})
		).data.templates as Array<templateType>;

		$templatesStore = templatesFromServer;
		console.log(templates);
		isLoading = false;
	};

	fetchTemplates();

	const handleSort = (e) => {
		$templatesStore = e.detail.items;
	};
</script>

<div class="pl-4 col-span-3 pt-8  " style="border:1px solid #2B2B2B">
	<div class="mb-8 flex justify-between pr-4 items-center">
		<p class="text-2xl">My templates</p>
		<a href="/template/new">
			<button>
				<Icon src={FolderAdd} class="h-6 w-6" />
			</button>
		</a>
	</div>
	{#if isLoading}
		<h1>Loading...</h1>
	{:else if templates?.length}
		<div class="min-h-[80vh]">
			{#each templates as template (template.templateId)}
				<TemplateCard {template} />
			{/each}
		</div>
	{:else}
		<div>Wow so empty 😶</div>
	{/if}
</div>
