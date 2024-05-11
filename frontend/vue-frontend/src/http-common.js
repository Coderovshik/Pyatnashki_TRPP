import axios from "axios";

const AXIOS = axios.create({
  baseURL: `http://localhost/:8080`,
  headers: {
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Methods": "DELETE, POST, GET, OPTIONS",
    "Access-Control-Allow-Headers":
      "Content-Type, Authorization, X-Requested-With",
  },
  methods: "*",
  crossDomain: true,
});
export default AXIOS;
