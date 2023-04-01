/** @type {import('tailwindcss').Config} */
module.exports = {
	darkMode: true,
	content: ["./src/**/*.{html,js,svelte,ts}"],
	theme: {
		extend: {
			colors: {
				sidebarBg: "#111536",
				backgroundDashboard: "#1c1c1c",
				borderColor: "#e5e5e5",
				highlightPurple: "#2e2e2e",
				supabaseGreen: "#3fcf8e",
			},
		},
	},
	plugins: [],
};
