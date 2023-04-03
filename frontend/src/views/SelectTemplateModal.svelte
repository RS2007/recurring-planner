<script lang="ts">
	import { DialogTitle } from "@rgossiaux/svelte-headlessui";
	import { toast } from "@zerodevx/svelte-toast";
	import { Icon, X } from "svelte-hero-icons";
	import Modal from "../components/Modal.svelte";
	import { templatesStore } from "../store/templateStore";
	import { userStore } from "../store/userStore";
	import { dateStringConstructor } from "../utils/dateUtils";
	import { getGoogleToken, userAuthToken } from "../utils/userUtils";
	import { _axios } from "../utils/_axios";
	import type { userCtx, templateType } from "../utils/types";
	export let isOpen: boolean;
	export let selectedDay: number;

	let selectedIndex: null | number = null;
	const selectTemplate = (index) => {
		selectedIndex = index;
	};
	const firstDayOfMonth = new Date(new Date().getFullYear(), new Date().getMonth(), 0).getDay();
	const submitForm = async (e) => {
		e.preventDefault();
		console.log(`gonna add template ${selectedIndex} to the day ${selectedDay}`);
		if ([null, undefined].includes(selectedIndex)) {
			toast.push("Select a template");
		}
		const { accessToken, idToken } = getGoogleToken();
		try {
			await _axios.post(
				`/template/day`,
				{
					dateString: dateStringConstructor(
						selectedDay - firstDayOfMonth + 1,
						new Date().getMonth() + 1,
						new Date().getFullYear()
					),
					templateId: templates[selectedIndex].templateId,
					accessToken,
					idToken,
				},
				{
					headers: {
						Authorization: `Bearer ${userAuthToken()}`,
					},
				}
			);
			toast.push("Event pushed to your google calendar succesfully");
		} catch (err) {
			console.log(err);
			toast.push(err?.response?.data?.error?.message || "Something went wrong");
		}
	};
	let user: userCtx;
	let templates: Array<templateType>;
	userStore.subscribe((value) => (user = value));
	templatesStore.subscribe((value) => (templates = value));
	console.log({ templates });
</script>

<Modal bind:isOpen>
	<div class="flex flex-col w-full mb-4">
		<div class="flex justify-between mb-4 w-full">
			<DialogTitle class="text-white text-xl font-extrabold mb-2">Select Template</DialogTitle
			>
			<button
				on:click={() => {
					isOpen = false;
				}}
			>
				<Icon src={X} class="h-6 w-6 text-white" />
			</button>
		</div>

		<hr class="border-white/10 w-full" />
	</div>

	<form on:submit={submitForm}>
		<div class="flex direction-column text-white gap-8 h-fit">
			{#each templates as template, index}
				<div class="w-full align-middle text-lg mb-8">
					<input
						type="checkbox"
						class="mt-4 accent-supabaseGreen"
						name={index.toString()}
						on:change={() => selectTemplate(index)}
					/>
					<label for={index.toString()}>&nbsp;&nbsp;{template.name}</label>
				</div>
			{/each}
		</div>

		<button class={`bg-supabaseGreen text-white px-4 py-2 w-full`} type="submit">Submit</button>
	</form>
</Modal>
