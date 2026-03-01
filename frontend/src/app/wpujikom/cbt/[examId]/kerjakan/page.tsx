"use client";

import { apiService } from "@/lib/api";
import Link from "next/link";
import { useParams, useRouter, useSearchParams } from "next/navigation";
import { useCallback, useEffect, useRef, useState } from "react";
interface CBTQuestionOption {
  id: string;
  option_key: string;
  option_text: string;
}

interface CBTQuestion {
  num: number;
  question_id: string;
  type: string;
  question_text: string;
  weight: number;
  options?: CBTQuestionOption[];
}

interface CBTQuestionsResponse {
  attempt_id: string;
  total: number;
  questions: CBTQuestion[];
}

function formatTime(seconds: number): string {
  const m = Math.floor(seconds / 60);
  const s = seconds % 60;
  return `${m}:${s.toString().padStart(2, "0")}`;
}

export default function CBTWorkspacePage() {
  const params = useParams();
  const router = useRouter();
  const searchParams = useSearchParams();
  const examId = params.examId as string;
  const attemptId = searchParams.get("attemptId") || "";
  const [questions, setQuestions] = useState<CBTQuestion[]>([]);
  const [currentIdx, setCurrentIdx] = useState(0);
  const [answers, setAnswers] = useState<
    Record<string, { option_id?: string; option_ids?: string[]; answer_text?: string }>
  >({});
  const [timeRemaining, setTimeRemaining] = useState(0);
  const [submitted, setSubmitted] = useState(false);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [submitConfirm, setSubmitConfirm] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const timerRef = useRef<ReturnType<typeof setInterval> | null>(null);
  const startedAtRef = useRef<string | null>(null);
  const durasiMenitRef = useRef(0);

  const saveAnswer = useCallback(
    async (
      questionId: string,
      value: { option_id?: string; option_ids?: string[]; answer_text?: string }
    ) => {
      if (!attemptId || submitted) return;
      setAnswers((prev) => ({ ...prev, [questionId]: value }));
      await apiService.post(`v1/cbt/attempts/${attemptId}/answers`, {
        question_id: questionId,
        option_id: value.option_id || undefined,
        option_ids: value.option_ids?.length ? value.option_ids : undefined,
        answer_text: value.answer_text || undefined,
      });
    },
    [attemptId, submitted]
  );

  const fetchQuestions = useCallback(async () => {
    if (!attemptId) {
      setError("Attempt ID tidak valid");
      setLoading(false);
      return;
    }
    const res = await apiService.get<CBTQuestionsResponse>(
      `v1/cbt/attempts/${attemptId}/questions`
    );
    if (res.success && res.data) {
      const d = res.data as CBTQuestionsResponse;
      setQuestions(d.questions ?? []);
    } else {
      setError(res.message ?? "Gagal memuat soal");
    }
    setLoading(false);
  }, [attemptId]);

  useEffect(() => {
    fetchQuestions();
  }, [fetchQuestions]);

  const startedAtParam = searchParams.get("startedAt");
  const durasiParam = searchParams.get("durasiMenit");
  const submittedRef = useRef(false);

  const handleAutoSubmit = useCallback(async () => {
    if (submittedRef.current) return;
    submittedRef.current = true;
    setSubmitted(true);
    const res = await apiService.post(`v1/cbt/attempts/${attemptId}/submit`, {});
    if (res.success) {
      router.push(`/wpujikom/cbt/selesai?attemptId=${attemptId}`);
    } else {
      router.push(`/wpujikom/cbt/selesai?attemptId=${attemptId}&error=1`);
    }
  }, [attemptId, router]);

  useEffect(() => {
    if (!attemptId) return;
    const durasi = durasiParam ? parseInt(durasiParam, 10) : 60;
    durasiMenitRef.current = durasi;
    const started = startedAtParam;
    if (started) {
      startedAtRef.current = started;
      const end = new Date(started).getTime() + durasi * 60 * 1000;
      const update = () => {
        const now = Date.now();
        const left = Math.max(0, Math.floor((end - now) / 1000));
        setTimeRemaining(left);
        if (left <= 0 && timerRef.current) {
          clearInterval(timerRef.current);
          handleAutoSubmit();
        }
      };
      update();
      timerRef.current = setInterval(update, 1000);
    } else {
      setTimeRemaining(durasi * 60);
    }
    return () => {
      if (timerRef.current) clearInterval(timerRef.current);
    };
  }, [attemptId, startedAtParam, durasiParam, handleAutoSubmit]);

  const handleSubmit = async () => {
    setSubmitting(true);
    const res = await apiService.post(`v1/cbt/attempts/${attemptId}/submit`, {});
    setSubmitting(false);
    setSubmitConfirm(false);
    if (res.success) {
      setSubmitted(true);
      router.push(`/wpujikom/cbt/selesai?attemptId=${attemptId}`);
    } else {
      setError(res.message ?? "Gagal submit");
    }
  };

  if (!attemptId) {
    return (
      <div className="p-8">
        <p className="text-red-600">Sesi ujian tidak valid.</p>
        <Link href="/wpujikom/cbt" className="text-blue-600 mt-2 inline-block">
          Kembali ke Daftar Ujian
        </Link>
      </div>
    );
  }

  if (loading) return <div className="p-8">Memuat soal...</div>;
  if (error && questions.length === 0) {
    return (
      <div className="p-8">
        <p className="text-red-600">{error}</p>
        <Link href="/wpujikom/cbt" className="text-blue-600 mt-2 inline-block">
          Kembali
        </Link>
      </div>
    );
  }

  const current = questions[currentIdx];
  const isLastFive = timeRemaining > 0 && timeRemaining <= 300;

  return (
    <div className="max-w-4xl mx-auto">
      <div className="flex flex-wrap items-center justify-between gap-4 mb-6">
        <div
          className={`text-lg font-mono font-semibold ${isLastFive ? "text-red-600" : "text-gray-900"}`}
        >
          Sisa: {formatTime(timeRemaining)}
        </div>
        <div className="flex flex-wrap gap-2">
          {questions.map((q, i) => (
            <button
              key={q.question_id}
              onClick={() => setCurrentIdx(i)}
              className={`w-9 h-9 rounded-lg border text-sm font-medium ${
                i === currentIdx
                  ? "bg-blue-600 text-white border-blue-600"
                  : answers[q.question_id]
                    ? "bg-green-100 border-green-300 text-green-800"
                    : "border-gray-300 transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
              }`}
            >
              {i + 1}
            </button>
          ))}
        </div>
      </div>

      {error && (
        <div className="mb-4 p-3 rounded-lg bg-red-50 text-red-700 text-sm">{error}</div>
      )}

      {current && (
        <div className="bg-white rounded-xl shadow-sm border border-gray-200 p-6 mb-6">
          <p className="text-sm text-gray-500 mb-2">
            Soal {current.num} dari {questions.length}
          </p>
          <div
            className="font-medium text-gray-900 mb-4 text-sm [&_h1]:text-xl [&_h1]:font-bold [&_h2]:text-lg [&_h2]:font-semibold [&_h3]:text-base [&_h3]:font-semibold [&_ul]:list-disc [&_ul]:pl-6 [&_ol]:list-decimal [&_ol]:pl-6 [&_blockquote]:border-l-2 [&_blockquote]:border-gray-300 [&_blockquote]:pl-4 [&_blockquote]:italic"
            dangerouslySetInnerHTML={{ __html: current.question_text || "" }}
          />
          {current.type === "PG" || current.type === "MRA" || current.type === "BENAR_SALAH" ? (
            <div className="space-y-2">
              {current.type === "MRA" && (
                <p className="text-sm text-gray-500 mb-2">Pilih semua jawaban yang benar (boleh lebih dari satu)</p>
              )}
              {current.options?.map((opt) => {
                const isSelected =
                  current.type === "MRA"
                    ? (answers[current.question_id]?.option_ids ?? []).includes(opt.id)
                    : answers[current.question_id]?.option_id === opt.id;
                const toggleMRA = () => {
                  const currentIds = answers[current.question_id]?.option_ids ?? [];
                  const next = currentIds.includes(opt.id)
                    ? currentIds.filter((id) => id !== opt.id)
                    : [...currentIds, opt.id];
                  saveAnswer(current.question_id, { option_ids: next });
                };
                return (
                  <label
                    key={opt.id}
                    className={`flex items-center gap-2 p-3 rounded-lg border cursor-pointer ${
                      isSelected
                        ? "border-blue-500 bg-blue-50"
                        : "border-gray-200 transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
                    }`}
                  >
                    <input
                      type={current.type === "MRA" ? "checkbox" : "radio"}
                      name={current.type === "MRA" ? undefined : `q-${current.question_id}`}
                      checked={isSelected}
                      onChange={() =>
                        current.type === "MRA"
                          ? toggleMRA()
                          : saveAnswer(current.question_id, { option_id: opt.id })
                      }
                    />
                    <span>
                      {opt.option_key}. {opt.option_text}
                    </span>
                  </label>
                );
              })}
            </div>
          ) : (
            <textarea
              value={answers[current.question_id]?.answer_text ?? ""}
              onChange={(e) =>
                saveAnswer(current.question_id, { answer_text: e.target.value })
              }
              rows={5}
              className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
              placeholder="Tulis jawaban Anda..."
            />
          )}
        </div>
      )}

      <div className="flex justify-between">
        <button
          onClick={() => setCurrentIdx((i) => Math.max(0, i - 1))}
          disabled={currentIdx === 0}
          className="px-4 py-2 border border-gray-300 rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50 disabled:opacity-50"
        >
          Sebelumnya
        </button>
        {currentIdx < questions.length - 1 ? (
          <button
            onClick={() => setCurrentIdx((i) => Math.min(questions.length - 1, i + 1))}
            className="px-4 py-2 bg-blue-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-blue-700"
          >
            Berikutnya
          </button>
        ) : (
          <button
            onClick={() => setSubmitConfirm(true)}
            disabled={submitted || submitting}
            className="px-4 py-2 bg-green-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-green-700 disabled:opacity-50"
          >
            Submit Ujian
          </button>
        )}
      </div>

      {submitConfirm && (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/50 p-4">
          <div className="bg-white rounded-lg shadow-xl p-6 max-w-md w-full">
            <h3 className="text-lg font-semibold">Konfirmasi Submit</h3>
            <p className="mt-2 text-gray-600">
              Anda yakin ingin mengirim jawaban? Setelah submit Anda tidak dapat mengubah jawaban.
            </p>
            <div className="mt-4 flex gap-2 justify-end">
              <button
                onClick={() => setSubmitConfirm(false)}
                className="px-4 py-2 border border-gray-300 rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-gray-50"
              >
                Batal
              </button>
              <button
                onClick={handleSubmit}
                disabled={submitting}
                className="px-4 py-2 bg-green-600 text-white rounded-lg transition-all duration-300 hover:scale-[1.02] hover:-translate-y-0.5 hover:bg-green-700 disabled:opacity-50"
              >
                {submitting ? "Mengirim..." : "Ya, Submit"}
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
