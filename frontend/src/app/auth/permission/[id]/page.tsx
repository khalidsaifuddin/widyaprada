"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

interface PermissionDetail {
  id: string;
  code: string;
  name: string;
  group?: string;
  description?: string;
  created_at?: string;
  updated_at?: string;
}

export default function PermissionDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [permission, setPermission] = useState<PermissionDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [deleteDialog, setDeleteDialog] = useState(false);
  const [deleteReason, setDeleteReason] = useState("");
  const [deleteLoading, setDeleteLoading] = useState(false);

  const fetchPermission = useCallback(async () => {
    if (!id) return;
    setLoading(true);
    const res = await apiService.get<PermissionDetail>("v1/rbac/permissions/" + id);
    if (res.success && res.data) {
      setPermission(res.data as PermissionDetail);
    } else {
      setError(res.message ?? "Permission tidak ditemukan");
    }
    setLoading(false);
  }, [id]);

  useEffect(() => {
    fetchPermission();
  }, [fetchPermission]);

  const handleDelete = async () => {
    if (!deleteReason.trim()) return;
    setDeleteLoading(true);
    const res = await apiService.delete("v1/rbac/permissions/" + id, {
      reason: deleteReason.trim(),
    });
    if (res.success) {
      router.push("/auth/permission");
      return;
    }
    setError(res.message ?? "Gagal menghapus permission");
    setDeleteLoading(false);
  };

  if (loading) return <div className="p-8">Memuat...</div>;
  if (error && !permission) {
    return (
      <div className="space-y-4">
        <Link href="/auth/permission" className="text-blue-600 hover:underline">
          ← Kembali
        </Link>
        <p className="text-red-600">{error}</p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <Link href="/auth/permission" className="text-gray-600 hover:text-gray-900">
          ← Kembali
        </Link>
        <div className="flex gap-2">
          <Link
            href={`/auth/permission/${id}/edit`}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
          >
            Edit
          </Link>
          <button
            onClick={() => setDeleteDialog(true)}
            className="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700"
          >
            Hapus
          </button>
        </div>
      </div>

      {error && (
        <div className="p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>
      )}

      {permission && (
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 max-w-2xl">
          <h1 className="text-xl font-bold text-gray-900 mb-6">Detail Permission</h1>
          <dl className="space-y-4">
            <div>
              <dt className="text-sm font-medium text-gray-500">Kode</dt>
              <dd className="mt-1 text-gray-900 font-mono">{permission.code}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Nama</dt>
              <dd className="mt-1 text-gray-900">{permission.name}</dd>
            </div>
            {permission.group && (
              <div>
                <dt className="text-sm font-medium text-gray-500">Modul/Group</dt>
                <dd className="mt-1">
                  <span className="inline-flex px-2 py-0.5 rounded text-sm font-medium bg-gray-100 text-gray-800">
                    {permission.group}
                  </span>
                </dd>
              </div>
            )}
            {permission.description && (
              <div>
                <dt className="text-sm font-medium text-gray-500">Deskripsi</dt>
                <dd className="mt-1 text-gray-600">{permission.description}</dd>
              </div>
            )}
            {permission.created_at && (
              <div>
                <dt className="text-sm font-medium text-gray-500">Dibuat</dt>
                <dd className="mt-1 text-gray-900">
                  {new Date(permission.created_at).toLocaleString("id-ID")}
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
              <h3 className="text-lg font-semibold text-gray-900">Konfirmasi Hapus Permission</h3>
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
                  className="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"
                >
                  Batal
                </button>
                <button
                  onClick={handleDelete}
                  disabled={!deleteReason.trim() || deleteLoading}
                  className="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 disabled:opacity-50"
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
