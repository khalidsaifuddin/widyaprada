"use client";

import { apiService } from "@/lib/api";
import { useCallback, useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import NewsCard, { type NewsCardProps } from "@/components/molecules/NewsCard";

interface ArticleItem {
  id: string;
  title: string;
  slug: string;
  excerpt?: string;
  thumbnail_url?: string;
  published_at?: string;
}

interface BeritaListResponse {
  items: ArticleItem[];
  total_page: number;
  total_data: number;
  page: number;
  page_size: number;
}

export default function BeritaList() {
  const { t } = useTranslation("common");
  const [items, setItems] = useState<ArticleItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [search, setSearch] = useState("");
  const [sort, setSort] = useState("terbaru");
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const pageSize = 9;

  const fetchBerita = useCallback(async () => {
    setLoading(true);
    setError("");
    const res = await apiService.get<BeritaListResponse>("v1/berita", {
      q: search || undefined,
      sort: sort || undefined,
      page,
      page_size: pageSize,
      status: "Published",
    });
    if (res.success && res.data) {
      const d = res.data as BeritaListResponse;
      setItems(d.items ?? []);
      setTotalPage(d.total_page ?? 1);
    } else {
      setError(res.message ?? t("news.loadFailed"));
    }
    setLoading(false);
  }, [search, sort, page, t]);

  useEffect(() => {
    fetchBerita();
  }, [fetchBerita]);

  return (
    <div className="py-12">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <h1 className="text-2xl font-bold text-gray-900 mb-8">{t("news.title")}</h1>

        <form
          onSubmit={(e) => {
            e.preventDefault();
            setPage(1);
            fetchBerita();
          }}
          className="mb-8 flex flex-wrap gap-3"
        >
          <input
            type="text"
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            placeholder={t("news.searchPlaceholder")}
            className="rounded-lg border border-gray-300 px-3 py-2 text-sm w-64 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
          <select
            value={sort}
            onChange={(e) => {
              setSort(e.target.value);
              setPage(1);
            }}
            className="rounded-lg border border-gray-300 px-3 py-2 text-sm"
          >
            <option value="terbaru">{t("news.sortNewest")}</option>
            <option value="terlama">{t("news.sortOldest")}</option>
          </select>
          <button
            type="submit"
            className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700"
          >
            {t("action.search")}
          </button>
        </form>

        {loading ? (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {[1, 2, 3, 4, 5, 6].map((i) => (
              <div key={i} className="h-56 rounded-xl bg-gray-100 animate-pulse" />
            ))}
          </div>
        ) : error ? (
          <p className="text-red-600">{error}</p>
        ) : items.length === 0 ? (
          <p className="text-gray-500 text-center py-16">{t("news.emptySearch")}</p>
        ) : (
          <>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {items.map((item) => (
                <NewsCard key={item.id} {...(item as NewsCardProps)} />
              ))}
            </div>
            {totalPage > 1 && (
              <div className="mt-8 flex justify-center gap-2">
                <button
                  type="button"
                  onClick={() => setPage((p) => Math.max(1, p - 1))}
                  disabled={page <= 1}
                  className="rounded-lg border border-gray-300 px-3 py-1.5 text-sm disabled:opacity-50"
                >
                  {t("news.prev")}
                </button>
                <span className="px-3 py-1.5 text-sm text-gray-600">
                  {page} / {totalPage}
                </span>
                <button
                  type="button"
                  onClick={() => setPage((p) => Math.min(totalPage, p + 1))}
                  disabled={page >= totalPage}
                  className="rounded-lg border border-gray-300 px-3 py-1.5 text-sm disabled:opacity-50"
                >
                  {t("news.next")}
                </button>
              </div>
            )}
          </>
        )}
      </div>
    </div>
  );
}
