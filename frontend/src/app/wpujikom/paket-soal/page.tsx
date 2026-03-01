"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";
import { getUserProfile } from "@/lib/auth";

interface PackageListItem {
  id: string;
  code: string;
  name: string;
  description?: string;
  status: string;
  verification_status: string;
  question_count: number;
  created_at?: string;
}

interface PackageListResponse {
  items: PackageListItem[];
  total_page: number;
  total_data: number;
  page: number;
  page_size: number;
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

export default function PaketSoalPage() {
  const [packages, setPackages] = useState<PackageListItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [search, setSearch] = useState("");
  const [status, setStatus] = useState("");
  const [statusVerifikasi, setStatusVerifikasi] = useState("");
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [totalData, setTotalData] = useState(0);
  const [canEditRole, setCanEditRole] = useState(false);
  const [canVerifyRole, setCanVerifyRole] = useState(false);
  const pageSize = 20;

  useEffect(() => {
    getUserProfile().then((p) => {
      setCanEditRole(canEdit(p?.role_user));
      setCanVerifyRole(canVerify(p?.role_user));
    });
  }, []);

  const fetchPackages = useCallback(async () => {
    setLoading(true);
    setError("");
    const res = await apiService.get<PackageListResponse>("v1/question-packages", {
      q: search || undefined,
      status: status || undefined,
      status_verifikasi: statusVerifikasi || undefined,
      page,
      page_size: pageSize,
    });
    if (res.success && res.data) {
      const d = res.data as PackageListResponse;
      setPackages(d.items ?? []);
      setTotalPage(d.total_page ?? 1);
      setTotalData(d.total_data ?? 0);
    } else {
      setError(res.message ?? "Gagal memuat daftar paket soal");
    }
    setLoading(false);
  }, [search, status, statusVerifikasi, page]);

  useEffect(() => {
    fetchPackages();
  }, [fetchPackages]);

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    setPage(1);
    fetchPackages();
  };

  const handleVerify = async (id: string, verify: boolean) => {
    const endpoint = verify
      ? `v1/question-packages/${id}/verify`
      : `v1/question-packages/${id}/unverify`;
    const res = await apiService.post(endpoint, {});
    if (res.success) fetchPackages();
    else setError(res.message ?? "Gagal verifikasi");
  };

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Paket Soal</h1>
          <p className="text-gray-600 mt-1">Kumpulan soal untuk ujian</p>
        </div>
        {canEditRole && (
          <Link
            href="/wpujikom/paket-soal/create"
            className="inline-flex items-center justify-center px-4 py-2 rounded-lg bg-blue-600 text-white transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 font-medium"
          >
            Tambah Paket
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
        ) : packages.length === 0 ? (
          <div className="p-12 text-center text-gray-500">Tidak ada paket soal</div>
        ) : (
          <>
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Kode</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nama</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Jumlah Soal</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Verifikasi</th>
                    <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">Aksi</th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {packages.map((p) => (
                    <tr key={p.id}>
                      <td className="px-4 py-3 text-sm font-mono text-gray-900">{p.code}</td>
                      <td className="px-4 py-3 text-sm text-gray-600">{p.name}</td>
                      <td className="px-4 py-3 text-sm text-gray-600">{p.question_count}</td>
                      <td className="px-4 py-3">
                        <span
                          className={`inline-flex px-2 py-0.5 rounded text-xs font-medium ${
                            p.status === "Aktif" ? "bg-green-100 text-green-800" : "bg-gray-100 text-gray-700"
                          }`}
                        >
                          {p.status}
                        </span>
                      </td>
                      <td className="px-4 py-3">
                        <span
                          className={`inline-flex px-2 py-0.5 rounded text-xs font-medium ${
                            p.verification_status === "Sudah"
                              ? "bg-blue-100 text-blue-800"
                              : "bg-amber-100 text-amber-800"
                          }`}
                        >
                          {p.verification_status}
                        </span>
                      </td>
                      <td className="px-4 py-3 text-right">
                        <Link
                          href={`/wpujikom/paket-soal/${p.id}`}
                          className="text-blue-600 hover:text-blue-800 text-sm font-medium mr-2"
                        >
                          Detail
                        </Link>
                        {canEditRole && (
                          <Link
                            href={`/wpujikom/paket-soal/${p.id}/edit`}
                            className="text-blue-600 hover:text-blue-800 text-sm font-medium mr-2"
                          >
                            Edit
                          </Link>
                        )}
                        {canVerifyRole && (
                          <button
                            onClick={() => handleVerify(p.id, p.verification_status !== "Sudah")}
                            className="text-amber-600 hover:text-amber-800 text-sm font-medium"
                          >
                            {p.verification_status === "Sudah" ? "Batal Verifikasi" : "Verifikasi"}
                          </button>
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
