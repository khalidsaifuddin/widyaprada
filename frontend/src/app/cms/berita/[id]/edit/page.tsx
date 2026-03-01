"use client";

import BeritaImageSlider from "@/components/molecules/BeritaImageSlider";
import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function BeritaCMSEditPage() {
  const params = useParams();
  const router = useRouter();
  const id = params?.id as string;
  const [judul, setJudul] = useState("");
  const [slug, setSlug] = useState("");
  const [konten, setKonten] = useState("");
  const [ringkasan, setRingkasan] = useState("");
  const [thumbnail, setThumbnail] = useState("");
  const [tanggalPublikasi, setTanggalPublikasi] = useState("");
  const [status, setStatus] = useState("Published");
  const [penulis, setPenulis] = useState("");
  const [kategori, setKategori] = useState("");
  const [galleryUrls, setGalleryUrls] = useState<string[]>([]);
  const [loading, setLoading] = useState(false);
  const [loadDetail, setLoadDetail] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!id) return;
    apiService.get<Record<string, unknown>>(`v1/cms/berita/${id}`).then((res) => {
      if (res.success && res.data) {
        const d = res.data as Record<string, unknown>;
        setJudul((d.title as string) ?? "");
        setSlug((d.slug as string) ?? "");
        setKonten((d.content as string) ?? "");
        setRingkasan((d.excerpt as string) ?? "");
        setThumbnail((d.thumbnail_url as string) ?? "");
        setTanggalPublikasi((d.published_at as string)?.slice(0, 16) ?? "");
        setStatus((d.status as string) ?? "Published");
        setPenulis((d.author_name as string) ?? "");
        setKategori((d.category as string) ?? "");
        const urls = d.gallery_urls as string[] | undefined;
        setGalleryUrls(Array.isArray(urls) ? urls : []);
      }
      setLoadDetail(false);
    });
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!judul.trim()) { setError("Judul wajib diisi"); return; }
    setLoading(true);
    const res = await apiService.put(`v1/cms/berita/${id}`, {
      judul: judul.trim(),
      slug: slug.trim(),
      konten: konten.trim(),
      ringkasan: ringkasan.trim(),
      thumbnail: thumbnail.trim() || undefined,
      gallery_urls: galleryUrls.filter((u) => u.trim()),
      tanggal_publikasi: tanggalPublikasi || undefined,
      status,
      penulis: penulis.trim() || undefined,
      kategori: kategori.trim() || undefined,
    });
    setLoading(false);
    if (res.success) router.push(`/cms/berita/${id}`);
    else setError(res.message ?? "Gagal menyimpan");
  };

  if (loadDetail) return <div className="animate-pulse h-32 bg-gray-100 rounded" />;

  return (
    <div className="space-y-6">
      <Link href={`/cms/berita/${id}`} className="text-blue-600 hover:underline text-sm">← Kembali</Link>
      <h1 className="text-2xl font-bold text-gray-900">Edit Berita</h1>
      <form onSubmit={handleSubmit} className="max-w-2xl space-y-4">
        {error && <p className="text-red-600 text-sm">{error}</p>}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Judul *</label>
          <input type="text" value={judul} onChange={(e) => setJudul(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" required />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Slug</label>
          <input type="text" value={slug} onChange={(e) => setSlug(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Ringkasan</label>
          <textarea value={ringkasan} onChange={(e) => setRingkasan(e.target.value)} rows={2} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Konten</label>
          <textarea value={konten} onChange={(e) => setKonten(e.target.value)} rows={10} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm font-mono" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">URL Thumbnail</label>
          <input type="url" value={thumbnail} onChange={(e) => setThumbnail(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" placeholder="https://..." />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Galeri Gambar (multiple)</label>
          <p className="text-xs text-gray-500 mb-2">URL gambar ditampilkan sebagai slider di halaman detail berita</p>
          {galleryUrls.map((url, i) => (
            <div key={i} className="flex gap-2 mb-2">
              <input
                type="url"
                value={url}
                onChange={(e) => {
                  const next = [...galleryUrls];
                  next[i] = e.target.value;
                  setGalleryUrls(next);
                }}
                className="flex-1 rounded-lg border border-gray-300 px-3 py-2 text-sm"
                placeholder={`URL gambar ${i + 1}`}
              />
              <button
                type="button"
                onClick={() => setGalleryUrls(galleryUrls.filter((_, j) => j !== i))}
                className="rounded-lg border border-red-200 px-3 py-2 text-sm text-red-600 hover:bg-red-50"
              >
                Hapus
              </button>
            </div>
          ))}
          <button
            type="button"
            onClick={() => setGalleryUrls([...galleryUrls, ""])}
            className="rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50"
          >
            + Tambah Gambar
          </button>
          {(() => {
            const imgs: string[] = [];
            if (thumbnail?.trim()) imgs.push(thumbnail);
            galleryUrls.filter((u) => u?.trim()).forEach((u) => imgs.push(u));
            return imgs.length > 0 ? (
              <div className="mt-4">
                <p className="text-xs text-gray-500 mb-2">Preview slider</p>
                <BeritaImageSlider images={imgs} title={judul || "Berita"} className="max-w-2xl" />
              </div>
            ) : null;
          })()}
        </div>
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Tanggal Publikasi</label>
            <input type="datetime-local" value={tanggalPublikasi} onChange={(e) => setTanggalPublikasi(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Status</label>
            <select value={status} onChange={(e) => setStatus(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm">
              <option value="Draft">Draft</option>
              <option value="Published">Published</option>
            </select>
          </div>
        </div>
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Penulis</label>
            <input type="text" value={penulis} onChange={(e) => setPenulis(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Kategori</label>
            <input type="text" value={kategori} onChange={(e) => setKategori(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
          </div>
        </div>
        <div className="flex gap-3">
          <button type="submit" disabled={loading} className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700 disabled:opacity-50">{loading ? "Menyimpan..." : "Simpan"}</button>
          <Link href={`/cms/berita/${id}`} className="rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">Batal</Link>
        </div>
      </form>
    </div>
  );
}
