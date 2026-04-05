import Link from "next/link";

export default function HomePage() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center bg-zinc-950 px-6 text-white">
      <div className="w-full max-w-xl text-center">
        <h1 className="mb-4 text-5xl font-bold tracking-tight">Rafty</h1>
        <p className="mb-8 text-lg text-zinc-400">
          Sua rede social moderna. Ou pelo menos tentando.
        </p>

        <div className="flex justify-center gap-4">
          <Link
            href="/login"
            className="rounded-xl bg-white px-6 py-3 font-medium text-black transition hover:opacity-90"
          >
            Entrar
          </Link>

          <Link
            href="/register"
            className="rounded-xl border border-zinc-700 px-6 py-3 font-medium text-white transition hover:bg-zinc-900"
          >
            Criar conta
          </Link>
        </div>
      </div>
    </main>
  );
}