<script lang="ts">
  import {
    Dialog,
    DialogOverlay,
    DialogTitle,
    DialogDescription,
  } from "@rgossiaux/svelte-headlessui";
  export let isOpen;
</script>

<Dialog
  open={isOpen}
  on:close={() => (isOpen = false)}
  class="fixed inset-0 p-4 pt-[25vh] "
>
  <DialogOverlay />
  <div class="bg-pink-200">
    <DialogTitle>Deactivate account</DialogTitle>
    <DialogDescription>
      This will permanently deactivate your account
    </DialogDescription>

    <p>
      Are you sure you want to deactivate your account? All of your data will be
      permanently removed. This action cannot be undone.
    </p>

    <button on:click={() => (isOpen = false)}>Deactivate</button>
    <button on:click={() => (isOpen = false)}>Cancel</button>
  </div>
</Dialog>
