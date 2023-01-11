<script>
  import { Route } from "svelte-navigator";
  import Login from "./Login.svelte";
  import Register from "./Register.svelte";
</script>

<div class="flex flex-col">
  <h2 class="text-3xl text-center mb-4">Log into Recurrent Planner</h2>
  <div
    class="bg-highlightPurple w-fit h-fit py-8 px-8 rounded-[12px] border-white/10 backdrop-saturate-[180%] backdrop-blur-sm"
  >
    <Route path="/login" component={Login} />
    <Route path="/register" component={Register} />
  </div>
  <div class="border-[2px] border-white/10 mt-4 p-4 rounded-lg">
    New to reccurent planner, <span class="text-blue-400 cursor-pointer"
      >create an account</span
    >
  </div>
</div>
