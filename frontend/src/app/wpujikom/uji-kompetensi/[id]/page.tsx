"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

interface ExamDetail {
  id: string;
  code: string;
  name: string;
  jadwal_mulai: string;
  jadwal_selesai: string;
  durasi_menit: number;
  status: string;
  verification_status: string;
  shuffle_questions: boolean;
  tampilkan_leaderboard: boolean;
  contents: { source_type: string; source_id: string; sort_order: number }[];
  participants: { user_id: string; user_name?: string }[];
  created_at?: string;
  updated_at?: string;
}

function formatDate(d: string): string {
  try {
    return new Date(d).toLocaleString("id-ID", {
      day: "numeric",
      month: "short",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      timeZone: "Asia/Jakarta",
    });
  } catch {
    return d;
  }
}

export default function UjianDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [exam, setExam] = useState<ExamDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [deleteDialog, setDeleteDialog] = useState(false);
  const [deleteReason, setDeleteReason] = useState("");
  const [deleteLoading, setDeleteLoading] = useState(false);

  const fetchExam = useCallback(async () => {
    const res = await apiService.get<ExamDetail>("v1/exams/" + id);
    if (res.success && res.data) {
      setExam(res.data as ExamDetail);
    } else {
      setError(res.message ?? "Ujian tidak ditemukan");
    }
    setLoading(false);
  }, [id]);

  useEffect(() => {
    fetchExam();
  }, [fetchExam]);

  const handlePublish = async () => {
    const res = await apiService.post(`v1/exams/${id}/publish`, {});
    if (res.success) fetchExam();
    else setError(res.message ?? "Gagal menerbitkan");
  };

  const handleVerify = async (verify: boolean) => {
    const endpoint = verify ? `v1/exams/${id}/verify` : `v1/exams/${id}/unverify`;
    const res = await apiService.post(endpoint, {});
    if (res.success) fetchExam();
    else setError(res.message ?? "Gagal verifikasi");
  };

  const handleDelete = async () => {
    if (!deleteReason.trim()) return;
    setDeleteLoading(true);
    const res = await apiService.delete(`v1/exams/${id}`, { reason: deleteReason.trim() });
    if (res.success) {
      router.push("/wpujikom/uji-kompetensi");
      return;
    }
    setError(res.message ?? "Gagal menghapus");
    setDeleteLoading(false);
  };

  if (loading) return <div className="p-8">Memuat...</div>;
  if (error && !exam) {
    return (
      <div className="space-y-4">
        <Link href="/wpujikom/uji-kompetensi" className="text-blue-600 hover:underline">
          ← Kembali
        </Link>
        <p className="text-red-600">{error}</p>
      </div>
    );
  }

  if (!exam) return null;

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <Link href="/wpujikom/uji-kompetensi" className="text-gray-600 hover:text-gray-900">
          ← Kembali
        </Link>
        <div className="flex gap-2">
          {exam.status === "Draft" && (
            <>
              <Link
                href={`/wpujikom/uji-kompetensi/${id}/edit`}
                className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
              >
                Edit
              </Link>
              <button
                onClick={handlePublish}
                className="px-4 py-2 bg-green-600 text-white rounded-lg hover:bg-green-700"
              >
                Terbitkan
              </button>
              <button
                onClick={() => setDeleteDialog(true)}
                className="px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700"
              >
                Hapus
              </button>
            </>
          )}
          <button
            onClick={() => handleVerify(exam.verification_status !== "Sudah")}
            className="px-4 py-2 bg-amber-600 text-white rounded-lg hover:bg-amber-700"
          >
            {exam.verification_status === "Sudah" ? "Batal Verifikasi" : "Verifikasi"}
          </button>
        </div>
      </div>

      {error && (
        <div className="p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>
      )}

      <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 max-w-3xl">
        <h1 className="text-xl font-bold text-gray-900 mb-6">Detail Ujian</h1>
        <dl className="space-y-4">
          <div>
            <dt className="text-sm font-medium text-gray-500">Kode</dt>
            <dd className="font-mono">{exam.code}</dd>
          </div>
          <div>
            <dt className="text-sm font-medium text-gray-500">Nama</dt>
            <dd>{exam.name}</dd>
          </div>
          <div>
            <dt className="text-sm font-medium text-gray-500">Jadwal</dt>
            <dd>
              {formatDate(exam.jadwal_mulai)} - {formatDate(exam.jadwal_selesai)}
            </dd>
          </div>
          <div>
            <dt className="text-sm font-medium text-gray-500">Durasi</dt>
            <dd>{exam.durasi_menit} menit</dd>
          </div>
          <div className="flex gap-4">
            <div>
              <dt className="text-sm font-medium text-gray-500">Status</dt>
              <dd>
                <span className="inline-flex px-2 py-0.5 rounded text-xs font-medium bg-gray-100">
                  {exam.status}
                </span>
              </dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Verifikasi</dt>
              <dd>
                <span className="inline-flex px-2 py-0.5 rounded text-xs font-medium bg-blue-100 text-blue-800">
                  {exam.verification_status}
                </span>
              </dd>
            </div>
          </div>
          <div>
            <dt className="text-sm font-medium text-gray-500">Konten</dt>
            <dd className="text-sm">
              {exam.contents?.length
                ? exam.contents
                    .sort((a, b) => a.sort_order - b.sort_order)
                    .map((c) => (
                      <span key={c.source_id} className="inline-block mr-2">
                        {c.source_type}: {c.source_id}
                      </span>
                    ))
                : "-"}
            </dd>
          </div>
          <div>
            <dt className="text-sm font-medium text-gray-500">Peserta</dt>
            <dd className="text-sm">
              {exam.participants?.length ?? 0} orang
            </dd>
          </div>
        </dl>
      </div>

      {deleteDialog && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4">
          <div className="bg-white rounded-lg shadow-xl p-6 max-w-md w-full">
            <h3 className="text-lg font-semibold">Konfirmasi Hapus</h3>
            <p className="mt-2 text-gray-600">Alasan penghapusan wajib diisi.</p>
            <textarea
              value={deleteReason}
              onChange={(e) => setDeleteReason(e.target.value)}
              placeholder="Alasan..."
              rows={3}
              className="mt-3 w-full px-3 py-2 border rounded-lg"
            />
            <div className="mt-4 flex justify-end gap-2">
              <button onClick={() => setDeleteDialog(false)} className="px-4 py-2 border rounded-lg">
                Batal
              </button>
              <button
                onClick={handleDelete}
                disabled={!deleteReason.trim() || deleteLoading}
                className="px-4 py-2 bg-red-600 text-white rounded-lg disabled:opacity-50"
              >
                {deleteLoading ? "Menghapus..." : "Hapus"}
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
