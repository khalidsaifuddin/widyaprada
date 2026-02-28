"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function TautanEditPage() {
  const params = useParams();
  const router = useRouter();
  const id = params?.id as string;
  const [judul, setJudul] = useState("");
  const [url, setUrl] = useState("");
  const [deskripsi, setDeskripsi] = useState("");
  const [status, setStatus] = useState("Aktif");
  const [bukaDiTabBaru, setBukaDiTabBaru] = useState(true);
  const [loading, setLoading] = useState(false);
  const [loadDetail, setLoadDetail] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!id) return;
    apiService.get<Record<string, unknown>>(`v1/cms/tautan/${id}`).then((res) => {
      if (res.success && res.data) {
        const d = res.data as Record<string, unknown>;
        setJudul((d.title as string) ?? "");
        setUrl((d.url as string) ?? "");
        setDeskripsi((d.description as string) ?? "");
        setStatus((d.status as string) ?? "Aktif");
        setBukaDiTabBaru((d.buka_di_tab_baru as boolean) ?? true);
      }
      setLoadDetail(false);
    });
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!judul.trim() || !url.trim()) { setError("Judul dan URL wajib diisi"); return; }
    setLoading(true);
    const res = await apiService.put(`v1/cms/tautan/${id}`, {
      judul: judul.trim(),
      url: url.trim(),
      deskripsi: deskripsi.trim() || undefined,
      status,
      buka_di_tab_baru: bukaDiTabBaru,
    });
    setLoading(false);
    if (res.success) router.push(`/cms/tautan/${id}`);
    else setError(res.message ?? "Gagal menyimpan");
  };

  if (loadDetail) return <div className="animate-pulse h-32 bg-gray-100 rounded" />;

  return (
    <div className="space-y-6">
      <Link href={`/cms/tautan/${id}`} className="text-blue-600 hover:underline text-sm">← Kembali</Link>
      <h1 className="text-2xl font-bold text-gray-900">Edit Tautan</h1>
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
          <Link href={`/cms/tautan/${id}`} className="rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">Batal</Link>
        </div>
      </form>
    </div>
  );
}
