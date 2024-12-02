// import { Server } from "@/api/server";

const SERVER_HOST = import.meta.env.VITE_SERVER_HOST;
const SERVER_URL = `${SERVER_HOST}`;

console.log("[SERVER] ", SERVER_URL);

// export const server = new Server({
//   baseURL: SERVER_URL,
//   withCredentials: true,
// });
