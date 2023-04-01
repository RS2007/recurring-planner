<script>
	import {
		Listbox,
		ListboxButton,
		ListboxOption,
		ListboxOptions,
	} from "@rgossiaux/svelte-headlessui";
	import { Icon, Mail, Tag, UserCircle, ChevronDown } from "svelte-hero-icons";
	import { accentColors } from "../utils/formUtils";
	import { fade } from "svelte/transition";
	import { getUser } from "../utils/userUtils";
	import { userStore } from "../store/userStore";
	let user;
	userStore.subscribe((value) => (user = value));
	let accentColor = accentColors.find((color) => color.label === getUser().accentColor.label);
	const changeAccentColor = (e) => {
		accentColor = accentColors.find((accentColor) => accentColor.label === e.detail);
	};
</script>

<div class="col-span-12 flex pt-8 pl-4 gap-6">
	<div
		class="pt-4 bg-highlightPurple h-[220px] w-fit  backdrop-saturate-[180%] backdrop-blur-sm text-md flex flex-col"
	>
		<h1 class="text-2xl px-4">Customize</h1>

		<hr class="border-white/10" />
		<div class="flex gap-4 px-4 pt-2 min-w-[450px] gap-8">
			<div class="flex flex-col">
				<label>Sort templates By</label>
				<Listbox bind:value={accentColor} on:change={changeAccentColor}>
					<ListboxButton
						class={`${accentColor.class} px-3 pt-2 pb-1 rounded-lg flex gap-2`}
						><span>{accentColor.label}</span><Icon
							src={ChevronDown}
							class="h-6 w-6"
						/></ListboxButton
					>
					<div transition:fade class="z-10">
						<ListboxOptions
							class={`bg-backgroundDashboard delay-100  rounded-lg border-[1px] border-white/10 backdrop-saturate-[180%] backdrop-blur-sm h-fit z-10`}
						>
							{#each accentColors as accentColor}
								<ListboxOption
									value={accentColor.label}
									class={({ active }) =>
										`p-2 ${active ? "bg-blue-500 text-white" : ""}`}
									>{accentColor.label}</ListboxOption
								>{/each}
						</ListboxOptions>
					</div>
				</Listbox>
			</div>
		</div>
	</div>
</div>
