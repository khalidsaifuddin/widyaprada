"use client";

import RichTextEditor from "@/components/molecules/RichTextEditor";
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
  option_weight?: number;
}

const OPTION_KEYS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ".split("");

function getNextOptionKey(current: string[]): string {
  const used = new Set(current);
  for (const k of OPTION_KEYS) {
    if (!used.has(k)) return k;
  }
  return String.fromCharCode(65 + current.length);
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
  const [bsCorrect, setBsCorrect] = useState(true); // BENAR_SALAH
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
    const stripped = questionText.replace(/<[^>]*>/g, "").trim();
    if (!stripped) {
      setError("Teks soal wajib diisi");
      return;
    }
    if (type === "PG" || type === "MRA" || type === "BENAR_SALAH") {
      const opts =
        type === "PG" || type === "MRA"
          ? options
          : [
              { option_key: "BENAR", option_text: "Benar", is_correct: bsCorrect },
              { option_key: "SALAH", option_text: "Salah", is_correct: !bsCorrect },
            ];
      const correctCount = opts.filter((o) => o.is_correct).length;
      if (correctCount === 0) {
        setError("Pilih minimal satu kunci jawaban");
        return;
      }
      if (type === "PG" && correctCount > 1) {
        setError("Pilihan Ganda hanya boleh satu kunci jawaban");
        return;
      }
      if ((type === "PG" || type === "MRA") && opts.some((o) => !o.option_text?.trim())) {
        setError("Semua opsi wajib diisi");
        return;
      }
    }

    setLoading(true);
    let opts: QuestionOptionInput[] = [];
    let key = answerKey;
    if (type === "PG") {
      opts = options.map((o) => ({ ...o, option_weight: 1 }));
      key = options.find((o) => o.is_correct)?.option_key ?? "";
    } else if (type === "MRA") {
      opts = options.map((o) => ({
        ...o,
        option_weight: (o.option_weight ?? 1) <= 0 ? 1 : (o.option_weight ?? 1),
      }));
      key = options
        .filter((o) => o.is_correct)
        .map((o) => o.option_key)
        .sort()
        .join(",");
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
    if (type === "PG") {
      setOptions((prev) => prev.map((o, i) => ({ ...o, is_correct: i === idx })));
    } else if (type === "MRA") {
      setOptions((prev) =>
        prev.map((o, i) => (i === idx ? { ...o, is_correct: !o.is_correct } : o))
      );
    }
  };

  const addOption = () => {
    const next = getNextOptionKey(options.map((o) => o.option_key));
    setOptions((prev) => [...prev, { option_key: next, option_text: "", is_correct: false }]);
  };

  const removeOption = (idx: number) => {
    if (options.length <= 2) return;
    setOptions((prev) => prev.filter((_, i) => i !== idx));
  };

  const setOptionWeight = (idx: number, value: number) => {
    setOptions((prev) =>
      prev.map((o, i) => (i === idx ? { ...o, option_weight: Math.max(0.1, value) } : o))
    );
  };

  const [draggedIdx, setDraggedIdx] = useState<number | null>(null);

  const moveOption = (fromIdx: number, toIdx: number) => {
    if (fromIdx === toIdx) return;
    setOptions((prev) => {
      const arr = [...prev];
      const [moved] = arr.splice(fromIdx, 1);
      arr.splice(toIdx, 0, moved);
      return arr.map((o, i) => ({ ...o, option_key: OPTION_KEYS[i] ?? String(i + 1) }));
    });
  };

  const handleDragStart = (e: React.DragEvent, idx: number) => {
    setDraggedIdx(idx);
    e.dataTransfer.effectAllowed = "move";
    e.dataTransfer.setData("text/plain", String(idx));
  };

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault();
  };

  const handleDrop = (e: React.DragEvent, toIdx: number) => {
    e.preventDefault();
    if (draggedIdx === null) return;
    if (draggedIdx !== toIdx) moveOption(draggedIdx, toIdx);
    setDraggedIdx(null);
  };

  const handleDragEnd = () => {
    setDraggedIdx(null);
  };

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
            <option value="MRA">Multiple Right Answer</option>
            <option value="BENAR_SALAH">Benar-Salah</option>
            <option value="ESSAY">Essay</option>
          </select>
          {type === "MRA" && (
            <p className="text-xs text-gray-500 mt-1">
              Beberapa jawaban benar. Masing-masing opsi bisa punya bobot berbeda.
            </p>
          )}
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
          <label className="block text-sm font-medium text-gray-700 mb-1">
            Tingkat Kesulitan
          </label>
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
          <RichTextEditor
            value={questionText}
            onChange={setQuestionText}
            placeholder="Tulis teks soal di sini..."
            minHeight="10rem"
          />
        </div>

        {(type === "PG" || type === "MRA") && (
          <div>
            <div className="flex items-center justify-between mb-2">
              <label className="block text-sm font-medium text-gray-700">
                Opsi Jawaban
                {type === "PG"
                  ? " (pilih satu kunci)"
                  : " (centang yang benar, atur bobot per opsi)"}
              </label>
              <button
                type="button"
                onClick={addOption}
                className="text-sm text-blue-600 hover:text-blue-800 font-medium"
              >
                + Tambah Opsi
              </button>
            </div>
            <p className="text-xs text-gray-500 mb-2">Geser untuk mengubah urutan</p>
            <div className="space-y-2">
              {options.map((o, i) => (
                <div
                  key={i}
                  draggable
                  onDragStart={(e) => handleDragStart(e, i)}
                  onDragOver={handleDragOver}
                  onDrop={(e) => handleDrop(e, i)}
                  onDragEnd={handleDragEnd}
                  className={`flex gap-2 items-center flex-wrap p-2 rounded-lg border transition-colors ${
                    draggedIdx === i ? "border-blue-400 bg-blue-50 opacity-75" : "border-gray-200 bg-white hover:border-gray-300"
                  }`}
                >
                  <div
                    className="flex-shrink-0 touch-none"
                    onMouseDown={(e) => e.stopPropagation()}
                  >
                    <GripIcon />
                  </div>
                  <span className="w-6 font-medium">{o.option_key}.</span>
                  <input
                    type="text"
                    value={o.option_text}
                    onChange={(e) =>
                      setOptions((prev) =>
                        prev.map((opt, j) =>
                          j === i ? { ...opt, option_text: e.target.value } : opt
                        )
                      )
                    }
                    className="flex-1 min-w-[120px] px-3 py-2 border border-gray-300 rounded-lg"
                    placeholder={`Opsi ${o.option_key}`}
                  />
                  {type === "MRA" && (
                    <div className="flex items-center gap-1">
                      <span className="text-xs text-gray-500">Bobot:</span>
                      <input
                        type="number"
                        value={o.option_weight ?? 1}
                        onChange={(e) =>
                          setOptionWeight(i, parseFloat(e.target.value) || 1)
                        }
                        min={0.1}
                        step={0.1}
                        className="w-16 px-2 py-1 border border-gray-300 rounded text-sm"
                      />
                    </div>
                  )}
                  <label className="flex items-center gap-1">
                    <input
                      type={type === "PG" ? "radio" : "checkbox"}
                      name={type === "PG" ? "correct" : undefined}
                      checked={o.is_correct}
                      onChange={() => setOptionCorrect(i)}
                    />
                    <span className="text-sm">Kunci</span>
                  </label>
                  {options.length > 2 && (
                    <button
                      type="button"
                      onClick={() => removeOption(i)}
                      className="text-red-600 hover:text-red-800 text-sm"
                    >
                      Hapus
                    </button>
                  )}
                </div>
              ))}
            </div>
          </div>
        )}

        {type === "BENAR_SALAH" && (
          <div>
            <label className="block text-sm font-medium text-gray-700 mb-2">
              Kunci Jawaban
            </label>
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
            <label className="block text-sm font-medium text-gray-700 mb-1">
              Model Jawaban / Rubrik
            </label>
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
          <label className="block text-sm font-medium text-gray-700 mb-1">Bobot Soal</label>
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
            className="px-4 py-2 bg-blue-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700 disabled:opacity-50"
          >
            {loading ? "Menyimpan..." : "Simpan"}
          </button>
          <Link
            href="/wpujikom/bank-soal"
            className="px-4 py-2 border border-gray-300 rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
          >
            Batal
          </Link>
        </div>
      </form>
    </div>
  );
}
