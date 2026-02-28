## SDD – Bank Soal

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom  
**Fitur**: Bank Soal (List, Detail, Create, Edit, Delete, Verifikasi)  

Dokumen ini menjelaskan **desain teknis backend** untuk PRD Fitur Bank Soal dengan stack: Go (Golang), Clean Architecture, PostgreSQL.

---

## 1. Arsitektur dan Konteks

- Usecase: BankSoalUsecase - List, Get, Create, Update, Delete, Verify.
- Delivery: REST di /api/v1/questions.
- Infrastructure: QuestionRepository, QuestionOptionRepository.

---

## 2. Kontrak API

- GET /api/v1/questions - List (q, tipe, kategori, status, status_verifikasi)
- GET /api/v1/questions/:id - Detail
- POST /api/v1/questions - Create
- PUT /api/v1/questions/:id - Update
- DELETE /api/v1/questions/:id - Delete (body: reason)
- POST /api/v1/questions/:id/verify - Verifikasi
- POST /api/v1/questions/:id/unverify - Batalkan verifikasi

---

## 3. Skema Database

- questions: id, code (unique), type, category_id, difficulty, question_text, answer_key, weight, status (Draft/Aktif), verification_status (Belum/Sudah).
- question_options: question_id, option_key, option_text, is_correct.
- question_categories: id, code, name.

---

## 4. Aturan Bisnis

- Delete: alasan wajib. Cek paket ujian yang memakai soal.
- Verifikasi: hanya Verifikator dan Super Admin.
- Create/Edit: kode unik, opsi dan kunci wajib untuk PG dan Benar-Salah.

---

## 5. RBAC

- **Super Admin**: CRUD (Create, Edit, Delete), migrasi/import soal. Hanya Super Admin yang berhak menginput dan migrasi soal.
- Admin Uji Kompetensi: Read-only (List, Detail); tidak Create/Edit/Delete.
- Verifikator, Super Admin: Verifikasi / Batalkan Verifikasi.
