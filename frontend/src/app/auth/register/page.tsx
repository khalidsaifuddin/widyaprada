"use client";

import { api, app, ui } from "@/config";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

function isValidEmail(email: string): boolean {
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email.trim());
}

export default function RegisterPage() {
  const router = useRouter();
  const [nama, setNama] = useState("");
  const [email, setEmail] = useState("");
  const [nip, setNip] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");

    if (!nama.trim()) {
      setError("Nama wajib diisi");
      return;
    }
    if (!email.trim()) {
      setError("Email wajib diisi");
      return;
    }
    if (!isValidEmail(email)) {
      setError("Format email tidak valid");
      return;
    }

    setLoading(true);
    try {
      const res = await fetch(`${api.baseUrl}/v1/auth/register`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify({
          name: nama.trim(),
          email: email.trim(),
          nip: nip.trim() || undefined,
        }),
      });
      const data = await res.json().catch(() => ({}));

      if (res.ok && (res.status === 201 || data.message)) {
        setSuccess(true);
      } else {
        const msg =
          data.status ??
          data.message ??
          (typeof data.data === "string" ? data.data : null) ??
          "Registrasi gagal. Silakan coba lagi.";
        setError(msg);
      }
    } catch {
      setError("Koneksi gagal. Periksa jaringan Anda.");
    } finally {
      setLoading(false);
    }
  };

  if (success) {
    return (
      <div
        className="min-h-screen flex items-center justify-center p-4"
        style={{
          background: `linear-gradient(to top, ${ui.theme.gradient.from}, ${ui.theme.gradient.to})`,
        }}
      >
        <div className="w-full max-w-md">
          <div className="bg-white rounded-2xl shadow-xl p-8 text-center">
            <div className="mb-6">
              <h2 className="text-2xl font-bold text-gray-900">Registrasi Berhasil</h2>
              <p className="text-gray-600 mt-2">
                Registrasi berhasil. Silakan cek email Anda untuk mendapatkan kata sandi awal.
              </p>
            </div>
            <Link
              href="/auth/login"
              className="inline-flex items-center justify-center w-full py-3 px-4 rounded-lg text-white bg-blue-600 transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 font-medium"
            >
              Kembali ke Login
            </Link>
          </div>
        </div>
      </div>
    );
  }

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
            <h2 className="text-2xl font-bold text-gray-900">Daftar</h2>
            <p className="text-gray-600 mt-1">Registrasi calon peserta</p>
          </div>
          <form onSubmit={handleSubmit} className="space-y-5">
            <div>
              <label htmlFor="nama" className="block text-sm font-medium text-gray-700 mb-1">
                Nama
              </label>
              <input
                id="nama"
                type="text"
                autoComplete="name"
                value={nama}
                onChange={(e) => setNama(e.target.value)}
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900"
                placeholder="Nama lengkap"
                disabled={loading}
              />
            </div>
            <div>
              <label htmlFor="email" className="block text-sm font-medium text-gray-700 mb-1">
                Email
              </label>
              <input
                id="email"
                type="email"
                autoComplete="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900"
                placeholder="Email"
                disabled={loading}
              />
            </div>
            <div>
              <label htmlFor="nip" className="block text-sm font-medium text-gray-700 mb-1">
                NIP <span className="text-gray-400">(opsional)</span>
              </label>
              <input
                id="nip"
                type="text"
                value={nip}
                onChange={(e) => setNip(e.target.value)}
                className="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900"
                placeholder="NIP"
                disabled={loading}
              />
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
            <button
              type="submit"
              disabled={loading}
              className="w-full flex justify-center items-center py-3 px-4 rounded-lg text-white bg-blue-600 transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:opacity-50 font-medium"
            >
              {loading ? (
                <span className="animate-spin h-5 w-5 border-2 border-white border-t-transparent rounded-full" />
              ) : (
                "Daftar"
              )}
            </button>
          </form>
          <p className="text-center text-sm text-gray-600 mt-6">
            Sudah punya akun?{" "}
            <Link href="/auth/login" className="text-blue-600 hover:text-blue-800 hover:underline font-medium">
              Masuk
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}
