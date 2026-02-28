"use client";

import { apiService } from "@/lib/api";
import { ConfirmDialog } from "@/components";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";

interface LinkItem {
  id: string;
  title: string;
  url: string;
  status: string;
}

interface LinkListResponse {
  items: LinkItem[];
  total_page: number;
  total_data: number;
}

export default function TautanListPage() {
  const [items, setItems] = useState<LinkItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [deleteId, setDeleteId] = useState<string | null>(null);

  const fetchList = useCallback(async () => {
    setLoading(true);
    const res = await apiService.get<LinkListResponse>("v1/cms/tautan", { page: 1, page_size: 50 });
    if (res.success && res.data) {
      const d = res.data as LinkListResponse;
      setItems(d.items ?? []);
    }
    setLoading(false);
  }, []);

  useEffect(() => {
    fetchList();
  }, [fetchList]);

  const handleDelete = async () => {
    if (!deleteId) return;
    const res = await apiService.delete(`v1/cms/tautan/${deleteId}`);
    setDeleteId(null);
    if (res.success) fetchList();
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">CMS Tautan</h1>
          <p className="text-gray-600 mt-1">Kelola tautan penting landing page</p>
        </div>
        <Link href="/cms/tautan/create" className="inline-flex px-4 py-2 rounded-lg bg-blue-600 text-white font-medium hover:bg-blue-700">
          Tambah Tautan
        </Link>
      </div>
      <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
        {loading ? (
          <div className="p-8 space-y-3">
            {[1, 2, 3].map((i) => (
              <div key={i} className="h-12 bg-gray-100 rounded animate-pulse" />
            ))}
          </div>
        ) : items.length === 0 ? (
          <p className="p-8 text-gray-500">Belum ada tautan.</p>
        ) : (
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Judul</th>
                <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">URL</th>
                <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">Aksi</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-200">
              {items.map((item) => (
                <tr key={item.id} className="hover:bg-gray-50">
                  <td className="px-4 py-3 font-medium text-gray-900">{item.title || "-"}</td>
                  <td className="px-4 py-3 text-sm text-gray-600">
                    <a href={item.url} target="_blank" rel="noopener noreferrer" className="text-blue-600 hover:underline truncate max-w-xs block">
                      {item.url}
                    </a>
                  </td>
                  <td className="px-4 py-3">
                    <span className={`px-2 py-0.5 rounded text-xs font-medium ${item.status === "Aktif" ? "bg-green-100 text-green-800" : "bg-gray-100"}`}>
                      {item.status}
                    </span>
                  </td>
                  <td className="px-4 py-3 text-right space-x-2">
                    <Link href={`/cms/tautan/${item.id}`} className="text-blue-600 hover:underline text-sm">Detail</Link>
                    <Link href={`/cms/tautan/${item.id}/edit`} className="text-blue-600 hover:underline text-sm">Edit</Link>
                    <button type="button" onClick={() => setDeleteId(item.id)} className="text-red-600 hover:underline text-sm">Hapus</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
      <ConfirmDialog isOpen={!!deleteId} onClose={() => setDeleteId(null)} onConfirm={handleDelete} title="Hapus Tautan" message="Apakah Anda yakin?" confirmText="Ya, Hapus" type="danger" />
    </div>
  );
}
