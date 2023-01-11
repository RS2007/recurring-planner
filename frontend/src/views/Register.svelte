<script lang="ts">
  import {
    Listbox,
    ListboxButton,
    ListboxOptions,
    ListboxOption,
  } from "@rgossiaux/svelte-headlessui";
  import { Icon, ChevronDown } from "svelte-hero-icons";
  import { fade } from "svelte/transition";
  let accentColor = { label: "Blue", class: "bg-blue-600" };

  const accentColors = [
    { label: "Blue", class: "bg-blue-600" },
    { label: "Pink", class: "bg-pink-500" },
  ];

  const changeAccentColor = (e) => {
    accentColor = accentColors.find(
      (accentColor) => accentColor.label === e.detail
    );
  };
</script>

<form class="w-full">
  <div class="grid grid-cols-12 gap-y-2 h-fit">
    <div class="flex flex-col col-start-1 col-end-6">
      <label>Email</label>
      <input
        type="text"
        class="bg-backgroundDashboard w-[320px] min-h-[40px] rounded-lg"
      />
    </div>
    <div class="flex flex-col col-start-8 col-end-13">
      <label>Name</label>
      <input
        type="text"
        class="bg-backgroundDashboard w-[320px] min-h-[40px] rounded-lg"
      />
    </div>
    <div class="flex flex-col col-start-1 col-end-6 h-[80px]">
      <label>Password</label>
      <input class="w-[320px] h-[40px] rounded-lg bg-backgroundDashboard" />
    </div>
    <div class="flex flex-col col-start-8 col-end-13 h-[50px]">
      <label>Accent color</label>
      <Listbox value={accentColor} on:change={changeAccentColor}>
        <ListboxButton
          class={`${accentColor.class} px-3 pt-2 pb-1 rounded-lg flex gap-2`}
          ><span>{accentColor.label}</span><Icon
            src={ChevronDown}
            class="h-6 w-6"
          /></ListboxButton
        >
        <div transition:fade class="z-10">
          <ListboxOptions
            class={`bg-backgroundDashboard delay-100  rounded-lg border-[1px] border-white/10 backdrop-saturate-[180%] backdrop-blur-sm h-[40px] z-10`}
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
    <div class="col-span-12 mt-4">
      <button class="w-full bg-blue-600 col-span-12 py-2 rounded-lg"
        >Register</button
      >
    </div>
  </div>
</form>
