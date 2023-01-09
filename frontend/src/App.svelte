<script lang="ts">
  import Sidebar from "./components/Sidebar.svelte";
  import Calendar from "./views/Calendar.svelte";
  import TemplateBar from "./views/TemplateBar.svelte";
  import EventDetailsDrawer from "./views/EventDetailsDrawer.svelte";
</script>

<main
  class="grid grid-cols-[repeat(16,1fr)] w-[100%]  h-[100vh] bg-backgroundDashboard text-white"
>
  <Sidebar />
  <TemplateBar />
  <Calendar />
  <!-- <EventDetailsDrawer /> -->
</main>
