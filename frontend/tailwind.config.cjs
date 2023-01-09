/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: true,
  content: ["./src/**/*.{html,js,svelte,ts}"],
  theme: {
    extend: {
      colors: {
        sidebarBg: "#111536",
        backgroundDashboard: "#111536",
        borderColor: "#e5e5e5",
        highlightPurple: "#181e4c",
      },
    },
  },
  plugins: [],
};
