"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";
import AssignmentCardItem from "@/components/molecules/AssignmentCardItem";

interface AssignmentItem {
  exam_id: string;
  exam_name: string;
  deadline: string;
  status: "belum_dikerjakan" | "sudah_dikerjakan";
  score?: number;
  can_view_leaderboard: boolean;
}

interface AssignmentListResponse {
  items: AssignmentItem[];
  total_page: number;
  total_data: number;
  page: number;
  page_size: number;
}

export default function AssignmentListPage() {
  const [assignments, setAssignments] = useState<AssignmentItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [statusFilter, setStatusFilter] = useState("");
  const [page, setPage] = useState(1);
  const [totalPage, setTotalPage] = useState(1);
  const [totalData, setTotalData] = useState(0);
  const pageSize = 10;

  const fetchAssignments = useCallback(async () => {
    setLoading(true);
    setError("");
    const res = await apiService.get<AssignmentListResponse>("v1/assignments", {
      status: statusFilter || undefined,
      page,
      page_size: pageSize,
    });
    if (res.success && res.data) {
      const d = res.data as AssignmentListResponse;
      setAssignments(d.items ?? []);
      setTotalPage(d.total_page ?? 1);
      setTotalData(d.total_data ?? 0);
    } else {
      setError(res.message ?? "Gagal memuat tugas");
    }
    setLoading(false);
  }, [statusFilter, page]);

  useEffect(() => {
    fetchAssignments();
  }, [fetchAssignments]);

  return (
    <div className="space-y-6">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Tugas Saya</h1>
          <p className="text-gray-600 mt-1">Penugasan uji kompetensi</p>
        </div>
        <Link
          href="/wpujikom/assignment/apply"
          className="inline-flex items-center justify-center px-4 py-2 rounded-lg border border-gray-300 text-gray-700 transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50 font-medium"
        >
          Apply Ujikom
        </Link>
      </div>

      <div className="flex gap-2">
        <select
          value={statusFilter}
          onChange={(e) => {
            setStatusFilter(e.target.value);
            setPage(1);
          }}
          className="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
        >
          <option value="">Semua</option>
          <option value="belum_dikerjakan">Belum dikerjakan</option>
          <option value="sudah_dikerjakan">Sudah dikerjakan</option>
        </select>
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
      ) : assignments.length === 0 ? (
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-12 text-center">
          <p className="text-gray-500">Anda belum memiliki penugasan ujian.</p>
        </div>
      ) : (
        <div className="space-y-3">
          {assignments.map((a) => (
            <AssignmentCardItem
              key={a.exam_id}
              id={a.exam_id}
              exam_name={a.exam_name}
              deadline={a.deadline}
              status={a.status}
              score={a.score}
              can_start={a.status === "belum_dikerjakan"}
              can_view_result={a.status === "sudah_dikerjakan"}
              can_view_leaderboard={a.can_view_leaderboard}
            />
          ))}
          {totalPage > 1 && (
            <div className="flex justify-center gap-2 pt-4">
              <button
                onClick={() => setPage((p) => Math.max(1, p - 1))}
                disabled={page <= 1}
                className="px-3 py-1 rounded border border-gray-300 disabled:opacity-50 text-sm"
              >
                Sebelumnya
              </button>
              <button
                onClick={() => setPage((p) => Math.min(totalPage, p + 1))}
                disabled={page >= totalPage}
                className="px-3 py-1 rounded border border-gray-300 disabled:opacity-50 text-sm"
              >
                Selanjutnya
              </button>
            </div>
          )}
        </div>
      )}
    </div>
  );
}
