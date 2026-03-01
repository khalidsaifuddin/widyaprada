"use client";

import BeritaImageSlider from "@/components/molecules/BeritaImageSlider";
import { apiService } from "@/lib/api";
import { ConfirmDialog } from "@/components";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface ArticleDetail {
  id: string;
  title: string;
  slug: string;
  content: string;
  excerpt: string;
  thumbnail_url: string;
  gallery_urls?: string[];
  published_at: string;
  status: string;
  author_name: string;
  category: string;
}

export default function BeritaCMSDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params?.id as string;
  const [data, setData] = useState<ArticleDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [deleteOpen, setDeleteOpen] = useState(false);

  useEffect(() => {
    if (!id) return;
    apiService.get<ArticleDetail>(`v1/cms/berita/${id}`).then((res) => {
      if (res.success && res.data) setData(res.data as ArticleDetail);
      setLoading(false);
    });
  }, [id]);

  const handleDelete = async () => {
    const res = await apiService.delete(`v1/cms/berita/${id}`);
    setDeleteOpen(false);
    if (res.success) router.push("/cms/berita");
  };

  if (loading) return <div className="animate-pulse h-32 bg-gray-100 rounded" />;
  if (!data) return <p className="text-red-600">Berita tidak ditemukan</p>;

  const allImages: string[] = [];
  if (data.thumbnail_url?.trim()) allImages.push(data.thumbnail_url);
  (data.gallery_urls ?? []).filter((u) => u?.trim()).forEach((u) => allImages.push(u));

  return (
    <div className="space-y-6">
      <div className="flex justify-between">
        <Link href="/cms/berita" className="text-blue-600 hover:underline text-sm">← Kembali</Link>
        <div className="flex gap-2">
          <Link href={`/cms/berita/${id}/edit`} className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700">Edit</Link>
          <button type="button" onClick={() => setDeleteOpen(true)} className="rounded-lg border border-red-600 text-red-600 px-4 py-2 text-sm font-medium hover:bg-red-50">Hapus</button>
        </div>
      </div>
      <div className="bg-white rounded-xl border border-gray-200 p-6">
        <h1 className="text-xl font-bold text-gray-900 mb-4">{data.title}</h1>
        <div className="flex flex-wrap gap-4 text-sm text-gray-500 mb-4">
          {data.published_at && <span>{new Date(data.published_at).toLocaleDateString("id-ID")}</span>}
          {data.author_name && <span>Oleh: {data.author_name}</span>}
          {data.category && <span>Kategori: {data.category}</span>}
          <span className={`px-2 py-0.5 rounded text-xs ${data.status === "Published" ? "bg-green-100 text-green-800" : "bg-gray-100"}`}>{data.status}</span>
        </div>
        {allImages.length > 0 && <BeritaImageSlider images={allImages} title={data.title} className="mb-6" />}
        {data.excerpt && <p className="text-gray-600 mb-4">{data.excerpt}</p>}
        <div className="text-gray-700 text-sm [&_p]:mb-2" dangerouslySetInnerHTML={{ __html: data.content || "" }} />
      </div>
      <ConfirmDialog isOpen={deleteOpen} onClose={() => setDeleteOpen(false)} onConfirm={handleDelete} title="Hapus Berita" message="Apakah Anda yakin?" confirmText="Ya, Hapus" type="danger" />
    </div>
  );
}
