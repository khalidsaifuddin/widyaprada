"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useEffect, useState } from "react";

interface JurnalItem {
  id: string;
  title: string;
  author: string;
  abstract?: string;
  published_at?: string;
}

interface LandingHomeData {
  jurnal?: JurnalItem[];
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

export default function JournalPanel() {
  const [items, setItems] = useState<JurnalItem[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    apiService
      .get<LandingHomeData>("v1/landing/home")
      .then((res) => {
        if (res.success && res.data) {
          const d = res.data as LandingHomeData;
          setItems(d.jurnal ?? []);
        }
      })
      .catch(() => {})
      .finally(() => setLoading(false));
  }, []);

  return (
    <section className="py-12 bg-white border-t border-gray-200">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between mb-8">
          <h2 className="text-xl font-bold text-gray-900">Jurnal Terbaru</h2>
          <Link
            href="/jurnal"
            className="text-sm font-medium text-blue-600 hover:text-blue-800"
          >
            Lihat semua
          </Link>
        </div>
        {loading ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {[1, 2, 3].map((i) => (
              <div key={i} className="h-44 rounded-xl bg-gray-100 animate-pulse" />
            ))}
          </div>
        ) : items.length === 0 ? (
          <p className="text-gray-500 text-center py-12">Belum ada jurnal</p>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {items.slice(0, 6).map((item) => (
              <Link
                key={item.id}
                href={`/jurnal/${item.id}`}
                className="block rounded-xl border border-gray-200 bg-white p-5 hover:shadow-md hover:border-blue-200 transition-all"
              >
                <p className="text-xs text-gray-500 mb-1">
                  {item.published_at ? formatDate(item.published_at) : "-"} • {item.author}
                </p>
                <h3 className="font-semibold text-gray-900 line-clamp-2">{item.title}</h3>
                {item.abstract && (
                  <p className="mt-2 text-sm text-gray-600 line-clamp-2">{item.abstract}</p>
                )}
                <span className="mt-2 inline-block text-sm font-medium text-blue-600">
                  Baca detail →
                </span>
              </Link>
            ))}
          </div>
        )}
      </div>
    </section>
  );
}
