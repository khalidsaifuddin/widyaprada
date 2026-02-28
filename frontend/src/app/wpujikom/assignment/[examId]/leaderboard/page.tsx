"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
import { getUserProfile } from "@/lib/auth";

interface LeaderboardItem {
  rank: number;
  user_id: string;
  name: string;
  score: number;
}

interface LeaderboardResponse {
  exam_id: string;
  exam_name: string;
  items: LeaderboardItem[];
}

export default function LeaderboardPage() {
  const params = useParams();
  const examId = params.examId as string;
  const [data, setData] = useState<LeaderboardResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [currentUserId, setCurrentUserId] = useState<string | null>(null);

  useEffect(() => {
    getUserProfile().then((p) => setCurrentUserId(p?.user_id ?? null));
  }, []);

  const fetchLeaderboard = useCallback(async () => {
    const res = await apiService.get<LeaderboardResponse>(
      `v1/assignments/${examId}/leaderboard`
    );
    if (res.success && res.data) {
      setData(res.data as LeaderboardResponse);
    } else {
      setError(res.message ?? "Gagal memuat leaderboard. Mungkin bersifat privat.");
    }
    setLoading(false);
  }, [examId]);

  useEffect(() => {
    fetchLeaderboard();
  }, [fetchLeaderboard]);

  if (loading) return <div className="p-8">Memuat...</div>;
  if (error && !data) {
    return (
      <div className="space-y-4">
        <Link href={`/wpujikom/assignment/${examId}/hasil`} className="text-blue-600 hover:underline">
          ← Kembali
        </Link>
        <p className="text-red-600">{error}</p>
      </div>
    );
  }

  if (!data) return null;

  return (
    <div className="space-y-6">
      <Link href={`/wpujikom/assignment/${examId}/hasil`} className="text-gray-600 hover:text-gray-900">
        ← Kembali ke Hasil
      </Link>
      <h1 className="text-2xl font-bold text-gray-900">Leaderboard: {data.exam_name}</h1>
      <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                Peringkat
              </th>
              <th className="px-4 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                Nama
              </th>
              <th className="px-4 py-3 text-right text-xs font-medium text-gray-500 uppercase">
                Nilai
              </th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200">
            {data.items.map((row) => (
              <tr
                key={row.user_id}
                className={
                  row.user_id === currentUserId
                    ? "bg-blue-50 font-medium"
                    : ""
                }
              >
                <td className="px-4 py-3">{row.rank}</td>
                <td className="px-4 py-3">{row.name}</td>
                <td className="px-4 py-3 text-right font-mono">{row.score}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
