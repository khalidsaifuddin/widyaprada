## SDD Frontend – Paket Soal

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom  
**Fitur**: Paket Soal (List, Detail, Create, Edit, Delete)  
**PRD Terkait**: [PRD_Paket_Soal](../../prd/PRD_Paket_Soal.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk Paket Soal dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/wpujikom/paket-soal` (list), `/wpujikom/paket-soal/create`, `/wpujikom/paket-soal/[id]`, `/wpujikom/paket-soal/[id]/edit`
- **Role**: Admin Uji Kompetensi, Verifikator (verifikasi), Super Admin
- **Layout**: DashboardLayout
- **API**: `/api/v1/paket-soal`

---

## 2. Atomic Design – Komponen

### 2.1 Atoms

| Komponen | Deskripsi |
|----------|-----------|
| `Button` | Tambah, Simpan, Verifikasi, Hapus |
| `Input`, `Label`, `Badge` | Status (Draft/Aktif), Status Verifikasi |

### 2.2 Molecules

| Komponen | Deskripsi |
|----------|-----------|
| `FormField`, `SearchBar`, `FilterDropdown` | Filter: status, status verifikasi |
| `SoalSelector` | Pilih soal dari Bank Soal (search, filter, multi-select) |
| `DragDropList` | Urutan soal dalam paket (drag-and-drop atau nomor) |
| `DeleteConfirmDialog` | Konfirmasi + alasan penghapusan wajib |

### 2.3 Organisms

| Komponen | Deskripsi |
|----------|-----------|
| `PaketSoalListTable` | Tabel: kode, nama, jumlah soal, status, status verifikasi; search, filter; aksi Detail, Edit, Hapus, Verifikasi |
| `PaketSoalDetailCard` | Nama, deskripsi, daftar soal berurutan; tombol Edit, Verifikasi |
| `PaketSoalForm` | Create/Edit: kode, nama, deskripsi; tambah/hapus soal dari Bank Soal; atur urutan (drag-drop) |
| `PaketSoalDeleteDialog` | Konfirmasi + alasan |

### 2.4 Pages

| Route | Page |
|-------|------|
| `/wpujikom/paket-soal` | PaketSoalListPage |
| `/wpujikom/paket-soal/create` | PaketSoalCreatePage |
| `/wpujikom/paket-soal/[id]` | PaketSoalDetailPage |
| `/wpujikom/paket-soal/[id]/edit` | PaketSoalEditPage |

---

## 3. State & Validasi

- Form: minimal 1 soal dalam paket
- Edit: tambah/hapus soal; ubah urutan (drag-drop)
- Delete: cek apakah paket dipakai di ujian; jika iya, peringatan atau blok

---

## 4. Integrasi API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/paket-soal` | List |
| GET | `/api/v1/paket-soal/:id` | Detail + daftar soal |
| POST | `/api/v1/paket-soal` | Create |
| PUT | `/api/v1/paket-soal/:id` | Edit (termasuk urutan soal) |
| DELETE | `/api/v1/paket-soal/:id` | Delete (body: reason) |
| POST | `/api/v1/paket-soal/:id/verify` | Verifikasi |

---

## 5. File Lokasi

```
frontend/src/
├── app/wpujikom/paket-soal/
│   ├── page.tsx
│   ├── create/page.tsx
│   └── [id]/page.tsx, [id]/edit/page.tsx
├── components/organisms/PaketSoalListTable.tsx
├── components/organisms/PaketSoalDetailCard.tsx
├── components/organisms/PaketSoalForm.tsx
└── components/molecules/SoalSelector.tsx
```
