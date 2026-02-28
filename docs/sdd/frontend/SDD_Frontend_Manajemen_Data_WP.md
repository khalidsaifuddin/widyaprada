## SDD Frontend – Manajemen Data WP

**Aplikasi**: Widyaprada  
**Modul**: Manajemen Data WP  
**Fitur**: Manajemen Data Widyaprada (List, Detail, Create, Edit, Delete)  
**PRD Terkait**: [PRD_Manajemen_Data_WP](../../prd/PRD_Manajemen_Data_WP.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk Manajemen Data WP dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/wpdata` (list), `/wpdata/create`, `/wpdata/[id]`, `/wpdata/[id]/edit`
- **Role**: Admin Satker (scope satker), Super Admin (seluruh sistem)
- **Layout**: DashboardLayout
- **API**: `/api/v1/wp-data` atau `/api/v1/manajemen-data-wp`
- **Scope**: Admin Satker hanya data WP dalam satker/unit kerjanya

---

## 2. Atomic Design – Komponen

### 2.1 Atoms

| Komponen | Deskripsi |
|----------|-----------|
| `Button` | Tambah, Simpan, Batal, Hapus |
| `Input`, `Label`, `Select`, `Textarea` | Field form |
| `Badge` | Status (Aktif/Nonaktif) |

### 2.2 Molecules

| Komponen | Deskripsi |
|----------|-----------|
| `FormField` | Label + Input + error |
| `SearchBar` | Pencarian NIP, nama |
| `FilterDropdown` | Filter: satker, unit kerja, status |
| `DeleteConfirmDialog` | Konfirmasi + alasan penghapusan (jika kebijakan) |
| `FileUploadField` | Upload dokumen (untuk verifikasi calon peserta) |

### 2.3 Organisms

| Komponen | Deskripsi |
|----------|-----------|
| `WPDataListTable` | Tabel: NIP, Nama, Satker, Unit Kerja, Jenjang, Status; search, filter, paginasi; aksi Detail, Edit, Hapus |
| `WPDataDetailCard` | Detail lengkap data WP; sub-section riwayat (jika ada): riwayat pangkat, jabatan, satker |
| `WPDataForm` | Create/Edit: NIP, Nama, Jenis Kelamin, Golongan, Pangkat, Jenjang Jabatan Fungsional, Satker, Unit Kerja, Pendidikan, TMT, No. SK, No. HP, Email, Alamat, Status |
| `CalonPesertaList` | Daftar calon peserta apply; verifikasi dokumen; tolak dengan catatan |
| `DokumenViewer` | Lihat dokumen persyaratan yang diupload calon |
| `WPDataDeleteDialog` | Konfirmasi + alasan |

### 2.4 Pages

| Route | Page |
|-------|------|
| `/wpdata` | WPDataListPage |
| `/wpdata/create` | WPDataCreatePage |
| `/wpdata/[id]` | WPDataDetailPage (dengan tab: Info, Riwayat, Calon Peserta jika berlaku) |
| `/wpdata/[id]/edit` | WPDataEditPage |

---

## 3. State & Validasi

- NIP: unik, format 18 digit (jika diterapkan)
- Field wajib minimal: NIP, Nama, Satker, Status
- Admin Satker: filter otomatis by satker user; tidak bisa akses data WP satker lain

---

## 4. Integrasi API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/wp-data` | List (search, filter by satker, paginate) |
| GET | `/api/v1/wp-data/:id` | Detail (+ riwayat jika ada) |
| POST | `/api/v1/wp-data` | Create |
| PUT | `/api/v1/wp-data/:id` | Edit |
| DELETE | `/api/v1/wp-data/:id` | Delete (body: reason jika kebijakan) |
| GET | `/api/v1/wp-data/calon-peserta` | Daftar calon (untuk verifikasi) |
| POST | `/api/v1/wp-data/calon-peserta/:id/verify` | Verifikasi / tolak |

---

## 5. File Lokasi

```
frontend/src/
├── app/wpdata/
│   ├── page.tsx
│   ├── create/page.tsx
│   └── [id]/page.tsx, [id]/edit/page.tsx
├── components/organisms/WPDataListTable.tsx
├── components/organisms/WPDataDetailCard.tsx
├── components/organisms/WPDataForm.tsx
└── components/organisms/CalonPesertaList.tsx
```
