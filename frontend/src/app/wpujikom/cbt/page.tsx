"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";

interface CBTExamItem {
  id: string;
  code: string;
  name: string;
  jadwal_mulai: string;
  jadwal_selesai: string;
  durasi_menit: number;
}

interface CBTListResponse {
  items: CBTExamItem[];
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

export default function CBTListPage() {
  const [exams, setExams] = useState<CBTExamItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  const fetchExams = useCallback(async () => {
    setLoading(true);
    setError("");
    const res = await apiService.get<CBTListResponse>("v1/cbt/exams");
    if (res.success && res.data) {
      const d = res.data as CBTListResponse;
      setExams(d.items ?? []);
    } else {
      setError(res.message ?? "Gagal memuat daftar ujian");
    }
    setLoading(false);
  }, []);

  useEffect(() => {
    fetchExams();
  }, [fetchExams]);

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">CBT - Computer Based Test</h1>
          <p className="text-gray-600 mt-1">Daftar ujian yang tersedia untuk Anda</p>
        </div>
        <Link
          href="/wpujikom/cbt/riwayat"
          className="inline-flex items-center justify-center px-4 py-2 rounded-lg border border-gray-300 text-gray-700 hover:bg-gray-50 font-medium"
        >
          Riwayat Ujian
        </Link>
      </div>

      {error && (
        <div className="p-4 rounded-lg bg-red-50 text-red-700">{error}</div>
      )}

      {loading ? (
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          {[1, 2, 3].map((i) => (
            <div key={i} className="h-40 rounded-xl bg-gray-100 animate-pulse" />
          ))}
        </div>
      ) : exams.length === 0 ? (
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-12 text-center">
          <p className="text-gray-500">
            Tidak ada ujian yang tersedia untuk Anda saat ini.
          </p>
          <Link
            href="/wpujikom/assignment"
            className="mt-4 inline-block text-blue-600 hover:text-blue-800"
          >
            Lihat Tugas Saya
          </Link>
        </div>
      ) : (
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          {exams.map((exam) => (
            <div
              key={exam.id}
              className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:border-blue-200 transition-colors"
            >
              <h3 className="font-semibold text-gray-900">{exam.name}</h3>
              <p className="text-sm text-gray-500 mt-1 font-mono">{exam.code}</p>
              <p className="text-sm text-gray-600 mt-2">
                {formatDateTime(exam.jadwal_mulai)} - {formatDateTime(exam.jadwal_selesai)}
              </p>
              <p className="text-sm text-gray-600">Durasi: {exam.durasi_menit} menit</p>
              <Link
                href={`/wpujikom/cbt/${exam.id}/mulai`}
                className="mt-4 inline-flex px-4 py-2 rounded-lg bg-blue-600 text-white hover:bg-blue-700 font-medium text-sm"
              >
                Mulai Ujian
              </Link>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
