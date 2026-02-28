"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function WPDataEditPage() {
  const params = useParams();
  const router = useRouter();
  const id = params?.id as string;
  const [nip, setNip] = useState("");
  const [namaLengkap, setNamaLengkap] = useState("");
  const [jenisKelamin, setJenisKelamin] = useState("");
  const [golonganRuang, setGolonganRuang] = useState("");
  const [pangkat, setPangkat] = useState("");
  const [jenjangJabatanFungsional, setJenjangJabatanFungsional] = useState("");
  const [satkerId, setSatkerId] = useState("");
  const [unitKerja, setUnitKerja] = useState("");
  const [pendidikanTerakhir, setPendidikanTerakhir] = useState("");
  const [tmtGolongan, setTmtGolongan] = useState("");
  const [tmtJabatanFungsional, setTmtJabatanFungsional] = useState("");
  const [noSk, setNoSk] = useState("");
  const [noHp, setNoHp] = useState("");
  const [email, setEmail] = useState("");
  const [alamat, setAlamat] = useState("");
  const [status, setStatus] = useState("Aktif");
  const [keterangan, setKeterangan] = useState("");
  const [loading, setLoading] = useState(false);
  const [loadDetail, setLoadDetail] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    if (!id) return;
    apiService.get<Record<string, unknown>>(`v1/wp-data/${id}`).then((res) => {
      if (res.success && res.data) {
        const d = res.data as Record<string, unknown>;
        setNip((d.nip as string) ?? "");
        setNamaLengkap((d.nama_lengkap as string) ?? "");
        setJenisKelamin((d.jenis_kelamin as string) ?? "");
        setGolonganRuang((d.golongan_ruang as string) ?? "");
        setPangkat((d.pangkat as string) ?? "");
        setJenjangJabatanFungsional((d.jenjang_jabatan_fungsional as string) ?? "");
        setSatkerId((d.satker_id as string) ?? "");
        setUnitKerja((d.unit_kerja as string) ?? "");
        setPendidikanTerakhir((d.pendidikan_terakhir as string) ?? "");
        setTmtGolongan((d.tmt_golongan as string) ?? "");
        setTmtJabatanFungsional((d.tmt_jabatan_fungsional as string) ?? "");
        setNoSk((d.no_sk as string) ?? "");
        setNoHp((d.no_hp as string) ?? "");
        setEmail((d.email as string) ?? "");
        setAlamat((d.alamat as string) ?? "");
        setStatus((d.status as string) ?? "Aktif");
        setKeterangan((d.keterangan as string) ?? "");
      }
      setLoadDetail(false);
    });
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!nip.trim() || !namaLengkap.trim() || !satkerId.trim()) {
      setError("NIP, Nama Lengkap, dan Satker ID wajib diisi");
      return;
    }
    setLoading(true);
    const res = await apiService.put(`v1/wp-data/${id}`, {
      nip: nip.trim(),
      nama_lengkap: namaLengkap.trim(),
      jenis_kelamin: jenisKelamin || undefined,
      golongan_ruang: golonganRuang || undefined,
      pangkat: pangkat || undefined,
      jenjang_jabatan_fungsional: jenjangJabatanFungsional || undefined,
      satker_id: satkerId.trim(),
      unit_kerja: unitKerja || undefined,
      pendidikan_terakhir: pendidikanTerakhir || undefined,
      tmt_golongan: tmtGolongan || undefined,
      tmt_jabatan_fungsional: tmtJabatanFungsional || undefined,
      no_sk: noSk || undefined,
      no_hp: noHp || undefined,
      email: email || undefined,
      alamat: alamat || undefined,
      status: status || "Aktif",
      keterangan: keterangan || undefined,
    });
    setLoading(false);
    if (res.success) router.push(`/wpdata/${id}`);
    else setError(res.message ?? "Gagal menyimpan");
  };

  if (loadDetail) return <div className="animate-pulse h-32 bg-gray-100 rounded" />;

  return (
    <div className="space-y-6">
      <Link href={`/wpdata/${id}`} className="text-blue-600 hover:underline text-sm">← Kembali</Link>
      <h1 className="text-2xl font-bold text-gray-900">Edit Data WP</h1>
      <form onSubmit={handleSubmit} className="max-w-2xl space-y-4">
        {error && <p className="text-red-600 text-sm">{error}</p>}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">NIP *</label>
            <input type="text" value={nip} onChange={(e) => setNip(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" required />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Nama Lengkap *</label>
            <input type="text" value={namaLengkap} onChange={(e) => setNamaLengkap(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" required />
          </div>
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Jenis Kelamin</label>
            <select value={jenisKelamin} onChange={(e) => setJenisKelamin(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm">
              <option value="">-</option>
              <option value="L">Laki-laki</option>
              <option value="P">Perempuan</option>
            </select>
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Satker ID *</label>
            <input type="text" value={satkerId} onChange={(e) => setSatkerId(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" required />
          </div>
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Unit Kerja</label>
            <input type="text" value={unitKerja} onChange={(e) => setUnitKerja(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Status</label>
            <select value={status} onChange={(e) => setStatus(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm">
              <option value="Aktif">Aktif</option>
              <option value="Nonaktif">Nonaktif</option>
            </select>
          </div>
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">No. HP</label>
            <input type="text" value={noHp} onChange={(e) => setNoHp(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Email</label>
            <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
          </div>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Alamat</label>
          <textarea value={alamat} onChange={(e) => setAlamat(e.target.value)} rows={2} className="w-full rounded-lg border border-gray-300 px-3 py-2 text-sm" />
        </div>
        <div className="flex gap-3">
          <button type="submit" disabled={loading} className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700 disabled:opacity-50">{loading ? "Menyimpan..." : "Simpan"}</button>
          <Link href={`/wpdata/${id}`} className="rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-50">Batal</Link>
        </div>
      </form>
    </div>
  );
}
