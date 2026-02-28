"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

interface QuestionListItem {
  id: string;
  code: string;
}

interface PackageListItem {
  id: string;
  code: string;
  name: string;
}

interface CalonItem {
  id: string;
  user_id: string;
  user_name: string;
  status_kode: string;
}

function toISO_local(d: Date): string {
  return d.toISOString().slice(0, 16);
}

export default function UjianCreatePage() {
  const router = useRouter();
  const [code, setCode] = useState("");
  const [name, setName] = useState("");
  const [jadwalMulai, setJadwalMulai] = useState("");
  const [jadwalSelesai, setJadwalSelesai] = useState("");
  const [durasiMenit, setDurasiMenit] = useState(60);
  const [questionIds, setQuestionIds] = useState<string[]>([]);
  const [packageIds, setPackageIds] = useState<string[]>([]);
  const [participantIds, setParticipantIds] = useState<string[]>([]);
  const [shuffleQuestions, setShuffleQuestions] = useState(false);
  const [tampilkanLeaderboard, setTampilkanLeaderboard] = useState(false);
  const [questions, setQuestions] = useState<QuestionListItem[]>([]);
  const [packages, setPackages] = useState<PackageListItem[]>([]);
  const [calonLolos, setCalonLolos] = useState<CalonItem[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    const now = new Date();
    setJadwalMulai(toISO_local(now));
    setJadwalSelesai(toISO_local(new Date(now.getTime() + 2 * 60 * 60 * 1000)));
  }, []);

  useEffect(() => {
    apiService.get<{ items: QuestionListItem[] }>("v1/questions", { page_size: 500 }).then((res) => {
      if (res.success && res.data) {
        const d = res.data as { items: QuestionListItem[] };
        setQuestions(d.items ?? []);
      }
    });
  }, []);

  useEffect(() => {
    apiService.get<{ items: PackageListItem[] }>("v1/question-packages", { page_size: 100 }).then((res) => {
      if (res.success && res.data) {
        const d = res.data as { items: PackageListItem[] };
        setPackages(d.items ?? []);
      }
    });
  }, []);

  useEffect(() => {
    apiService
      .get<{ items: CalonItem[] }>("v1/wp-data/calon-peserta", {
        status_verifikasi: "ujikom_lolos",
        page_size: 500,
      })
      .then((res) => {
        if (res.success && res.data) {
          const d = res.data as { items: CalonItem[] };
          setCalonLolos(d.items ?? []);
        }
      });
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!code.trim() || !name.trim() || !jadwalMulai || !jadwalSelesai) {
      setError("Kode, nama, dan jadwal wajib diisi");
      return;
    }
    if (questionIds.length === 0 && packageIds.length === 0) {
      setError("Minimal 1 soal (individu atau paket)");
      return;
    }
    if (participantIds.length === 0) {
      setError("Minimal 1 peserta (pilih dari calon yang lolos verifikasi)");
      return;
    }
    if (durasiMenit < 1) {
      setError("Durasi minimal 1 menit");
      return;
    }

    const jadwalMulaiISO = new Date(jadwalMulai).toISOString();
    const jadwalSelesaiISO = new Date(jadwalSelesai).toISOString();

    setLoading(true);
    const res = await apiService.post("v1/exams", {
      code: code.trim(),
      name: name.trim(),
      jadwal_mulai: jadwalMulaiISO,
      jadwal_selesai: jadwalSelesaiISO,
      durasi_menit: durasiMenit,
      question_ids: questionIds,
      package_ids: packageIds,
      participant_ids: participantIds,
      shuffle_questions: shuffleQuestions,
      tampilkan_leaderboard: tampilkanLeaderboard,
    });
    if (res.success) {
      router.push("/wpujikom/uji-kompetensi");
      return;
    }
    setError(res.message ?? "Gagal membuat ujian");
    setLoading(false);
  };

  const toggleQuestion = (id: string) => {
    setQuestionIds((p) => (p.includes(id) ? p.filter((x) => x !== id) : [...p, id]));
  };
  const togglePackage = (id: string) => {
    setPackageIds((p) => (p.includes(id) ? p.filter((x) => x !== id) : [...p, id]));
  };
  const toggleParticipant = (userId: string) => {
    setParticipantIds((p) => (p.includes(userId) ? p.filter((x) => x !== userId) : [...p, userId]));
  };

  return (
    <div className="space-y-6 max-w-2xl">
      <Link href="/wpujikom/uji-kompetensi" className="text-gray-600 hover:text-gray-900">
        ← Kembali
      </Link>
      <h1 className="text-2xl font-bold text-gray-900">Tambah Ujian</h1>

      <form onSubmit={handleSubmit} className="space-y-5 bg-white rounded-xl shadow-sm border p-6">
        {error && <div className="p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>}

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Kode *</label>
          <input
            type="text"
            value={code}
            onChange={(e) => setCode(e.target.value)}
            className="w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Nama *</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>
        <div className="grid grid-cols-2 gap-4">
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Jadwal Mulai *</label>
            <input
              type="datetime-local"
              value={jadwalMulai}
              onChange={(e) => setJadwalMulai(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500"
              required
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Jadwal Selesai *</label>
            <input
              type="datetime-local"
              value={jadwalSelesai}
              onChange={(e) => setJadwalSelesai(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500"
              required
            />
          </div>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Durasi (menit) *</label>
          <input
            type="number"
            value={durasiMenit}
            onChange={(e) => setDurasiMenit(Number(e.target.value) || 60)}
            min={1}
            className="w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500"
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">Soal Individu</label>
          <div className="flex flex-wrap gap-2 max-h-32 overflow-y-auto border rounded p-2">
            {questions.map((q) => (
              <label key={q.id} className="flex items-center gap-1 text-sm">
                <input
                  type="checkbox"
                  checked={questionIds.includes(q.id)}
                  onChange={() => toggleQuestion(q.id)}
                />
                {q.code}
              </label>
            ))}
          </div>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">Paket Soal</label>
          <div className="flex flex-wrap gap-2 max-h-32 overflow-y-auto border rounded p-2">
            {packages.map((p) => (
              <label key={p.id} className="flex items-center gap-1 text-sm">
                <input
                  type="checkbox"
                  checked={packageIds.includes(p.id)}
                  onChange={() => togglePackage(p.id)}
                />
                {p.name}
              </label>
            ))}
          </div>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">Peserta (calon lolos) *</label>
          <div className="max-h-40 overflow-y-auto border rounded p-2 space-y-1">
            {calonLolos.length === 0 ? (
              <p className="text-sm text-amber-600">Tidak ada calon lolos. Verifikasi calon peserta terlebih dahulu.</p>
            ) : (
              calonLolos.map((c) => (
                <label key={c.id} className="flex items-center gap-2 text-sm">
                  <input
                    type="checkbox"
                    checked={participantIds.includes(c.user_id)}
                    onChange={() => toggleParticipant(c.user_id)}
                  />
                  {c.user_name} ({c.user_id})
                </label>
              ))
            )}
          </div>
        </div>
        <div className="flex gap-4">
          <label className="flex items-center gap-2">
            <input
              type="checkbox"
              checked={shuffleQuestions}
              onChange={(e) => setShuffleQuestions(e.target.checked)}
            />
            Acak urutan soal
          </label>
          <label className="flex items-center gap-2">
            <input
              type="checkbox"
              checked={tampilkanLeaderboard}
              onChange={(e) => setTampilkanLeaderboard(e.target.checked)}
            />
            Tampilkan Leaderboard
          </label>
        </div>
        <div className="flex gap-2">
          <button
            type="submit"
            disabled={loading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
          >
            {loading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link href="/wpujikom/uji-kompetensi" className="px-4 py-2 border rounded-lg hover:bg-gray-50">
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
