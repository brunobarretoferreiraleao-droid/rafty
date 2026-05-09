"use client";

import { useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";

import { api } from "@/services/api";

export function CreatePostCard() {
  const queryClient = useQueryClient();

  const [content, setContent] = useState("");

  const mutation = useMutation({
    mutationFn: async () => {
      return api.post("/posts", {
        content,
        visibility: "public",
      });
    },

    onSuccess: () => {
      setContent("");

      queryClient.invalidateQueries({
        queryKey: ["feed"],
      });
    },
  });

  return (
    <div className="rounded-3xl border border-white/10 bg-white/3 p-6 backdrop-blur-xl">
      <textarea
        placeholder="What is flowing through your mind?"
        value={content}
        onChange={(e) => setContent(e.target.value)}
        className="min-h-30 w-full resize-none bg-transparent outline-none"
      />

      <div className="mt-4 flex justify-end">
        <button
          onClick={() => mutation.mutate()}
          disabled={!content || mutation.isPending}
          className="rounded-2xl bg-blue-500 px-5 py-2 font-medium transition hover:bg-blue-400 disabled:opacity-50"
        >
          Post
        </button>
      </div>
    </div>
  );
}