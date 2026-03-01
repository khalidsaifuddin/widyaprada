"use client";

import { apiService } from "@/lib/api";
import { ConfirmDialog } from "@/components";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";

interface WPDataItem {
  id: string;
  nip: string;
  nama_lengkap: string;
  satker_id: string;
  unit_kerja: string;
  status: string;
  pendidikan_terakhir: string;
}

interface WPDataListResponse {
  items: WPDataItem[];
  total_page: number;
  total_data: number;
  page: number;
  page_size: number;
}

export default function WPDataListPage() {
  const [items, setItems] = useState<WPDataItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [search, setSearch] = useState("");
  const [status, setStatus] = useState("");
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [deleteId, setDeleteId] = useState<string | null>(null);

  const fetchList = useCallback(async () => {
    setLoading(true);
    setError("");
    const res = await apiService.get<WPDataListResponse>("v1/wp-data", {
      q: search || undefined,
      status: status || undefined,
      page,
      page_size: 10,
    });
    if (res.success && res.data) {
      const d = res.data as WPDataListResponse;
      setItems(d.items ?? []);
      setTotalPage(d.total_page ?? 1);
    } else {
      setError(res.message ?? "Gagal memuat data");
    }
    setLoading(false);
  }, [search, status, page]);

  useEffect(() => {
    fetchList();
  }, [fetchList]);

  const handleDelete = async () => {
    if (!deleteId) return;
    const res = await apiService.delete(`v1/wp-data/${deleteId}`, { reason: "Dihapus oleh admin" });
    setDeleteId(null);
    if (res.success) fetchList();
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Manajemen Data WP</h1>
          <p className="text-gray-600 mt-1">Data Widyaprada - NIP, Nama, Satker, Unit Kerja</p>
        </div>
        <Link href="/wpdata/create" className="inline-flex px-4 py-2 rounded-lg bg-blue-600 text-white font-medium transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700">
          Tambah Data WP
        </Link>
      </div>
      <form onSubmit={(e) => { e.preventDefault(); setPage(1); fetchList(); }} className="flex flex-wrap gap-2">
        <input type="text" value={search} onChange={(e) => setSearch(e.target.value)} placeholder="Cari NIP, nama..." className="rounded-lg border border-gray-300 px-3 py-2 text-sm w-56" />
        <select value={status} onChange={(e) => { setStatus(e.target.value); setPage(1); }} className="rounded-lg border border-gray-300 px-3 py-2 text-sm">
          <option value="">Semua Status</option>
          <option value="Aktif">Aktif</option>
          <option value="Nonaktif">Nonaktif</option>
        </select>
        <button type="submit" className="rounded-lg bg-gray-100 px-4 py-2 text-sm font-medium">Cari</button>
      </form>
      <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
        {loading ? (
          <div className="p-8 space-y-3">
            {[1, 2, 3].map((i) => (
              <div key={i} className="h-12 bg-gray-100 rounded animate-pulse" />
            ))}
          </div>
        ) : error ? (
          <p className="p-8 text-red-600">{error}</p>
        ) : items.length === 0 ? (
          <p className="p-8 text-gray-500">Belum ada data WP.</p>
        ) : (
          <>
            <table className="min-w-full divide-y divide-gray-200">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">NIP</th>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nama</th>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Satker</th>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Unit Kerja</th>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                  <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">Aksi</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-200">
                {items.map((item) => (
                  <tr key={item.id} className="transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50">
                    <td className="px-4 py-3 text-sm font-mono text-gray-900">{item.nip}</td>
                    <td className="px-4 py-3 text-sm font-medium text-gray-900">{item.nama_lengkap}</td>
                    <td className="px-4 py-3 text-sm text-gray-600">{item.satker_id || "-"}</td>
                    <td className="px-4 py-3 text-sm text-gray-600">{item.unit_kerja || "-"}</td>
                    <td className="px-4 py-3">
                      <span className={`px-2 py-0.5 rounded text-xs font-medium ${item.status === "Aktif" ? "bg-green-100 text-green-800" : "bg-gray-100"}`}>{item.status}</span>
                    </td>
                    <td className="px-4 py-3 text-right space-x-2">
                      <Link href={`/wpdata/${item.id}`} className="text-blue-600 hover:underline text-sm">Detail</Link>
                      <Link href={`/wpdata/${item.id}/edit`} className="text-blue-600 hover:underline text-sm">Edit</Link>
                      <button type="button" onClick={() => setDeleteId(item.id)} className="text-red-600 hover:underline text-sm">Hapus</button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
            {totalPage > 1 && (
              <div className="px-4 py-3 border-t border-gray-200 flex justify-center gap-2">
                <button type="button" onClick={() => setPage((p) => Math.max(1, p - 1))} disabled={page <= 1} className="rounded border px-2 py-1 text-sm disabled:opacity-50">Prev</button>
                <span className="px-2 py-1 text-sm">{page} / {totalPage}</span>
                <button type="button" onClick={() => setPage((p) => Math.min(totalPage, p + 1))} disabled={page >= totalPage} className="rounded border px-2 py-1 text-sm disabled:opacity-50">Next</button>
              </div>
            )}
          </>
        )}
      </div>
      <ConfirmDialog isOpen={!!deleteId} onClose={() => setDeleteId(null)} onConfirm={handleDelete} title="Hapus Data WP" message="Apakah Anda yakin? Masukkan alasan penghapusan jika diperlukan." confirmText="Ya, Hapus" type="danger" />
    </div>
  );
}
