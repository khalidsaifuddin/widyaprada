"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useEffect, useState } from "react";
import JournalCardItem from "@/components/molecules/JournalCardItem";

interface JournalItem {
  id: string;
  title: string;
  submitted_at: string;
  status: string;
}

export default function JournalBlock() {
  const [journals, setJournals] = useState<JournalItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService
      .get<{ data?: JournalItem[] }>("v1/dashboard/journals", { limit: 10, page: 1 })
      .then((res) => {
        if (res.success) {
          const raw = res.data as { data?: JournalItem[] } | JournalItem[];
          const items = Array.isArray(raw) ? raw : raw?.data ?? [];
          setJournals(items);
        } else {
          setError(res.message ?? "Gagal memuat jurnal");
        }
      })
      .catch(() => setError("Gagal memuat jurnal"))
      .finally(() => setLoading(false));
  }, []);

  return (
    <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
      <div className="flex items-center justify-between p-6 border-b border-gray-200">
        <h2 className="text-lg font-semibold text-gray-900">Jurnal Saya</h2>
        <Link
          href="/wpjurnal"
          className="text-sm font-medium text-blue-600 hover:text-blue-800"
        >
          Lihat semua
        </Link>
      </div>
      <div className="p-6">
        {loading ? (
          <div className="space-y-3">
            {[1, 2, 3].map((i) => (
              <div key={i} className="h-24 rounded-lg bg-gray-100 animate-pulse" />
            ))}
          </div>
        ) : error ? (
          <p className="text-sm text-red-600">{error}</p>
        ) : journals.length === 0 ? (
          <p className="text-gray-500 text-center py-8">
            Anda belum mengirimkan jurnal.
          </p>
        ) : (
          <div className="space-y-3">
            {journals.map((j) => (
              <JournalCardItem key={j.id} {...j} />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
