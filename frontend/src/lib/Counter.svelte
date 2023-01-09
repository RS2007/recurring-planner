<script lang="ts">
  let count: number = 0;
  const increment = () => {
    count += 1;
  };
  let todos: string[] = ["First todo"];
  let newTodo: string = "";
  const addTodo = () => {
    todos = [...todos, newTodo];
    newTodo = "";
  };
</script>

<div class="w-full py-3">
  <input
    bind:value={newTodo}
    type="text"
    placeholder="Enter todo"
    class="h-12 p-2"
  />
  <button on:click={addTodo} class="h-12 ml-2">Add todo</button>
</div>
{newTodo}
<div class="flex flex-col text-white">
  {#each todos as item,index}
    <span>{index+1}{"."}{item}</span>
  {/each}
</div>

<button on:click={increment}>
  count is {count}
</button>
