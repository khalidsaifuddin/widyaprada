"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useEffect, useState } from "react";
import NewsCard, { type NewsCardProps } from "@/components/molecules/NewsCard";

interface ArticlePublicItem {
  id: string;
  title: string;
  slug: string;
  excerpt?: string;
  thumbnail_url?: string;
  published_at?: string;
}

interface LandingHomeData {
  berita?: ArticlePublicItem[];
}

export default function NewsPanel() {
  const [items, setItems] = useState<ArticlePublicItem[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    apiService
      .get<LandingHomeData>("v1/landing/home")
      .then((res) => {
        if (res.success && res.data) {
          const d = res.data as LandingHomeData;
          setItems(d.berita ?? []);
        }
      })
      .catch(() => {})
      .finally(() => setLoading(false));
  }, []);

  return (
    <section className="py-12 bg-white">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between mb-8">
          <h2 className="text-xl font-bold text-gray-900">Berita Terbaru</h2>
          <Link
            href="/berita"
            className="text-sm font-medium text-blue-600 hover:text-blue-800"
          >
            Lihat semua
          </Link>
        </div>
        {loading ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {[1, 2, 3].map((i) => (
              <div key={i} className="h-56 rounded-xl bg-gray-100 animate-pulse" />
            ))}
          </div>
        ) : items.length === 0 ? (
          <p className="text-gray-500 text-center py-12">Belum ada berita</p>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {items.slice(0, 6).map((item) => (
              <NewsCard key={item.id} {...(item as NewsCardProps)} />
            ))}
          </div>
        )}
      </div>
    </section>
  );
}
