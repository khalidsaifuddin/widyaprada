"use client";

import { ErrorDialog } from "@/components";
import { api, app, ui } from "@/config";
import { storeAuthData } from "@/lib/auth";
import { ArrowRightIcon, EyeIcon, EyeSlashIcon } from "@heroicons/react/24/outline";
import { useRouter, useSearchParams } from "next/navigation";
import { Suspense, useEffect, useState } from "react";

function LoginForm() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const redirect = searchParams.get("redirect") || "/";

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [successOpen, setSuccessOpen] = useState(false);

  useEffect(() => {
    const check = async () => {
      try {
        const res = await fetch(`${api.baseUrl}/v1/auth/me`, {
          credentials: "include",
          headers: { Accept: "application/json" },
        });
        if (res.ok) router.replace(redirect);
      } catch {
        // not logged in
      }
    };
    check();
  }, [api.baseUrl, redirect, router]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setLoading(true);
    try {
      const res = await fetch(`${api.baseUrl}/v1/auth/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
          "X-Requested-With": "XMLHttpRequest",
        },
        body: JSON.stringify({
          user_name: username.trim(),
          user_password: password,
        }),
        credentials: "include",
      });
      const data = await res.json().catch(() => ({}));

      if (res.ok && (data.code === 200 || data.data?.access_token)) {
        const authData = {
          access_token: data.data?.access_token ?? data.data?.token ?? data.access_token,
          user_id: String(data.data?.user_id ?? data.data?.id ?? ""),
          user_name: data.data?.user_name ?? data.data?.username ?? username,
          user_nik: data.data?.user_nik ?? data.data?.nik ?? "",
          user_fullname: data.data?.user_fullname ?? data.data?.fullname ?? data.data?.name ?? username,
          expiry: data.data?.expires_at ?? data.data?.expiry ?? new Date(Date.now() + 86400000).toISOString(),
          role_user: data.data?.role_user ?? [],
        };
        await storeAuthData(authData);
        setSuccessOpen(true);
        setTimeout(() => {
          window.location.href = redirect;
        }, 800);
      } else {
        setError(data.message ?? data.status ?? "Login gagal. Periksa username dan password.");
      }
    } catch {
      setError("Koneksi gagal. Periksa jaringan Anda.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div
      className="min-h-screen flex items-center justify-center p-4"
      style={{
        background: `linear-gradient(to top, ${ui.theme.gradient.from}, ${ui.theme.gradient.to})`,
      }}
    >
      <div className="w-full max-w-md">
        <div className="text-center mb-8 md:hidden">
          <img src={ui.logo.src} alt={ui.logo.alt} className="h-14 w-auto mx-auto mb-2" />
          <h2 className="text-xl font-bold text-white">{app.name}</h2>
        </div>
        <div className="bg-white rounded-2xl shadow-xl p-8">
          <div className="text-center mb-6">
            <h2 className="text-2xl font-bold text-gray-900">Masuk</h2>
            <p className="text-gray-600 mt-1">Gunakan kredensial Anda</p>
          </div>
          <form onSubmit={handleSubmit} className="space-y-5">
            <div>
              <label htmlFor="username" className="block text-sm font-medium text-gray-700 mb-1">
                Username
              </label>
              <input
                id="username"
                type="text"
                required
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900"
                placeholder="Username"
                disabled={loading}
              />
            </div>
            <div>
              <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-1">
                Password
              </label>
              <div className="relative">
                <input
                  id="password"
                  type={showPassword ? "text" : "password"}
                  required
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="w-full px-4 py-3 pr-12 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900"
                  placeholder="Password"
                  disabled={loading}
                />
                <button
                  type="button"
                  onClick={() => setShowPassword((s) => !s)}
                  className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
                >
                  {showPassword ? <EyeSlashIcon className="h-5 w-5" /> : <EyeIcon className="h-5 w-5" />}
                </button>
              </div>
            </div>
            {error && (
              <div className="rounded-lg bg-red-50 border border-red-200 p-3 text-sm text-red-700">
                {error}
              </div>
            )}
            <button
              type="submit"
              disabled={loading}
              className="w-full flex justify-center items-center py-3 px-4 rounded-lg text-white bg-blue-600 hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 font-medium"
            >
              {loading ? (
                <span className="animate-spin h-5 w-5 border-2 border-white border-t-transparent rounded-full" />
              ) : (
                <>
                  Masuk
                  <ArrowRightIcon className="ml-2 h-5 w-5" />
                </>
              )}
            </button>
          </form>
          <p className="text-center text-xs text-gray-500 mt-6">
            © {new Date().getFullYear()} {app.name}
          </p>
        </div>
      </div>
      <ErrorDialog
        isOpen={successOpen}
        onClose={() => setSuccessOpen(false)}
        title="Berhasil"
        message="Login berhasil. Mengalihkan..."
        type="info"
        disableClose
      />
    </div>
  );
}

export default function LoginPage() {
  return (
    <Suspense
      fallback={
        <div className="min-h-screen flex items-center justify-center bg-gray-100">
          <div className="animate-spin rounded-full h-12 w-12 border-2 border-blue-600 border-t-transparent" />
        </div>
      }
    >
      <LoginForm />
    </Suspense>
  );
}
