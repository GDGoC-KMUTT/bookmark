import { Server } from "@/api/api"

export const SERVER_HOST = import.meta.env.VITE_SERVER_HOST
const ENVIRONMENT = import.meta.env.VITE_ENVIRONMENT
const SERVER_URL = `${SERVER_HOST}`

if (ENVIRONMENT === "0") {
    console.log("[SERVER] ", SERVER_URL)
}

export const server = new Server({
    baseURL: SERVER_URL,
    withCredentials: true,
})

