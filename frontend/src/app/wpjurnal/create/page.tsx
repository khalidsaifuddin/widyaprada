"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

const currentYear = new Date().getFullYear();

export default function CreateJurnalPage() {
  const router = useRouter();
  const [title, setTitle] = useState("");
  const [author, setAuthor] = useState("");
  const [abstract, setAbstract] = useState("");
  const [content, setContent] = useState("");
  const [category, setCategory] = useState("");
  const [year, setYear] = useState(currentYear);
  const [pdfFile, setPdfFile] = useState<File | null>(null);
  const [pdfUrlDirect, setPdfUrlDirect] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!title.trim()) {
      setError("Judul wajib diisi");
      return;
    }
    setLoading(true);

    let pdfUrl = "";
    if (pdfFile) {
      const formData = new FormData();
      formData.append("file", pdfFile);
      const uploadRes = await apiService.uploadFile<{ url?: string }>("v1/wpjurnal/upload-pdf", formData);
      if (!uploadRes.success || !uploadRes.data?.url) {
        setError(uploadRes.message ?? "Gagal mengunggah file PDF");
        setLoading(false);
        return;
      }
      pdfUrl = uploadRes.data.url;
    } else if (pdfUrlDirect.trim()) {
      pdfUrl = pdfUrlDirect.trim();
    }

    const res = await apiService.post<{ id?: string }>("v1/wpjurnal", {
      title: title.trim(),
      author: author.trim(),
      abstract: abstract.trim(),
      content: content.trim(),
      category: category.trim(),
      year,
      pdf_url: pdfUrl,
    });

    if (res.success && res.data?.id) {
      router.push(`/wpjurnal/${res.data.id}/edit`);
    } else {
      setError(res.message ?? "Gagal menyimpan jurnal");
    }
    setLoading(false);
  };

  return (
    <div className="max-w-2xl space-y-6">
      <Link href="/wpjurnal" className="text-gray-600 hover:text-gray-900">
        ← Kembali ke Manajemen Jurnal
      </Link>
      <h1 className="text-2xl font-bold text-gray-900">Buat Jurnal Baru</h1>

      <form onSubmit={handleSubmit} className="space-y-6">
        {error && (
          <div className="p-4 rounded-lg bg-red-50 text-red-700">{error}</div>
        )}

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Judul *</label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="Judul jurnal"
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Penulis</label>
          <input
            type="text"
            value={author}
            onChange={(e) => setAuthor(e.target.value)}
            placeholder="Nama penulis"
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Abstrak</label>
          <textarea
            value={abstract}
            onChange={(e) => setAbstract(e.target.value)}
            rows={4}
            placeholder="Ringkasan jurnal"
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">File PDF</label>
          <p className="text-sm text-gray-500 mb-2">
            Unggah file (simpan ke folder server) atau masukkan URL langsung
          </p>
          <div className="space-y-3">
            <div>
              <span className="text-xs font-medium text-gray-500 block mb-1">Unggah file (path relatif ke backend)</span>
              <input
                type="file"
                accept=".pdf,application/pdf"
                onChange={(e) => {
                  setPdfFile(e.target.files?.[0] ?? null);
                  if (e.target.files?.[0]) setPdfUrlDirect("");
                }}
                className="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded file:border-0 file:bg-blue-50 file:text-blue-700"
              />
              {pdfFile && <p className="mt-1 text-sm text-green-600">Dipilih: {pdfFile.name}</p>}
            </div>
            <div>
              <span className="text-xs font-medium text-gray-500 block mb-1">Atau masukkan URL langsung</span>
              <input
                type="url"
                value={pdfUrlDirect}
                onChange={(e) => {
                  setPdfUrlDirect(e.target.value);
                  if (e.target.value) setPdfFile(null);
                }}
                placeholder="https://example.com/jurnal.pdf"
                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
              />
            </div>
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Kategori</label>
          <input
            type="text"
            value={category}
            onChange={(e) => setCategory(e.target.value)}
            placeholder="Contoh: Pendidikan Dasar"
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Tahun</label>
          <input
            type="number"
            value={year}
            onChange={(e) => setYear(parseInt(e.target.value, 10) || currentYear)}
            min={2000}
            max={currentYear + 1}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Konten (opsional)</label>
          <textarea
            value={content}
            onChange={(e) => setContent(e.target.value)}
            rows={6}
            placeholder="Isi tambahan atau catatan"
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          />
        </div>

        <div className="flex gap-3">
          <button
            type="submit"
            disabled={loading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
          >
            {loading ? "Menyimpan..." : "Simpan sebagai Draft"}
          </button>
          <Link
            href="/wpjurnal"
            className="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
