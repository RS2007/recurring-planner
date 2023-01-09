import { writable } from "svelte/store";

export const todos = writable([]);

export const addTodo = (todo) => {
  todos.update((todo) => [...todos.get(), todo]);
};
