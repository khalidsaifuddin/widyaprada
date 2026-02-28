"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function SliderEditPage() {
  const params = useParams();
  const router = useRouter();
  const id = params?.id as string;
  const [imageUrl, setImageUrl] = useState("");
  const [title, setTitle] = useState("");
  const [subtitle, setSubtitle] = useState("");
  const [linkUrl, setLinkUrl] = useState("");
  const [ctaLabel, setCtaLabel] = useState("");
  const [sortOrder, setSortOrder] = useState(0);
  const [status, setStatus] = useState("Published");
  const [tanggalMulai, setTanggalMulai] = useState("");
  const [tanggalSelesai, setTanggalSelesai] = useState("");
  const [loading, setLoading] = useState(false);
  const [loadDetail, setLoadDetail] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!id) return;
    apiService.get<Record<string, unknown>>(`v1/cms/slider/${id}`).then((res) => {
      if (res.success && res.data) {
        const d = res.data as Record<string, unknown>;
        setImageUrl((d.image_url as string) ?? "");
        setTitle((d.title as string) ?? "");
        setSubtitle((d.subtitle as string) ?? "");
        setLinkUrl((d.link_url as string) ?? "");
        setCtaLabel((d.cta_label as string) ?? "");
        setSortOrder((d.sort_order as number) ?? 0);
        setStatus((d.status as string) ?? "Published");
        setTanggalMulai((d.date_start as string)?.slice(0, 16) ?? "");
        setTanggalSelesai((d.date_end as string)?.slice(0, 16) ?? "");
      }
      setLoadDetail(false);
    });
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!imageUrl.trim()) {
      setError("URL gambar wajib diisi");
      return;
    }
    setLoading(true);
    const res = await apiService.put(`v1/cms/slider/${id}`, {
      image_url: imageUrl.trim(),
      title: title.trim(),
      subtitle: subtitle.trim(),
      link_url: linkUrl.trim() || undefined,
      cta_label: ctaLabel.trim() || undefined,
      sort_order: sortOrder,
      status,
      tanggal_mulai_tampil: tanggalMulai || undefined,
      tanggal_selesai_tampil: tanggalSelesai || undefined,
    });
    setLoading(false);
    if (res.success) {
      router.push(`/cms/slider/${id}`);
    } else {
      setError(res.message ?? "Gagal menyimpan");
    }
  };

  if (loadDetail) return <div className="animate-pulse h-32 bg-gray-100 rounded" />;

  return (
    <div className="space-y-6">
      <Link href={`/cms/slider/${id}`} className="text-blue-600 hover:underline text-sm inline-block">
        ← Kembali
      </Link>
      <h1 className="text-2xl font-bold text-gray-900">Edit Slide</h1>
      <form onSubmit={handleSubmit} className="max-w-xl space-y-4">
        {error && <p className="text-red-600 text-sm">{error}</p>}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">URL Gambar *</label>
          <input type="url" value={imageUrl} onChange={(e) => setImageUrl(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" required />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Judul</label>
          <input type="text" value={title} onChange={(e) => setTitle(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Subjudul</label>
          <input type="text" value={subtitle} onChange={(e) => setSubtitle(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">URL Link</label>
          <input type="url" value={linkUrl} onChange={(e) => setLinkUrl(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Label CTA</label>
          <input type="text" value={ctaLabel} onChange={(e) => setCtaLabel(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Urutan</label>
          <input type="number" value={sortOrder} onChange={(e) => setSortOrder(parseInt(e.target.value, 10) || 0)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Status</label>
          <select value={status} onChange={(e) => setStatus(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm">
            <option value="Draft">Draft</option>
            <option value="Published">Published</option>
          </select>
        </div>
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Tanggal Mulai</label>
            <input type="datetime-local" value={tanggalMulai} onChange={(e) => setTanggalMulai(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Tanggal Selesai</label>
            <input type="datetime-local" value={tanggalSelesai} onChange={(e) => setTanggalSelesai(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
          </div>
        </div>
        <div className="flex gap-3">
          <button type="submit" disabled={loading} className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700 disabled:opacity-50">
            {loading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link href={`/cms/slider/${id}`} className="rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
