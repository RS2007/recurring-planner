<script lang="ts">
  import { Icon, FolderAdd, ChevronUp, ChevronDown } from "svelte-hero-icons";
  import Modal from "../components/Modal.svelte";
  import TemplateCard from "./TemplateCard.svelte";

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

  let isOpen = false;

  const templates: Array<templateType> = [
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
  ];
</script>

<div class="pl-4 col-span-3 pt-8  " style="border:1px solid #2B2B2B">
  <div class="mb-8 flex justify-between pr-4 items-center">
    <p class="text-2xl">My templates</p>
    <button
      on:click={() => {
        isOpen = true;
      }}
    >
      <Icon src={FolderAdd} class="h-6 w-6" />
    </button>
  </div>
  <Modal {isOpen} />
  {#each templates as template}
    <TemplateCard {template} />
  {/each}
</div>
