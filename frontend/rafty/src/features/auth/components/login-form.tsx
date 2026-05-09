"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

import { login } from "../api/login";
import { useAuthStore } from "../store/auth-store";

import axios from "axios";

export function LoginForm() {
  const router = useRouter();

  const setAuth = useAuthStore((state) => state.setAuth);

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();

    try {
      setLoading(true);
      setError("");

      const response = await login({
        email,
        password,
      });

      setAuth(response.data.token, response.data.user);

      router.push("/feed");
    } catch (err: unknown) {
        if (axios.isAxiosError(err)) {
          setError(err.response?.data?.error || "Erro ao logar");
        } else {
          setError("Erro ao logar");
        }  
    } finally {
      setLoading(false);
    }
  }

  return (
    <form
      onSubmit={handleSubmit}
      className="w-full max-w-md rounded-3xl border border-white/10 bg-white/5 p-8 backdrop-blur-xl"
    >
      <h1 className="text-3xl font-bold">🌊 Login</h1>

      <div className="mt-8 space-y-4">
        <input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="w-full rounded-2xl border border-white/10 bg-black/20 px-4 py-3 outline-none"
        />

        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="w-full rounded-2xl border border-white/10 bg-black/20 px-4 py-3 outline-none"
        />
      </div>

      {error && (
        <p className="mt-4 text-sm text-red-400">
          {error}
        </p>
      )}

      <button
        disabled={loading}
        className="mt-6 w-full rounded-2xl bg-blue-500 py-3 font-semibold transition hover:bg-blue-400 disabled:opacity-50"
      >
        {loading ? "Loading..." : "Sign In"}
      </button>
    </form>
  );
}