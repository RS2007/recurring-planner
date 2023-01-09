<script lang="ts">
  import {
    Menu,
    MenuButton,
    MenuItems,
    MenuItem,
  } from "@rgossiaux/svelte-headlessui";
  import { Icon, ChevronDown } from "svelte-hero-icons";
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
  const numDays = [...Array(31).keys()];
</script>

<div class="col-span-11 px-4 flex flex-col">
  <!-- // TODO: create the basic calendar UI using grids Calendar -->
  <div
    class="flex justify-between w-[100%] border-[1px] border-white/10 bg-highlightPurple mt-8 p-4 rounded-lg overflow-y-visible h-16"
  >
    <div>January 2022</div>
    <Menu>
      <MenuButton class="bg-blue-600 px-4 py-1 rounded-lg flex gap-2"
        ><span>Month View </span>
        <Icon src={ChevronDown} class="h-6 w-6" />
      </MenuButton>
      <MenuItems
        class="z-10 bg-highlightPurple rounded-[12px] border-[1px] border-white/10 backdrop-saturate-[180%] backdrop-blur-sm pt-2"
      >
        {#each views as view}
          <MenuItem let:active>
            <div class="py-2 hover:bg-blue-700 pl-4">
              <p>{view}</p>
            </div>
          </MenuItem>
        {/each}
      </MenuItems>
    </Menu>
  </div>
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
</div>
