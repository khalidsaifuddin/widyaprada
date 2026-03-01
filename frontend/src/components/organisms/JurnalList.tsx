"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";
import { useTranslation } from "react-i18next";

interface JurnalItem {
  id: string;
  title: string;
  author: string;
  abstract?: string;
  published_at?: string;
  year?: number;
  category?: string;
}

interface JurnalListResponse {
  items: JurnalItem[];
  total_page: number;
  total_data: number;
  page: number;
  page_size: number;
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

export default function JurnalList() {
  const { t } = useTranslation("common");
  const [items, setItems] = useState<JurnalItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [search, setSearch] = useState("");
  const [tahun, setTahun] = useState("");
  const [sort, setSort] = useState("terbaru");
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const pageSize = 9;

  const fetchJurnal = useCallback(async () => {
    setLoading(true);
    setError("");
    const res = await apiService.get<JurnalListResponse>("v1/jurnal", {
      q: search || undefined,
      tahun: tahun || undefined,
      sort: sort || undefined,
      page,
      page_size: pageSize,
    });
    if (res.success && res.data) {
      const d = res.data as JurnalListResponse;
      setItems(d.items ?? []);
      setTotalPage(d.total_page ?? 1);
    } else {
      setError(res.message ?? t("journal.loadFailed"));
    }
    setLoading(false);
  }, [search, tahun, sort, page, t]);

  useEffect(() => {
    fetchJurnal();
  }, [fetchJurnal]);

  return (
    <div className="py-12">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <h1 className="text-2xl font-bold text-gray-900 mb-8">{t("journal.title")}</h1>

        <form
          onSubmit={(e) => {
            e.preventDefault();
            setPage(1);
            fetchJurnal();
          }}
          className="mb-8 flex flex-wrap gap-3"
        >
          <input
            type="text"
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            placeholder={t("journal.searchPlaceholder")}
            className="rounded-lg border border-gray-300 px-3 py-2 text-sm w-64 focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
          />
          <input
            type="text"
            value={tahun}
            onChange={(e) => setTahun(e.target.value)}
            placeholder={t("journal.year")}
            className="rounded-lg border border-gray-300 px-3 py-2 text-sm w-24"
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
              <div key={i} className="h-44 rounded-xl bg-gray-100 animate-pulse" />
            ))}
          </div>
        ) : error ? (
          <p className="text-red-600">{error}</p>
        ) : items.length === 0 ? (
          <p className="text-gray-500 text-center py-16">{t("journal.emptySearch")}</p>
        ) : (
          <>
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {items.map((item) => (
                <Link
                  key={item.id}
                  href={`/jurnal/${item.id}`}
                  className="block rounded-xl border border-gray-200 bg-white p-5 hover:shadow-md hover:border-blue-200 transition-all"
                >
                  <p className="text-xs text-gray-500 mb-1">
                    {item.published_at ? formatDate(item.published_at) : item.year || "-"} • {item.author}
                  </p>
                  <h3 className="font-semibold text-gray-900 line-clamp-2">{item.title}</h3>
                  {item.abstract && (
                    <p className="mt-2 text-sm text-gray-600 line-clamp-2">{item.abstract}</p>
                  )}
                  <span className="mt-2 inline-block text-sm font-medium text-blue-600">
                    {t("journal.readMore")}
                  </span>
                </Link>
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
