## SDD Frontend – Manajemen Uji Kompetensi

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom  
**Fitur**: Manajemen Uji Kompetensi (List, Detail, Create, Edit, Delete, Jadwal, Paket Soal, Peserta)  
**PRD Terkait**: [PRD_Manajemen_Uji_Kompetensi](../../prd/PRD_Manajemen_Uji_Kompetensi.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk Manajemen Uji Kompetensi dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/wpujikom/uji-kompetensi` (list), `/wpujikom/uji-kompetensi/create`, `/wpujikom/uji-kompetensi/[id]`, `/wpujikom/uji-kompetensi/[id]/edit`
- **Role**: Admin Uji Kompetensi, Verifikator (verifikasi), Super Admin
- **Layout**: DashboardLayout
- **API**: `/api/v1/ujian` atau `/api/v1/uji-kompetensi`

---

## 2. Atomic Design – Komponen

### 2.1 Atoms

| Komponen | Deskripsi |
|----------|-----------|
| `Button` | Tambah, Simpan, Terbitkan, Verifikasi, Hapus |
| `Input`, `Label`, `Badge`, `Select` | Status (Draft, Diterbitkan, Berlangsung, Selesai), Status Verifikasi |
| `DateTimePicker` | Jadwal Mulai, Jadwal Selesai |

### 2.2 Molecules

| Komponen | Deskripsi |
|----------|-----------|
| `FormField`, `SearchBar`, `FilterDropdown` | Filter: status, status verifikasi |
| `SoalIndividuSelector` | Pilih soal dari Bank Soal (multi-select) |
| `PaketSoalSelector` | Pilih paket soal (bisa lebih dari satu) |
| `PesertaSelector` | Daftar calon peserta (hasil validasi); atur jadwal Ujikom |
| `LeaderboardToggle` | Tampilkan Leaderboard: Ya/Tidak (default Tidak) |
| `DeleteConfirmDialog` | Konfirmasi + alasan |
| `VerificationAction` | Tombol Verifikasi / Batalkan Verifikasi |

### 2.3 Organisms

| Komponen | Deskripsi |
|----------|-----------|
| `UjianListTable` | Tabel: kode, nama, jadwal, durasi, status, status verifikasi; search, filter; aksi Detail, Edit, Terbitkan, Verifikasi, Hapus |
| `UjianDetailCard` | Detail lengkap: jadwal, durasi, konten (soal individu + paket), peserta, rekap hasil; tombol Edit, Terbitkan, Verifikasi |
| `UjianForm` | Create/Edit: kode, nama, deskripsi; Jadwal Mulai, Selesai; Durasi; Konten (soal individu + paket soal); Peserta (dari validasi); Tampilkan Leaderboard (Ya/Tidak) |
| `UjianKontenEditor` | Tambah soal individu, tambah paket soal; pengacakan urutan opsional |
| `PesertaList` | Daftar peserta (calon lolos validasi); atur jadwal Ujikom; rekap nilai |
| `RekapHasilTable` | Daftar peserta, status (sudah/belum submit), nilai |
| `CalonPesertaList` | Daftar calon apply; verifikasi/tolak dokumen; catatan jika tolak |
| `UjianDeleteDialog` | Konfirmasi + alasan |

### 2.4 Pages

| Route | Page |
|-------|------|
| `/wpujikom/uji-kompetensi` | UjianListPage |
| `/wpujikom/uji-kompetensi/create` | UjianCreatePage |
| `/wpujikom/uji-kompetensi/[id]` | UjianDetailPage (tab: Info, Konten, Peserta, Rekap Hasil) |
| `/wpujikom/uji-kompetensi/[id]/edit` | UjianEditPage |

---

## 3. State & Validasi

- Form: Jadwal Selesai > Jadwal Mulai; Durasi > 0; minimal 1 soal (individu atau paket)
- Peserta: hanya calon yang lolos validasi; tidak assign manual
- Terbitkan: hanya ujian Draft; setelah Terbitkan, edit terbatas
- Verifikasi: hanya Verifikator; tombol Verifikasi/Batalkan Verifikasi

---

## 4. Integrasi API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/ujian` | List |
| GET | `/api/v1/ujian/:id` | Detail (konten, peserta, rekap) |
| POST | `/api/v1/ujian` | Create |
| PUT | `/api/v1/ujian/:id` | Edit |
| POST | `/api/v1/ujian/:id/publish` | Terbitkan |
| DELETE | `/api/v1/ujian/:id` | Delete (body: reason) |
| POST | `/api/v1/ujian/:id/verify` | Verifikasi |
| GET | `/api/v1/ujian/:id/calon-peserta` | Daftar calon apply |
| POST | `/api/v1/ujian/:id/calon-peserta/:userId/approve` | Setuju |
| POST | `/api/v1/ujian/:id/calon-peserta/:userId/reject` | Tolak (body: catatan) |
| PUT | `/api/v1/ujian/:id/jadwal-peserta` | Atur jadwal Ujikom |

---

## 5. File Lokasi

```
frontend/src/
├── app/wpujikom/uji-kompetensi/
│   ├── page.tsx
│   ├── create/page.tsx
│   └── [id]/page.tsx, [id]/edit/page.tsx
├── components/organisms/UjianListTable.tsx
├── components/organisms/UjianDetailCard.tsx
├── components/organisms/UjianForm.tsx
├── components/organisms/UjianKontenEditor.tsx
├── components/organisms/CalonPesertaList.tsx
└── components/organisms/RekapHasilTable.tsx
```
