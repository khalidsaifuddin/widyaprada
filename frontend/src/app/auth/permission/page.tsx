"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";

interface PermissionListItem {
  id: string;
  code: string;
  name: string;
  group?: string;
  description?: string;
}

interface PermissionListResponse {
  items: PermissionListItem[];
  total_page: number;
  total_data: number;
  page: number;
  page_size: number;
}

export default function PermissionListPage() {
  const [permissions, setPermissions] = useState<PermissionListItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [search, setSearch] = useState("");
  const [groupFilter, setGroupFilter] = useState("");
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [totalData, setTotalData] = useState(0);
  const pageSize = 10;

  const fetchPermissions = useCallback(async () => {
    setLoading(true);
    setError("");
    const res = await apiService.get<PermissionListResponse>("v1/rbac/permissions", {
      q: search || undefined,
      group: groupFilter || undefined,
      page,
      page_size: pageSize,
    });
    if (res.success && res.data) {
      const d = res.data as PermissionListResponse;
      setPermissions(d.items ?? []);
      setTotalPage(d.total_page ?? 1);
      setTotalData(d.total_data ?? 0);
    } else {
      setError(res.message ?? "Gagal memuat daftar permission");
    }
    setLoading(false);
  }, [search, groupFilter, page]);

  useEffect(() => {
    fetchPermissions();
  }, [fetchPermissions]);

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    setPage(1);
    fetchPermissions();
  };

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Kelola Permission</h1>
          <p className="text-gray-600 mt-1">Daftar permission aplikasi</p>
        </div>
        <Link
          href="/auth/permission/create"
          className="inline-flex items-center justify-center px-4 py-2 rounded-lg bg-blue-600 text-white transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 font-medium"
        >
          Tambah Permission
        </Link>
      </div>

      <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
        <form onSubmit={handleSearch} className="p-4 border-b border-gray-200 space-y-3">
          <div className="flex flex-wrap gap-2">
            <input
              type="text"
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              placeholder="Cari kode, nama..."
              className="flex-1 min-w-[200px] px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            />
            <input
              type="text"
              value={groupFilter}
              onChange={(e) => setGroupFilter(e.target.value)}
              placeholder="Filter modul/group"
              className="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 w-48"
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
        ) : permissions.length === 0 ? (
          <div className="p-12 text-center text-gray-500">Tidak ada data permission</div>
        ) : (
          <>
            <div className="overflow-x-auto">
              <table className="min-w-full divide-y divide-gray-200">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Kode</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nama</th>
                    <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Modul/Group</th>
                    <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">Aksi</th>
                  </tr>
                </thead>
                <tbody className="bg-white divide-y divide-gray-200">
                  {permissions.map((p) => (
                    <tr key={p.id}>
                      <td className="px-4 py-3 text-sm font-medium text-gray-900 font-mono">{p.code}</td>
                      <td className="px-4 py-3 text-sm text-gray-600">{p.name}</td>
                      <td className="px-4 py-3">
                        {p.group ? (
                          <span className="inline-flex px-2 py-0.5 rounded text-xs font-medium bg-gray-100 text-gray-800">
                            {p.group}
                          </span>
                        ) : (
                          <span className="text-gray-400">-</span>
                        )}
                      </td>
                      <td className="px-4 py-3 text-right">
                        <Link
                          href={`/auth/permission/${p.id}`}
                          className="text-blue-600 hover:text-blue-800 text-sm font-medium mr-3"
                        >
                          Detail
                        </Link>
                        <Link
                          href={`/auth/permission/${p.id}/edit`}
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
