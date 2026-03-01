"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

interface ResultResponse {
  exam_id: string;
  exam_name: string;
  score?: number;
  submitted_at?: string;
}

export default function ExamResultPage() {
  const params = useParams();
  const examId = params.examId as string;
  const [result, setResult] = useState<ResultResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const fetchResult = useCallback(async () => {
    const res = await apiService.get<ResultResponse>(
      `v1/assignments/${examId}/result`
    );
    if (res.success && res.data) {
      setResult(res.data as ResultResponse);
    } else {
      setError(res.message ?? "Gagal memuat hasil");
    }
    setLoading(false);
  }, [examId]);

  useEffect(() => {
    fetchResult();
  }, [fetchResult]);

  if (loading) return <div className="p-8">Memuat...</div>;
  if (error && !result) {
    return (
      <div className="space-y-4">
        <Link href="/wpujikom/assignment" className="text-blue-600 hover:underline">
          ← Kembali
        </Link>
        <p className="text-red-600">{error}</p>
      </div>
    );
  }

  if (!result) return null;

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      <Link href="/wpujikom/assignment" className="text-gray-600 hover:text-gray-900">
        ← Kembali ke Tugas Saya
      </Link>
      <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-8 text-center">
        <h1 className="text-xl font-bold text-gray-900">Hasil Ujian</h1>
        <p className="mt-2 text-gray-600">{result.exam_name}</p>
        {result.score !== undefined && result.score !== null ? (
          <p className="mt-4 text-3xl font-bold text-blue-600">Nilai: {result.score}</p>
        ) : (
          <p className="mt-4 text-gray-500">Nilai akan ditampilkan setelah dinilai.</p>
        )}
        {result.submitted_at && (
          <p className="mt-2 text-sm text-gray-500">
            Dikirim: {new Date(result.submitted_at).toLocaleString("id-ID")}
          </p>
        )}
        <div className="mt-6">
          <Link
            href={`/wpujikom/assignment/${examId}/leaderboard`}
            className="inline-flex px-4 py-2 rounded-lg border border-gray-300 text-gray-700 transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50 font-medium"
          >
            Lihat Leaderboard (jika tersedia)
          </Link>
        </div>
      </div>
    </div>
  );
}
