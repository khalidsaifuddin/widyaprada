"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

interface CBTExamDetail {
  id: string;
  code: string;
  name: string;
  jadwal_mulai: string;
  jadwal_selesai: string;
  durasi_menit: number;
  dapat_mulai: boolean;
  alasan?: string;
}

export default function CBTInstructionsPage() {
  const params = useParams();
  const router = useRouter();
  const examId = params.examId as string;
  const [exam, setExam] = useState<CBTExamDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [starting, setStarting] = useState(false);
  const [error, setError] = useState("");
  const [notFoundMessage, setNotFoundMessage] = useState("");

  const fetchExam = useCallback(async () => {
    const res = await apiService.get<CBTExamDetail>(`v1/cbt/exams/${examId}`);
    if (res.success && res.data) {
      setExam(res.data);
      setNotFoundMessage("");
    } else {
      setExam(null);
      setNotFoundMessage(res.message ?? "Ujian tidak ditemukan.");
    }
    setLoading(false);
  }, [examId]);

  useEffect(() => {
    fetchExam();
  }, [fetchExam]);

  const handleMulai = async () => {
    setStarting(true);
    setError("");
    const res = await apiService.post<{
      attempt_id: string;
      started_at: string;
      durasi_menit: number;
      jadwal_selesai?: string;
    }>(`v1/cbt/exams/${examId}/start`, {});
    if (res.success && res.data) {
      const data = res.data as {
        attempt_id: string;
        started_at: string;
        durasi_menit: number;
      };
      const params = new URLSearchParams({
        attemptId: data.attempt_id,
        startedAt: data.started_at,
        durasiMenit: String(data.durasi_menit),
      });
      router.push(`/wpujikom/cbt/${examId}/kerjakan?${params}`);
      return;
    }
    setError(res.message ?? "Gagal memulai ujian");
    setStarting(false);
  };

  if (loading) return <div className="p-8">Memuat...</div>;
  if (!exam) {
    return (
      <div className="space-y-4">
        <Link href="/wpujikom/cbt" className="text-blue-600 hover:underline">
          ← Kembali
        </Link>
        <p className="text-red-600">{notFoundMessage || "Ujian tidak ditemukan."}</p>
      </div>
    );
  }

  const canStart = exam.dapat_mulai;

  return (
    <div className="max-w-2xl mx-auto space-y-6">
      <Link href="/wpujikom/cbt" className="text-gray-600 hover:text-gray-900">
        ← Batal
      </Link>
      <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-6">
        <h1 className="text-xl font-bold text-gray-900">Petunjuk Ujian</h1>
        <div className="mt-4 space-y-3 text-gray-600">
          <p><strong>Ujian:</strong> {exam.name}</p>
          <p><strong>Durasi:</strong> {exam.durasi_menit} menit</p>
          <p>Pastikan koneksi internet Anda stabil. Setelah memulai, timer akan berjalan.</p>
          <p>Jawaban otomatis tersimpan. Klik &quot;Submit Ujian&quot; untuk mengirim jawaban final.</p>
        </div>
        {exam.alasan && (
          <div className="mt-4 p-3 rounded-lg bg-amber-50 text-amber-800 text-sm">{exam.alasan}</div>
        )}
        {error && (
          <div className="mt-4 p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>
        )}
        <div className="mt-6 flex gap-3">
          <button
            onClick={handleMulai}
            disabled={starting || !canStart}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 disabled:opacity-50 font-medium"
          >
            {starting ? "Memulai..." : "Mulai Ujian"}
          </button>
          <Link
            href="/wpujikom/cbt"
            className="px-4 py-2 border border-gray-300 rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </div>
    </div>
  );
}
