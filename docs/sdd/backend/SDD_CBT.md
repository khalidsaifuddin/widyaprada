## SDD – CBT (Computer-Based Test)

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom  
**Fitur**: Mengerjakan Ujian dan Melihat Hasil  

Dokumen ini menjelaskan **desain teknis backend** untuk PRD `[PRD] Fitur CBT` dengan stack: **Go (Golang)**, Clean Architecture, PostgreSQL.

---

## 1. Arsitektur & Konteks

- `usecase`: `CBTUsecase` — ListUjianTersedia, MulaiUjian, GetSoal, SimpanJawaban, SubmitUjian, GetRiwayatHasil.
- `delivery/http`: REST di `/api/v1/cbt`.
- `infrastructure`: ExamRepository, ExamAttemptRepository, AnswerRepository.

---

## 2. Kontrak API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/cbt/exams` | Daftar ujian tersedia |
| POST | `/api/v1/cbt/exams/:id/start` | Mulai ujian |
| GET | `/api/v1/cbt/attempts/:attemptId/questions` | Daftar soal |
| GET | `/api/v1/cbt/attempts/:attemptId/questions/:num` | Soal per nomor |
| POST | `/api/v1/cbt/attempts/:attemptId/answers` | Simpan jawaban |
| POST | `/api/v1/cbt/attempts/:attemptId/submit` | Submit ujian |
| GET | `/api/v1/cbt/history` | Riwayat ujian + nilai |

---

## 3. Skema Database

- `exams`, `exam_participants`, `exam_contents`.
- `exam_attempts`: id, exam_id, user_id, started_at, submitted_at, score.
- `exam_answers`: attempt_id, question_id, answer_value.

---

## 4. Aturan Bisnis

- Ujian tersedia: status Diterbitkan, jadwal aktif, user = peserta.
- Satu peserta satu attempt per ujian.
- Timer: started_at + durasi_menit. Auto-submit jika waktu habis.
- Nilai: PG & B-S otomatis; Essay manual.

---

## 5. RBAC

- Role **Widyaprada** saja.
