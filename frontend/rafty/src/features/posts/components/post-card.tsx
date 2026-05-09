type PostCardProps = {
  post: {
    id: string;
    content?: string;
    title?: string;
    created_at: string;
  };
};

export function PostCard({ post }: PostCardProps) {
  return (
    <article className="rounded-3xl border border-white/10 bg-white/3 p-6 backdrop-blur-xl">
      {post.title && (
        <h2 className="text-xl font-bold">
          {post.title}
        </h2>
      )}

      {post.content && (
        <p className="mt-3 whitespace-pre-wrap text-zinc-300">
          {post.content}
        </p>
      )}

      <div className="mt-6 text-xs text-zinc-500">
        {new Date(post.created_at).toLocaleString()}
      </div>
    </article>
  );
}