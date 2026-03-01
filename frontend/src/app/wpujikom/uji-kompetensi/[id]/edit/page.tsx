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

type ContentItem = { sourceType: "question" | "package"; sourceId: string; label: string };

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
  const [contents, setContents] = useState<ContentItem[]>([]);
  const [participantIds, setParticipantIds] = useState<string[]>([]);
  const [shuffleQuestions, setShuffleQuestions] = useState(false);
  const [tampilkanLeaderboard, setTampilkanLeaderboard] = useState(false);
  const [questions, setQuestions] = useState<QuestionListItem[]>([]);
  const [packages, setPackages] = useState<PackageListItem[]>([]);
  const [calonLolos, setCalonLolos] = useState<CalonItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [submitLoading, setSubmitLoading] = useState(false);
  const [error, setError] = useState("");
  const [draggedIdx, setDraggedIdx] = useState<number | null>(null);

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
        const contentsData = (e.contents ?? []) as { source_type: string; source_id: string; sort_order?: number }[];
        const sorted = [...contentsData].sort((a, b) => (a.sort_order ?? 0) - (b.sort_order ?? 0));
        const loaded: ContentItem[] = sorted.map((c) => ({
          sourceType: c.source_type === "package" ? "package" : "question",
          sourceId: c.source_id,
          label: c.source_id, // will be replaced when questions/packages load
        }));
        setContents(loaded);
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

  // Update labels when questions/packages loaded
  useEffect(() => {
    if (questions.length === 0 && packages.length === 0) return;
    setContents((prev) =>
      prev.map((c) => {
        if (c.sourceType === "question") {
          const q = questions.find((x) => x.id === c.sourceId);
          return { ...c, label: q?.code ?? c.sourceId };
        }
        const p = packages.find((x) => x.id === c.sourceId);
        return { ...c, label: p?.name || p?.code ?? c.sourceId };
      })
    );
  }, [questions, packages]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (contents.length === 0) {
      setError("Minimal 1 soal");
      return;
    }
    setSubmitLoading(true);
    const res = await apiService.put("v1/exams/" + id, {
      code: code.trim(),
      name: name.trim(),
      jadwal_mulai: new Date(jadwalMulai).toISOString(),
      jadwal_selesai: new Date(jadwalSelesai).toISOString(),
      durasi_menit: durasiMenit,
      contents: contents.map((c) => ({
        source_type: c.sourceType,
        source_id: c.sourceId,
      })),
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

  const addContent = (val: string) => {
    if (!val) return;
    const [type, sourceId] = val.split(":");
    if (type === "q") {
      const q = questions.find((x) => x.id === sourceId);
      if (q && !contents.some((c) => c.sourceType === "question" && c.sourceId === sourceId)) {
        setContents((prev) => [...prev, { sourceType: "question", sourceId, label: q.code }]);
      }
    } else if (type === "p") {
      const p = packages.find((x) => x.id === sourceId);
      if (p && !contents.some((c) => c.sourceType === "package" && c.sourceId === sourceId)) {
        setContents((prev) => [...prev, { sourceType: "package", sourceId, label: p.name || p.code }]);
      }
    }
  };

  const removeContent = (idx: number) => {
    setContents((prev) => prev.filter((_, i) => i !== idx));
  };

  const moveContent = (fromIdx: number, toIdx: number) => {
    if (fromIdx === toIdx) return;
    setContents((prev) => {
      const arr = [...prev];
      const [moved] = arr.splice(fromIdx, 1);
      arr.splice(toIdx, 0, moved);
      return arr;
    });
  };

  const handleDragStart = (e: React.DragEvent, idx: number) => {
    setDraggedIdx(idx);
    e.dataTransfer.effectAllowed = "move";
    e.dataTransfer.setData("text/plain", String(idx));
  };
  const handleDragOver = (e: React.DragEvent) => e.preventDefault();
  const handleDrop = (e: React.DragEvent, toIdx: number) => {
    e.preventDefault();
    if (draggedIdx === null) return;
    if (draggedIdx !== toIdx) moveContent(draggedIdx, toIdx);
    setDraggedIdx(null);
  };
  const handleDragEnd = () => setDraggedIdx(null);

  const GripIcon = () => (
    <svg className="w-5 h-5 text-gray-400 cursor-grab active:cursor-grabbing" viewBox="0 0 24 24" fill="currentColor">
      <circle cx="9" cy="6" r="1.5" />
      <circle cx="9" cy="12" r="1.5" />
      <circle cx="9" cy="18" r="1.5" />
      <circle cx="15" cy="6" r="1.5" />
      <circle cx="15" cy="12" r="1.5" />
      <circle cx="15" cy="18" r="1.5" />
    </svg>
  );

  const addedQuestionIds = new Set(contents.filter((c) => c.sourceType === "question").map((c) => c.sourceId));
  const addedPackageIds = new Set(contents.filter((c) => c.sourceType === "package").map((c) => c.sourceId));
  const availableQuestions = questions.filter((q) => !addedQuestionIds.has(q.id));
  const availablePackages = packages.filter((p) => !addedPackageIds.has(p.id));

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
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Jadwal Selesai *</label>
            <input
              type="datetime-local"
              value={jadwalSelesai}
              onChange={(e) => setJadwalSelesai(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500"
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
          <div className="flex items-center justify-between mb-2">
            <label className="block text-sm font-medium text-gray-700">Soal / Paket (urutkan ke bawah)</label>
            <select
              value=""
              onChange={(e) => {
                addContent(e.target.value);
                e.target.value = "";
              }}
              className="px-3 py-2 border rounded-lg focus:ring-2 focus:ring-blue-500 text-sm"
            >
              <option value="">+ Tambah soal atau paket</option>
              {availableQuestions.length > 0 && (
                <optgroup label="Soal Individu">
                  {availableQuestions.map((q) => (
                    <option key={q.id} value={`q:${q.id}`}>
                      {q.code}
                    </option>
                  ))}
                </optgroup>
              )}
              {availablePackages.length > 0 && (
                <optgroup label="Paket Soal">
                  {availablePackages.map((p) => (
                    <option key={p.id} value={`p:${p.id}`}>
                      {p.name || p.code}
                    </option>
                  ))}
                </optgroup>
              )}
              {availableQuestions.length === 0 && availablePackages.length === 0 && contents.length > 0 && (
                <option disabled>Semua sudah ditambahkan</option>
              )}
            </select>
          </div>
          <p className="text-xs text-gray-500 mb-2">Geser untuk mengubah urutan</p>
          <div className="space-y-2 border rounded-lg p-2 min-h-[80px]">
            {contents.length === 0 ? (
              <p className="text-sm text-gray-500 py-4 text-center">Belum ada soal. Pilih dari dropdown di atas.</p>
            ) : (
              contents.map((item, i) => (
                <div
                  key={`${item.sourceType}-${item.sourceId}-${i}`}
                  draggable
                  onDragStart={(e) => handleDragStart(e, i)}
                  onDragOver={handleDragOver}
                  onDrop={(e) => handleDrop(e, i)}
                  onDragEnd={handleDragEnd}
                  className={`flex items-center gap-2 p-2 rounded-lg border transition-colors ${
                    draggedIdx === i
                      ? "border-blue-400 bg-blue-50 opacity-75"
                      : "border-gray-200 bg-white hover:border-gray-300"
                  }`}
                >
                  <div className="flex-shrink-0">
                    <GripIcon />
                  </div>
                  <span className="text-xs px-2 py-0.5 rounded bg-gray-100 text-gray-600">
                    {item.sourceType === "question" ? "Soal" : "Paket"}
                  </span>
                  <span className="flex-1 font-medium">{item.label}</span>
                  <button
                    type="button"
                    onClick={() => removeContent(i)}
                    className="text-red-600 hover:text-red-800 text-sm"
                  >
                    Hapus
                  </button>
                </div>
              ))
            )}
          </div>
        </div>

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">Peserta (opsional)</label>
          <div className="max-h-40 overflow-y-auto border rounded p-2 space-y-1">
            {calonLolos.map((c) => (
              <label key={c.id} className="flex items-center gap-2 text-sm">
                <input
                  type="checkbox"
                  checked={participantIds.includes(c.user_id)}
                  onChange={() =>
                    setParticipantIds((p) =>
                      p.includes(c.user_id) ? p.filter((x) => x !== c.user_id) : [...p, c.user_id]
                    )
                  }
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
            className="px-4 py-2 bg-blue-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 disabled:opacity-50"
          >
            {submitLoading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link href={`/wpujikom/uji-kompetensi/${id}`} className="px-4 py-2 border rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50">
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
