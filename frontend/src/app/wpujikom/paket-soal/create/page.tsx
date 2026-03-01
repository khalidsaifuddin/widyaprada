"use client";

import RichTextEditor from "@/components/molecules/RichTextEditor";
import { apiService } from "@/lib/api";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface QuestionListItem {
  id: string;
  code: string;
  type: string;
  question_text: string;
}

interface QuestionListResponse {
  items: QuestionListItem[];
}

export default function PaketSoalCreatePage() {
  const router = useRouter();
  const [code, setCode] = useState("");
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [status, setStatus] = useState("Draft");
  const [questionIds, setQuestionIds] = useState<string[]>([]);
  const [questions, setQuestions] = useState<QuestionListItem[]>([]);
  const [qSearch, setQSearch] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService
      .get<QuestionListResponse>("v1/questions", { page_size: 200, q: qSearch || undefined })
      .then((res) => {
        if (res.success && res.data) {
          const d = res.data as QuestionListResponse;
          setQuestions(d.items ?? []);
        }
      })
      .catch(() => {});
  }, [qSearch]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!code.trim()) {
      setError("Kode wajib diisi");
      return;
    }
    if (!name.trim()) {
      setError("Nama wajib diisi");
      return;
    }
    if (questionIds.length === 0) {
      setError("Minimal 1 soal dalam paket");
      return;
    }

    setLoading(true);
    const res = await apiService.post("v1/question-packages", {
      code: code.trim(),
      name: name.trim(),
      description: description?.replace(/<[^>]*>/g, "").trim() ? description : undefined,
      status,
      question_ids: questionIds,
    });
    if (res.success) {
      router.push("/wpujikom/paket-soal");
      return;
    }
    setError(res.message ?? "Gagal menambah paket soal");
    setLoading(false);
  };

  const toggleQuestion = (id: string) => {
    setQuestionIds((prev) =>
      prev.includes(id) ? prev.filter((q) => q !== id) : [...prev, id]
    );
  };

  const selectAll = () => {
    if (questionIds.length === questions.length) {
      setQuestionIds([]);
    } else {
      setQuestionIds(questions.map((q) => q.id));
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <Link href="/wpujikom/paket-soal" className="text-gray-600 hover:text-gray-900">
          ← Kembali
        </Link>
        <h1 className="text-2xl font-bold text-gray-900">Tambah Paket Soal</h1>
      </div>

      <form
        onSubmit={handleSubmit}
        className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 max-w-2xl space-y-5"
      >
        {error && (
          <div className="p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>
        )}
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Kode *</label>
          <input
            type="text"
            value={code}
            onChange={(e) => setCode(e.target.value)}
            placeholder="contoh: PKG-001"
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Nama *</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Nama paket soal"
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Deskripsi (opsional)</label>
          <RichTextEditor
            value={description}
            onChange={setDescription}
            placeholder="Tulis deskripsi paket soal..."
            minHeight="8rem"
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Status</label>
          <select
            value={status}
            onChange={(e) => setStatus(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          >
            <option value="Draft">Draft</option>
            <option value="Aktif">Aktif</option>
          </select>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-2">Pilih Soal (min. 1) *</label>
          <input
            type="text"
            value={qSearch}
            onChange={(e) => setQSearch(e.target.value)}
            placeholder="Cari soal..."
            className="w-full mb-2 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          />
          <button
            type="button"
            onClick={selectAll}
            className="text-sm text-blue-600 hover:text-blue-800 mb-2"
          >
            {questionIds.length === questions.length ? "Batal pilih semua" : "Pilih semua"}
          </button>
          <div className="border border-gray-200 rounded-lg p-4 max-h-64 overflow-y-auto space-y-2">
            {questions.length === 0 ? (
              <p className="text-gray-500 text-sm">Tidak ada soal. Tambah soal di Bank Soal terlebih dahulu.</p>
            ) : (
              questions.map((q) => (
                <label
                  key={q.id}
                  className="flex items-start gap-2 p-2 rounded transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50 cursor-pointer"
                >
                  <input
                    type="checkbox"
                    checked={questionIds.includes(q.id)}
                    onChange={() => toggleQuestion(q.id)}
                    className="mt-1 rounded border-gray-300"
                  />
                  <div className="flex-1 min-w-0">
                    <span className="font-mono text-sm">{q.code}</span>
                    <span className="text-gray-500 text-sm ml-2">({q.type})</span>
                    <p className="text-sm text-gray-600 truncate">{q.question_text?.replace(/<[^>]*>/g, "") ?? ""}</p>
                  </div>
                </label>
              ))
            )}
          </div>
          <p className="text-xs text-gray-500 mt-1">Terpilih: {questionIds.length} soal</p>
        </div>
        <div className="flex gap-2">
          <button
            type="submit"
            disabled={loading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 disabled:opacity-50"
          >
            {loading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link
            href="/wpujikom/paket-soal"
            className="px-4 py-2 border border-gray-300 rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
