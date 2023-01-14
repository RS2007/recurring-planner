<script lang="ts">
  import { Icon, ChevronUp, ChevronDown } from "svelte-hero-icons";
  import { fade } from "svelte/transition";
  import { themeColor } from "../utils/userUtils";
  type schedule = {
    title: string;
    startTime: string;
    endTime: string;
  };
  type templateType = {
    title: string;
    schedule: Array<schedule>;
    expanded: boolean;
  };

  export let template: templateType;

  const collapseTemplate = () => {
    console.log("TEmplate");
    template.expanded = !template.expanded;
  };
</script>

<div
  class="mb-4 bg-highlightPurple mr-4  backdrop-saturate-[180%] backdrop-blur-sm"
  style={`background: linear-gradient(105.69deg, #181E4C 1.05%, rgba(110, 123, 248, 0.59) 200.32%);
border-radius: 10px;`}
>
  <button on:click={collapseTemplate} class="w-full">
    <div class="flex justify-between pr-4 pl-4 py-3">
      <p class="text-lg">{template.title}</p>
      <Icon src={template.expanded ? ChevronUp : ChevronDown} class="h-6 w-6" />
    </div>
  </button>
  {#if template.expanded}
    <div class={` mr-4 rounded-lg`} transition:fade>
      {#each template.schedule as event}
        <div class="pl-4 py-2 flex flex-col mb-2">
          <div class=" flex justify-between">
            <p class="text-md">{event.title}</p>
            <p class="text-md text-gray-400">
              {event.startTime} - {event.endTime}
            </p>
          </div>
          <!-- {#if index != template.schedule.length - 1} -->
          <!--   <hr class="opacity-50" /> -->
          <!-- {/if} -->
          <div class={`${themeColor()} rounded-lg px-2 py-1 w-fit text-base`}>
            First Floor
          </div>
        </div>
      {/each}
    </div>
  {/if}
</div>
