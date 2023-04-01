<script lang="ts" defer>
	import { _axios } from "../utils/_axios";
	import { userStore } from "../store/userStore";
	import type { userCtx } from "../utils/types";
	import { toast } from "@zerodevx/svelte-toast";

	let user: userCtx;
	userStore.subscribe((value) => (user = value));

	// @ts-expect-error
	const client = window.google.accounts.oauth2.initCodeClient({
		client_id: import.meta.env.VITE_GOOGLE_CLIENT_ID,
		scope: import.meta.env.VITE_GOOGLE_SCOPE,
		redirect_uri: import.meta.env.VITE_GOOGLE_REDIRECT_URI,
		ux_mode: "redirect",
	});
</script>

<div class="flex flex-col">
	<h2 class="text-3xl text-center mb-4">Log into Recurrent Planner</h2>
	<div
		class="bg-highlightPurple w-full h-fit py-8 px-8 rounded-[12px] border-white/10 backdrop-saturate-[180%] backdrop-blur-sm"
	>
		<button
			class="w-full bg-supabaseGreen py-2 rounded-sm"
			on:click={() => {
				client.requestCode();
			}}
		>
			Hey
		</button>
	</div>
	<div class="border-[2px] border-white/10 mt-4 p-4 rounded-lg">
		New to reccurent planner, <span class="text-blue-400 cursor-pointer">create an account</span
		>
	</div>
</div>
