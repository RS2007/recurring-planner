<script lang="ts">
  import { Icon, FolderAdd } from "svelte-hero-icons";
  import { useNavigate } from "svelte-navigator";
  import TemplateCard from "./TemplateCard.svelte";
  import { dndzone } from "svelte-dnd-action";
  import { _axios } from "../utils/_axios";

  type schedule = {
    title: string;
    startTime: string;
    endTime: string;
  };

  type templateType = {
    id: number;
    title: string;
    schedule: Array<schedule>;
    expanded: boolean;
  };

  const navigate = useNavigate();
  let templates: Array<templateType>;
  let isLoading = false;

  const fetchTemplates = async () => {
    isLoading = true;
    templates = (
      await _axios.get("/template/all", {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("TOKEN_RP")}`,
        },
      })
    ).data.templates as any;
    templates = templates.map((templateObject) => {
      templateObject.id = templateObject.templateId;
      delete templateObject["templateId"];
      return templateObject;
    });
    console.log(templates);
    isLoading = false;
  };

  fetchTemplates();

  const handleSort = (e) => {
    templates = e.detail.items;
  };
</script>

<div class="pl-4 col-span-3 pt-8  " style="border:1px solid #2B2B2B">
  <div class="mb-8 flex justify-between pr-4 items-center">
    <p class="text-2xl">My templates</p>
    <button
      on:click={() => {
        navigate("/dashboard/template/new");
      }}
    >
      <Icon src={FolderAdd} class="h-6 w-6" />
    </button>
  </div>
  {#if isLoading}
    <h1>Loading...</h1>
  {:else}
    <div
      class="min-h-[80vh]"
      use:dndzone={{ items: templates, type: "templates" }}
      on:consider={handleSort}
      on:finalize={handleSort}
    >
      {#each templates as template (template.id)}
        <TemplateCard {template} />
      {/each}
    </div>
  {/if}
</div>
