"use client";

import { ErrorDialog } from "@/components";
import { api, app, ui } from "@/config";
import { getAuthCookie, storeAuthData } from "@/lib/auth";
import { ArrowRightIcon, EyeIcon, EyeSlashIcon } from "@heroicons/react/24/outline";
import Link from "next/link";
import { useRouter, useSearchParams } from "next/navigation";
import { Suspense, useEffect, useState } from "react";

// Backend LoginResponse: { access_token, token_type, expires_in, user: { id, name, email, username, roles, default_home_path } }
interface LoginResponse {
  access_token?: string;
  token_type?: string;
  expires_in?: number;
  user?: {
    id: string;
    name: string;
    email: string;
    username: string;
    roles?: { id: string; code: string; name: string }[];
    default_home_path?: string;
  };
}

function LoginForm() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const redirectParam = searchParams.get("redirect");
  const sessionExpired = searchParams.get("session") === "expired";
  const logoutSuccess = searchParams.get("message") === "logout_success";
  const redirect = redirectParam || "/dashboard";

  const [identifier, setIdentifier] = useState("");
  const [password, setPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(
    sessionExpired ? "Sesi Anda telah berakhir. Silakan masuk kembali." : ""
  );
  const [infoMessage, setInfoMessage] = useState(logoutSuccess ? "Anda telah keluar." : "");
  const [successOpen, setSuccessOpen] = useState(false);

  useEffect(() => {
    if (getAuthCookie()) {
      router.replace(redirect);
      return;
    }
  }, [redirect, router]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    if (!identifier.trim()) {
      setError("Email atau username wajib diisi");
      return;
    }
    if (!password) {
      setError("Kata sandi wajib diisi");
      return;
    }

    setLoading(true);
    try {
      const res = await fetch(`${api.baseUrl}/v1/auth/login`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({
          identifier: identifier.trim(),
          password,
        }),
      });
      const data: LoginResponse & { message?: string } = await res.json().catch(() => ({}));

      if (res.ok && data.access_token && data.user) {
        const u = data.user;
        const expiresIn = data.expires_in ?? 86400;
        const expiry = new Date(Date.now() + expiresIn * 1000).toISOString();
        const authData = {
          access_token: data.access_token,
          user_id: u.id,
          user_name: u.username,
          user_nik: "",
          user_fullname: u.name,
          default_home_path: u.default_home_path || "/dashboard",
          expiry,
          role_user: (u.roles ?? []).map((r) => ({
            role_aplikasi_id: r.id,
            role_aplikasi: r.code,
            user_id: u.id,
            instansi_id: "",
            nama_instansi: "",
          })),
        };
        await storeAuthData(authData);
        setSuccessOpen(true);
        const homePath = u.default_home_path || "/dashboard";
        setTimeout(() => {
          window.location.href = redirectParam || homePath;
        }, 800);
      } else {
        setError(data.message ?? "Email/username atau kata sandi salah.");
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
          {infoMessage && (
            <div
              role="status"
              className="rounded-lg bg-green-50 border border-green-200 p-3 text-sm text-green-800 mb-4"
            >
              {infoMessage}
            </div>
          )}
          <form onSubmit={handleSubmit} className="space-y-5">
            <div>
              <label htmlFor="identifier" className="block text-sm font-medium text-gray-700 mb-1">
                Email atau Username
              </label>
              <input
                id="identifier"
                type="text"
                autoComplete="username"
                value={identifier}
                onChange={(e) => setIdentifier(e.target.value)}
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900"
                placeholder="Email atau username"
                disabled={loading}
                aria-describedby={error ? "identifier-error" : undefined}
              />
            </div>
            <div>
              <label htmlFor="password" className="block text-sm font-medium text-gray-700 mb-1">
                Kata sandi
              </label>
              <div className="relative">
                <input
                  id="password"
                  type={showPassword ? "text" : "password"}
                  autoComplete="current-password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="w-full px-4 py-3 pr-12 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900"
                  placeholder="Kata sandi"
                  disabled={loading}
                  aria-describedby={error ? "password-error" : undefined}
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
              <div
                id="form-error"
                role="alert"
                className="rounded-lg bg-red-50 border border-red-200 p-3 text-sm text-red-700"
              >
                {error}
              </div>
            )}
            <div className="flex justify-between items-center">
              <Link
                href="/auth/register"
                className="text-sm text-blue-600 hover:text-blue-800 hover:underline"
              >
                Belum punya akun? Daftar
              </Link>
              <Link
                href="/auth/forgot-password"
                className="text-sm text-blue-600 hover:text-blue-800 hover:underline"
              >
                Lupa kata sandi?
              </Link>
            </div>
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
