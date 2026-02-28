"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";

interface CBTHistoryItem {
  attempt_id: string;
  exam_id: string;
  exam_name: string;
  started_at: string;
  submitted_at: string;
  score?: number;
}

interface CBTHistoryResponse {
  items: CBTHistoryItem[];
}

function formatDateTime(d: string): string {
  try {
    return new Date(d).toLocaleString("id-ID", {
      day: "numeric",
      month: "short",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      timeZone: "Asia/Jakarta",
    });
  } catch {
    return d;
  }
}

export default function CBTRiwayatPage() {
  const [items, setItems] = useState<CBTHistoryItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  const fetchHistory = useCallback(async () => {
    setLoading(true);
    setError("");
    const res = await apiService.get<CBTHistoryResponse>("v1/cbt/history");
    if (res.success && res.data) {
      const d = res.data as CBTHistoryResponse;
      setItems(d.items ?? []);
    } else {
      setError(res.message ?? "Gagal memuat riwayat");
    }
    setLoading(false);
  }, []);

  useEffect(() => {
    fetchHistory();
  }, [fetchHistory]);

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Riwayat Ujian</h1>
          <p className="text-gray-600 mt-1">Daftar ujian yang sudah dikerjakan</p>
        </div>
        <Link
          href="/wpujikom/cbt"
          className="inline-flex px-4 py-2 rounded-lg border border-gray-300 text-gray-700 hover:bg-gray-50 font-medium"
        >
          ← Daftar Ujian
        </Link>
      </div>

      {error && (
        <div className="p-4 rounded-lg bg-red-50 text-red-700">{error}</div>
      )}

      {loading ? (
        <div className="space-y-3">
          {[1, 2, 3].map((i) => (
            <div key={i} className="h-24 rounded-lg bg-gray-100 animate-pulse" />
          ))}
        </div>
      ) : items.length === 0 ? (
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-12 text-center">
          <p className="text-gray-500">Belum ada riwayat ujian.</p>
          <Link href="/wpujikom/cbt" className="mt-4 inline-block text-blue-600 hover:text-blue-800">
            Daftar Ujian
          </Link>
        </div>
      ) : (
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
          <div className="overflow-x-auto">
            <table className="min-w-full divide-y divide-gray-200">
              <thead className="bg-gray-50">
                <tr>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                    Ujian
                  </th>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                    Submit
                  </th>
                  <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                    Nilai
                  </th>
                  <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">
                    Aksi
                  </th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-200">
                {items.map((item) => (
                  <tr key={item.attempt_id}>
                    <td className="px-4 py-3">
                      <span className="font-medium text-gray-900">{item.exam_name}</span>
                    </td>
                    <td className="px-4 py-3 text-sm text-gray-600">
                      {formatDateTime(item.submitted_at)}
                    </td>
                    <td className="px-4 py-3">
                      {item.score !== undefined && item.score !== null ? (
                        <span className="font-medium">{item.score}</span>
                      ) : (
                        <span className="text-gray-400">-</span>
                      )}
                    </td>
                    <td className="px-4 py-3 text-right">
                      <Link
                        href={`/wpujikom/assignment/${item.exam_id}/hasil`}
                        className="text-blue-600 hover:text-blue-800 text-sm font-medium"
                      >
                        Lihat Hasil
                      </Link>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      )}
    </div>
  );
}
