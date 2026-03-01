"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useEffect, useState } from "react";
import JournalCardItem from "@/components/molecules/JournalCardItem";
import { Card } from "@/components";

interface JournalItem {
  id: string;
  title: string;
  submitted_at: string;
  status: string;
}

export default function WPJurnalPage() {
  const [journals, setJournals] = useState<JournalItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService
      .get<{ data?: JournalItem[] }>("v1/dashboard/journals", { limit: 50, page: 1 })
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
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <h1 className="text-2xl font-bold text-gray-900">Manajemen Jurnal</h1>
        <Link
          href="/wpjurnal/create"
          className="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-blue-600 text-white transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 font-medium"
        >
          <span>+</span> Buat Jurnal
        </Link>
      </div>

      <Card title="Jurnal Saya" subtitle="Daftar jurnal yang telah dibuat">
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
            Anda belum membuat jurnal. Klik &quot;Buat Jurnal&quot; untuk menambahkan jurnal dengan upload file PDF.
          </p>
        ) : (
          <div className="space-y-3">
            {journals.map((j) => (
              <JournalCardItem key={j.id} {...j} />
            ))}
          </div>
        )}
      </Card>
    </div>
  );
}
