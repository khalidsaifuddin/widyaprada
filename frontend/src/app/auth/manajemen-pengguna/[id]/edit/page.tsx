"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface RoleItem {
  id: string;
  code: string;
  name: string;
}

export default function EditUserPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [nip, setNip] = useState("");
  const [roleIds, setRoleIds] = useState<string[]>([]);
  const [isActive, setIsActive] = useState(true);
  const [roles, setRoles] = useState<RoleItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [submitLoading, setSubmitLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService.get(`v1/users/${id}`).then((res) => {
      if (res.success && res.data) {
        const u = res.data as Record<string, unknown>;
        setName(String(u.name ?? ""));
        setEmail(String(u.email ?? ""));
        setUsername(String(u.username ?? ""));
        setNip(String(u.nip ?? ""));
        setIsActive(Boolean(u.is_active ?? true));
        const r = (u.roles ?? []) as { id: string }[];
        setRoleIds(r.map((x) => x.id));
      }
      setLoading(false);
    });
  }, [id]);

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
    if (roleIds.length === 0) {
      setError("Pilih minimal satu role");
      return;
    }

    setSubmitLoading(true);
    const body: Record<string, unknown> = {
      name: name.trim(),
      email: email.trim(),
      username: username.trim(),
      nip: nip.trim() || undefined,
      role_ids: roleIds,
      is_active: isActive,
    };
    if (password && password.length >= 8) body.password = password;

    const res = await apiService.put(`v1/users/${id}`, body);
    if (res.success) {
      router.push(`/auth/manajemen-pengguna/${id}`);
      return;
    }
    setError(res.message ?? "Gagal memperbarui pengguna");
    setSubmitLoading(false);
  };

  const toggleRole = (rid: string) => {
    setRoleIds((prev) =>
      prev.includes(rid) ? prev.filter((r) => r !== rid) : [...prev, rid]
    );
  };

  if (loading) return <div className="p-8">Memuat...</div>;

  return (
    <div className="space-y-6">
      <Link href={`/auth/manajemen-pengguna/${id}`} className="text-gray-600 hover:text-gray-900">
        ← Kembali
      </Link>
      <h1 className="text-2xl font-bold text-gray-900">Edit Pengguna</h1>

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
          <label className="block text-sm font-medium text-gray-700 mb-1">Kata sandi baru (kosongkan jika tidak diubah)</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            minLength={8}
          />
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
            disabled={submitLoading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
          >
            {submitLoading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link
            href={`/auth/manajemen-pengguna/${id}`}
            className="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
