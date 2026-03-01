"use client";

import BeritaImageSlider from "@/components/molecules/BeritaImageSlider";
import { apiService } from "@/lib/api";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

function slugify(s: string): string {
  return s
    .toLowerCase()
    .replace(/\s+/g, "-")
    .replace(/[^a-z0-9-]/g, "");
}

export default function BeritaCMSCreatePage() {
  const router = useRouter();
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
  const [error, setError] = useState("");

  const handleJudulChange = (v: string) => {
    setJudul(v);
    if (!slug || slug === slugify(judul)) setSlug(slugify(v));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!judul.trim()) { setError("Judul wajib diisi"); return; }
    setLoading(true);
    const res = await apiService.post("v1/cms/berita", {
      judul: judul.trim(),
      slug: slug.trim() || slugify(judul),
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
    if (res.success && res.data) {
      const d = res.data as { id?: string };
      router.push(`/cms/berita/${d.id ?? ""}`);
    } else {
      setError(res.message ?? "Gagal membuat berita");
    }
  };

  return (
    <div className="space-y-6">
      <Link href="/cms/berita" className="text-blue-600 hover:underline text-sm">← Kembali</Link>
      <h1 className="text-2xl font-bold text-gray-900">Tambah Berita</h1>
      <form onSubmit={handleSubmit} className="max-w-2xl space-y-4">
        {error && <p className="text-red-600 text-sm">{error}</p>}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Judul *</label>
          <input type="text" value={judul} onChange={(e) => handleJudulChange(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" required />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Slug</label>
          <input type="text" value={slug} onChange={(e) => setSlug(e.target.value)} placeholder="url-friendly" className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Ringkasan</label>
          <textarea value={ringkasan} onChange={(e) => setRingkasan(e.target.value)} rows={2} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Konten (HTML)</label>
          <textarea value={konten} onChange={(e) => setKonten(e.target.value)} rows={10} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm font-mono" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">URL Thumbnail</label>
          <input type="url" value={thumbnail} onChange={(e) => setThumbnail(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" placeholder="https://..." />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Galeri Gambar (multiple)</label>
          <p className="text-xs text-gray-500 mb-2">Tambahkan URL gambar untuk ditampilkan sebagai slider di halaman detail berita</p>
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
          <Link href="/cms/berita" className="rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">Batal</Link>
        </div>
      </form>
    </div>
  );
}
