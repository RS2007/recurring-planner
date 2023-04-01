import Axios from "axios";

const baseURL = "http://localhost:5000/api";
export const _axios = Axios.create({ baseURL, withCredentials: true });
