"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";

interface RoleListItem {
  id: string;
  code: string;
  name: string;
  description?: string;
  permissions?: string[];
}

interface RoleListResponse {
  items: RoleListItem[];
  total_page: number;
  total_data: number;
  page: number;
  page_size: number;
}

export default function RoleListPage() {
  const [roles, setRoles] = useState<RoleListItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [search, setSearch] = useState("");
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [totalData, setTotalData] = useState(0);
  const pageSize = 10;

  const fetchRoles = useCallback(async () => {
    setLoading(true);
    setError("");
    const res = await apiService.get<RoleListResponse>("v1/rbac/roles", {
      q: search || undefined,
      page,
      page_size: pageSize,
    });
    if (res.success && res.data) {
      const d = res.data as RoleListResponse;
      setRoles(d.items ?? []);
      setTotalPage(d.total_page ?? 1);
      setTotalData(d.total_data ?? 0);
    } else {
      setError(res.message ?? "Gagal memuat daftar role");
    }
    setLoading(false);
  }, [search, page]);

  useEffect(() => {
    fetchRoles();
  }, [fetchRoles]);

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    setPage(1);
    fetchRoles();
  };

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Kelola Role</h1>
          <p className="text-gray-600 mt-1">Daftar role dan permission</p>
        </div>
        <Link
          href="/auth/role/create"
          className="inline-flex items-center justify-center px-4 py-2 rounded-lg bg-blue-600 text-white transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 font-medium"
        >
          Tambah Role
        </Link>
      </div>

      <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
        <form onSubmit={handleSearch} className="p-4 border-b border-gray-200">
          <div className="flex gap-2">
            <input
              type="text"
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              placeholder="Cari kode, nama role..."
              className="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
            />
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
        ) : roles.length === 0 ? (
          <div className="p-12 text-center text-gray-500">Tidak ada data role</div>
        ) : (
          <>
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Kode</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nama</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Jumlah Permission</th>
                    <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">Aksi</th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {roles.map((r) => (
                    <tr key={r.id}>
                      <td className="px-4 py-3 text-sm font-medium text-gray-900">{r.code}</td>
                      <td className="px-4 py-3 text-sm text-gray-600">{r.name}</td>
                      <td className="px-4 py-3">
                        <span className="inline-flex px-2 py-0.5 rounded text-xs font-medium bg-blue-100 text-blue-800">
                          {r.permissions?.length ?? 0} permission
                        </span>
                      </td>
                      <td className="px-4 py-3 text-right">
                        <Link
                          href={`/auth/role/${r.id}`}
                          className="text-blue-600 hover:text-blue-800 text-sm font-medium mr-3"
                        >
                          Detail
                        </Link>
                        <Link
                          href={`/auth/role/${r.id}/edit`}
                          className="text-blue-600 hover:text-blue-800 text-sm font-medium"
                        >
                          Edit
                        </Link>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
            {totalPage > 1 && (
              <div className="px-4 py-3 border-t border-gray-200 flex items-center justify-between">
                <p className="text-sm text-gray-600">
                  Menampilkan {(page - 1) * pageSize + 1} - {Math.min(page * pageSize, totalData)} dari{" "}
                  {totalData}
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
