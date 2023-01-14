<script lang="ts">
  import {
    Listbox,
    ListboxButton,
    ListboxOptions,
    ListboxOption,
  } from "@rgossiaux/svelte-headlessui";
  import { ChevronDown, Icon } from "svelte-hero-icons";
  import { fade } from "svelte/transition";
  import Empty from "../components/Empty.svelte";
  import { accentColors } from "../utils/formUtils";
  import { getUser, themeColor } from "../utils/userUtils";
  const views = ["Day View", "Month View", "Year View"];
  const days = [
    { abbrev: "MON", name: "Monday" },
    { abbrev: "TUE", name: "Tuesday" },
    { abbrev: "WED", name: "Wednesday" },
    { abbrev: "THU", name: "Thursday" },
    { abbrev: "FRI", name: "Friday" },
    { abbrev: "SAT", name: "Saturday" },
    { abbrev: "SUN", name: "Sunday" },
  ];
  const numDays = [...Array(31).keys()].splice(1);
  const hours = [null, ...Array(24).keys()];
  let selectedView = "Day View";
  const changeViewState = (e) => {
    console.log("in");
    selectedView = e.detail;
  };
</script>

<div class="col-span-11 px-4 flex flex-col">
  <div
    class="flex justify-between w-[100%] border-[1px] border-white/10 bg-highlightPurple mt-8 p-4 rounded-lg overflow-y-visible h-16"
  >
    <div>January 2022</div>
    <Listbox value={selectedView} on:change={changeViewState}>
      <ListboxButton
        class={`${themeColor()} px-3 pt-2 pb-1 rounded-lg flex gap-2`}
        ><span>{selectedView}</span><Icon
          src={ChevronDown}
          class="h-6 w-6"
        /></ListboxButton
      >
      <div transition:fade>
        <ListboxOptions
          class={`bg-highlightPurple delay-100  rounded-lg border-[1px] border-white/10 backdrop-saturate-[180%] backdrop-blur-sm`}
        >
          {#each views as view}
            <ListboxOption
              value={view}
              class={({ active }) =>
                `p-2 ${active ? "bg-blue-500 text-white" : ""}`}
              >{view}</ListboxOption
            >{/each}
        </ListboxOptions>
      </div>
    </Listbox>
  </div>
  {#if selectedView == "Month View"}
    <div class="grid grid-cols-7 grid-rows-[repeat(22,1fr)]">
      {#each days as { abbrev }}
        <div
          class="col-span-1 flex justify-center border-[1px] border-white/10 row-span-1 py-1"
        >
          {abbrev}
        </div>
      {/each}

      {#each numDays as numDay}
        <div
          class="col-span-1 pl-6 pt-2  border-[1px] border-white/10 row-span-4"
        >
          {numDay}
        </div>
      {/each}
    </div>
  {:else if selectedView === "Day View"}
    <div class="flex w-full">
      <div class="flex flex-col w-[5%] ">
        {#each hours as hour}
          <span
            class="border-[1px] border-white/10 px-6 h-[100px] text-gray-400"
            >{hour === null ? "" : hour}
          </span>
        {/each}
      </div>
      <div class="grid grid-rows-[repeat(34,50px)] grid-cols-7 w-[95%]">
        {#each days as { abbrev }}
          <div
            class="col-span-1 flex justify-center border-[1px] border-white/10 row-span-2 py-1  flex-col items-center"
          >
            <div>{abbrev}</div>
            <div>24</div>
          </div>
        {/each}
        {#each days as { abbrev }}
          <div
            class="col-span-1 flex justify-center border-[1px] border-white/10 row-span-1 py-1  "
          />
        {/each}
        {#each days as { abbrev }}
          <div
            class="col-span-1 flex justify-center border-[1px] border-white/10 row-span-1 py-1"
          >
            <div
              class="bg-[#1e4040] rounded-lg w-full h-[100px]  backdrop-saturate-[180%] backdrop-blur-sm "
            >
              Go to library
            </div>
          </div>
        {/each}
        {#each days as { abbrev }}
          <div
            class="col-span-1 flex justify-center border-[1px] border-white/10 row-span-1 py-1"
          >
            <div
              class="bg-highlightPurple  w-full h-[100px]  backdrop-saturate-[180%] backdrop-blur-sm "
            />
          </div>
        {/each}
      </div>
    </div>
  {:else}
    <Empty />
  {/if}
</div>
