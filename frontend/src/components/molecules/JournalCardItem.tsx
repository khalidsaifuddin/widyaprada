"use client";

import Link from "next/link";

export interface JournalCardItemProps {
  id: string;
  title: string;
  submitted_at: string;
  status: string; // Draft|Menunggu Verifikasi|Diverifikasi|Ditolak|Published
}

function formatDate(d: string): string {
  try {
    return new Date(d).toLocaleDateString("id-ID", {
      day: "numeric",
      month: "short",
      year: "numeric",
      timeZone: "Asia/Jakarta",
    });
  } catch {
    return d;
  }
}

function getStatusStyle(status: string): string {
  const s = status?.toLowerCase() ?? "";
  if (s.includes("diverifikasi") || s.includes("published")) return "bg-green-100 text-green-800";
  if (s.includes("menunggu")) return "bg-amber-100 text-amber-800";
  if (s.includes("ditolak")) return "bg-red-100 text-red-800";
  if (s.includes("draft")) return "bg-gray-100 text-gray-700";
  return "bg-gray-100 text-gray-600";
}

export default function JournalCardItem({
  id,
  title,
  submitted_at,
  status,
}: JournalCardItemProps) {
  const canEdit = status?.toLowerCase().includes("draft");

  return (
    <div className="p-4 rounded-lg border border-gray-200 bg-white hover:border-blue-200 hover:shadow-md hover:-translate-y-0.5 transition-all duration-300">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3">
        <div className="flex-1 min-w-0">
          <h4 className="font-medium text-gray-900 truncate">{title || "Tanpa judul"}</h4>
          <p className="text-sm text-gray-500 mt-0.5">Submit: {formatDate(submitted_at)}</p>
          <span
            className={`inline-flex mt-2 px-2 py-0.5 rounded text-xs font-medium ${getStatusStyle(status)}`}
          >
            {status || "Draft"}
          </span>
        </div>
        <div className="flex gap-2">
          <Link
            href={`/wpjurnal/${id}`}
            className="inline-flex px-3 py-1.5 text-sm font-medium rounded-lg border border-gray-300 text-gray-700 transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
          >
            Lihat
          </Link>
          {canEdit && (
            <Link
              href={`/wpjurnal/${id}/edit`}
              className="inline-flex px-3 py-1.5 text-sm font-medium rounded-lg bg-blue-600 text-white transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700"
            >
              Edit
            </Link>
          )}
        </div>
      </div>
    </div>
  );
}
