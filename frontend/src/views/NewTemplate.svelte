<script>
  import {
    Icon,
    LocationMarker,
    ViewList,
    Calendar,
    Plus,
  } from "svelte-hero-icons";
  import CreateTemplateModal from "./CreateTemplateModal.svelte";
  let isOpen = false;

  let events = [
    {
      summary: "Google I/O 2015",
      description: "800 Howard St., San Francisco, CA 94103",
      location: "A chance to hear more about Google's developer products.",
      startTime: "9 am",
      endTime: "5 pm",
    },
    {
      summary: "Google I/O 2015",
      description: "800 Howard St., San Francisco, CA 94103",
      location: "A chance to hear more about Google's developer products.",
      startTime: "9 am",
      endTime: "5 pm",
    },
  ];

  const openModal = (e) => {
    e.preventDefault();
    isOpen = true;
  };
</script>

<div class="col-span-12  pl-4 pt-8">
  <h1 class="text-3xl mb-8">Create a new template</h1>
  <form class="flex flex-col">
    <div class="flex gap-6">
      <div class="flex flex-col">
        <label>Name</label>
        <input
          type="text"
          class="pl-2 bg-highlightPurple w-[320px] min-h-[40px] rounded-lg"
        />
      </div>
      <div class="flex flex-col">
        <label>Tags</label>
        <input
          type="text"
          class="pl-2 bg-highlightPurple w-[320px] min-h-[40px] rounded-lg"
        />
      </div>
    </div>
    <div class="flex text-[24px] mt-2 mb-4 gap-4 items-center">
      <h1>Events</h1>
      <button class="bg-blue-600 px-2 h-6 rounded-md" on:click={openModal}
        ><Icon src={Plus} class="h-4 w-4" /></button
      >
      <CreateTemplateModal bind:isOpen bind:events />
    </div>
    <div class="flex gap-6">
      {#each events as event}
        <div
          class="pt-4 bg-highlightPurple h-[220px] w-[400px]  backdrop-saturate-[180%] backdrop-blur-sm text-md"
          style={`background: linear-gradient(105.69deg, #181E4C 1.05%, rgba(110, 123, 248, 0.59) 161.32%);
border-radius: 10px;`}
        >
          <h1 class="ml-4 text-[22px] font-extrabold">{event.summary}</h1>
          <hr class="border-white/10" />
          <div class="pt-4 pl-4">
            <div class="flex gap-2">
              <Icon src={LocationMarker} class="h-6 w-6 stroke-green-400" />
              <div class="flex items-center">
                <p class="text-[12px]">{event.location}</p>
              </div>
            </div>
          </div>
          <div class="pt-4 pl-4">
            <div class="flex gap-2">
              <Icon src={ViewList} class="h-6 w-6 stroke-fuchsia-400" />
              <div class="flex items-center">
                <p class="text-[12px]">
                  {event.description}
                </p>
              </div>
            </div>
          </div>
          <div class="pt-4 pl-4">
            <div class="flex gap-2">
              <Icon src={Calendar} class="h-5 w-5 stroke-red-400" />
              <div class="flex items-center">
                <p class="text-[12px]">{event.startTime} - {event.endTime}</p>
              </div>
            </div>
          </div>
        </div>
      {/each}
    </div>
    <button type="submit" class="py-2 bg-blue-600 mt-4 max-w-[200px] rounded-lg"
      >Create Template</button
    >
  </form>
</div>
