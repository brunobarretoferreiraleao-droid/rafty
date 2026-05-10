export type Post = {
  id: string;
  author_id: string;

  title?: string | null;
  content?: string | null;

  visibility: "public" | "friends" | "custom" | "private";

  created_at: string;
  updated_at: string;
};