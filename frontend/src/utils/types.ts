import type { IconSource } from "svelte-hero-icons";

export type schedule = {
	summary: string;
	startTime: number;
	endTime: number;
	location: string;
};

export type templateType = {
	name: string;
	events: Array<schedule>;
	templateId: number;
	expanded: boolean;
};

export type userCtx = {
	email: string;
	name: string;
	accentColor: { label: string; class: "bg-blue-600" | "bg-pink-500" };
};

export type optionsType = {
	title: string;
	icon: IconSource;
	active?: boolean;
	link?: string;
	action?: () => void;
};

export type EventType = {
	description: string;
	endTime: number;
	location: string;
	startTime: number;
	summary: number;
	uuserId: string;
};

type _days =
	| "Monday"
	| "Tuesday"
	| "Wednesday"
	| "Thursday"
	| "Friday"
	| "Saturday"
	| "Sunday"
	| "";

export type eventFromServerType = {
	dayName: _days;
	backgroundColor?: string;
	day: number;
	startHour: number;
	endHour: number;
	startMinute: number;
	endMinute: number;
	title?: string;
	id: number;
	events: Array<EventType>;
};
