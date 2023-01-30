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
  const numDays = [...Array(new Date(2022, 1, 0).getDate()).keys()].splice(1);
  const hours = [...Array(24).keys()];
  const testArray = new Array(48).fill(new Array(8));
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
    <div class=" w-full h-full grid-cols-8 grid-rows-[repeat(48,50px)] grid">
      <div class="col-span-1 row-span-2 border-[1px] border-white/10 px-2">
        Time
      </div>
      {#each days as day}
        <div class="col-span-1 row-span-2 border-[1px] border-white/10 px-2">
          {day.abbrev}
        </div>
      {/each}
      {#each testArray as halfHour, halfHourIndex}
        {#each halfHour as block, blockIndex}
          {#if halfHourIndex % 2 == 0 && blockIndex == 0}
            <div class="col-span-1 row-span-2 border-[1px] border-white/10">
              {halfHourIndex / 2}
            </div>
          {:else if blockIndex != 0}
            <div class="col-span-1 row-span-1 border-[1px] border-white/10" />
          {/if}
        {/each}
      {/each}
    </div>
  {:else}
    <Empty />
  {/if}
</div>
