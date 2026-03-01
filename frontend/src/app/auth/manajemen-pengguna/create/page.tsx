"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface RoleItem {
  id: string;
  code: string;
  name: string;
}

export default function CreateUserPage() {
  const router = useRouter();
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [nip, setNip] = useState("");
  const [roleIds, setRoleIds] = useState<string[]>([]);
  const [isActive, setIsActive] = useState(true);
  const [roles, setRoles] = useState<RoleItem[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService
      .get<{ items?: RoleItem[] }>("v1/rbac/roles", { page_size: 100 })
      .then((res) => {
        if (res.success && res.data) {
          const d = res.data as { items?: RoleItem[] };
          setRoles(d.items ?? []);
        }
      })
      .catch(() => {});
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!name.trim()) {
      setError("Nama wajib diisi");
      return;
    }
    if (!email.trim()) {
      setError("Email wajib diisi");
      return;
    }
    if (!username.trim()) {
      setError("Username wajib diisi");
      return;
    }
    if (!password || password.length < 8) {
      setError("Kata sandi minimal 8 karakter");
      return;
    }
    if (roleIds.length === 0) {
      setError("Pilih minimal satu role");
      return;
    }

    setLoading(true);
    const res = await apiService.post("v1/users", {
      name: name.trim(),
      email: email.trim(),
      username: username.trim(),
      password,
      nip: nip.trim() || undefined,
      role_ids: roleIds,
      is_active: isActive,
    });
    if (res.success) {
      router.push("/auth/manajemen-pengguna");
      return;
    }
    setError(res.message ?? "Gagal menambah pengguna");
    setLoading(false);
  };

  const toggleRole = (id: string) => {
    setRoleIds((prev) =>
      prev.includes(id) ? prev.filter((r) => r !== id) : [...prev, id]
    );
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <Link
          href="/auth/manajemen-pengguna"
          className="text-gray-600 hover:text-gray-900"
        >
          ← Kembali
        </Link>
        <h1 className="text-2xl font-bold text-gray-900">Tambah Pengguna</h1>
      </div>

      <form onSubmit={handleSubmit} className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 max-w-2xl space-y-5">
        {error && (
          <div className="p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>
        )}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Nama *</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Email *</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Username *</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Kata sandi *</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
            minLength={8}
          />
          <p className="text-xs text-gray-500 mt-1">Minimal 8 karakter</p>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">NIP (opsional)</label>
          <input
            type="text"
            value={nip}
            onChange={(e) => setNip(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Role *</label>
          <div className="flex flex-wrap gap-2">
            {roles.map((r) => (
              <label key={r.id} className="inline-flex items-center gap-2">
                <input
                  type="checkbox"
                  checked={roleIds.includes(r.id)}
                  onChange={() => toggleRole(r.id)}
                  className="rounded border-gray-300"
                />
                <span>{r.name || r.code}</span>
              </label>
            ))}
          </div>
          {roles.length === 0 && (
            <p className="text-xs text-amber-600">Tidak ada role tersedia. Pastikan Anda punya akses RBAC.</p>
          )}
        </div>
        <div>
          <label className="inline-flex items-center gap-2">
            <input
              type="checkbox"
              checked={isActive}
              onChange={(e) => setIsActive(e.target.checked)}
              className="rounded border-gray-300"
            />
            <span className="text-sm font-medium text-gray-700">Aktif</span>
          </label>
        </div>
        <div className="flex gap-2">
          <button
            type="submit"
            disabled={loading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 disabled:opacity-50"
          >
            {loading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link
            href="/auth/manajemen-pengguna"
            className="px-4 py-2 border border-gray-300 rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
