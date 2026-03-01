"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";

interface JurnalData {
  id: string;
  title: string;
  author: string;
  abstract: string;
  content: string;
  pdf_url?: string;
  category: string;
  year: number;
  status: string;
  created_at?: string;
  updated_at?: string;
}

function formatDate(d: string): string {
  try {
    return new Date(d).toLocaleDateString("id-ID", {
      day: "numeric",
      month: "long",
      year: "numeric",
      timeZone: "Asia/Jakarta",
    });
  } catch {
    return d;
  }
}

export default function WPJurnalDetailPage() {
  const params = useParams();
  const id = params?.id as string;
  const [data, setData] = useState<JurnalData | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!id) {
      setLoading(false);
      return;
    }
    apiService
      .get<JurnalData>(`v1/wpjurnal/${id}`)
      .then((res) => {
        if (res.success && res.data) {
          setData(res.data as JurnalData);
        } else {
          setError(res.message ?? "Jurnal tidak ditemukan");
        }
      })
      .catch(() => setError("Gagal memuat jurnal"))
      .finally(() => setLoading(false));
  }, [id]);

  if (loading) {
    return (
      <div className="max-w-3xl space-y-6">
        <div className="h-8 w-48 bg-gray-100 rounded animate-pulse" />
        <div className="h-32 bg-gray-100 rounded animate-pulse" />
      </div>
    );
  }

  if (error || !data) {
    return (
      <div className="max-w-3xl space-y-6">
        <Link href="/wpjurnal" className="text-gray-600 hover:text-gray-900">← Kembali</Link>
        <p className="text-red-600">{error || "Jurnal tidak ditemukan"}</p>
      </div>
    );
  }

  const pdfFullUrl = data.pdf_url?.startsWith("http")
    ? data.pdf_url
    : data.pdf_url
      ? ((process.env.NEXT_PUBLIC_API_BASE_URL?.replace(/\/api\/?$/, "") ?? "http://localhost:8080") + data.pdf_url)
      : "";

  return (
    <div className="max-w-3xl space-y-6">
      <Link href="/wpjurnal" className="text-gray-600 hover:text-gray-900">
        ← Kembali ke Manajemen Jurnal
      </Link>

      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <h1 className="text-2xl font-bold text-gray-900">{data.title}</h1>
        <Link
          href={`/wpjurnal/${id}/edit`}
          className="inline-flex px-4 py-2 rounded-lg bg-blue-600 text-white transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700"
        >
          Edit
        </Link>
      </div>

      <div className="flex flex-wrap gap-4 text-sm text-gray-500">
        {data.author && <span>Penulis: {data.author}</span>}
        {data.year && <span>Tahun: {data.year}</span>}
        {data.category && <span>Kategori: {data.category}</span>}
        <span
          className={`px-2 py-0.5 rounded text-xs font-medium ${
            data.status === "Published" ? "bg-green-100 text-green-800" : "bg-gray-100 text-gray-700"
          }`}
        >
          {data.status}
        </span>
      </div>

      {data.abstract && (
        <div className="rounded-lg bg-gray-50 p-4">
          <h3 className="font-semibold text-gray-900 mb-2">Abstrak</h3>
          <p className="text-gray-700 leading-relaxed">{data.abstract}</p>
        </div>
      )}

      {data.pdf_url && (
        <div>
          <a
            href={pdfFullUrl}
            target="_blank"
            rel="noopener noreferrer"
            className="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-red-600 text-white transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-red-700 font-medium"
          >
            <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 20 20">
              <path fillRule="evenodd" d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4zm2 6a1 1 0 011-1h6a1 1 0 110 2H7a1 1 0 01-1-1zm1 3a1 1 0 100 2h6a1 1 0 100-2H7z" clipRule="evenodd" />
            </svg>
            Unduh PDF
          </a>
        </div>
      )}

      {data.content && (
        <div
          className="text-gray-700 leading-relaxed [&_p]:mb-4"
          dangerouslySetInnerHTML={{ __html: data.content }}
        />
      )}

      {(data.created_at || data.updated_at) && (
        <p className="text-sm text-gray-500">
          {data.created_at && `Dibuat: ${formatDate(data.created_at)}`}
          {data.created_at && data.updated_at && " • "}
          {data.updated_at && `Diperbarui: ${formatDate(data.updated_at)}`}
        </p>
      )}
    </div>
  );
}
