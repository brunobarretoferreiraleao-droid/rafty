import { useQuery } from "@tanstack/react-query";

import type { Post } from "@/types/post";

import { api } from "@/services/api";

export function useFeed(enabled = true) {
  return useQuery<Post[]>({
    queryKey: ["feed"],

    enabled,

    queryFn: async () => {
      const response = await api.get<{ data: Post[] }>("/posts/feed");

      return response.data.data;
    },
  });
}