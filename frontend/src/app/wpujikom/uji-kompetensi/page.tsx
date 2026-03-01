"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";

interface ExamListItem {
  id: string;
  code: string;
  name: string;
  jadwal_mulai: string;
  jadwal_selesai: string;
  durasi_menit: number;
  status: string;
  verification_status: string;
  participant_count: number;
  created_at?: string;
}

interface ExamListResponse {
  items: ExamListItem[];
  total_page: number;
  total_data: number;
  page: number;
  page_size: number;
}

function formatDate(d: string): string {
  try {
    return new Date(d).toLocaleDateString("id-ID", {
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

export default function UjiKompetensiPage() {
  const [exams, setExams] = useState<ExamListItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [search, setSearch] = useState("");
  const [status, setStatus] = useState("");
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [totalData, setTotalData] = useState(0);
  const pageSize = 20;

  const fetchExams = useCallback(async () => {
    setLoading(true);
    setError("");
    const res = await apiService.get<ExamListResponse>("v1/exams", {
      q: search || undefined,
      status: status || undefined,
      page,
      page_size: pageSize,
    });
    if (res.success && res.data) {
      const d = res.data as ExamListResponse;
      setExams(d.items ?? []);
      setTotalPage(d.total_page ?? 1);
      setTotalData(d.total_data ?? 0);
    } else {
      setError(res.message ?? "Gagal memuat daftar ujian");
    }
    setLoading(false);
  }, [search, status, page]);

  useEffect(() => {
    fetchExams();
  }, [fetchExams]);

  const handlePublish = async (id: string) => {
    const res = await apiService.post(`v1/exams/${id}/publish`, {});
    if (res.success) fetchExams();
    else setError(res.message ?? "Gagal menerbitkan");
  };

  const handleVerify = async (id: string, verify: boolean) => {
    const endpoint = verify ? `v1/exams/${id}/verify` : `v1/exams/${id}/unverify`;
    const res = await apiService.post(endpoint, {});
    if (res.success) fetchExams();
    else setError(res.message ?? "Gagal verifikasi");
  };

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Manajemen Uji Kompetensi</h1>
          <p className="text-gray-600 mt-1">Kelola ujian dan jadwal</p>
        </div>
        <Link
          href="/wpujikom/uji-kompetensi/create"
          className="inline-flex items-center justify-center px-4 py-2 rounded-lg bg-blue-600 text-white transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 font-medium"
        >
          Tambah Ujian
        </Link>
      </div>

      <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
        <form
          onSubmit={(e) => {
            e.preventDefault();
            setPage(1);
            fetchExams();
          }}
          className="p-4 border-b border-gray-200 space-y-3"
        >
          <div className="flex flex-wrap gap-2">
            <input
              type="text"
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              placeholder="Cari kode, nama..."
              className="flex-1 min-w-[180px] px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            />
            <select
              value={status}
              onChange={(e) => setStatus(e.target.value)}
              className="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            >
              <option value="">Semua Status</option>
              <option value="Draft">Draft</option>
              <option value="Diterbitkan">Diterbitkan</option>
              <option value="Berlangsung">Berlangsung</option>
              <option value="Selesai">Selesai</option>
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
        ) : exams.length === 0 ? (
          <div className="p-12 text-center text-gray-500">Tidak ada ujian</div>
        ) : (
          <>
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Kode</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nama</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Jadwal</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Durasi</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Peserta</th>
                    <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">Aksi</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-200">
                  {exams.map((e) => (
                    <tr key={e.id}>
                      <td className="px-4 py-3 text-sm font-mono text-gray-900">{e.code}</td>
                      <td className="px-4 py-3 text-sm text-gray-600">{e.name}</td>
                      <td className="px-4 py-3 text-sm text-gray-600">
                        {formatDate(e.jadwal_mulai)} - {formatDate(e.jadwal_selesai)}
                      </td>
                      <td className="px-4 py-3 text-sm text-gray-600">{e.durasi_menit} mnt</td>
                      <td className="px-4 py-3">
                        <span
                          className={`inline-flex px-2 py-0.5 rounded text-xs font-medium ${
                            e.status === "Diterbitkan" || e.status === "Berlangsung"
                              ? "bg-green-100 text-green-800"
                              : e.status === "Selesai"
                                ? "bg-gray-100 text-gray-700"
                                : "bg-amber-100 text-amber-800"
                          }`}
                        >
                          {e.status}
                        </span>
                      </td>
                      <td className="px-4 py-3 text-sm">{e.participant_count}</td>
                      <td className="px-4 py-3 text-right">
                        <Link
                          href={`/wpujikom/uji-kompetensi/${e.id}`}
                          className="text-blue-600 hover:text-blue-800 text-sm font-medium mr-2"
                        >
                          Detail
                        </Link>
                        {e.status === "Draft" && (
                          <>
                            <Link
                              href={`/wpujikom/uji-kompetensi/${e.id}/edit`}
                              className="text-blue-600 hover:text-blue-800 text-sm font-medium mr-2"
                            >
                              Edit
                            </Link>
                            <button
                              onClick={() => handlePublish(e.id)}
                              className="text-green-600 hover:text-green-800 text-sm font-medium mr-2"
                            >
                              Terbitkan
                            </button>
                          </>
                        )}
                        <button
                          onClick={() => handleVerify(e.id, e.verification_status !== "Sudah")}
                          className="text-amber-600 hover:text-amber-800 text-sm font-medium"
                        >
                          {e.verification_status === "Sudah" ? "Batal Verifikasi" : "Verifikasi"}
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
            {totalPage > 1 && (
              <div className="px-4 py-3 border-t border-gray-200 flex justify-between">
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
