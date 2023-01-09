<script lang="ts">
  import { Icon, FolderAdd, ChevronUp, ChevronDown } from "svelte-hero-icons";
  type schedule = {
    title: string;
    startTime: string;
    endTime: string;
  };
  type template = {
    title: string;
    schedule: Array<schedule>;
    expanded: boolean;
  };
  const templates: Array<template> = [
    {
      title: "Example Template 1",
      schedule: [
        { title: "Go to library", startTime: "19:30", endTime: "20:30" },
        { title: "Go to library", startTime: "19:30", endTime: "20:30" },
        { title: "Go to library", startTime: "19:30", endTime: "20:30" },
        { title: "Go to library", startTime: "19:30", endTime: "20:30" },
      ],
      expanded: true,
    },
    {
      title: "Example Template 1",
      schedule: [
        { title: "Go to library", startTime: "19:30", endTime: "20:30" },
        { title: "Go to library", startTime: "19:30", endTime: "20:30" },
        { title: "Go to library", startTime: "19:30", endTime: "20:30" },
        { title: "Go to library", startTime: "19:30", endTime: "20:30" },
      ],
      expanded: false,
    },
  ];
</script>

<div class="pl-4 col-span-3 pt-8  " style="border:1px solid #2B2B2B">
  <div class="mb-8 flex justify-between pr-4 items-center">
    <p class="text-2xl">My templates</p>
    <Icon src={FolderAdd} class="h-6 w-6" />
  </div>
  {#each templates as template}
    <div
      class="mb-4 bg-highlightPurple mr-4 rounded-[12px] border-[1px] border-white/10 backdrop-saturate-[180%] backdrop-blur-sm"
    >
      <div class="flex justify-between pr-4 pl-4 py-3">
        <p class="text-lg">{template.title}</p>
        <Icon
          src={template.expanded ? ChevronUp : ChevronDown}
          class="h-6 w-6"
        />
      </div>
      <div
        class={`bg-highlightPurple mr-4 rounded-lg ${
          template.expanded ? "block" : "hidden"
        }`}
      >
        {#each template.schedule as event, index}
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
            <div class="bg-blue-600/70 rounded-lg px-2 py-1 w-fit text-base">
              First Floor
            </div>
          </div>
        {/each}
      </div>
    </div>
  {/each}
</div>
