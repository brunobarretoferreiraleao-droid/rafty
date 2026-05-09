import { useQuery } from "@tanstack/react-query";

import { api } from "@/services/api";

export function useFeed() {
  return useQuery({
    queryKey: ["feed"],

    queryFn: async () => {
      const response = await api.get("/posts/feed");

      return response.data.data;
    },
  });
}