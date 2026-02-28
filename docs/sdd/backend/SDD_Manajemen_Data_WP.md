## SDD – Manajemen Data WP

**Aplikasi**: Widyaprada  
**Modul**: Manajemen Data WP  
**Fitur**: CRUD Data Widyaprada  

Dokumen ini menjelaskan **desain teknis backend** untuk PRD Fitur Manajemen Data WP dengan stack: Go (Golang), Clean Architecture, PostgreSQL.

---

## 1. Arsitektur dan Konteks

- Usecase: WidyapradaDataUsecase - List, Get, Create, Update, Delete.
- Delivery: REST di /api/v1/widyaprada-data atau /api/v1/wp-data.
- Scope: Admin Satker (satker sendiri), Super Admin (semua).

---

## 2. Kontrak API

- GET /api/v1/wp-data/calon-peserta - List calon peserta (yang apply). Query: q, status_verifikasi, page, page_size.
- GET /api/v1/wp-data/calon-peserta/:id - Detail calon + dokumen. Aksi: verifikasi (setuju/tolak dengan catatan).
- POST /api/v1/wp-data/calon-peserta/:id/verify - Verifikasi (body: approved: bool, catatan: string jika tolak).
- GET /api/v1/wp-data - List. Query: q, satker_id, unit_kerja, status, page, page_size, sort.
- GET /api/v1/wp-data/:id - Detail
- POST /api/v1/wp-data - Create
- PUT /api/v1/wp-data/:id - Update
- DELETE /api/v1/wp-data/:id - Delete (body: reason)

Body Create/Update: nip, nama_lengkap, jenis_kelamin, golongan_ruang, pangkat, jenjang_jabatan_fungsional, satker_id, unit_kerja, pendidikan_terakhir, tmt_golongan, tmt_jabatan_fungsional, no_sk, no_hp, email, alamat, status (Aktif/Nonaktif), keterangan.

---

## 3. Skema Database

- `ujikom_applications`, `ujikom_application_documents` — lihat SDD Assignment (apply pendaftaran, calon peserta).
- widyaprada_data: id, nip (unique), nama_lengkap, jenis_kelamin, golongan_ruang, pangkat, jenjang_jabatan_fungsional, satker_id, unit_kerja, pendidikan_terakhir, tmt_golongan, tmt_jabatan_fungsional, no_sk_pengangkatan, no_hp, email, alamat, status, keterangan, user_id (relasi ke user jika ada), created_at, updated_at, deleted_at, deleted_reason.
- Tabel referensi: satkers, unit_kerja.
- Tabel historis (opsional): riwayat_pangkat, riwayat_jabatan, dll (one-to-many).

---

## 4. Aturan Bisnis

- NIP unik. Validasi format 18 digit (jika diterapkan).
- Field wajib minimal: nip, nama_lengkap, satker_id, status.
- Delete: alasan wajib. Soft delete disarankan.
- Admin Satker: filter by satker_id.

---

## 5. RBAC

- Permission: WP_DATA_READ, WP_DATA_CREATE, WP_DATA_UPDATE, WP_DATA_DELETE.
- Scope satker untuk Admin Satker.
