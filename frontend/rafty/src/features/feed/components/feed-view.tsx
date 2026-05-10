"use client";

import { CreatePostCard } from "./create-post-card";

import { useFeed } from "../hooks/use-feed";

import { PostCard } from "@/features/posts/components/post-card";

export function FeedView() {
  const { data, isLoading } = useFeed();

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