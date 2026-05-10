"use client";

import Link from "next/link";

import { useRouter } from "next/navigation";

import { useAuthStore } from "@/features/auth/store/auth-store";

export function Navbar() {
  const router = useRouter();

  const logout = useAuthStore((state) => state.logout);

  function handleLogout() {
    logout();
    router.push("/login");
  }

  return (
    <header className="sticky top-0 z-50 border-b border-white/10 bg-[#071018]/80 backdrop-blur-xl">
      <div className="mx-auto flex h-16 max-w-6xl items-center justify-between px-6">
        <Link
          href="/feed"
          className="text-xl font-black"
        >
          🌊 Rafty
        </Link>

        <nav className="flex items-center gap-6 text-sm text-zinc-300">
          <Link href="/feed">Feed</Link>

          <button
            onClick={handleLogout}
            className="rounded-xl border border-white/10 px-4 py-2 transition hover:bg-white/5"
          >
            Logout
          </button>
        </nav>
      </div>
    </header>
  );
}