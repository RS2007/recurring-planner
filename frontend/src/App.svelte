<script lang="ts">
  import { SvelteToast } from "@zerodevx/svelte-toast";
  import Sidebar from "./components/Sidebar.svelte";
  import Calendar from "./views/Calendar.svelte";
  import TemplateBar from "./views/TemplateBar.svelte";
  import EventDetailsDrawer from "./views/EventDetailsDrawer.svelte";
  import Auth from "./views/Auth.svelte";
  import { Router, Route } from "svelte-navigator";
  import NewTemplate from "./views/NewTemplate.svelte";
  import Customize from "./views/Customize.svelte";
</script>

<Router>
  <SvelteToast />
  <Route path="/dashboard/customize">
    <div
      class="grid grid-cols-[repeat(16,1fr)] w-[100%] min-h-screen  h-fit bg-backgroundDashboard text-white"
    >
      <Sidebar />
      <TemplateBar />
      <Customize />
    </div>
  </Route>

  <Route path="/dashboard/template/new">
    <div
      class="grid grid-cols-[repeat(16,1fr)] w-[100%] min-h-screen  h-fit bg-backgroundDashboard text-white"
    >
      <Sidebar />
      <TemplateBar />
      <NewTemplate />
    </div>
  </Route>
  <Route exact path="/dashboard">
    <div
      class="grid grid-cols-[repeat(16,1fr)] w-[100%] min-h-screen  h-fit bg-backgroundDashboard text-white"
    >
      <Sidebar />
      <TemplateBar />
      <Calendar />
      <!-- <EventDetailsDrawer /> -->
    </div>
  </Route>
  <Route path="/auth/*">
    <div
      class="min-h-screen h-fit bg-backgroundDashboard text-white flex  items-center w-full flex-col justify-center"
    >
      <Auth />
    </div>
  </Route>
</Router>
