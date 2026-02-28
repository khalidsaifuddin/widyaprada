"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
import { getUserProfile } from "@/lib/auth";

interface PackageQuestionInfo {
  question_id: string;
  question_code: string;
  sort_order: number;
}

interface PackageDetail {
  id: string;
  code: string;
  name: string;
  description: string;
  status: string;
  verification_status: string;
  questions: PackageQuestionInfo[];
  created_at?: string;
  updated_at?: string;
}

function canEdit(roleUser: { role_aplikasi?: string }[] | undefined): boolean {
  return roleUser?.some((r) =>
    r.role_aplikasi === "SUPER_ADMIN" || r.role_aplikasi === "ADMIN_UJIKOM"
  ) ?? false;
}

function canVerify(roleUser: { role_aplikasi?: string }[] | undefined): boolean {
  return roleUser?.some((r) =>
    r.role_aplikasi === "SUPER_ADMIN" || r.role_aplikasi === "VERIFIKATOR"
  ) ?? false;
}

export default function PaketSoalDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [pkg, setPkg] = useState<PackageDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [deleteDialog, setDeleteDialog] = useState(false);
  const [deleteReason, setDeleteReason] = useState("");
  const [deleteLoading, setDeleteLoading] = useState(false);
  const [canEditRole, setCanEditRole] = useState(false);
  const [canVerifyRole, setCanVerifyRole] = useState(false);

  useEffect(() => {
    getUserProfile().then((p) => {
      setCanEditRole(canEdit(p?.role_user));
      setCanVerifyRole(canVerify(p?.role_user));
    });
  }, []);

  const fetchPackage = useCallback(async () => {
    if (!id) return;
    setLoading(true);
    const res = await apiService.get<PackageDetail>("v1/question-packages/" + id);
    if (res.success && res.data) {
      setPkg(res.data as PackageDetail);
    } else {
      setError(res.message ?? "Paket tidak ditemukan");
    }
    setLoading(false);
  }, [id]);

  useEffect(() => {
    fetchPackage();
  }, [fetchPackage]);

  const handleDelete = async () => {
    if (!deleteReason.trim()) return;
    setDeleteLoading(true);
    const res = await apiService.delete("v1/question-packages/" + id, {
      reason: deleteReason.trim(),
    });
    if (res.success) {
      router.push("/wpujikom/paket-soal");
      return;
    }
    setError(res.message ?? "Gagal menghapus paket");
    setDeleteLoading(false);
  };

  const handleVerify = async (verify: boolean) => {
    const endpoint = verify
      ? `v1/question-packages/${id}/verify`
      : `v1/question-packages/${id}/unverify`;
    const res = await apiService.post(endpoint, {});
    if (res.success) fetchPackage();
    else setError(res.message ?? "Gagal verifikasi");
  };

  if (loading) return <div className="p-8">Memuat...</div>;
  if (error && !pkg) {
    return (
      <div className="space-y-4">
        <Link href="/wpujikom/paket-soal" className="text-blue-600 hover:underline">
          ← Kembali
        </Link>
        <p className="text-red-600">{error}</p>
      </div>
    );
  }

  const sortedQuestions = [...(pkg?.questions ?? [])].sort(
    (a, b) => a.sort_order - b.sort_order
  );

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <Link href="/wpujikom/paket-soal" className="text-gray-600 hover:text-gray-900">
          ← Kembali
        </Link>
        <div className="flex gap-2">
          {canEditRole && (
            <>
              <Link
                href={`/wpujikom/paket-soal/${id}/edit`}
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
            </>
          )}
          {canVerifyRole && (
            <button
              onClick={() => handleVerify(pkg!.verification_status !== "Sudah")}
              className="px-4 py-2 bg-amber-600 text-white rounded-lg hover:bg-amber-700"
            >
              {pkg!.verification_status === "Sudah" ? "Batal Verifikasi" : "Verifikasi"}
            </button>
          )}
        </div>
      </div>

      {error && (
        <div className="p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>
      )}

      {pkg && (
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 max-w-3xl space-y-6">
          <h1 className="text-xl font-bold text-gray-900">Detail Paket Soal</h1>
          <dl className="space-y-4">
            <div>
              <dt className="text-sm font-medium text-gray-500">Kode</dt>
              <dd className="mt-1 font-mono text-gray-900">{pkg.code}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Nama</dt>
              <dd className="mt-1 text-gray-900">{pkg.name}</dd>
            </div>
            {pkg.description && (
              <div>
                <dt className="text-sm font-medium text-gray-500">Deskripsi</dt>
                <dd className="mt-1 text-gray-600">{pkg.description}</dd>
              </div>
            )}
            <div className="flex gap-4">
              <div>
                <dt className="text-sm font-medium text-gray-500">Status</dt>
                <dd>
                  <span
                    className={`inline-flex px-2 py-0.5 rounded text-xs font-medium ${
                      pkg.status === "Aktif" ? "bg-green-100 text-green-800" : "bg-gray-100 text-gray-700"
                    }`}
                  >
                    {pkg.status}
                  </span>
                </dd>
              </div>
              <div>
                <dt className="text-sm font-medium text-gray-500">Verifikasi</dt>
                <dd>
                  <span
                    className={`inline-flex px-2 py-0.5 rounded text-xs font-medium ${
                      pkg.verification_status === "Sudah"
                        ? "bg-blue-100 text-blue-800"
                        : "bg-amber-100 text-amber-800"
                    }`}
                  >
                    {pkg.verification_status}
                  </span>
                </dd>
              </div>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500 mb-2">Daftar Soal ({sortedQuestions.length})</dt>
              <dd className="space-y-2">
                {sortedQuestions.length === 0 ? (
                  <p className="text-gray-500">Belum ada soal dalam paket ini.</p>
                ) : (
                  sortedQuestions.map((q, i) => (
                    <div
                      key={q.question_id}
                      className="flex items-center gap-2 p-2 rounded border border-gray-100"
                    >
                      <span className="text-gray-500 w-6">{i + 1}.</span>
                      <Link
                        href={`/wpujikom/bank-soal/${q.question_id}`}
                        className="font-mono text-blue-600 hover:text-blue-800"
                      >
                        {q.question_code}
                      </Link>
                    </div>
                  ))
                )}
              </dd>
            </div>
          </dl>
        </div>
      )}

      {deleteDialog && (
        <div className="fixed inset-0 z-50 overflow-y-auto">
          <div className="flex min-h-full items-end justify-center p-4 sm:items-center sm:p-0">
            <div className="fixed inset-0 bg-black/50" onClick={() => setDeleteDialog(false)} />
            <div className="relative bg-white rounded-lg shadow-xl p-6 max-w-md w-full mx-4">
              <h3 className="text-lg font-semibold text-gray-900">Konfirmasi Hapus Paket Soal</h3>
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
