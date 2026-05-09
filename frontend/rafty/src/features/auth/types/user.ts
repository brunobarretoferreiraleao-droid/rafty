export type User = {
  id: string;
  username: string;
  email: string;

  display_name?: string;
  bio?: string;
  avatar_url?: string;

  is_private: boolean;

  created_at: string;
  updated_at: string;
};