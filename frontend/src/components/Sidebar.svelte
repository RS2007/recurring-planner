<script lang="ts">
	import { Icon, Calendar, Cog, UserCircle, Logout } from "svelte-hero-icons";
	import { logOut } from "../utils/userUtils";
	import type { optionsType } from "../utils/types";

	const options: Array<optionsType> = [
		{ title: "Calendar", icon: Calendar, active: true, link: "/" },
		{ title: "Log out", icon: Logout, action: logOut },
	];
</script>

<div class="col-span-1 w-full bg-backgroundDashboard h-full pt-12">
	{#each options as option}
		<div
			class={`flex flex-col items-center w-[100%] mb-6 cursor-pointer ${
				option.active ? `text-[#3fcf8e]` : "text-gray-400"
			}`}
		>
			<a href={`${option.link ?? "/"}`}>
				<button
					class="w-full flex flex-col items-center "
					on:click={() => {
						option.action();
					}}
				>
					<div>
						<Icon src={option.icon} solid={false} class="h-8 w-8" />
					</div>
					<div>
						<p class={`align-middle text-sm `}>
							{option.title}
						</p>
					</div>
				</button>
			</a>
		</div>
	{/each}
</div>
