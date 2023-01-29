<script>
  import {
    Listbox,
    ListboxButton,
    ListboxOption,
    ListboxOptions,
  } from "@rgossiaux/svelte-headlessui";
  import { Icon, Mail, Tag, UserCircle, ChevronDown } from "svelte-hero-icons";
  import { accentColors } from "../utils/formUtils";
  import { fade } from "svelte/transition";
  import { getUser } from "../utils/userUtils";
  let accentColor = accentColors.find(
    (color) => color.label === getUser().accentColor
  );
  const changeAccentColor = (e) => {
    accentColor = accentColors.find(
      (accentColor) => accentColor.label === e.detail
    );
  };
</script>

<div class="col-span-12 flex pt-8 pl-4 gap-6">
  <div
    class="pt-4 bg-highlightPurple h-[220px] w-[450px]  backdrop-saturate-[180%] backdrop-blur-sm text-md flex flex-col"
    style={`background: linear-gradient(105.69deg, #181E4C 1.05%, rgba(110, 123, 248, 0.59) 161.32%);
border-radius: 12px;`}
  >
    <div class="flex flex-col">
      <h1 class="text-2xl font-bold pl-4">Profile</h1>

      <hr class="border-white/10" />
    </div>
    <div class="flex">
      <img
        src="https://api.dicebear.com/5.x/adventurer/svg?seed=Charlie"
        class="h-[40%]"
        alt="a dude"
      />
      <div class="flex flex-col gap-2 justify-center h-[40%]">
        <div class="flex gap-2 ">
          <Icon src={UserCircle} class="h-6 w-6 stroke-red-400" />
          <div>RS2007</div>
        </div>
        <div class="flex gap-2 ">
          <Icon src={Mail} class="h-6 w-6 stroke-green-400" />
          <div>mario@gmail.com</div>
        </div>
        <div class="flex gap-2 ">
          <Icon src={Tag} class="h-6 w-6 stroke-fuchsia-400" />
          <div>15</div>
        </div>
      </div>
    </div>
  </div>
  <div
    class="pt-4 bg-highlightPurple h-[220px] w-fit  backdrop-saturate-[180%] backdrop-blur-sm text-md flex flex-col"
    style={`background: linear-gradient(105.69deg, #181E4C 1.05%, rgba(110, 123, 248, 0.59) 161.32%);
border-radius: 12px;`}
  >
    <h1 class="text-2xl px-4">Customize</h1>

    <hr class="border-white/10" />
    <div class="flex gap-4 px-4 pt-2 min-w-[450px] gap-8">
      <div class="flex flex-col">
        <label>Accent color</label>
        <Listbox bind:value={accentColor} on:change={changeAccentColor}>
          <ListboxButton
            class={`${accentColor.class} px-3 pt-2 pb-1 rounded-lg flex gap-2`}
            ><span>{accentColor.label}</span><Icon
              src={ChevronDown}
              class="h-6 w-6"
            /></ListboxButton
          >
          <div transition:fade class="z-10">
            <ListboxOptions
              class={`bg-backgroundDashboard delay-100  rounded-lg border-[1px] border-white/10 backdrop-saturate-[180%] backdrop-blur-sm h-fit z-10`}
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
      <div class="flex flex-col">
        <label>Sort templates By</label>
        <Listbox bind:value={accentColor} on:change={changeAccentColor}>
          <ListboxButton
            class={`${accentColor.class} px-3 pt-2 pb-1 rounded-lg flex gap-2`}
            ><span>{accentColor.label}</span><Icon
              src={ChevronDown}
              class="h-6 w-6"
            /></ListboxButton
          >
          <div transition:fade class="z-10">
            <ListboxOptions
              class={`bg-backgroundDashboard delay-100  rounded-lg border-[1px] border-white/10 backdrop-saturate-[180%] backdrop-blur-sm h-fit z-10`}
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
    </div>
  </div>
</div>
