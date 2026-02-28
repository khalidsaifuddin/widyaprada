"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

interface Category {
  id: string;
  code: string;
  name: string;
}

interface QuestionOptionInput {
  option_key: string;
  option_text: string;
  is_correct: boolean;
}

interface QuestionDetail {
  id: string;
  code: string;
  type: string;
  category_id: string;
  difficulty: string;
  question_text: string;
  answer_key: string;
  weight: number;
  status: string;
  options: { option_key: string; option_text: string; is_correct: boolean }[];
}

export default function BankSoalEditPage() {
  const params = useParams();
  const router = useRouter();
  const id = params.id as string;
  const [code, setCode] = useState("");
  const [type, setType] = useState("PG");
  const [categoryId, setCategoryId] = useState("");
  const [difficulty, setDifficulty] = useState("");
  const [questionText, setQuestionText] = useState("");
  const [answerKey, setAnswerKey] = useState("");
  const [weight, setWeight] = useState(1);
  const [status, setStatus] = useState("Draft");
  const [options, setOptions] = useState<QuestionOptionInput[]>([
    { option_key: "A", option_text: "", is_correct: false },
    { option_key: "B", option_text: "", is_correct: false },
    { option_key: "C", option_text: "", is_correct: false },
    { option_key: "D", option_text: "", is_correct: false },
  ]);
  const [bsCorrect, setBsCorrect] = useState(true);
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(true);
  const [submitLoading, setSubmitLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService.get<Category[]>("v1/questions/categories").then((res) => {
      if (res.success && res.data) {
        const raw = res.data;
        setCategories(Array.isArray(raw) ? raw : []);
      }
    });
  }, []);

  useEffect(() => {
    apiService.get<QuestionDetail>("v1/questions/" + id).then((res) => {
      if (res.success && res.data) {
        const q = res.data as QuestionDetail;
        setCode(q.code ?? "");
        setType(q.type ?? "PG");
        setCategoryId(q.category_id ?? "");
        setDifficulty(q.difficulty ?? "");
        setQuestionText(q.question_text ?? "");
        setAnswerKey(q.answer_key ?? "");
        setWeight(q.weight ?? 1);
        setStatus(q.status ?? "Draft");
        if (q.type === "BENAR_SALAH") {
          setBsCorrect(q.answer_key === "BENAR");
        } else if (q.type === "PG" && q.options?.length) {
          const pgOpts = ["A", "B", "C", "D"].map((key) => {
            const o = q.options.find((x) => x.option_key === key);
            return o
              ? { ...o, is_correct: o.is_correct }
              : { option_key: key, option_text: "", is_correct: false };
          });
          setOptions(pgOpts);
        }
      }
      setLoading(false);
    });
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!code.trim() || !questionText.trim()) {
      setError("Kode dan teks soal wajib diisi");
      return;
    }
    if (type === "PG") {
      const key = options.find((o) => o.is_correct)?.option_key;
      if (!key || options.some((o) => !o.option_text?.trim())) {
        setError("Pilih kunci dan isi semua opsi");
        return;
      }
    }

    setSubmitLoading(true);
    let opts: QuestionOptionInput[] = [];
    let key = answerKey;
    if (type === "PG") {
      opts = options;
      key = options.find((o) => o.is_correct)?.option_key ?? "";
    } else if (type === "BENAR_SALAH") {
      opts = [
        { option_key: "BENAR", option_text: "Benar", is_correct: bsCorrect },
        { option_key: "SALAH", option_text: "Salah", is_correct: !bsCorrect },
      ];
      key = bsCorrect ? "BENAR" : "SALAH";
    }

    const res = await apiService.put("v1/questions/" + id, {
      code: code.trim(),
      type,
      category_id: categoryId || undefined,
      difficulty: difficulty || undefined,
      question_text: questionText.trim(),
      answer_key: key,
      weight,
      status,
      options: opts,
    });
    if (res.success) {
      router.push(`/wpujikom/bank-soal/${id}`);
      return;
    }
    setError(res.message ?? "Gagal memperbarui soal");
    setSubmitLoading(false);
  };

  const setOptionCorrect = (idx: number) => {
    setOptions((prev) => prev.map((o, i) => ({ ...o, is_correct: i === idx })));
  };

  if (loading) return <div className="p-8">Memuat...</div>;

  return (
    <div className="space-y-6">
      <Link href={`/wpujikom/bank-soal/${id}`} className="text-gray-600 hover:text-gray-900">
        ← Kembali
      </Link>
      <h1 className="text-2xl font-bold text-gray-900">Edit Soal</h1>

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
          <label className="block text-sm font-medium text-gray-700 mb-1">Tipe *</label>
          <select
            value={type}
            onChange={(e) => setType(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          >
            <option value="PG">Pilihan Ganda</option>
            <option value="BENAR_SALAH">Benar-Salah</option>
            <option value="ESSAY">Essay</option>
          </select>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Kategori</label>
          <select
            value={categoryId}
            onChange={(e) => setCategoryId(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          >
            <option value="">-- Pilih --</option>
            {categories.map((c) => (
              <option key={c.id} value={c.id}>
                {c.name || c.code}
              </option>
            ))}
          </select>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Tingkat Kesulitan</label>
          <select
            value={difficulty}
            onChange={(e) => setDifficulty(e.target.value)}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
          >
            <option value="">-- Pilih --</option>
            <option value="Mudah">Mudah</option>
            <option value="Sedang">Sedang</option>
            <option value="Sulit">Sulit</option>
          </select>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Teks Soal *</label>
          <textarea
            value={questionText}
            onChange={(e) => setQuestionText(e.target.value)}
            rows={4}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>

        {type === "PG" && (
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">Opsi Jawaban</label>
            <div className="space-y-2">
              {options.map((o, i) => (
                <div key={i} className="flex gap-2 items-center">
                  <span className="w-6 font-medium">{o.option_key}.</span>
                  <input
                    type="text"
                    value={o.option_text}
                    onChange={(e) =>
                      setOptions((prev) =>
                        prev.map((opt, j) => (j === i ? { ...opt, option_text: e.target.value } : opt))
                      )
                    }
                    className="flex-1 px-3 py-2 border border-gray-300 rounded-lg"
                  />
                  <label className="flex items-center gap-1">
                    <input
                      type="radio"
                      name="correct"
                      checked={o.is_correct}
                      onChange={() => setOptionCorrect(i)}
                    />
                    <span className="text-sm">Kunci</span>
                  </label>
                </div>
              ))}
            </div>
          </div>
        )}

        {type === "BENAR_SALAH" && (
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">Kunci Jawaban</label>
            <div className="flex gap-4">
              <label className="flex items-center gap-2">
                <input
                  type="radio"
                  name="bs"
                  checked={bsCorrect}
                  onChange={() => setBsCorrect(true)}
                />
                Benar
              </label>
              <label className="flex items-center gap-2">
                <input
                  type="radio"
                  name="bs"
                  checked={!bsCorrect}
                  onChange={() => setBsCorrect(false)}
                />
                Salah
              </label>
            </div>
          </div>
        )}

        {type === "ESSAY" && (
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-1">Model Jawaban</label>
            <input
              type="text"
              value={answerKey}
              onChange={(e) => setAnswerKey(e.target.value)}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
            />
          </div>
        )}

        <div>
          <label className="block text-sm font-medium text-gray-700 mb-1">Bobot</label>
          <input
            type="number"
            value={weight}
            onChange={(e) => setWeight(Number(e.target.value) || 1)}
            min={0.5}
            step={0.5}
            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
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
        <div className="flex gap-2">
          <button
            type="submit"
            disabled={submitLoading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
          >
            {submitLoading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link
            href={`/wpujikom/bank-soal/${id}`}
            className="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
