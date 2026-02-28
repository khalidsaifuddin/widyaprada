"use client";

import { apiService } from "@/lib/api";
import { ConfirmDialog } from "@/components";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface LinkDetail {
  id: string;
  title: string;
  url: string;
  description: string;
  status: string;
  buka_di_tab_baru: boolean;
}

export default function TautanDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params?.id as string;
  const [data, setData] = useState<LinkDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [deleteOpen, setDeleteOpen] = useState(false);

  useEffect(() => {
    if (!id) return;
    apiService.get<LinkDetail>(`v1/cms/tautan/${id}`).then((res) => {
      if (res.success && res.data) setData(res.data as LinkDetail);
      setLoading(false);
    });
  }, [id]);

  const handleDelete = async () => {
    const res = await apiService.delete(`v1/cms/tautan/${id}`);
    setDeleteOpen(false);
    if (res.success) router.push("/cms/tautan");
  };

  if (loading) return <div className="animate-pulse h-32 bg-gray-100 rounded" />;
  if (!data) return <p className="text-red-600">Tautan tidak ditemukan</p>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between">
        <Link href="/cms/tautan" className="text-blue-600 hover:underline text-sm">← Kembali</Link>
        <div className="flex gap-2">
          <Link href={`/cms/tautan/${id}/edit`} className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700">Edit</Link>
          <button type="button" onClick={() => setDeleteOpen(true)} className="rounded-lg border border-red-600 text-red-600 px-4 py-2 text-sm font-medium hover:bg-red-50">Hapus</button>
        </div>
      </div>
      <div className="bg-white rounded-xl border border-gray-200 p-6">
        <h1 className="text-xl font-bold text-gray-900 mb-4">{data.title}</h1>
        <dl className="space-y-2 text-sm">
          <div>
            <dt className="text-gray-500">URL</dt>
            <dd><a href={data.url} target="_blank" rel="noopener noreferrer" className="text-blue-600 hover:underline">{data.url}</a></dd>
          </div>
          {data.description && (
            <div>
              <dt className="text-gray-500">Deskripsi</dt>
              <dd>{data.description}</dd>
            </div>
          )}
          <div>
            <dt className="text-gray-500">Status</dt>
            <dd><span className={`px-2 py-0.5 rounded text-xs ${data.status === "Aktif" ? "bg-green-100 text-green-800" : "bg-gray-100"}`}>{data.status}</span></dd>
          </div>
          <div>
            <dt className="text-gray-500">Buka di tab baru</dt>
            <dd>{data.buka_di_tab_baru ? "Ya" : "Tidak"}</dd>
          </div>
        </dl>
      </div>
      <ConfirmDialog isOpen={deleteOpen} onClose={() => setDeleteOpen(false)} onConfirm={handleDelete} title="Hapus Tautan" message="Apakah Anda yakin?" confirmText="Ya, Hapus" type="danger" />
    </div>
  );
}
