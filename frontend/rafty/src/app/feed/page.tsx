import { FeedView } from "@/features/feed/components/feed-view";

import { Navbar } from "@/components/layout/navbar";
import { Sidebar } from "@/components/layout/sidebar";

export default function FeedPage() {
  return (
    <main className="min-h-screen bg-[#071018] text-white">
      <Navbar />

      <div className="mx-auto flex max-w-7xl">
        <Sidebar />

        <FeedView />
      </div>
    </main>
  );
}