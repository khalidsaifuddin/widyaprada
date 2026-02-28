"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useEffect, useState } from "react";
import AssignmentCardItem from "@/components/molecules/AssignmentCardItem";

interface AssignmentItem {
  id: string;
  exam_name: string;
  deadline: string;
  status: "belum_dikerjakan" | "sudah_dikerjakan";
  score?: number;
  can_start: boolean;
  can_view_result: boolean;
  can_view_leaderboard: boolean;
}

export default function AssignmentBlock() {
  const [assignments, setAssignments] = useState<AssignmentItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService
      .get<{ data?: AssignmentItem[] }>("v1/dashboard/assignments", { limit: 10, page: 1 })
      .then((res) => {
        if (res.success) {
          const raw = res.data as { data?: AssignmentItem[] } | AssignmentItem[];
          const items = Array.isArray(raw) ? raw : raw?.data ?? [];
          setAssignments(items);
        } else {
          setError(res.message ?? "Gagal memuat tugas");
        }
      })
      .catch(() => setError("Gagal memuat tugas"))
      .finally(() => setLoading(false));
  }, []);

  return (
    <div className="bg-white rounded-xl shadow-sm border border-gray-200 overflow-hidden">
      <div className="flex items-center justify-between p-6 border-b border-gray-200">
        <h2 className="text-lg font-semibold text-gray-900">Tugas Saya</h2>
        <Link
          href="/wpujikom/assignment"
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
        ) : assignments.length === 0 ? (
          <p className="text-gray-500 text-center py-8">
            Anda belum memiliki penugasan ujian.
          </p>
        ) : (
          <div className="space-y-3">
            {assignments.map((a) => (
              <AssignmentCardItem key={a.id} {...a} />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
