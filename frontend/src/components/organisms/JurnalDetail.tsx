"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useEffect, useState } from "react";

interface JurnalDetailData {
  id: string;
  title: string;
  author: string;
  abstract?: string;
  content?: string;
  published_at?: string;
  year?: number;
  category?: string;
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

export default function JurnalDetail({ id }: { id: string }) {
  const [data, setData] = useState<JurnalDetailData | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!id) {
      setLoading(false);
      return;
    }
    apiService
      .get<JurnalDetailData>(`v1/jurnal/${id}`)
      .then((res) => {
        if (res.success && res.data) {
          setData(res.data as JurnalDetailData);
        } else {
          setError(res.message ?? "Jurnal tidak ditemukan");
        }
      })
      .catch(() => setError("Gagal memuat jurnal"))
      .finally(() => setLoading(false));
  }, [id]);

  if (loading) {
    return (
      <div className="py-12">
        <div className="mx-auto max-w-3xl px-4 sm:px-6 lg:px-8">
          <div className="h-8 w-48 bg-gray-100 rounded animate-pulse mb-4" />
          <div className="h-4 w-full bg-gray-100 rounded animate-pulse mb-2" />
          <div className="h-32 bg-gray-100 rounded animate-pulse" />
        </div>
      </div>
    );
  }

  if (error || !data) {
    return (
      <div className="py-12">
        <div className="mx-auto max-w-3xl px-4 sm:px-6 lg:px-8">
          <p className="text-red-600">{error || "Jurnal tidak ditemukan"}</p>
          <Link href="/jurnal" className="mt-4 inline-block text-blue-600 hover:underline">
            ← Kembali ke Daftar Jurnal
          </Link>
        </div>
      </div>
    );
  }

  return (
    <article className="py-12">
      <div className="mx-auto max-w-3xl px-4 sm:px-6 lg:px-8">
        <Link href="/jurnal" className="text-sm text-blue-600 hover:underline mb-6 inline-block">
          ← Kembali ke Daftar Jurnal
        </Link>
        <h1 className="text-2xl md:text-3xl font-bold text-gray-900 mb-4">{data.title}</h1>
        <div className="flex flex-wrap gap-4 text-sm text-gray-500 mb-6">
          {data.author && <span>Penulis: {data.author}</span>}
          {data.published_at && <span>Publikasi: {formatDate(data.published_at)}</span>}
          {data.year && <span>Tahun: {data.year}</span>}
          {data.category && <span>Kategori: {data.category}</span>}
        </div>
        {data.abstract && (
          <div className="rounded-lg bg-gray-50 p-4 mb-8">
            <h3 className="font-semibold text-gray-900 mb-2">Abstrak</h3>
            <p className="text-gray-700 leading-relaxed">{data.abstract}</p>
          </div>
        )}
        {data.content && (
          <div
            className="text-gray-700 leading-relaxed [&_p]:mb-4 [&_ul]:list-disc [&_ul]:pl-6 [&_ol]:list-decimal [&_ol]:pl-6 [&_h2]:text-xl [&_h2]:font-semibold [&_h2]:mt-6"
            dangerouslySetInnerHTML={{ __html: data.content }}
          />
        )}
      </div>
    </article>
  );
}
