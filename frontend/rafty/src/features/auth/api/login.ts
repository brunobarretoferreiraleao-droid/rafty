import { api } from "@/services/api";

export async function login(data: {
  email: string;
  password: string;
}) {
  const response = await api.post("/auth/login", data);

  return response.data;
}