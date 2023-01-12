<script lang="ts">
  import { toast } from "@zerodevx/svelte-toast";
  import {
    Listbox,
    ListboxButton,
    ListboxOptions,
    ListboxOption,
  } from "@rgossiaux/svelte-headlessui";
  import { Icon, ChevronDown } from "svelte-hero-icons";
  import { fade } from "svelte/transition";
  import { _axios } from "../utils/_axios";
  import * as yup from "yup";
  import {
    accentColors,
    isEmailError,
    isNameError,
    isPasswordError,
  } from "../utils/formUtils";
  import { useNavigate } from "svelte-navigator";

  const navigate = useNavigate();

  const schema = yup.object().shape({
    email: yup.string().required("Email is required").email("Email is invalid"),
    password: yup.string().required("Password is required"),
    name: yup.string().required("Name is required"),
    accentColor: yup.object().shape({
      label: yup.string().oneOf(["Blue", "Pink"]),
      class: yup.string().oneOf(["bg-blue-600", "bg-pink=500"]),
    }),
  });

  let name = "";
  let email = "";
  let password = "";
  let accentColor = { label: "Blue", class: "bg-blue-600" };
  let nameError = "";
  let emailError = "";
  let passwordError = "";

  const changeAccentColor = (e) => {
    accentColor = accentColors.find(
      (accentColor) => accentColor.label === e.detail
    );
  };

  const submitData = async (e) => {
    e.preventDefault();
    try {
      await schema.validate({
        email,
        name,
        password,
        accentColor,
      });
      const res = await _axios.post("/auth/register", {
        name,
        email,
        password,
        accentColor: accentColor.label,
      });
      console.log(res);
      toast.push("Succesful Registration");
      navigate("/dashboard");
    } catch (e: any) {
      if (e.errors) {
        e.errors.forEach((error) => {
          if (isEmailError(error)) emailError = error;
          if (isPasswordError(error)) passwordError = error;
          if (isNameError(error)) nameError = error;
        });
      }
      console.log(`Error is ${e.message}`);
      toast.push("Error");
    }
  };
</script>

<form class="w-full" on:submit={submitData}>
  <div class="grid grid-cols-12 gap-y-2 h-fit">
    <div class="flex flex-col col-start-1 col-end-6">
      <label>Email</label>
      <input
        type="text"
        class="pl-2 bg-backgroundDashboard w-[320px] min-h-[40px] rounded-lg"
        bind:value={email}
      />
    </div>
    <div class="flex flex-col col-start-8 col-end-13">
      <label>Name</label>
      <input
        type="text"
        class="pl-2 bg-backgroundDashboard w-[320px] min-h-[40px] rounded-lg"
        bind:value={name}
      />
    </div>
    <div class="flex flex-col col-start-1 col-end-6 h-[80px]">
      <label>Password</label>
      <input
        class="pl-2 w-[320px] h-[40px] rounded-lg bg-backgroundDashboard"
        type="password"
        bind:value={password}
      />
    </div>
    <div class="flex flex-col col-start-8 col-end-13 h-[50px]">
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
      <button
        type="submit"
        class={`w-full ${accentColor.class} col-span-12 py-2 rounded-lg`}
        >Register</button
      >
    </div>
  </div>
</form>
