"use client";

import { apiService } from "@/lib/api";
import { ArrowTopRightOnSquareIcon, LinkIcon } from "@heroicons/react/24/outline";
import { useEffect, useState } from "react";

interface LinkItem {
  id: string;
  title: string;
  url: string;
  description?: string;
  sort_order?: number;
  buka_di_tab_baru?: boolean;
}

interface LandingHomeData {
  tautan?: LinkItem[];
}

export default function LinksPanel() {
  const [items, setItems] = useState<LinkItem[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    apiService
      .get<LandingHomeData>("v1/landing/home")
      .then((res) => {
        if (res.success && res.data) {
          const d = res.data as LandingHomeData;
          setItems(d.tautan ?? []);
        }
      })
      .catch(() => {})
      .finally(() => setLoading(false));
  }, []);

  if (loading || items.length === 0) return null;

  return (
    <section className="py-12 bg-gray-50 border-t border-gray-200">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <h2 className="text-xl font-bold text-gray-900 mb-6">Tautan Penting</h2>
        <div className="flex flex-wrap gap-3">
          {items.map((item) => (
            <a
              key={item.id}
              href={item.url}
              target={item.buka_di_tab_baru ? "_blank" : undefined}
              rel={item.buka_di_tab_baru ? "noopener noreferrer" : undefined}
              className="inline-flex items-center gap-2 rounded-lg border border-gray-200 bg-white px-4 py-2.5 text-sm font-medium text-gray-700 hover:bg-gray-50 hover:border-blue-200 transition-colors"
            >
              <LinkIcon className="h-4 w-4 text-gray-400" />
              {item.title}
              {item.buka_di_tab_baru && (
                <ArrowTopRightOnSquareIcon className="h-4 w-4 text-gray-400" />
              )}
            </a>
          ))}
        </div>
      </div>
    </section>
  );
}
