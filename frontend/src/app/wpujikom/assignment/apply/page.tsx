"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useRouter, useCallback, useEffect, useState } from "react";

interface DokumenItem {
  id: string;
  kode: string;
  nama: string;
  urutan: number;
  tipe_input: string;
  batasan?: string;
  untuk_jenis_ujikom?: string;
}

interface DokumenListResponse {
  items: DokumenItem[];
}

function fileToBase64(file: File): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => {
      const result = reader.result as string;
      resolve(result);
    };
    reader.onerror = reject;
    reader.readAsDataURL(file);
  });
}

export default function ApplyUjikomPage() {
  const router = useRouter();
  const [jenisUjikom, setJenisUjikom] = useState("");
  const [documents, setDocuments] = useState<DokumenItem[]>([]);
  const [formData, setFormData] = useState<Record<string, { file?: File; text?: string }>>({});
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);
  const [existingStatus, setExistingStatus] = useState<string | null>(null);

  const fetchDokumen = useCallback(async () => {
    if (!jenisUjikom) return;
    const res = await apiService.get<DokumenListResponse>(
      "v1/ujikom/dokumen-persyaratan",
      { jenis_ujikom: jenisUjikom }
    );
    if (res.success && res.data) {
      const d = res.data as DokumenListResponse;
      setDocuments(d.items ?? []);
      setFormData({});
    }
  }, [jenisUjikom]);

  useEffect(() => {
    fetchDokumen();
  }, [fetchDokumen]);

  useEffect(() => {
    if (!jenisUjikom) return;
    apiService.get<{ status_kode?: string }>("v1/ujikom/apply/status").then((res) => {
      if (res.success && res.data) {
        const d = res.data as { status_kode?: string; jenis_ujikom?: string };
        if (d.jenis_ujikom === jenisUjikom && d.status_kode) {
          setExistingStatus(d.status_kode);
        }
      }
    });
  }, [jenisUjikom]);

  const handleFileChange = (kode: string, file: File | null) => {
    setFormData((prev) => ({
      ...prev,
      [kode]: { ...prev[kode], file: file ?? undefined },
    }));
  };

  const handleTextChange = (kode: string, text: string) => {
    setFormData((prev) => ({
      ...prev,
      [kode]: { ...prev[kode], text },
    }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!jenisUjikom) {
      setError("Pilih jenis ujikom");
      return;
    }
    if (existingStatus) {
      setError("Anda sudah pernah apply untuk jenis ujikom ini.");
      return;
    }
    setError("");
    setLoading(true);

    const docs: { document_type: string; file_path?: string; portofolio_text?: string }[] = [];

    for (const doc of documents) {
      const data = formData[doc.kode];
      if (doc.tipe_input === "file") {
        if (!data?.file) {
          setError(`${doc.nama} wajib diisi (unggah file)`);
          setLoading(false);
          return;
        }
        try {
          const base64 = await fileToBase64(data.file);
          docs.push({ document_type: doc.kode, file_path: base64 });
        } catch {
          setError(`Gagal membaca file ${doc.nama}`);
          setLoading(false);
          return;
        }
      } else {
        if (!data?.text?.trim()) {
          setError(`${doc.nama} wajib diisi`);
          setLoading(false);
          return;
        }
        docs.push({ document_type: doc.kode, portofolio_text: data.text.trim() });
      }
    }

    const res = await apiService.post("v1/ujikom/apply", {
      jenis_ujikom: jenisUjikom,
      documents: docs,
    });

    if (res.success) {
      setSuccess(true);
    } else {
      setError(res.message ?? "Gagal mengirim pendaftaran");
    }
    setLoading(false);
  };

  if (success) {
    return (
      <div className="max-w-2xl mx-auto space-y-6">
        <div className="bg-white rounded-xl shadow-sm border border-green-200 p-8 text-center">
          <h1 className="text-xl font-bold text-green-800">Pendaftaran berhasil dikirim</h1>
          <p className="mt-2 text-gray-600">Status: Menunggu verifikasi.</p>
          <Link
            href="/wpujikom/assignment"
            className="mt-6 inline-block px-4 py-2 rounded-lg bg-blue-600 text-white transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700"
          >
            Kembali ke Tugas Saya
          </Link>
        </div>
      </div>
    );
  }

  return (
    <div className="max-w-2xl space-y-6">
      <Link href="/wpujikom/assignment" className="text-gray-600 hover:text-gray-900">
        ← Kembali
      </Link>
      <h1 className="text-2xl font-bold text-gray-900">Pendaftaran Uji Kompetensi</h1>

      <form onSubmit={handleSubmit} className="space-y-6">
        {error && (
          <div className="p-4 rounded-lg bg-red-50 text-red-700">{error}</div>
        )}

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Jenis Ujikom *</label>
          <select
            value={jenisUjikom}
            onChange={(e) => setJenisUjikom(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          >
            <option value="">-- Pilih --</option>
            <option value="perpindahan_jabatan">Perpindahan Jabatan</option>
            <option value="kenaikan_tingkat">Kenaikan Tingkat</option>
          </select>
        </div>

        {existingStatus && (
          <div className="p-4 rounded-lg bg-amber-50 text-amber-800">
            Anda sudah pernah mendaftar untuk jenis ini. Status: {existingStatus}
          </div>
        )}

        {documents.length > 0 && (
          <div className="space-y-4">
            <h2 className="font-semibold text-gray-900">Dokumen Persyaratan</h2>
            {documents.map((doc) => (
              <div key={doc.id} className="border border-gray-200 rounded-lg p-4">
                <label className="block text-sm font-medium text-gray-700 mb-1">
                  {doc.nama} {doc.batasan && `(${doc.batasan})`}
                </label>
                {doc.tipe_input === "file" ? (
                  <input
                    type="file"
                    onChange={(e) => handleFileChange(doc.kode, e.target.files?.[0] ?? null)}
                    className="block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded file:border-0 file:bg-blue-50 file:text-blue-700"
                  />
                ) : (
                  <textarea
                    value={formData[doc.kode]?.text ?? ""}
                    onChange={(e) => handleTextChange(doc.kode, e.target.value)}
                    rows={4}
                    placeholder={doc.batasan || ""}
                    className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
                  />
                )}
              </div>
            ))}
          </div>
        )}

        <div className="flex gap-3">
          <button
            type="submit"
            disabled={loading || !jenisUjikom || existingStatus !== null}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 disabled:opacity-50"
          >
            {loading ? "Mengirim..." : "Kirim Pendaftaran"}
          </button>
          <Link
            href="/wpujikom/assignment"
            className="px-4 py-2 border border-gray-300 rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
