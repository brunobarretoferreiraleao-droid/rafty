"use client";

import { CreatePostCard } from "./create-post-card";

import { useFeed } from "../hooks/use-feed";

import { PostCard } from "@/features/posts/components/post-card";

import { useHydration } from "@/hooks/use-hydration";

export function FeedView() {
  const hydrated = useHydration();

  const { data, isLoading, error } = useFeed(hydrated);

  if (!hydrated) {
    return null;
  }

  console.log(error);

  return (
    <div className="mx-auto flex w-full max-w-3xl flex-col gap-6 px-6 py-8">
      <CreatePostCard />

      {isLoading && (
        <p className="text-zinc-400">
          Loading feed...
        </p>
      )}

      {data?.map((post) => (
        <PostCard
          key={post.id}
          post={post}
        />
      ))}
    </div>
  );
}