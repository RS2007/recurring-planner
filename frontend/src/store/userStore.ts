import { writable, type Writable } from "svelte/store";

export const userStore: Writable<{
	name: string;
	email: string;
	accentColor: { label: string; class: "bg-blue-600" | "bg-pink-500" };
}> = writable({
	name: "",
	email: "",
	accentColor: { label: "Blue", class: "bg-blue-600" },
});
