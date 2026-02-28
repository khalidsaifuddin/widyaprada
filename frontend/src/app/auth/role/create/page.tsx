"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface PermissionItem {
  id: string;
  code: string;
  name: string;
  group?: string;
}

export default function RoleCreatePage() {
  const router = useRouter();
  const [code, setCode] = useState("");
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [permissionIds, setPermissionIds] = useState<string[]>([]);
  const [permissions, setPermissions] = useState<PermissionItem[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService
      .get<{ items?: PermissionItem[] }>("v1/rbac/permissions", { page_size: 200 })
      .then((res) => {
        if (res.success && res.data) {
          const d = res.data as { items?: PermissionItem[] };
          setPermissions(d.items ?? []);
        }
      })
      .catch(() => {});
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!code.trim()) {
      setError("Kode wajib diisi");
      return;
    }
    if (!name.trim()) {
      setError("Nama wajib diisi");
      return;
    }

    setLoading(true);
    const res = await apiService.post("v1/rbac/roles", {
      code: code.trim(),
      name: name.trim(),
      description: description.trim() || undefined,
      permission_ids: permissionIds,
    });
    if (res.success) {
      router.push("/auth/role");
      return;
    }
    setError(res.message ?? "Gagal menambah role");
    setLoading(false);
  };

  const togglePermission = (id: string) => {
    setPermissionIds((prev) =>
      prev.includes(id) ? prev.filter((p) => p !== id) : [...prev, id]
    );
  };

  const toggleAllInGroup = (group: string) => {
    const inGroup = permissions.filter((p) => (p.group || "") === group);
    const allSelected = inGroup.every((p) => permissionIds.includes(p.id));
    if (allSelected) {
      setPermissionIds((prev) => prev.filter((id) => !inGroup.some((p) => p.id === id)));
    } else {
      setPermissionIds((prev) => [...new Set([...prev, ...inGroup.map((p) => p.id)])]);
    }
  };

  const grouped = permissions.reduce<Record<string, PermissionItem[]>>((acc, p) => {
    const g = p.group || "Lainnya";
    if (!acc[g]) acc[g] = [];
    acc[g].push(p);
    return acc;
  }, {});

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <Link href="/auth/role" className="text-gray-600 hover:text-gray-900">
          ← Kembali
        </Link>
        <h1 className="text-2xl font-bold text-gray-900">Tambah Role</h1>
      </div>

      <form
        onSubmit={handleSubmit}
        className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 max-w-2xl space-y-5"
      >
        {error && (
          <div className="p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>
        )}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Kode *</label>
          <input
            type="text"
            value={code}
            onChange={(e) => setCode(e.target.value)}
            placeholder="contoh: ADMIN_CUSTOM"
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Nama *</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Nama role"
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Deskripsi (opsional)</label>
          <textarea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            rows={2}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">Permission</label>
          <div className="border border-gray-200 rounded-lg p-4 max-h-64 overflow-y-auto space-y-4">
            {Object.entries(grouped).map(([group, items]) => (
              <div key={group}>
                <label className="flex items-center gap-2 font-medium text-gray-700 mb-2 cursor-pointer">
                  <input
                    type="checkbox"
                    checked={items.every((p) => permissionIds.includes(p.id))}
                    onChange={() => toggleAllInGroup(group)}
                    className="rounded border-gray-300"
                  />
                  {group}
                </label>
                <div className="flex flex-wrap gap-2 pl-6">
                  {items.map((p) => (
                    <label key={p.id} className="inline-flex items-center gap-2 text-sm">
                      <input
                        type="checkbox"
                        checked={permissionIds.includes(p.id)}
                        onChange={() => togglePermission(p.id)}
                        className="rounded border-gray-300"
                      />
                      {p.name || p.code}
                    </label>
                  ))}
                </div>
              </div>
            ))}
          </div>
        </div>
        <div className="flex gap-2">
          <button
            type="submit"
            disabled={loading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
          >
            {loading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link
            href="/auth/role"
            className="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
