"use client";

import { apiService } from "@/lib/api";
import { ConfirmDialog } from "@/components";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";

interface ArticleItem {
  id: string;
  title: string;
  slug: string;
  published_at?: string;
  status: string;
}

interface ArticleListResponse {
  items: ArticleItem[];
  total_page: number;
  total_data: number;
}

export default function BeritaCMSListPage() {
  const [items, setItems] = useState<ArticleItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [deleteId, setDeleteId] = useState<string | null>(null);
  const [search, setSearch] = useState("");
  const [page, setPage] = useState(1);

  const fetchList = useCallback(async () => {
    setLoading(true);
    const res = await apiService.get<ArticleListResponse>("v1/cms/berita", {
      q: search || undefined,
      page,
      page_size: 20,
    });
    if (res.success && res.data) {
      const d = res.data as ArticleListResponse;
      setItems(d.items ?? []);
    }
    setLoading(false);
  }, [search, page]);

  useEffect(() => {
    fetchList();
  }, [fetchList]);

  const handleDelete = async () => {
    if (!deleteId) return;
    const res = await apiService.delete(`v1/cms/berita/${deleteId}`);
    setDeleteId(null);
    if (res.success) fetchList();
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">CMS Berita</h1>
          <p className="text-gray-600 mt-1">Kelola berita landing page</p>
        </div>
        <Link href="/cms/berita/create" className="inline-flex px-4 py-2 rounded-lg bg-blue-600 text-white font-medium transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700">
          Tambah Berita
        </Link>
      </div>
      <form onSubmit={(e) => { e.preventDefault(); setPage(1); fetchList(); }} className="flex gap-2">
        <input type="text" value={search} onChange={(e) => setSearch(e.target.value)} placeholder="Cari judul..." className="rounded-lg border border-gray-300 px-3 py-2 text-sm w-64" />
        <button type="submit" className="rounded-lg bg-gray-100 px-4 py-2 text-sm font-medium">Cari</button>
      </form>
      <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
        {loading ? (
          <div className="p-8 space-y-3">
            {[1, 2, 3].map((i) => (
              <div key={i} className="h-12 bg-gray-100 rounded animate-pulse" />
            ))}
          </div>
        ) : items.length === 0 ? (
          <p className="p-8 text-gray-500">Belum ada berita.</p>
        ) : (
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Judul</th>
                <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">Aksi</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-200">
              {items.map((item) => (
                <tr key={item.id} className="transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50">
                  <td className="px-4 py-3">
                    <Link href={`/cms/berita/${item.id}`} className="font-medium text-gray-900 hover:text-blue-600">
                      {item.title || "-"}
                    </Link>
                    {item.published_at && (
                      <p className="text-xs text-gray-500 mt-0.5">{new Date(item.published_at).toLocaleDateString("id-ID")}</p>
                    )}
                  </td>
                  <td className="px-4 py-3">
                    <span className={`px-2 py-0.5 rounded text-xs font-medium ${item.status === "Published" ? "bg-green-100 text-green-800" : "bg-gray-100"}`}>
                      {item.status}
                    </span>
                  </td>
                  <td className="px-4 py-3 text-right space-x-2">
                    <Link href={`/cms/berita/${item.id}`} className="text-blue-600 hover:underline text-sm">Detail</Link>
                    <Link href={`/cms/berita/${item.id}/edit`} className="text-blue-600 hover:underline text-sm">Edit</Link>
                    <button type="button" onClick={() => setDeleteId(item.id)} className="text-red-600 hover:underline text-sm">Hapus</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
      <ConfirmDialog isOpen={!!deleteId} onClose={() => setDeleteId(null)} onConfirm={handleDelete} title="Hapus Berita" message="Apakah Anda yakin?" confirmText="Ya, Hapus" type="danger" />
    </div>
  );
}
