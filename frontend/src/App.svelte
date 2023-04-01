<script>
	import { SvelteToast } from "@zerodevx/svelte-toast";
	import Sidebar from "./components/Sidebar.svelte";
	import Calendar from "./views/Calendar.svelte";
	import TemplateBar from "./views/TemplateBar.svelte";
	import EventDetailsDrawer from "./views/EventDetailsDrawer.svelte";
	import Auth from "./views/Auth.svelte";
	import { Route } from "tinro";
	import NewTemplate from "./views/NewTemplate.svelte";
	import Customize from "./views/Customize.svelte";
	import { getUser, isLoggedIn } from "./utils/userUtils";
</script>

<div>
	<SvelteToast />
	<Route path="/customize">
		<div
			class="grid grid-cols-[repeat(16,1fr)] min-w-screen w-full min-h-screen  h-fit bg-backgroundDashboard text-white"
		>
			<Sidebar />
			<TemplateBar />
			<Customize />
		</div>
	</Route>

	<Route path="/template/new">
		<div
			class="grid grid-cols-[repeat(16,1fr)] w-[100%] min-h-screen  h-fit bg-backgroundDashboard text-white"
		>
			<Sidebar />
			<TemplateBar />
			<NewTemplate />
		</div>
	</Route>
	<Route path="/">
		{#if isLoggedIn()}
			<div
				class="grid grid-cols-[repeat(16,1fr)] min-w-full w-fit min-h-screen  h-fit bg-backgroundDashboard text-white"
			>
				<Sidebar />
				<TemplateBar />
				<Calendar />
				<!-- <EventDetailsDrawer /> -->
			</div>
		{:else}
			<div
				class="min-h-screen h-fit bg-backgroundDashboard text-white flex  items-center w-full flex-col justify-center"
			>
				<Auth />
			</div>
		{/if}
	</Route>
</div>
