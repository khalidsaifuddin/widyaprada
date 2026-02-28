"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";
import { getUserProfile } from "@/lib/auth";

interface QuestionListItem {
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
  created_at?: string;
}

interface QuestionListResponse {
  items: QuestionListItem[];
  total_page: number;
  total_data: number;
  page: number;
  page_size: number;
}

function isSuperAdmin(roleUser: { role_aplikasi?: string }[] | undefined): boolean {
  return roleUser?.some((r) => r.role_aplikasi === "SUPER_ADMIN") ?? false;
}

export default function BankSoalPage() {
  const [questions, setQuestions] = useState<QuestionListItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [search, setSearch] = useState("");
  const [tipe, setTipe] = useState("");
  const [kategoriId, setKategoriId] = useState("");
  const [status, setStatus] = useState("");
  const [statusVerifikasi, setStatusVerifikasi] = useState("");
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [totalData, setTotalData] = useState(0);
  const [categories, setCategories] = useState<{ id: string; code: string; name: string }[]>([]);
  const [canCreate, setCanCreate] = useState(false);
  const pageSize = 10;

  useEffect(() => {
    getUserProfile().then((p) => setCanCreate(isSuperAdmin(p?.role_user)));
  }, []);

  useEffect(() => {
    apiService
      .get<{ id: string; code: string; name: string }[]>("v1/questions/categories")
      .then((res) => {
        if (res.success && res.data) {
          const raw = res.data;
          setCategories(Array.isArray(raw) ? raw : []);
        }
      })
      .catch(() => {});
  }, []);

  const fetchQuestions = useCallback(async () => {
    setLoading(true);
    setError("");
    const res = await apiService.get<QuestionListResponse>("v1/questions", {
      q: search || undefined,
      tipe: tipe || undefined,
      kategori_id: kategoriId || undefined,
      status: status || undefined,
      status_verifikasi: statusVerifikasi || undefined,
      page,
      page_size: pageSize,
    });
    if (res.success && res.data) {
      const d = res.data as QuestionListResponse;
      setQuestions(d.items ?? []);
      setTotalPage(d.total_page ?? 1);
      setTotalData(d.total_data ?? 0);
    } else {
      setError(res.message ?? "Gagal memuat daftar soal");
    }
    setLoading(false);
  }, [search, tipe, kategoriId, status, statusVerifikasi, page]);

  useEffect(() => {
    fetchQuestions();
  }, [fetchQuestions]);

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    setPage(1);
    fetchQuestions();
  };

  const handleVerify = async (id: string, verify: boolean) => {
    const endpoint = verify ? `v1/questions/${id}/verify` : `v1/questions/${id}/unverify`;
    const res = await apiService.post(endpoint, {});
    if (res.success) fetchQuestions();
    else setError(res.message ?? "Gagal verifikasi");
  };

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Bank Soal</h1>
          <p className="text-gray-600 mt-1">Kelola soal PG, Benar–Salah, Essay</p>
        </div>
        {canCreate && (
          <Link
            href="/wpujikom/bank-soal/create"
            className="inline-flex items-center justify-center px-4 py-2 rounded-lg bg-blue-600 text-white hover:bg-blue-700 font-medium"
          >
            Tambah Soal
          </Link>
        )}
      </div>

      <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
        <form onSubmit={handleSearch} className="p-4 border-b border-gray-200 space-y-3">
          <div className="flex flex-wrap gap-2">
            <input
              type="text"
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              placeholder="Cari kode, teks soal..."
              className="flex-1 min-w-[180px] px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            />
            <select
              value={tipe}
              onChange={(e) => setTipe(e.target.value)}
              className="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            >
              <option value="">Semua Tipe</option>
              <option value="PG">Pilihan Ganda</option>
              <option value="BENAR_SALAH">Benar-Salah</option>
              <option value="ESSAY">Essay</option>
            </select>
            <select
              value={kategoriId}
              onChange={(e) => setKategoriId(e.target.value)}
              className="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            >
              <option value="">Semua Kategori</option>
              {categories.map((c) => (
                <option key={c.id} value={c.id}>
                  {c.name || c.code}
                </option>
              ))}
            </select>
            <select
              value={status}
              onChange={(e) => setStatus(e.target.value)}
              className="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            >
              <option value="">Semua Status</option>
              <option value="Draft">Draft</option>
              <option value="Aktif">Aktif</option>
            </select>
            <select
              value={statusVerifikasi}
              onChange={(e) => setStatusVerifikasi(e.target.value)}
              className="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            >
              <option value="">Verifikasi</option>
              <option value="Belum">Belum</option>
              <option value="Sudah">Sudah</option>
            </select>
            <button
              type="submit"
              className="px-4 py-2 bg-gray-100 hover:bg-gray-200 rounded-lg font-medium"
            >
              Cari
            </button>
          </div>
        </form>

        {error && (
          <div className="mx-4 mt-4 p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>
        )}

        {loading ? (
          <div className="p-12 text-center text-gray-500">Memuat...</div>
        ) : questions.length === 0 ? (
          <div className="p-12 text-center text-gray-500">Tidak ada soal</div>
        ) : (
          <>
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Kode</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Tipe</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Kategori</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase max-w-[200px]">Teks</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Verifikasi</th>
                    <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">Aksi</th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {questions.map((q) => (
                    <tr key={q.id}>
                      <td className="px-4 py-3 text-sm font-mono text-gray-900">{q.code}</td>
                      <td className="px-4 py-3 text-sm text-gray-600">{q.type}</td>
                      <td className="px-4 py-3 text-sm text-gray-600">{q.category_name || "-"}</td>
                      <td className="px-4 py-3 text-sm text-gray-600 truncate max-w-[200px]">{q.question_text}</td>
                      <td className="px-4 py-3">
                        <span
                          className={`inline-flex px-2 py-0.5 rounded text-xs font-medium ${
                            q.status === "Aktif" ? "bg-green-100 text-green-800" : "bg-gray-100 text-gray-700"
                          }`}
                        >
                          {q.status}
                        </span>
                      </td>
                      <td className="px-4 py-3">
                        <span
                          className={`inline-flex px-2 py-0.5 rounded text-xs font-medium ${
                            q.verification_status === "Sudah"
                              ? "bg-blue-100 text-blue-800"
                              : "bg-amber-100 text-amber-800"
                          }`}
                        >
                          {q.verification_status}
                        </span>
                      </td>
                      <td className="px-4 py-3 text-right">
                        <Link
                          href={`/wpujikom/bank-soal/${q.id}`}
                          className="text-blue-600 hover:text-blue-800 text-sm font-medium mr-2"
                        >
                          Detail
                        </Link>
                        {canCreate && (
                          <>
                            <Link
                              href={`/wpujikom/bank-soal/${q.id}/edit`}
                              className="text-blue-600 hover:text-blue-800 text-sm font-medium mr-2"
                            >
                              Edit
                            </Link>
                            <button
                              onClick={() => handleVerify(q.id, q.verification_status !== "Sudah")}
                              className="text-amber-600 hover:text-amber-800 text-sm font-medium"
                            >
                              {q.verification_status === "Sudah" ? "Batal Verifikasi" : "Verifikasi"}
                            </button>
                          </>
                        )}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
            {totalPage > 1 && (
              <div className="px-4 py-3 border-t border-gray-200 flex items-center justify-between">
                <p className="text-sm text-gray-600">
                  {(page - 1) * pageSize + 1} - {Math.min(page * pageSize, totalData)} dari {totalData}
                </p>
                <div className="flex gap-2">
                  <button
                    onClick={() => setPage((p) => Math.max(1, p - 1))}
                    disabled={page <= 1}
                    className="px-3 py-1 rounded border border-gray-300 disabled:opacity-50 text-sm"
                  >
                    Sebelumnya
                  </button>
                  <button
                    onClick={() => setPage((p) => Math.min(totalPage, p + 1))}
                    disabled={page >= totalPage}
                    className="px-3 py-1 rounded border border-gray-300 disabled:opacity-50 text-sm"
                  >
                    Selanjutnya
                  </button>
                </div>
              </div>
            )}
          </>
        )}
      </div>
    </div>
  );
}
