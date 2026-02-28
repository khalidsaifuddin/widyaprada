"use client";

import Image from "next/image";
import Link from "next/link";

export interface NewsCardProps {
  id: string;
  title: string;
  slug: string;
  excerpt?: string;
  thumbnail_url?: string;
  published_at?: string;
}

function formatDate(d: string): string {
  try {
    return new Date(d).toLocaleDateString("id-ID", {
      day: "numeric",
      month: "short",
      year: "numeric",
      timeZone: "Asia/Jakarta",
    });
  } catch {
    return d;
  }
}

export default function NewsCard({
  title,
  slug,
  excerpt,
  thumbnail_url,
  published_at,
}: NewsCardProps) {
  return (
    <Link
      href={`/berita/${slug}`}
      className="block rounded-xl border border-gray-200 bg-white overflow-hidden hover:shadow-md hover:border-blue-200 transition-all"
    >
      {thumbnail_url && (
        <div className="relative h-40 bg-gray-100">
          <Image
            src={thumbnail_url}
            alt={title}
            fill
            className="object-cover"
            sizes="(max-width: 768px) 100vw, 33vw"
            unoptimized={thumbnail_url.startsWith("http")}
          />
        </div>
      )}
      <div className="p-4">
        <p className="text-xs text-gray-500 mb-1">
          {published_at ? formatDate(published_at) : "-"}
        </p>
        <h3 className="font-semibold text-gray-900 line-clamp-2">{title}</h3>
        {excerpt && (
          <p className="mt-2 text-sm text-gray-600 line-clamp-2">{excerpt}</p>
        )}
        <span className="mt-2 inline-block text-sm font-medium text-blue-600 hover:text-blue-800">
          Selengkapnya →
        </span>
      </div>
    </Link>
  );
}
