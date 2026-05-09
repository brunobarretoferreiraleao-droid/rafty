import axios from "axios";

export const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
});

api.interceptors.request.use((config) => {
  if (typeof window !== "undefined") {
    const storage = localStorage.getItem("rafty-auth-store");

    if (storage) {
      const parsed = JSON.parse(storage);

      const token = parsed.token;

      if (token) {
        config.headers.Authorization = `Bearer ${parsed.token}`;
      }
    }
  }
  
  return config;
});