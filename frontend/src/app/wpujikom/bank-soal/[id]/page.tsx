"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
import { getUserProfile } from "@/lib/auth";

interface QuestionOption {
  id: string;
  option_key: string;
  option_text: string;
  is_correct: boolean;
}

interface QuestionDetail {
  id: string;
  code: string;
  type: string;
  category_id: string;
  category_name?: string;
  difficulty: string;
  question_text: string;
  answer_key: string;
  weight: number;
  status: string;
  verification_status: string;
  options: QuestionOption[];
  created_at?: string;
  updated_at?: string;
}

function isSuperAdmin(roleUser: { role_aplikasi?: string }[] | undefined): boolean {
  return roleUser?.some((r) => r.role_aplikasi === "SUPER_ADMIN") ?? false;
}

function canVerify(roleUser: { role_aplikasi?: string }[] | undefined): boolean {
  return roleUser?.some((r) =>
    r.role_aplikasi === "SUPER_ADMIN" || r.role_aplikasi === "VERIFIKATOR"
  ) ?? false;
}

export default function BankSoalDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [question, setQuestion] = useState<QuestionDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [deleteDialog, setDeleteDialog] = useState(false);
  const [deleteReason, setDeleteReason] = useState("");
  const [deleteLoading, setDeleteLoading] = useState(false);
  const [canEdit, setCanEdit] = useState(false);
  const [canVerifyRole, setCanVerifyRole] = useState(false);

  useEffect(() => {
    getUserProfile().then((p) => {
      setCanEdit(isSuperAdmin(p?.role_user));
      setCanVerifyRole(canVerify(p?.role_user));
    });
  }, []);

  const fetchQuestion = useCallback(async () => {
    if (!id) return;
    setLoading(true);
    const res = await apiService.get<QuestionDetail>("v1/questions/" + id);
    if (res.success && res.data) {
      setQuestion(res.data as QuestionDetail);
    } else {
      setError(res.message ?? "Soal tidak ditemukan");
    }
    setLoading(false);
  }, [id]);

  useEffect(() => {
    fetchQuestion();
  }, [fetchQuestion]);

  const handleDelete = async () => {
    if (!deleteReason.trim()) return;
    setDeleteLoading(true);
    const res = await apiService.delete("v1/questions/" + id, { reason: deleteReason.trim() });
    if (res.success) {
      router.push("/wpujikom/bank-soal");
      return;
    }
    setError(res.message ?? "Gagal menghapus soal");
    setDeleteLoading(false);
  };

  const handleVerify = async (verify: boolean) => {
    const endpoint = verify ? `v1/questions/${id}/verify` : `v1/questions/${id}/unverify`;
    const res = await apiService.post(endpoint, {});
    if (res.success) fetchQuestion();
    else setError(res.message ?? "Gagal verifikasi");
  };

  if (loading) return <div className="p-8">Memuat...</div>;
  if (error && !question) {
    return (
      <div className="space-y-4">
        <Link href="/wpujikom/bank-soal" className="text-blue-600 hover:underline">
          ← Kembali
        </Link>
        <p className="text-red-600">{error}</p>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <Link href="/wpujikom/bank-soal" className="text-gray-600 hover:text-gray-900">
          ← Kembali
        </Link>
        <div className="flex gap-2">
          {canEdit && (
            <Link
              href={`/wpujikom/bank-soal/${id}/edit`}
              className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
            >
              Edit
            </Link>
          )}
          {canVerifyRole && (
            <button
              onClick={() => handleVerify(question!.verification_status !== "Sudah")}
              className="px-4 py-2 bg-amber-600 text-white rounded-lg hover:bg-amber-700"
            >
              {question!.verification_status === "Sudah" ? "Batal Verifikasi" : "Verifikasi"}
            </button>
          )}
          {canEdit && (
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

      {question && (
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 max-w-3xl space-y-6">
          <h1 className="text-xl font-bold text-gray-900">Detail Soal</h1>
          <dl className="space-y-4">
            <div>
              <dt className="text-sm font-medium text-gray-500">Kode</dt>
              <dd className="mt-1 font-mono text-gray-900">{question.code}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Tipe</dt>
              <dd className="mt-1">{question.type}</dd>
            </div>
            {question.category_name && (
              <div>
                <dt className="text-sm font-medium text-gray-500">Kategori</dt>
                <dd className="mt-1">{question.category_name}</dd>
              </div>
            )}
            {question.difficulty && (
              <div>
                <dt className="text-sm font-medium text-gray-500">Tingkat Kesulitan</dt>
                <dd className="mt-1">{question.difficulty}</dd>
              </div>
            )}
            <div>
              <dt className="text-sm font-medium text-gray-500">Teks Soal</dt>
              <dd className="mt-1 whitespace-pre-wrap text-gray-900">{question.question_text}</dd>
            </div>
            <div>
              <dt className="text-sm font-medium text-gray-500">Kunci Jawaban</dt>
              <dd className="mt-1 font-medium">{question.answer_key}</dd>
            </div>
            {question.options && question.options.length > 0 && (
              <div>
                <dt className="text-sm font-medium text-gray-500">Opsi</dt>
                <dd className="mt-1 space-y-1">
                  {question.options
                    .sort((a, b) => a.option_key.localeCompare(b.option_key))
                    .map((o) => (
                      <div
                        key={o.id}
                        className={`flex gap-2 ${o.is_correct ? "bg-green-50 p-2 rounded" : ""}`}
                      >
                        <span className="font-medium">{o.option_key}.</span>
                        <span>{o.option_text}</span>
                        {o.is_correct && (
                          <span className="text-green-600 text-sm font-medium">(Benar)</span>
                        )}
                      </div>
                    ))}
                </dd>
              </div>
            )}
            <div>
              <dt className="text-sm font-medium text-gray-500">Bobot</dt>
              <dd className="mt-1">{question.weight}</dd>
            </div>
            <div className="flex gap-4">
              <div>
                <dt className="text-sm font-medium text-gray-500">Status</dt>
                <dd>
                  <span
                    className={`inline-flex px-2 py-0.5 rounded text-xs font-medium ${
                      question.status === "Aktif" ? "bg-green-100 text-green-800" : "bg-gray-100 text-gray-700"
                    }`}
                  >
                    {question.status}
                  </span>
                </dd>
              </div>
              <div>
                <dt className="text-sm font-medium text-gray-500">Verifikasi</dt>
                <dd>
                  <span
                    className={`inline-flex px-2 py-0.5 rounded text-xs font-medium ${
                      question.verification_status === "Sudah"
                        ? "bg-blue-100 text-blue-800"
                        : "bg-amber-100 text-amber-800"
                    }`}
                  >
                    {question.verification_status}
                  </span>
                </dd>
              </div>
            </div>
          </dl>
        </div>
      )}

      {deleteDialog && (
        <div className="fixed inset-0 z-50 overflow-y-auto">
          <div className="flex min-h-full items-end justify-center p-4 sm:items-center sm:p-0">
            <div className="fixed inset-0 bg-black/50" onClick={() => setDeleteDialog(false)} />
            <div className="relative bg-white rounded-lg shadow-xl p-6 max-w-md w-full mx-4">
              <h3 className="text-lg font-semibold text-gray-900">Konfirmasi Hapus Soal</h3>
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
