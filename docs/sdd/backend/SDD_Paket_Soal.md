## SDD – Paket Soal

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom  
**Fitur**: Paket Soal (List, Detail, Create, Edit, Delete, Verifikasi)  

Dokumen ini menjelaskan **desain teknis backend** untuk PRD `[PRD] Fitur Paket Soal` dengan stack: **Go (Golang)**, Clean Architecture, PostgreSQL.

---

## 1. Arsitektur & Konteks

- `usecase`: `PaketSoalUsecase` — List, Get, Create, Update, Delete, Verify.
- `delivery/http`: REST di `/api/v1/question-packages`.
- `infrastructure`: `QuestionPackageRepository`, `PackageQuestionItemRepository`.

---

## 2. Kontrak API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/question-packages` | List (q, status, status_verifikasi) |
| GET | `/api/v1/question-packages/:id` | Detail + daftar soal |
| POST | `/api/v1/question-packages` | Create |
| PUT | `/api/v1/question-packages/:id` | Update (nama, daftar soal + urutan) |
| DELETE | `/api/v1/question-packages/:id` | Delete (body: reason) |
| POST | `/api/v1/question-packages/:id/verify` | Verifikasi |
| POST | `/api/v1/question-packages/:id/unverify` | Batalkan verifikasi |

---

## 3. Skema Database

- `question_packages`: id, code (unique), name, description, status (Draft|Aktif), verification_status (Belum|Sudah).
- `package_question_items`: package_id, question_id, sort_order.

---

## 4. Aturan Bisnis

- Minimal 1 soal dalam paket. Kode unik.
- Delete: alasan wajib. Cek ujian yang memakai paket.
- Hapus soal dari paket hanya memutus relasi; soal di Bank Soal tetap ada.
- Verifikasi: hanya Verifikator & Super Admin.

---

## 5. RBAC

- Admin Uji Kompetensi, Super Admin: CRUD.
- Verifikator, Super Admin: Verifikasi / Batalkan Verifikasi.
