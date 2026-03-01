"use client";

import RichTextEditor from "@/components/molecules/RichTextEditor";
import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
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

interface PackageDetail {
  id: string;
  code: string;
  name: string;
  description: string;
  status: string;
  questions: { question_id: string; question_code: string; sort_order: number }[];
}

export default function PaketSoalEditPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [code, setCode] = useState("");
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [status, setStatus] = useState("Draft");
  const [questionIds, setQuestionIds] = useState<string[]>([]);
  const [questions, setQuestions] = useState<QuestionListItem[]>([]);
  const [qSearch, setQSearch] = useState("");
  const [loading, setLoading] = useState(true);
  const [submitLoading, setSubmitLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService.get<PackageDetail>("v1/question-packages/" + id).then((res) => {
      if (res.success && res.data) {
        const p = res.data as PackageDetail;
        setCode(p.code ?? "");
        setName(p.name ?? "");
        setDescription(p.description ?? "");
        setStatus(p.status ?? "Draft");
        const ids = (p.questions ?? [])
          .sort((a, b) => a.sort_order - b.sort_order)
          .map((q) => q.question_id);
        setQuestionIds(ids);
      }
      setLoading(false);
    });
  }, [id]);

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
    if (!code.trim() || !name.trim()) {
      setError("Kode dan nama wajib diisi");
      return;
    }
    if (questionIds.length === 0) {
      setError("Minimal 1 soal dalam paket");
      return;
    }

    setSubmitLoading(true);
    const res = await apiService.put("v1/question-packages/" + id, {
      code: code.trim(),
      name: name.trim(),
      description: description?.replace(/<[^>]*>/g, "").trim() ? description : undefined,
      status,
      question_ids: questionIds,
    });
    if (res.success) {
      router.push(`/wpujikom/paket-soal/${id}`);
      return;
    }
    setError(res.message ?? "Gagal memperbarui paket soal");
    setSubmitLoading(false);
  };

  const toggleQuestion = (qid: string) => {
    setQuestionIds((prev) =>
      prev.includes(qid) ? prev.filter((q) => q !== qid) : [...prev, qid]
    );
  };

  const moveUp = (idx: number) => {
    if (idx <= 0) return;
    const qid = questionIds[idx];
    const newIds = [...questionIds];
    newIds.splice(idx, 1);
    newIds.splice(idx - 1, 0, qid);
    setQuestionIds(newIds);
  };

  const moveDown = (idx: number) => {
    if (idx >= questionIds.length - 1) return;
    const qid = questionIds[idx];
    const newIds = [...questionIds];
    newIds.splice(idx, 1);
    newIds.splice(idx + 1, 0, qid);
    setQuestionIds(newIds);
  };

  const getOrderedQuestions = () => {
    const byId = new Map(questions.map((q) => [q.id, q]));
    return questionIds
      .map((qid) => byId.get(qid))
      .filter(Boolean) as QuestionListItem[];
  };

  const availableToAdd = questions.filter((q) => !questionIds.includes(q.id));
  const orderedSelected = getOrderedQuestions();

  return (
    <div className="space-y-6">
      <Link href={`/wpujikom/paket-soal/${id}`} className="text-gray-600 hover:text-gray-900">
        ← Kembali
      </Link>
      <h1 className="text-2xl font-bold text-gray-900">Edit Paket Soal</h1>

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
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Deskripsi</label>
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
          <label className="block text-sm font-medium text-gray-700 mb-2">Soal dalam paket (urutkan)</label>
          <div className="space-y-2 border border-gray-200 rounded-lg p-4">
            {loading ? (
              <p className="text-gray-500">Memuat...</p>
            ) : orderedSelected.length === 0 ? (
              <p className="text-gray-500">Belum ada soal. Tambah dari daftar di bawah.</p>
            ) : (
              orderedSelected.map((q, i) => (
                <div
                  key={q.id}
                  className="flex items-center gap-2 p-2 rounded bg-gray-50"
                >
                  <span className="text-gray-500 w-6">{i + 1}.</span>
                  <span className="font-mono flex-1">{q.code}</span>
                  <button
                    type="button"
                    onClick={() => moveUp(i)}
                    disabled={i === 0}
                    className="text-gray-500 hover:text-gray-700 disabled:opacity-30"
                  >
                    ↑
                  </button>
                  <button
                    type="button"
                    onClick={() => moveDown(i)}
                    disabled={i === orderedSelected.length - 1}
                    className="text-gray-500 hover:text-gray-700 disabled:opacity-30"
                  >
                    ↓
                  </button>
                  <button
                    type="button"
                    onClick={() => toggleQuestion(q.id)}
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
          <label className="block text-sm font-medium text-gray-700 mb-2">Tambah soal</label>
          <input
            type="text"
            value={qSearch}
            onChange={(e) => setQSearch(e.target.value)}
            placeholder="Cari soal..."
            className="w-full mb-2 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          />
          <div className="border border-gray-200 rounded-lg p-4 max-h-48 overflow-y-auto space-y-2">
            {availableToAdd.length === 0 ? (
              <p className="text-gray-500 text-sm">
                {qSearch ? "Tidak ada soal yang cocok." : "Semua soal sudah ditambahkan."}
              </p>
            ) : (
              availableToAdd.map((q) => (
                <button
                  key={q.id}
                  type="button"
                  onClick={() => toggleQuestion(q.id)}
                  className="w-full flex items-center gap-2 p-2 rounded transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50 text-left"
                >
                  <span className="font-mono text-sm">{q.code}</span>
                  <span className="text-gray-500 text-sm">({q.type})</span>
                  <span className="text-sm text-gray-600 truncate flex-1">{q.question_text?.replace(/<[^>]*>/g, "") ?? ""}</span>
                  <span className="text-blue-600 text-sm">+ Tambah</span>
                </button>
              ))
            )}
          </div>
        </div>
        <div className="flex gap-2">
          <button
            type="submit"
            disabled={submitLoading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 disabled:opacity-50"
          >
            {submitLoading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link
            href={`/wpujikom/paket-soal/${id}`}
            className="px-4 py-2 border border-gray-300 rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
