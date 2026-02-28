"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import Image from "next/image";
import { useEffect, useState } from "react";

interface BeritaDetailData {
  id: string;
  title: string;
  slug: string;
  content: string;
  excerpt?: string;
  thumbnail_url?: string;
  published_at?: string;
  author_name?: string;
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

export default function BeritaDetail({ slug }: { slug: string }) {
  const [data, setData] = useState<BeritaDetailData | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!slug) {
      setLoading(false);
      return;
    }
    apiService
      .get<BeritaDetailData>(`v1/berita/${encodeURIComponent(slug)}`)
      .then((res) => {
        if (res.success && res.data) {
          setData(res.data as BeritaDetailData);
        } else {
          setError(res.message ?? "Berita tidak ditemukan");
        }
      })
      .catch(() => setError("Gagal memuat berita"))
      .finally(() => setLoading(false));
  }, [slug]);

  if (loading) {
    return (
      <div className="py-12">
        <div className="mx-auto max-w-3xl px-4 sm:px-6 lg:px-8">
          <div className="h-8 w-48 bg-gray-100 rounded animate-pulse mb-4" />
          <div className="h-4 w-full bg-gray-100 rounded animate-pulse mb-2" />
          <div className="h-64 bg-gray-100 rounded animate-pulse" />
        </div>
      </div>
    );
  }

  if (error || !data) {
    return (
      <div className="py-12">
        <div className="mx-auto max-w-3xl px-4 sm:px-6 lg:px-8">
          <p className="text-red-600">{error || "Berita tidak ditemukan"}</p>
          <Link href="/berita" className="mt-4 inline-block text-blue-600 hover:underline">
            ← Kembali ke Daftar Berita
          </Link>
        </div>
      </div>
    );
  }

  return (
    <article className="py-12">
      <div className="mx-auto max-w-3xl px-4 sm:px-6 lg:px-8">
        <Link href="/berita" className="text-sm text-blue-600 hover:underline mb-6 inline-block">
          ← Kembali ke Daftar Berita
        </Link>
        <h1 className="text-2xl md:text-3xl font-bold text-gray-900 mb-4">{data.title}</h1>
        <div className="flex flex-wrap gap-4 text-sm text-gray-500 mb-6">
          {data.published_at && <span>{formatDate(data.published_at)}</span>}
          {data.author_name && <span>Oleh: {data.author_name}</span>}
          {data.category && <span>Kategori: {data.category}</span>}
        </div>
        {data.thumbnail_url && (
          <div className="relative h-64 md:h-96 rounded-xl overflow-hidden mb-8 bg-gray-100">
            <Image
              src={data.thumbnail_url}
              alt={data.title}
              fill
              className="object-cover"
              sizes="(max-width: 768px) 100vw, 672px"
              unoptimized={data.thumbnail_url.startsWith("http")}
            />
          </div>
        )}
        <div
          className="text-gray-700 leading-relaxed [&_p]:mb-4 [&_ul]:list-disc [&_ul]:pl-6 [&_ol]:list-decimal [&_ol]:pl-6 [&_h2]:text-xl [&_h2]:font-semibold [&_h2]:mt-6 [&_img]:rounded-lg [&_img]:max-w-full"
          dangerouslySetInnerHTML={{ __html: data.content || "" }}
        />
      </div>
    </article>
  );
}
