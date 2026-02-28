## SDD – Manajemen Uji Kompetensi

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom  
**Fitur**: CRUD Ujian, Jadwal, Konten, Peserta, Terbitkan, Verifikasi  

Dokumen ini menjelaskan **desain teknis backend** untuk PRD `[PRD] Fitur Manajemen Uji Kompetensi` dengan stack: **Go (Golang)**, Clean Architecture, PostgreSQL.

---

## 1. Arsitektur & Konteks

- `usecase`: `ExamUsecase` — List, Get, Create, Update, Delete, Publish, Verify, GetRekap.
- `delivery/http`: REST di `/api/v1/exams`.
- `infrastructure`: ExamRepository, ExamContentRepository, ExamParticipantRepository.

---

## 2. Kontrak API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/exams` | List (q, status, status_verifikasi) |
| GET | `/api/v1/exams/:id` | Detail + konten + peserta + rekap |
| POST | `/api/v1/exams` | Create (Draft) |
| PUT | `/api/v1/exams/:id` | Update (hanya Draft) |
| DELETE | `/api/v1/exams/:id` | Delete (body: reason) |
| POST | `/api/v1/exams/:id/publish` | Terbitkan |
| POST | `/api/v1/exams/:id/verify` | Verifikasi |
| POST | `/api/v1/exams/:id/unverify` | Batalkan verifikasi |

**Body Create/Update**: kode, nama, jadwal, durasi, question_ids, package_ids, shuffle_questions, tampilkan_leaderboard. participant_ids tidak diisi manual — peserta berasal dari hasil validasi dokumen (ujikom_applications yang lolos).

---

## 3. Skema Database

- `exams`: id, kode, nama, jadwal_mulai, jadwal_selesai, durasi_menit, status (Draft|Diterbitkan|Berlangsung|Selesai), verification_status, shuffle_questions, tampilkan_leaderboard.
- `exam_contents`: exam_id, source_type (question|package), source_id, sort_order.
- `exam_participants`: exam_id, user_id. Peserta = user yang lolos validasi dari ujikom_applications; sync saat verifikasi.

---

## 4. Aturan Bisnis

- Konten = soal individu + paket soal; minimal 1 soal, minimal 1 peserta.
- Edit penuh hanya untuk status Draft.
- Terbitkan: Draft → Diterbitkan.

---

## 5. RBAC

- Admin Uji Kompetensi, Super Admin: CRUD, Terbitkan.
- Verifikator, Super Admin: Verifikasi / Batalkan Verifikasi.
