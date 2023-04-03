import Axios from "axios";

const baseURL = import.meta.env.PROD
	? "https://recurring-planner-monorepo.onrender.com/api"
	: "http://localhost:5000/api";
export const _axios = Axios.create({ baseURL, withCredentials: true });
