"use client";

import Link from "next/link";

export interface AssignmentCardItemProps {
  id: string;
  exam_name: string;
  deadline: string;
  status: "belum_dikerjakan" | "sudah_dikerjakan";
  score?: number;
  can_start: boolean;
  can_view_result: boolean;
  can_view_leaderboard: boolean;
}

function formatDeadline(d: string): string {
  try {
    return new Date(d).toLocaleDateString("id-ID", {
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

export default function AssignmentCardItem({
  id,
  exam_name,
  deadline,
  status,
  score,
  can_start,
  can_view_result,
  can_view_leaderboard,
}: AssignmentCardItemProps) {
  const statusLabel = status === "sudah_dikerjakan" ? "Sudah dikerjakan" : "Belum dikerjakan";
  const statusClass =
    status === "sudah_dikerjakan"
      ? "bg-green-100 text-green-800"
      : "bg-amber-100 text-amber-800";

  return (
    <div className="p-4 rounded-lg border border-gray-200 bg-white hover:border-blue-200 transition-colors">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3">
        <div className="flex-1 min-w-0">
          <h4 className="font-medium text-gray-900 truncate">{exam_name}</h4>
          <p className="text-sm text-gray-500 mt-0.5">Batas: {formatDeadline(deadline)}</p>
          <div className="flex items-center gap-2 mt-2">
            <span className={`inline-flex px-2 py-0.5 rounded text-xs font-medium ${statusClass}`}>
              {statusLabel}
            </span>
            {score !== undefined && status === "sudah_dikerjakan" && (
              <span className="text-sm text-gray-600">Skor: {score}</span>
            )}
          </div>
        </div>
        <div className="flex flex-wrap gap-2">
          {can_start && (
            <Link
              href={`/wpujikom/cbt/${id}/mulai`}
              className="inline-flex px-3 py-1.5 text-sm font-medium rounded-lg bg-blue-600 text-white hover:bg-blue-700"
            >
              Mulai Ujian
            </Link>
          )}
          {can_view_result && (
            <Link
              href={`/wpujikom/assignment/${id}/hasil`}
              className="inline-flex px-3 py-1.5 text-sm font-medium rounded-lg border border-gray-300 text-gray-700 hover:bg-gray-50"
            >
              Lihat Hasil
            </Link>
          )}
          {can_view_leaderboard && (
            <Link
              href={`/wpujikom/assignment/${id}/leaderboard`}
              className="inline-flex px-3 py-1.5 text-sm font-medium rounded-lg border border-gray-300 text-gray-700 hover:bg-gray-50"
            >
              Lihat Leaderboard
            </Link>
          )}
        </div>
      </div>
    </div>
  );
}
