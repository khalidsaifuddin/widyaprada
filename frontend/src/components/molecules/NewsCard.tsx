"use client";

import Image from "next/image";
import Link from "next/link";
import { resolveImageUrl } from "@/lib/image";

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
  const imageUrl = thumbnail_url?.trim() || undefined;

  return (
    <Link
      href={`/berita/${slug}`}
      className="block rounded-xl border border-gray-200 bg-white overflow-hidden hover:shadow-md hover:border-blue-200 transition-all"
    >
      <div className="relative h-40 bg-gray-100">
        {imageUrl ? (
          <Image
            src={resolveImageUrl(imageUrl)}
            alt={title}
            fill
            className="object-cover"
            sizes="(max-width: 768px) 100vw, 33vw"
            unoptimized
          />
        ) : (
          <div className="absolute inset-0 flex items-center justify-center text-gray-400">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-16 w-16" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
            </svg>
          </div>
        )}
      </div>
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
