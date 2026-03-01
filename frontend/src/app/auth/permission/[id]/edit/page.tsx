"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface PermissionDetail {
  id: string;
  code: string;
  name: string;
  group?: string;
  description?: string;
}

export default function PermissionEditPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [code, setCode] = useState("");
  const [name, setName] = useState("");
  const [group, setGroup] = useState("");
  const [description, setDescription] = useState("");
  const [loading, setLoading] = useState(true);
  const [submitLoading, setSubmitLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService.get<PermissionDetail>("v1/rbac/permissions/" + id).then((res) => {
      if (res.success && res.data) {
        const p = res.data as PermissionDetail;
        setCode(p.code ?? "");
        setName(p.name ?? "");
        setGroup(p.group ?? "");
        setDescription(p.description ?? "");
      }
      setLoading(false);
    });
  }, [id]);

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
    const res = await apiService.put("v1/rbac/permissions/" + id, {
      code: code.trim(),
      name: name.trim(),
      group: group.trim() || undefined,
      description: description.trim() || undefined,
    });
    if (res.success) {
      router.push(`/auth/permission/${id}`);
      return;
    }
    setError(res.message ?? "Gagal memperbarui permission");
    setSubmitLoading(false);
  };

  if (loading) return <div className="p-8">Memuat...</div>;

  return (
    <div className="space-y-6">
      <Link href={`/auth/permission/${id}`} className="text-gray-600 hover:text-gray-900">
        ← Kembali
      </Link>
      <h1 className="text-2xl font-bold text-gray-900">Edit Permission</h1>

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
          <label className="block text-sm font-medium text-gray-700 mb-1">Modul/Group (opsional)</label>
          <input
            type="text"
            value={group}
            onChange={(e) => setGroup(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Deskripsi (opsional)</label>
          <textarea
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            rows={3}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div className="flex gap-2">
          <button
            type="submit"
            disabled={submitLoading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 disabled:opacity-50"
          >
            {submitLoading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link
            href={`/auth/permission/${id}`}
            className="px-4 py-2 border border-gray-300 rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
