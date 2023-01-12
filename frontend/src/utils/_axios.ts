import Axios from "axios";

const baseURL = "http://localhost:5000";
export const _axios = Axios.create({ baseURL, withCredentials: true });
