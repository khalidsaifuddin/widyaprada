"use client";

import Link from "next/link";
import { useSearchParams } from "next/navigation";

export default function CBTSelesaiPage() {
  const searchParams = useSearchParams();
  const error = searchParams.get("error") === "1";

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-8 text-center">
        <h1 className="text-xl font-bold text-gray-900">
          {error ? "Waktu habis" : "Ujian telah berhasil dikirim"}
        </h1>
        <p className="mt-2 text-gray-600">
          {error
            ? "Jawaban Anda telah disimpan secara otomatis."
            : "Nilai akan tampil di Riwayat Ujian setelah dinilai."}
        </p>
        <div className="mt-6 flex flex-wrap gap-3 justify-center">
          <Link
            href="/wpujikom/cbt"
            className="inline-flex px-4 py-2 rounded-lg bg-blue-600 text-white hover:bg-blue-700 font-medium"
          >
            Kembali ke Daftar Ujian
          </Link>
          <Link
            href="/wpujikom/cbt/riwayat"
            className="inline-flex px-4 py-2 rounded-lg border border-gray-300 text-gray-700 hover:bg-gray-50 font-medium"
          >
            Lihat Riwayat
          </Link>
        </div>
      </div>
    </div>
  );
}
