"use client";

import { apiService } from "@/lib/api";
import { ConfirmDialog } from "@/components";
import Link from "next/link";
import Image from "next/image";
import { useCallback, useEffect, useState } from "react";

interface SlideItem {
  id: string;
  image_url: string;
  title: string;
  subtitle: string;
  sort_order: number;
  status: string;
}

interface SlideListResponse {
  items: SlideItem[];
  total_page: number;
  total_data: number;
}

export default function SliderListPage() {
  const [items, setItems] = useState<SlideItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [deleteId, setDeleteId] = useState<string | null>(null);

  const fetchList = useCallback(async () => {
    setLoading(true);
    const res = await apiService.get<SlideListResponse>("v1/cms/slider", {
      page: 1,
      page_size: 50,
    });
    if (res.success && res.data) {
      const d = res.data as SlideListResponse;
      setItems(d.items ?? []);
    }
    setLoading(false);
  }, []);

  useEffect(() => {
    fetchList();
  }, [fetchList]);

  const handleDelete = async () => {
    if (!deleteId) return;
    const res = await apiService.delete(`v1/cms/slider/${deleteId}`);
    setDeleteId(null);
    if (res.success) fetchList();
  };

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">CMS Slider</h1>
          <p className="text-gray-600 mt-1">Kelola slider landing page</p>
        </div>
        <Link
          href="/cms/slider/create"
          className="inline-flex px-4 py-2 rounded-lg bg-blue-600 text-white font-medium hover:bg-blue-700"
        >
          Tambah Slide
        </Link>
      </div>
      <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
        {loading ? (
          <div className="p-8 space-y-3">
            {[1, 2, 3].map((i) => (
              <div key={i} className="h-16 bg-gray-100 rounded animate-pulse" />
            ))}
          </div>
        ) : items.length === 0 ? (
          <p className="p-8 text-gray-500">Belum ada slide.</p>
        ) : (
          <table className="min-w-full divide-y divide-gray-200">
            <thead className="bg-gray-50">
              <tr>
                <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Thumbnail
                </th>
                <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Judul
                </th>
                <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Urutan
                </th>
                <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                  Status
                </th>
                <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">
                  Aksi
                </th>
              </tr>
            </thead>
            <tbody className="divide-y divide-gray-200">
              {items.map((item) => (
                <tr key={item.id} className="hover:bg-gray-50">
                  <td className="px-4 py-3">
                    {item.image_url ? (
                      <div className="relative h-12 w-20 rounded overflow-hidden bg-gray-100">
                        <Image
                          src={item.image_url}
                          alt=""
                          fill
                          className="object-cover"
                          sizes="80px"
                          unoptimized={item.image_url.startsWith("http")}
                        />
                      </div>
                    ) : (
                      <span className="text-gray-400 text-sm">-</span>
                    )}
                  </td>
                  <td className="px-4 py-3 text-sm font-medium text-gray-900">
                    {item.title || "-"}
                  </td>
                  <td className="px-4 py-3 text-sm text-gray-600">{item.sort_order}</td>
                  <td className="px-4 py-3">
                    <span
                      className={`px-2 py-0.5 rounded text-xs font-medium ${
                        item.status === "Published" ? "bg-green-100 text-green-800" : "bg-gray-100 text-gray-700"
                      }`}
                    >
                      {item.status}
                    </span>
                  </td>
                  <td className="px-4 py-3 text-right space-x-2">
                    <Link
                      href={`/cms/slider/${item.id}`}
                      className="text-blue-600 hover:underline text-sm"
                    >
                      Detail
                    </Link>
                    <Link
                      href={`/cms/slider/${item.id}/edit`}
                      className="text-blue-600 hover:underline text-sm"
                    >
                      Edit
                    </Link>
                    <button
                      type="button"
                      onClick={() => setDeleteId(item.id)}
                      className="text-red-600 hover:underline text-sm"
                    >
                      Hapus
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        )}
      </div>
      <ConfirmDialog
        isOpen={!!deleteId}
        onClose={() => setDeleteId(null)}
        onConfirm={handleDelete}
        title="Hapus Slide"
        message="Apakah Anda yakin ingin menghapus slide ini?"
        confirmText="Ya, Hapus"
        cancelText="Batal"
        type="danger"
      />
    </div>
  );
}
