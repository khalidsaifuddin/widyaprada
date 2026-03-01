"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
interface PermissionInfo {
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
  permissions: PermissionInfo[];
  created_at?: string;
  updated_at?: string;
}

export default function RoleDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [role, setRole] = useState<RoleDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [deleteDialog, setDeleteDialog] = useState(false);
  const [deleteReason, setDeleteReason] = useState("");
  const [deleteLoading, setDeleteLoading] = useState(false);

  const fetchRole = useCallback(async () => {
    if (!id) return;
    setLoading(true);
    const res = await apiService.get<RoleDetail>("v1/rbac/roles/" + id);
    if (res.success && res.data) {
      setRole(res.data as RoleDetail);
    } else {
      setError(res.message ?? "Role tidak ditemukan");
    }
    setLoading(false);
  }, [id]);

  useEffect(() => {
    fetchRole();
  }, [fetchRole]);

  const handleDelete = async () => {
    if (!deleteReason.trim()) return;
    setDeleteLoading(true);
    const res = await apiService.delete("v1/rbac/roles/" + id, { reason: deleteReason.trim() });
    if (res.success) {
      router.push("/auth/role");
      return;
    }
    setError(res.message ?? "Gagal menghapus role");
    setDeleteLoading(false);
  };

  if (loading) return <div className="p-8">Memuat...</div>;
  if (error && !role) {
    return (
      <div className="space-y-4">
        <Link href="/auth/role" className="text-blue-600 hover:underline">
          ← Kembali
        </Link>
        <p className="text-red-600">{error}</p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <Link href="/auth/role" className="text-gray-600 hover:text-gray-900">
          ← Kembali
        </Link>
        <div className="flex gap-2">
          <Link
            href={`/auth/role/${id}/edit`}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700"
          >
            Edit
          </Link>
          <button
            onClick={() => setDeleteDialog(true)}
            className="px-4 py-2 bg-red-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-red-700"
          >
            Hapus
          </button>
        </div>
      </div>

      {error && (
        <div className="p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>
      )}

      {role && (
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 max-w-2xl space-y-6">
          <h1 className="text-xl font-bold text-gray-900">Detail Role</h1>
          <dl className="space-y-4">
            <div>
              <dt className="text-sm font-medium text-gray-500">Kode</dt>
              <dd className="mt-1 text-gray-900 font-mono">{role.code}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Nama</dt>
              <dd className="mt-1 text-gray-900">{role.name}</dd>
            </div>
            {role.description && (
              <div>
                <dt className="text-sm font-medium text-gray-500">Deskripsi</dt>
                <dd className="mt-1 text-gray-600">{role.description}</dd>
              </div>
            )}
            <div>
              <dt className="text-sm font-medium text-gray-500">Permission ({role.permissions?.length ?? 0})</dt>
              <dd className="mt-1 flex flex-wrap gap-1">
                {role.permissions?.length ? (
                  role.permissions.map((p) => (
                    <span
                      key={p.id}
                      className="inline-flex px-2 py-0.5 rounded text-xs font-medium bg-blue-100 text-blue-800"
                    >
                      {p.name || p.code}
                    </span>
                  ))
                ) : (
                  <span className="text-gray-500">Tidak ada permission</span>
                )}
              </dd>
            </div>
            {role.created_at && (
              <div>
                <dt className="text-sm font-medium text-gray-500">Dibuat</dt>
                <dd className="mt-1 text-gray-900">
                  {new Date(role.created_at).toLocaleString("id-ID")}
                </dd>
              </div>
            )}
          </dl>
        </div>
      )}

      {deleteDialog && (
        <div className="fixed inset-0 z-50 overflow-y-auto">
          <div className="flex min-h-full items-end justify-center p-4 sm:items-center sm:p-0">
            <div className="fixed inset-0 bg-black/50" onClick={() => setDeleteDialog(false)} />
            <div className="relative bg-white rounded-lg shadow-xl p-6 max-w-md w-full mx-4">
              <h3 className="text-lg font-semibold text-gray-900">Konfirmasi Hapus Role</h3>
              <p className="mt-2 text-gray-600">Alasan penghapusan wajib diisi.</p>
              <textarea
                value={deleteReason}
                onChange={(e) => setDeleteReason(e.target.value)}
                placeholder="Alasan penghapusan..."
                rows={3}
                className="mt-3 w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
              />
              <div className="mt-4 flex justify-end gap-2">
                <button
                  onClick={() => setDeleteDialog(false)}
                  className="px-4 py-2 border border-gray-300 rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
                >
                  Batal
                </button>
                <button
                  onClick={handleDelete}
                  disabled={!deleteReason.trim() || deleteLoading}
                  className="px-4 py-2 bg-red-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-red-700 disabled:opacity-50"
                >
                  {deleteLoading ? "Menghapus..." : "Hapus"}
                </button>
              </div>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
