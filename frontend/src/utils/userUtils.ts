import { userStore } from "../store/userStore";
import { getCookie } from "svelte-cookie";
import type { userCtx } from "./types";

export const getUser = (): userCtx => {
	let $user: userCtx;
	userStore.subscribe(($) => ($user = $ as userCtx))();
	return $user;
};

export const isLoggedIn = (): boolean => {
	const authCookie = getCookie("RP_USER_DETAILS");
	console.log({ authCookie });
	return !!authCookie;
};

export const getGoogleToken = () => {
	const { googleAccessToken: accessToken, googleIdToken: idToken } = JSON.parse(
		getCookie("RP_USER_DETAILS")
	);
	return { accessToken, idToken };
};

export const setUser = () => {
	const { email, name } = JSON.parse(getCookie("RP_USER_DETAILS"));
	userStore.set({ email, name, accentColor: { label: "Blue", class: "bg-blue-600" } });
};

export const userAuthToken = () => {
	const authCookie = JSON.parse(getCookie("RP_USER_DETAILS"));
	return authCookie?.token;
};

export const logOut = () => {
	document.cookie = "RP_USER_DETAILS" + "=;expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;";
	window.location.reload();
};
