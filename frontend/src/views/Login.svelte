<script lang="ts">
  import { toast } from "@zerodevx/svelte-toast";
  import * as yup from "yup";
  import { user } from "../store/userStore";
  import {
    accentColors,
    isEmailError,
    isPasswordError,
  } from "../utils/formUtils";
  import { useNavigate } from "svelte-navigator";
  import { _axios } from "../utils/_axios";
  type userCtx = {
    email: string;
    name: string;
    accentColor: "Blue" | "Pink";
  };

  function getUser(): userCtx {
    let $user: userCtx;
    user.subscribe(($) => ($user = $ as userCtx))();
    return $user;
  }
  const navigate = useNavigate();

  const schema = yup.object().shape({
    email: yup.string().required("Email is required").email("Email is invalid"),
    password: yup.string().required("Password is required"),
  });
  let email = "";
  let password = "";

  let emailError = "";
  let passwordError = "";
  const submitData = async (e) => {
    e.preventDefault();
    try {
      const validationResult = await schema.validate(
        { email, password },
        { abortEarly: false }
      );
      const res = await _axios.post("/auth/login", validationResult);
      if (res.status === 200) {
        user.set({
          email,
          name: res.data.name,
          accentColor: res.data.accentColor,
        });
        console.log(getUser());
        localStorage.setItem("TOKEN_RP", res.data.token);
        toast.push("Login succesful");
        navigate("/dashboard");
      }
    } catch (e: any) {
      if (e.errors) {
        e.errors.forEach((error) => {
          if (isEmailError(error)) emailError = error;
          if (isPasswordError(error)) passwordError = error;
        });
      } else {
        toast.push(e.message);
      }
    }
  };
</script>

<form on:submit={submitData}>
  <div class="flex flex-col mb-2 gap-2">
    <label>Email</label>
    <input
      type="text"
      class="bg-backgroundDashboard min-w-[320px] min-h-[40px] rounded-lg pl-2"
      bind:value={email}
    />
    {#if emailError.length != 0}
      <div class="text-red-500">{emailError}</div>
    {/if}
  </div>
  <div class="flex flex-col gap-2 mb-6">
    <label>Password</label>
    <input
      type="password"
      class="bg-backgroundDashboard min-w-[320px] min-h-[40px] rounded-lg pl-2"
      bind:value={password}
    />
    {#if passwordError.length != 0}
      <div class="text-red-500">{passwordError}</div>
    {/if}
  </div>
  <button
    class={`w-full ${
      accentColors.find(
        (accentColor) => accentColor.label === getUser().accentColor
      ).class
    } py-2 rounded-md`}>Login</button
  >
</form>
