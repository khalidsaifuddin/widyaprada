"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useRouter } from "next/navigation";
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

export default function BankSoalCreatePage() {
  const router = useRouter();
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
  const [bsCorrect, setBsCorrect] = useState(true); // BENAR_SALAH: true = Benar, false = Salah
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    apiService.get<Category[]>("v1/questions/categories").then((res) => {
      if (res.success && res.data) {
        const raw = res.data;
        const arr = Array.isArray(raw) ? raw : [];
        setCategories(arr);
        if (arr.length && !categoryId) setCategoryId(arr[0].id);
      }
    });
  }, []);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    if (!code.trim()) {
      setError("Kode wajib diisi");
      return;
    }
    if (!questionText.trim()) {
      setError("Teks soal wajib diisi");
      return;
    }
    if (type === "PG" || type === "BENAR_SALAH") {
      const opts = type === "PG" ? options : [
        { option_key: "BENAR", option_text: "Benar", is_correct: bsCorrect },
        { option_key: "SALAH", option_text: "Salah", is_correct: !bsCorrect },
      ];
      const key = opts.find((o) => o.is_correct)?.option_key;
      if (!key) {
        setError("Pilih kunci jawaban");
        return;
      }
      if (type === "PG" && opts.some((o) => !o.option_text?.trim())) {
        setError("Semua opsi PG wajib diisi");
        return;
      }
    }

    setLoading(true);
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

    const res = await apiService.post("v1/questions", {
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
      router.push("/wpujikom/bank-soal");
      return;
    }
    setError(res.message ?? "Gagal menambah soal");
    setLoading(false);
  };

  const setOptionCorrect = (idx: number) => {
    setOptions((prev) =>
      prev.map((o, i) => ({ ...o, is_correct: i === idx }))
    );
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center gap-4">
        <Link href="/wpujikom/bank-soal" className="text-gray-600 hover:text-gray-900">
          ← Kembali
        </Link>
        <h1 className="text-2xl font-bold text-gray-900">Tambah Soal</h1>
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
            placeholder="contoh: PG-001"
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
            <label className="block text-sm font-medium text-gray-700 mb-2">Opsi Jawaban (pilih kunci)</label>
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
                    placeholder={`Opsi ${o.option_key}`}
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
            <label className="block text-sm font-medium text-gray-700 mb-1">Model Jawaban / Rubrik</label>
            <input
              type="text"
              value={answerKey}
              onChange={(e) => setAnswerKey(e.target.value)}
              placeholder="Petunjuk penilaian atau model jawaban"
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
            disabled={loading}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 disabled:opacity-50"
          >
            {loading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link
            href="/wpujikom/bank-soal"
            className="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
