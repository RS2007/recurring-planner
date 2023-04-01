import { writable, type Writable } from "svelte/store";
import type { templateType } from "../utils/types";

export const templatesStore: Writable<Array<templateType>> = writable([]);
