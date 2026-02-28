"use client";

import { apiService } from "@/lib/api";
import { getUserProfile } from "@/lib/auth";
import { PlayIcon } from "@heroicons/react/24/outline";
import Link from "next/link";
import { useEffect, useState } from "react";

interface JadwalUjikomItem {
  exam_id: string;
  exam_name: string;
  jadwal_mulai: string;
  jadwal_selesai: string;
}

interface PengumumanData {
  hasil_seleksi_administrasi?: { status?: string; message?: string };
  info_jadwal_ujikom?: JadwalUjikomItem[];
  can_start_ujikom?: boolean;
  exams_tersedia?: { id: string; name: string }[];
}

function formatDate(d: string): string {
  try {
    return new Date(d).toLocaleDateString("id-ID", {
      day: "numeric",
      month: "long",
      year: "numeric",
      hour: "2-digit",
      minute: "2-digit",
      timeZone: "Asia/Jakarta",
    });
  } catch {
    return d;
  }
}

export default function AnnouncementPanel() {
  const [data, setData] = useState<PengumumanData | null>(null);
  const [loading, setLoading] = useState(true);
  const [show, setShow] = useState(false);

  useEffect(() => {
    getUserProfile()
      .then((user) => {
        const roles = user?.role_user ?? [];
        const isPeserta = roles.some(
          (r) =>
            r.role_aplikasi === "PESERTA" ||
            r.role_aplikasi === "SUPER_ADMIN" ||
            r.role_aplikasi === "ADMIN_UJIKOM"
        );
        if (!isPeserta) {
          setLoading(false);
          return;
        }
        setShow(true);
        return apiService.get<PengumumanData>("v1/beranda/pengumuman");
      })
      .then((res) => {
        if (res?.success && res.data) {
          setData(res.data as PengumumanData);
        }
      })
      .catch(() => {})
      .finally(() => setLoading(false));
  }, []);

  if (!show || loading) return null;

  return (
    <section className="py-12 bg-blue-50 border-t border-blue-100">
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <h2 className="text-xl font-bold text-gray-900 mb-6">Pengumuman</h2>
        <div className="rounded-xl border border-blue-200 bg-white p-6 shadow-sm">
          {data?.hasil_seleksi_administrasi && (
            <div className="mb-6 pb-6 border-b border-gray-200">
              <h3 className="font-semibold text-gray-900 mb-2">
                Hasil Seleksi Administrasi
              </h3>
              <p className="text-gray-600 text-sm">
                {data.hasil_seleksi_administrasi.message ||
                  `Status: ${data.hasil_seleksi_administrasi.status || "-"}`}
              </p>
            </div>
          )}
          {data?.info_jadwal_ujikom && data.info_jadwal_ujikom.length > 0 && (
            <div className="space-y-4">
              <h3 className="font-semibold text-gray-900">Info Jadwal Uji Kompetensi</h3>
              <ul className="space-y-3">
                {data.info_jadwal_ujikom.map((j) => (
                  <li
                    key={j.exam_id}
                    className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-2 rounded-lg border border-gray-200 p-4"
                  >
                    <div>
                      <p className="font-medium text-gray-900">{j.exam_name}</p>
                      <p className="text-sm text-gray-500">
                        {formatDate(j.jadwal_mulai)} – {formatDate(j.jadwal_selesai)}
                      </p>
                    </div>
                    {data.can_start_ujikom && (
                      <Link
                        href={`/wpujikom/cbt/${j.exam_id}/mulai`}
                        className="inline-flex items-center gap-1.5 rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700"
                      >
                        <PlayIcon className="h-4 w-4" />
                        Mulai Ujikom
                      </Link>
                    )}
                  </li>
                ))}
              </ul>
            </div>
          )}
          {(!data || (!data.hasil_seleksi_administrasi && (!data.info_jadwal_ujikom || data.info_jadwal_ujikom.length === 0))) && (
            <p className="text-gray-500 text-sm">Belum ada pengumuman.</p>
          )}
        </div>
      </div>
    </section>
  );
}
