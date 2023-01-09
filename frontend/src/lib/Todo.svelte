<script>
import { todoStore } from "../store/todoStore";
let todo = "";
const createTodo = ()=>{
    todoStore.update(todos => {
        console.log(todos);
        return [...todos,todo]
    });
    todo = "";
}
</script>


<input type="text" placeholder="Enter todo" bind:value={todo}/>
<button on:click={createTodo} >Submit</button>
<ul class="text-white flex flex-col">
    {#each $todoStore as todo,index}
    <l1>{index+1}{"."}{todo}</l1>
    {/each}
    <li></li>
</ul>