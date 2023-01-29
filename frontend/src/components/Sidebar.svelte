<script lang="ts">
  import {
    Icon,
    Calendar,
    Cog,
    UserCircle,
    Logout,
    type IconSource,
  } from "svelte-hero-icons";
  import { useNavigate } from "svelte-navigator";
  import { themeColor } from "../utils/userUtils";
  type optionsType = {
    title: string;
    icon: IconSource;
    active?: boolean;
    link?: string;
  };
  const options: Array<optionsType> = [
    { title: "Calendar", icon: Calendar, active: true, link: "/dashboard" },
    { title: "Customize", icon: Cog, link: "/dashboard/customize" },
    { title: "Log out", icon: Logout },
  ];
  const navigate = useNavigate();
</script>

<div class="col-span-1 w-full bg-sidebarBg h-full pt-12">
  {#each options as option}
    <div
      class={`flex flex-col items-center w-[100%] mb-6 cursor-pointer ${
        option.active ? `text-blue-400` : "text-gray-400"
      }`}
    >
      <button
        class="w-full flex flex-col items-center "
        on:click={() => {
          navigate(option.link);
        }}
      >
        <div>
          <Icon src={option.icon} solid={false} class="h-8 w-8" />
        </div>
        <div>
          <p class={`align-middle text-sm `}>
            {option.title}
          </p>
        </div>
      </button>
    </div>
  {/each}
</div>
