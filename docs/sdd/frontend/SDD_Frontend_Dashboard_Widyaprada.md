## SDD Frontend – Dashboard Widyaprada

**Aplikasi**: Widyaprada  
**Modul**: Dashboard User  
**Fitur**: Dashboard Widyaprada – Assignment Uji Kompetensi & Jurnal Di-Submit  
**PRD Terkait**: [PRD_Dashboard_Widyaprada](../../prd/PRD_Dashboard_Widyaprada.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk Dashboard Widyaprada dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/dashboard`
- **Role**: Widyaprada (minimal)
- **Layout**: DashboardLayout (Sidebar + content)
- **API**: `GET /api/v1/dashboard/assignments`, `GET /api/v1/dashboard/journals`

---

## 2. Atomic Design – Komponen

### 2.1 Atoms

| Komponen | Deskripsi |
|----------|-----------|
| `Button` | Mulai Ujian, Lihat Hasil, Lihat Leaderboard, Lihat, Edit |
| `Badge` | Status (Belum dikerjakan, Sudah dikerjakan; Draft, Menunggu Verifikasi, dll) |
| `Skeleton` | Loading placeholder |

### 2.2 Molecules

| Komponen | Deskripsi |
|----------|-----------|
| `CardHeader` | Judul blok + link "Lihat semua" |
| `EmptyState` | Pesan "Anda belum memiliki penugasan ujian." / "Anda belum mengirimkan jurnal." |
| `AssignmentCardItem` | Satu item: nama ujian, batas waktu, status, hasil, tombol aksi |
| `JournalCardItem` | Satu item: judul, tanggal submit, status, tombol Lihat/Edit |

### 2.3 Organisms

| Komponen | Deskripsi |
|----------|-----------|
| `AssignmentBlock` | Blok "Tugas Saya": header + list AssignmentCardItem (5–10 item) + EmptyState; loading skeleton |
| `JournalBlock` | Blok "Jurnal Saya": header + list JournalCardItem + EmptyState; loading skeleton |
| `DashboardHeader` | Sapaan "Selamat datang, [Nama]" atau judul "Dashboard" |

### 2.4 Templates

| Komponen | Deskripsi |
|----------|-----------|
| `DashboardLayout` | Sidebar + content area |

### 2.5 Pages

| Route | Page |
|-------|------|
| `/dashboard` | DashboardPage = DashboardHeader + AssignmentBlock + JournalBlock |

---

## 3. Layout Urutan

1. **DashboardHeader** – sapaan singkat
2. **AssignmentBlock** (Blok 1) – di atas atau kiri (prioritas tinggi)
3. **JournalBlock** (Blok 2) – di bawah atau kanan

Grid responsif: di mobile stacking vertikal; di desktop 2 kolom atau stacked.

---

## 4. State & Fetch

### 4.1 AssignmentBlock
- Fetch: `GET /api/v1/dashboard/assignments?limit=10&page=1`
- State: assignments[], loading, error
- Aksi: Mulai Ujian → navigate `/wpujikom/cbt`; Lihat Hasil → hasil ujian; Lihat Leaderboard → `/wpujikom/assignment/[examId]/leaderboard`
- Link "Lihat semua" → `/wpujikom/assignment`

### 4.2 JournalBlock
- Fetch: `GET /api/v1/dashboard/journals?limit=10&page=1`
- State: journals[], loading, error
- Aksi: Lihat → detail jurnal; Edit → edit jurnal
- Link "Lihat semua" → `/wpjurnal`; opsional "Buat Jurnal"

---

## 5. Umpan Balik

- Skeleton/spinner per blok saat loading; jangan block seluruh halaman jika satu sumber lambat
- State kosong per blok: EmptyState dengan pesan sesuai
- Toast sukses setelah kembali dari CBT (opsional)

---

## 6. File Lokasi

```
frontend/src/
├── app/dashboard/page.tsx
├── components/organisms/AssignmentBlock.tsx
├── components/organisms/JournalBlock.tsx
├── components/molecules/AssignmentCardItem.tsx
└── components/molecules/JournalCardItem.tsx
```
