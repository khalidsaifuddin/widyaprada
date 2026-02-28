## SDD Frontend – Bank Soal

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom  
**Fitur**: Bank Soal (List, Detail, Create, Edit, Delete)  
**PRD Terkait**: [PRD_Bank_Soal](../../prd/PRD_Bank_Soal.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk Bank Soal dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/wpujikom/bank-soal` (list), `/wpujikom/bank-soal/create`, `/wpujikom/bank-soal/[id]`, `/wpujikom/bank-soal/[id]/edit`
- **Role**: Super Admin (CRUD), Admin Uji Kompetensi & Verifikator (view, Verifikator: verifikasi)
- **Layout**: DashboardLayout
- **API**: `/api/v1/bank-soal`

---

## 2. Atomic Design – Komponen

### 2.1 Atoms

| Komponen | Deskripsi |
|----------|-----------|
| `Button` | Tambah, Simpan, Batal, Hapus, Verifikasi |
| `Input` | Text, textarea, select |
| `Label`, `Badge` | Status (Draft/Aktif), Status Verifikasi |
| `RadioGroup` | PG: opsi A/B/C/D; B-S: Benar/Salah |

### 2.2 Molecules

| Komponen | Deskripsi |
|----------|-----------|
| `FormField` | Label + Input + error |
| `SearchBar` | Pencarian kode/teks soal |
| `FilterDropdown` | Filter: tipe soal, kategori kompetensi, tingkat kesulitan, status, status verifikasi |
| `QuestionOptionsEditor` | PG: daftar opsi + kunci; B-S: Benar/Salah; Essay: rubrik |
| `DeleteConfirmDialog` | Konfirmasi + alasan penghapusan wajib (Super Admin) |

### 2.3 Organisms

| Komponen | Deskripsi |
|----------|-----------|
| `SoalListTable` | Tabel: kode, tipe, kategori, tingkat kesulitan, status, status verifikasi; search, filter, sort, paginasi; aksi Detail, Edit (Super Admin), Hapus (Super Admin), Verifikasi (Verifikator) |
| `SoalDetailCard` | Detail lengkap soal (teks, opsi, kunci, bobot); tombol Edit (Super Admin), Verifikasi/Batalkan Verifikasi (Verifikator) |
| `SoalForm` | Create/Edit: kode, tipe, kategori, teks soal, opsi (PG), kunci (PG/B-S/Essay), bobot, status; validasi |
| `SoalDeleteDialog` | Konfirmasi + textarea alasan wajib |
| `VerificationBadge` | Badge status verifikasi + tombol Verifikasi (Verifikator) |

### 2.4 Pages

| Route | Page |
|-------|------|
| `/wpujikom/bank-soal` | SoalListPage |
| `/wpujikom/bank-soal/create` | SoalCreatePage (Super Admin only) |
| `/wpujikom/bank-soal/[id]` | SoalDetailPage |
| `/wpujikom/bank-soal/[id]/edit` | SoalEditPage (Super Admin only) |

---

## 3. State & Validasi

- List: search, filter (tipe, kategori, tingkat kesulitan, status, status verifikasi), sort, paginasi
- Form: tipe soal menentukan field (PG → opsi; B-S → Benar/Salah; Essay → rubrik/model jawaban)
- Delete: wajib alasan penghapusan

---

## 4. Integrasi API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/bank-soal` | List (search, filter, paginate) |
| GET | `/api/v1/bank-soal/:id` | Detail |
| POST | `/api/v1/bank-soal` | Create (Super Admin) |
| PUT | `/api/v1/bank-soal/:id` | Edit (Super Admin) |
| DELETE | `/api/v1/bank-soal/:id` | Delete (Super Admin, body: reason) |
| POST | `/api/v1/bank-soal/:id/verify` | Verifikasi (Verifikator) |

---

## 5. File Lokasi

```
frontend/src/
├── app/wpujikom/bank-soal/
│   ├── page.tsx
│   ├── create/page.tsx
│   └── [id]/page.tsx, [id]/edit/page.tsx
├── components/organisms/SoalListTable.tsx
├── components/organisms/SoalDetailCard.tsx
├── components/organisms/SoalForm.tsx
└── components/molecules/QuestionOptionsEditor.tsx
```
