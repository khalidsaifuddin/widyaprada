"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface PermissionItem {
  id: string;
  code: string;
  name: string;
  group?: string;
}

interface RoleDetail {
  id: string;
  code: string;
  name: string;
  description?: string;
  permissions: { id: string }[];
}

export default function RoleEditPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [code, setCode] = useState("");
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [permissionIds, setPermissionIds] = useState<string[]>([]);
  const [permissions, setPermissions] = useState<PermissionItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [submitLoading, setSubmitLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService.get<RoleDetail>("v1/rbac/roles/" + id).then((res) => {
      if (res.success && res.data) {
        const r = res.data as RoleDetail;
        setCode(r.code ?? "");
        setName(r.name ?? "");
        setDescription(r.description ?? "");
        setPermissionIds(r.permissions?.map((p) => p.id) ?? []);
      }
      setLoading(false);
    });
  }, [id]);

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

    setSubmitLoading(true);
    const res = await apiService.put("v1/rbac/roles/" + id, {
      code: code.trim(),
      name: name.trim(),
      description: description.trim() || undefined,
      permission_ids: permissionIds,
    });
    if (res.success) {
      router.push(`/auth/role/${id}`);
      return;
    }
    setError(res.message ?? "Gagal memperbarui role");
    setSubmitLoading(false);
  };

  const togglePermission = (pid: string) => {
    setPermissionIds((prev) =>
      prev.includes(pid) ? prev.filter((p) => p !== pid) : [...prev, pid]
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

  if (loading) return <div className="p-8">Memuat...</div>;

  return (
    <div className="space-y-6">
      <Link href={`/auth/role/${id}`} className="text-gray-600 hover:text-gray-900">
        ← Kembali
      </Link>
      <h1 className="text-2xl font-bold text-gray-900">Edit Role</h1>

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
            disabled={submitLoading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
          >
            {submitLoading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link
            href={`/auth/role/${id}`}
            className="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
