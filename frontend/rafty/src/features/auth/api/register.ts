import { api } from "@/services/api";

export async function register(data: {
  username: string;
  email: string;
  password: string;
}) {
  const response = await api.post("/auth/register", data);

  return response.data;
}