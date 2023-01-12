export const isEmailError = (errorMsg: string) =>
  errorMsg.toLowerCase().includes("email");

export const isPasswordError = (errorMsg: string) =>
  errorMsg.toLowerCase().includes("password");

export const isNameError = (errorMsg: string) =>
  errorMsg.toLowerCase().includes("name");

export const accentColors = [
  { label: "Blue", class: "bg-blue-600" },
  { label: "Pink", class: "bg-pink-500" },
];
