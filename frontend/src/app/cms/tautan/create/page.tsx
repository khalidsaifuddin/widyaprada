"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function TautanCreatePage() {
  const router = useRouter();
  const [judul, setJudul] = useState("");
  const [url, setUrl] = useState("");
  const [deskripsi, setDeskripsi] = useState("");
  const [status, setStatus] = useState("Aktif");
  const [bukaDiTabBaru, setBukaDiTabBaru] = useState(true);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!judul.trim() || !url.trim()) {
      setError("Judul dan URL wajib diisi");
      return;
    }
    setLoading(true);
    const res = await apiService.post("v1/cms/tautan", {
      judul: judul.trim(),
      url: url.trim(),
      deskripsi: deskripsi.trim() || undefined,
      status,
      buka_di_tab_baru: bukaDiTabBaru,
    });
    setLoading(false);
    if (res.success && res.data) {
      const d = res.data as { id?: string };
      router.push(`/cms/tautan/${d.id ?? ""}`);
    } else {
      setError(res.message ?? "Gagal membuat tautan");
    }
  };

  return (
    <div className="space-y-6">
      <Link href="/cms/tautan" className="text-blue-600 hover:underline text-sm">← Kembali</Link>
      <h1 className="text-2xl font-bold text-gray-900">Tambah Tautan</h1>
      <form onSubmit={handleSubmit} className="max-w-xl space-y-4">
        {error && <p className="text-red-600 text-sm">{error}</p>}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Judul *</label>
          <input type="text" value={judul} onChange={(e) => setJudul(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" required />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">URL *</label>
          <input type="url" value={url} onChange={(e) => setUrl(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" required />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Deskripsi</label>
          <textarea value={deskripsi} onChange={(e) => setDeskripsi(e.target.value)} rows={2} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Status</label>
          <select value={status} onChange={(e) => setStatus(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm">
            <option value="Aktif">Aktif</option>
            <option value="Nonaktif">Nonaktif</option>
          </select>
        </div>
        <div className="flex items-center gap-2">
          <input type="checkbox" id="bukaTab" checked={bukaDiTabBaru} onChange={(e) => setBukaDiTabBaru(e.target.checked)} />
          <label htmlFor="bukaTab" className="text-sm text-gray-700">Buka di tab baru</label>
        </div>
        <div className="flex gap-3">
          <button type="submit" disabled={loading} className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700 disabled:opacity-50">{loading ? "Menyimpan..." : "Simpan"}</button>
          <Link href="/cms/tautan" className="rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">Batal</Link>
        </div>
      </form>
    </div>
  );
}
