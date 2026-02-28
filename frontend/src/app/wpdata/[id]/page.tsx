"use client";

import { apiService } from "@/lib/api";
import { ConfirmDialog } from "@/components";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface WPDataDetail {
  id: string;
  nip: string;
  nama_lengkap: string;
  jenis_kelamin: string;
  golongan_ruang: string;
  pangkat: string;
  jenjang_jabatan_fungsional: string;
  satker_id: string;
  unit_kerja: string;
  pendidikan_terakhir: string;
  tmt_golongan: string;
  tmt_jabatan_fungsional: string;
  no_sk: string;
  no_hp: string;
  email: string;
  alamat: string;
  status: string;
  keterangan: string;
}

export default function WPDataDetailPage() {
  const params = useParams();
  const router = useRouter();
  const id = params?.id as string;
  const [data, setData] = useState<WPDataDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [deleteOpen, setDeleteOpen] = useState(false);

  useEffect(() => {
    if (!id) return;
    apiService.get<WPDataDetail>(`v1/wp-data/${id}`).then((res) => {
      if (res.success && res.data) setData(res.data as WPDataDetail);
      setLoading(false);
    });
  }, [id]);

  const handleDelete = async () => {
    const res = await apiService.delete(`v1/wp-data/${id}`, { reason: "Dihapus oleh admin" });
    setDeleteOpen(false);
    if (res.success) router.push("/wpdata");
  };

  if (loading) return <div className="animate-pulse h-32 bg-gray-100 rounded" />;
  if (!data) return <p className="text-red-600">Data tidak ditemukan</p>;

  const rows = [
    { label: "NIP", value: data.nip },
    { label: "Nama Lengkap", value: data.nama_lengkap },
    { label: "Jenis Kelamin", value: data.jenis_kelamin === "L" ? "Laki-laki" : data.jenis_kelamin === "P" ? "Perempuan" : data.jenis_kelamin || "-" },
    { label: "Satker ID", value: data.satker_id },
    { label: "Unit Kerja", value: data.unit_kerja || "-" },
    { label: "Golongan/Ruang", value: data.golongan_ruang || "-" },
    { label: "Pangkat", value: data.pangkat || "-" },
    { label: "Jenjang Jabatan Fungsional", value: data.jenjang_jabatan_fungsional || "-" },
    { label: "Pendidikan Terakhir", value: data.pendidikan_terakhir || "-" },
    { label: "No. HP", value: data.no_hp || "-" },
    { label: "Email", value: data.email || "-" },
    { label: "Alamat", value: data.alamat || "-" },
    { label: "Status", value: data.status },
  ];

  return (
    <div className="space-y-6">
      <div className="flex justify-between">
        <Link href="/wpdata" className="text-blue-600 hover:underline text-sm">← Kembali</Link>
        <div className="flex gap-2">
          <Link href={`/wpdata/${id}/edit`} className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700">Edit</Link>
          <button type="button" onClick={() => setDeleteOpen(true)} className="rounded-lg border border-red-600 text-red-600 px-4 py-2 text-sm font-medium hover:bg-red-50">Hapus</button>
        </div>
      </div>
      <div className="bg-white rounded-xl border border-gray-200 p-6">
        <h1 className="text-xl font-bold text-gray-900 mb-6">{data.nama_lengkap}</h1>
        <dl className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {rows.map((r) => (
            <div key={r.label}>
              <dt className="text-sm text-gray-500">{r.label}</dt>
              <dd className="text-sm font-medium text-gray-900">{r.value || "-"}</dd>
            </div>
          ))}
        </dl>
      </div>
      <ConfirmDialog isOpen={deleteOpen} onClose={() => setDeleteOpen(false)} onConfirm={handleDelete} title="Hapus Data WP" message="Apakah Anda yakin?" confirmText="Ya, Hapus" type="danger" />
    </div>
  );
}
