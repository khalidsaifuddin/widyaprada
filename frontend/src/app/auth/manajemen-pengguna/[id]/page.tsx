"use client";

import { apiService } from "@/lib/api";
import { getAuthData } from "@/lib/auth";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

interface RoleInfo {
  id: string;
  code: string;
  name: string;
}

interface UserDetail {
  id: string;
  name: string;
  email: string;
  username: string;
  nip?: string;
  roles: RoleInfo[];
  satker_id?: string;
  is_active: boolean;
  created_at?: string;
  updated_at?: string;
}

export default function UserDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [user, setUser] = useState<UserDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [deleteDialog, setDeleteDialog] = useState(false);
  const [deleteReason, setDeleteReason] = useState("");
  const [deleteLoading, setDeleteLoading] = useState(false);
  const [currentUserId, setCurrentUserId] = useState<string | null>(null);

  const fetchUser = useCallback(async () => {
    if (!id) return;
    setLoading(true);
    const res = await apiService.get<UserDetail>("v1/users/" + id);
    if (res.success && res.data) {
      setUser(res.data as UserDetail);
    } else {
      setError(res.message ?? "Pengguna tidak ditemukan");
    }
    setLoading(false);
  }, [id]);

  useEffect(() => {
    getAuthData().then((a) => setCurrentUserId(a?.user_id ?? null));
  }, []);

  useEffect(() => {
    fetchUser();
  }, [fetchUser]);

  const handleDelete = async () => {
    if (!deleteReason.trim()) return;
    setDeleteLoading(true);
    const res = await apiService.delete("v1/users/" + id, { reason: deleteReason.trim() });
    if (res.success) {
      router.push("/auth/manajemen-pengguna");
      return;
    }
    setError(res.message ?? "Gagal menghapus pengguna");
    setDeleteLoading(false);
  };

  const isSelf = currentUserId && user && currentUserId === user.id;

  if (loading) return <div className="p-8">Memuat...</div>;
  if (error && !user) {
    return (
      <div className="space-y-4">
        <Link href="/auth/manajemen-pengguna" className="text-blue-600 hover:underline">
          ← Kembali
        </Link>
        <p className="text-red-600">{error}</p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <Link href="/auth/manajemen-pengguna" className="text-gray-600 hover:text-gray-900">
          ← Kembali
        </Link>
        <div className="flex gap-2">
          <Link
            href={"/auth/manajemen-pengguna/" + id + "/edit"}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
          >
            Edit
          </Link>
          {!isSelf && (
            <button
              onClick={() => setDeleteDialog(true)}
              className="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700"
            >
              Hapus
            </button>
          )}
        </div>
      </div>

      {error && (
        <div className="p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>
      )}

      {user && (
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 max-w-2xl">
          <h1 className="text-xl font-bold text-gray-900 mb-6">Detail Pengguna</h1>
          <dl className="space-y-4">
            <div>
              <dt className="text-sm font-medium text-gray-500">Nama</dt>
              <dd className="mt-1 text-gray-900">{user.name}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Email</dt>
              <dd className="mt-1 text-gray-900">{user.email}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Username</dt>
              <dd className="mt-1 text-gray-900">{user.username}</dd>
            </div>
            {user.nip && (
              <div>
                <dt className="text-sm font-medium text-gray-500">NIP</dt>
                <dd className="mt-1 text-gray-900">{user.nip}</dd>
              </div>
            )}
            <div>
              <dt className="text-sm font-medium text-gray-500">Role</dt>
              <dd className="mt-1 flex flex-wrap gap-1">
                {user.roles?.map((r) => (
                  <span
                    key={r.id}
                    className="inline-flex px-2 py-0.5 rounded text-xs font-medium bg-blue-100 text-blue-800"
                  >
                    {r.name || r.code}
                  </span>
                )) ?? "-"}
              </dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Status</dt>
              <dd>
                <span
                  className={
                    user.is_active
                      ? "inline-flex px-2 py-0.5 rounded text-xs font-medium bg-green-100 text-green-800"
                      : "inline-flex px-2 py-0.5 rounded text-xs font-medium bg-gray-100 text-gray-600"
                  }
                >
                  {user.is_active ? "Aktif" : "Nonaktif"}
                </span>
              </dd>
            </div>
            {user.created_at && (
              <div>
                <dt className="text-sm font-medium text-gray-500">Dibuat</dt>
                <dd className="mt-1 text-gray-900">{new Date(user.created_at).toLocaleString("id-ID")}</dd>
              </div>
            )}
          </dl>
        </div>
      )}

      {deleteDialog && (
        <div className="fixed inset-0 z-50 overflow-y-auto">
          <div className="flex min-h-full items-center justify-center p-4">
            <div className="fixed inset-0 bg-black/50" onClick={() => setDeleteDialog(false)} />
            <div className="relative bg-white rounded-lg shadow-xl p-6 max-w-md w-full">
              <h3 className="text-lg font-semibold text-gray-900">Konfirmasi Hapus</h3>
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
