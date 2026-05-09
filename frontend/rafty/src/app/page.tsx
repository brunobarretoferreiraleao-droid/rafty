import Link from "next/link";

export default function HomePage() {
  return (
    <main className="relative flex min-h-screen items-center justify-center overflow-hidden bg-[#071018] px-6">
      <div className="absolute inset-0 bg-[radial-gradient(circle_at_top,#12345633,transparent_60%)]" />

      <div className="z-10 flex max-w-2xl flex-col items-center text-center">
        <h1 className="text-6xl font-black tracking-tight">
          🌊 Rafty
        </h1>

        <p className="mt-6 text-lg text-zinc-300">
          The internet, flowing better.
        </p>

        <p className="mt-3 max-w-xl text-sm leading-relaxed text-zinc-500">
          Uma rede social feita para navegar conteúdo,
          não sobreviver a ele.
        </p>

        <div className="mt-10 flex gap-4">
          <Link
            href="/register"
            className="rounded-2xl bg-blue-500 px-6 py-3 font-medium transition hover:scale-[1.02] hover:bg-blue-400"
          >
            Enter the Ocean
          </Link>

          <Link
            href="/login"
            className="rounded-2xl border border-white/10 bg-white/5 px-6 py-3 font-medium backdrop-blur-md transition hover:bg-white/10"
          >
            Login
          </Link>
        </div>
      </div>
    </main>
  );
}