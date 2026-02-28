"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

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

function toISO_local(d: string): string {
  try {
    const date = new Date(d);
    return date.toISOString().slice(0, 16);
  } catch {
    return "";
  }
}

export default function UjianEditPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
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
  const [loading, setLoading] = useState(true);
  const [submitLoading, setSubmitLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService.get("v1/exams/" + id).then((res) => {
      if (res.success && res.data) {
        const e = res.data as Record<string, unknown>;
        setCode(String(e.code ?? ""));
        setName(String(e.name ?? ""));
        setJadwalMulai(toISO_local(String(e.jadwal_mulai ?? "")));
        setJadwalSelesai(toISO_local(String(e.jadwal_selesai ?? "")));
        setDurasiMenit(Number(e.durasi_menit) || 60);
        setShuffleQuestions(Boolean(e.shuffle_questions));
        setTampilkanLeaderboard(Boolean(e.tampilkan_leaderboard));
        const contents = (e.contents ?? []) as { source_type: string; source_id: string }[];
        const qids: string[] = [];
        const pids: string[] = [];
        contents.forEach((c) => {
          if (c.source_type === "question") qids.push(c.source_id);
          else if (c.source_type === "package") pids.push(c.source_id);
        });
        setQuestionIds(qids);
        setPackageIds(pids);
        const parts = (e.participants ?? []) as { user_id: string }[];
        setParticipantIds(parts.map((p) => p.user_id));
      }
      setLoading(false);
    });
  }, [id]);

  useEffect(() => {
    apiService.get<{ items: QuestionListItem[] }>("v1/questions", { page_size: 500 }).then((res) => {
      if (res.success && res.data) setQuestions((res.data as { items: QuestionListItem[] }).items ?? []);
    });
    apiService.get<{ items: PackageListItem[] }>("v1/question-packages", { page_size: 100 }).then((res) => {
      if (res.success && res.data) setPackages((res.data as { items: PackageListItem[] }).items ?? []);
    });
    apiService
      .get<{ items: CalonItem[] }>("v1/wp-data/calon-peserta", { status_verifikasi: "ujikom_lolos", page_size: 500 })
      .then((res) => {
        if (res.success && res.data) setCalonLolos((res.data as { items: CalonItem[] }).items ?? []);
      });
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (questionIds.length === 0 && packageIds.length === 0) {
      setError("Minimal 1 soal");
      return;
    }
    if (participantIds.length === 0) {
      setError("Minimal 1 peserta");
      return;
    }
    setSubmitLoading(true);
    const res = await apiService.put("v1/exams/" + id, {
      code: code.trim(),
      name: name.trim(),
      jadwal_mulai: new Date(jadwalMulai).toISOString(),
      jadwal_selesai: new Date(jadwalSelesai).toISOString(),
      durasi_menit: durasiMenit,
      question_ids: questionIds,
      package_ids: packageIds,
      participant_ids: participantIds,
      shuffle_questions: shuffleQuestions,
      tampilkan_leaderboard: tampilkanLeaderboard,
    });
    if (res.success) {
      router.push(`/wpujikom/uji-kompetensi/${id}`);
      return;
    }
    setError(res.message ?? "Gagal memperbarui");
    setSubmitLoading(false);
  };

  const toggleQuestion = (qid: string) => {
    setQuestionIds((p) => (p.includes(qid) ? p.filter((x) => x !== qid) : [...p, qid]));
  };
  const togglePackage = (pid: string) => {
    setPackageIds((p) => (p.includes(pid) ? p.filter((x) => x !== pid) : [...p, pid]));
  };
  const toggleParticipant = (uid: string) => {
    setParticipantIds((p) => (p.includes(uid) ? p.filter((x) => x !== uid) : [...p, uid]));
  };

  if (loading) return <div className="p-8">Memuat...</div>;

  return (
    <div className="space-y-6 max-w-2xl">
      <Link href={`/wpujikom/uji-kompetensi/${id}`} className="text-gray-600 hover:text-gray-900">
        ← Kembali
      </Link>
      <h1 className="text-2xl font-bold">Edit Ujian</h1>

      <form onSubmit={handleSubmit} className="space-y-5 bg-white rounded-xl shadow-sm border p-6">
        {error && <div className="p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Kode *</label>
          <input
            type="text"
            value={code}
            onChange={(e) => setCode(e.target.value)}
            className="w-full px-3 py-2 border rounded-lg"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Nama *</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-full px-3 py-2 border rounded-lg"
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
              className="w-full px-3 py-2 border rounded-lg"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Jadwal Selesai *</label>
            <input
              type="datetime-local"
              value={jadwalSelesai}
              onChange={(e) => setJadwalSelesai(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg"
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
            className="w-full px-3 py-2 border rounded-lg"
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
          <label className="block text-sm font-medium text-gray-700 mb-2">Peserta *</label>
          <div className="max-h-40 overflow-y-auto border rounded p-2 space-y-1">
            {calonLolos.map((c) => (
              <label key={c.id} className="flex items-center gap-2 text-sm">
                <input
                  type="checkbox"
                  checked={participantIds.includes(c.user_id)}
                  onChange={() => toggleParticipant(c.user_id)}
                />
                {c.user_name}
              </label>
            ))}
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
            disabled={submitLoading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
          >
            {submitLoading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link href={`/wpujikom/uji-kompetensi/${id}`} className="px-4 py-2 border rounded-lg hover:bg-gray-50">
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
