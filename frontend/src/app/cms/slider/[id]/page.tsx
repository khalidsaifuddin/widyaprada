"use client";

import { apiService } from "@/lib/api";
import { ConfirmDialog } from "@/components";
import Link from "next/link";
import Image from "next/image";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface SlideDetail {
  id: string;
  image_url: string;
  title: string;
  subtitle: string;
  link_url: string;
  cta_label: string;
  sort_order: number;
  status: string;
  date_start: string;
  date_end: string;
  created_at?: string;
}

export default function SliderDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params?.id as string;
  const [data, setData] = useState<SlideDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [deleteOpen, setDeleteOpen] = useState(false);

  useEffect(() => {
    if (!id) return;
    apiService.get<SlideDetail>(`v1/cms/slider/${id}`).then((res) => {
      if (res.success && res.data) setData(res.data as SlideDetail);
      setLoading(false);
    });
  }, [id]);

  const handleDelete = async () => {
    const res = await apiService.delete(`v1/cms/slider/${id}`);
    setDeleteOpen(false);
    if (res.success) router.push("/cms/slider");
  };

  if (loading) return <div className="animate-pulse h-32 bg-gray-100 rounded" />;
  if (!data) return <p className="text-red-600">Slide tidak ditemukan</p>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <Link href="/cms/slider" className="text-blue-600 hover:underline text-sm">
          ← Kembali
        </Link>
        <div className="flex gap-2">
          <Link
            href={`/cms/slider/${id}/edit`}
            className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700"
          >
            Edit
          </Link>
          <button
            type="button"
            onClick={() => setDeleteOpen(true)}
            className="rounded-lg border border-red-600 text-red-600 px-4 py-2 text-sm font-medium hover:bg-red-50"
          >
            Hapus
          </button>
        </div>
      </div>
      <div className="bg-white rounded-xl border border-gray-200 p-6">
        <h1 className="text-xl font-bold text-gray-900 mb-4">{data.title || "Tanpa judul"}</h1>
        {data.image_url && (
          <div className="relative h-48 rounded-lg overflow-hidden mb-4">
            <Image
              src={data.image_url}
              alt=""
              fill
              className="object-cover"
              unoptimized={data.image_url.startsWith("http")}
            />
          </div>
        )}
        <dl className="grid grid-cols-1 sm:grid-cols-2 gap-3 text-sm">
          <dt className="text-gray-500">Subjudul</dt>
          <dd>{data.subtitle || "-"}</dd>
          <dt className="text-gray-500">Link URL</dt>
          <dd>{data.link_url ? <a href={data.link_url} className="text-blue-600 hover:underline">{data.link_url}</a> : "-"}</dd>
          <dt className="text-gray-500">Label CTA</dt>
          <dd>{data.cta_label || "-"}</dd>
          <dt className="text-gray-500">Urutan</dt>
          <dd>{data.sort_order}</dd>
          <dt className="text-gray-500">Status</dt>
          <dd><span className={`px-2 py-0.5 rounded text-xs ${data.status === "Published" ? "bg-green-100 text-green-800" : "bg-gray-100"}`}>{data.status}</span></dd>
        </dl>
      </div>
      <ConfirmDialog
        isOpen={deleteOpen}
        onClose={() => setDeleteOpen(false)}
        onConfirm={handleDelete}
        title="Hapus Slide"
        message="Apakah Anda yakin ingin menghapus slide ini?"
        confirmText="Ya, Hapus"
        type="danger"
      />
    </div>
  );
}
