# [PRD] Dashboard Widyaprada
## Product Requirements Document | Halaman Utama User Widyaprada

**Aplikasi**: Widyaprada  
**Modul**: Dashboard User  
**Fitur**: Dashboard Widyaprada – Daftar Assignment Uji Kompetensi & Daftar Jurnal yang Di-Submit  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Dashboard Widyaprada
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Dashboard Widyaprada adalah **halaman utama** yang ditampilkan kepada pengguna dengan role **Widyaprada** setelah login. Isi utamanya adalah **(1) Daftar Assignment Uji Kompetensi** (Tugas Saya / Penugasan Ujian) dan **(2) Daftar Jurnal yang Di-Submit** (Jurnal Saya). Dari dashboard ini user dapat melihat penugasan ujian yang harus dikerjakan beserta batas waktu, status pengerjaan, dan aksi Mulai Ujian/Lihat Hasil; serta melihat daftar jurnal yang telah user submit beserta status (Draft, Menunggu Verifikasi, Diverifikasi, Ditolak, Published). Dashboard berfungsi sebagai **titik masuk** dan **ringkasan tugas** agar Widyaprada dapat mengatur prioritas tanpa harus masuk ke masing-masing modul (WPUjikom, WPJurnal) terlebih dahulu.

### 1.3 Role yang Mengakses
- **Widyaprada**: Satu-satunya role yang melihat **Dashboard Widyaprada** ini sebagai halaman utama setelah login (atau sebagai menu "Dashboard" di area user). Admin, Verifikator, Super Admin memiliki menu dan titik masuk masing-masing (mis. Manajemen Uji Kompetensi, CMS); mereka tidak memakai "Dashboard Widyaprada" kecuali jika user yang sama juga punya role Widyaprada.
- **Catatan**: Satu pengguna dapat memiliki lebih dari satu role. Jika user punya role Widyaprada, maka akses ke Dashboard Widyaprada (dan isinya: Tugas Saya + Jurnal Saya) tersedia.

### 1.4 Relasi ke PRD Lain
- **Daftar Assignment Uji Kompetensi**: Data dan aturan tampil mengacu pada [PRD_Assignment.md](PRD_Assignment.md) (Tugas Saya: ujian yang di-assign ke user, batas waktu, status, Mulai Ujian, Lihat Hasil).
- **Daftar Jurnal yang Di-Submit**: Data bersumber dari modul **WPJurnal** (Manajemen Jurnal: Create/Submit, Edit, status verifikasi). PRD WPJurnal (ongoing); di dashboard hanya menampilkan **jurnal yang di-submit oleh user yang login** dengan status masing-masing.
- **CBT, Manajemen Uji Kompetensi, Jurnal (List/Detail publik)**: Lihat PRD_CBT, PRD_Manajemen_Uji_Kompetensi, PRD_Jurnal untuk detail fitur terkait.

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Widyaprada | pengguna dashboard | setelah login melihat **dashboard** berisi ringkasan tugas saya (penugasan ujian + jurnal saya) | langsung tahu apa yang harus dikerjakan dan status submitan jurnal tanpa buka banyak menu |
| 2 | Widyaprada | peserta uji kompetensi | di dashboard melihat **daftar assignment uji kompetensi** (ujian yang ditugaskan ke saya) dengan nama ujian, batas waktu, status (belum/sudah dikerjakan), dan tombol Mulai Ujian / Lihat Hasil | mengatur prioritas ujian dan mengerjakan sebelum deadline |
| 3 | Widyaprada | peserta uji kompetensi | mengklik "Mulai Ujian" atau "Lihat Hasil" dari dashboard ke alur CBT / Hasil | mengerjakan ujian atau melihat nilai tanpa navigasi panjang |
| 4 | Widyaprada | penulis jurnal | di dashboard melihat **daftar jurnal yang saya submit** (judul, tanggal submit, status: Draft / Menunggu Verifikasi / Diverifikasi / Ditolak / Published) | memantau status jurnal saya dan melanjutkan edit atau submit baru |
| 5 | Widyaprada | penulis jurnal | mengklik satu jurnal dari daftar untuk ke Detail atau Edit jurnal saya | mengedit draft atau melihat feedback jurnal yang ditolak/diverifikasi |
| 6 | Widyaprada | pengguna dashboard | melihat pesan/state kosong yang jelas jika belum ada penugasan ujian atau belum ada jurnal yang di-submit | tidak bingung saat data kosong |

**Keterangan**: Assignment = penugasan ujian (sumber: Manajemen Uji Kompetensi – daftar peserta). Jurnal yang di-submit = jurnal yang **user ini** buat/submit (sumber: WPJurnal – filter by current user).

---

## 3. Isi Utama Dashboard

### 3.1 Blok 1: Daftar Assignment Uji Kompetensi (Tugas Saya)

**Deskripsi**: Bagian dashboard yang menampilkan **daftar ujian yang di-assign ke user** (user termasuk peserta ujian). Sumber data dan aturan sama dengan fitur "Tugas Saya" di PRD Assignment; di dashboard bisa ditampilkan sebagai **ringkasan** (mis. 5–10 terdekat) dengan link "Lihat semua" ke halaman Tugas Saya penuh.

**Elemen yang Harus Ada**:
- **Judul blok**: "Tugas Saya" atau "Penugasan Uji Kompetensi" atau "Ujian Saya".
- **Daftar (kartu atau baris) per assignment**: **Nama Ujian**, **Batas Waktu** (Jadwal Selesai), **Status** (Belum dikerjakan / Sudah dikerjakan), **Hasil** (nilai jika sudah ada; atau "-" / "Belum keluar").
- **Aksi per item**: "Mulai Ujian" (jika masih dalam periode dan belum submit) → ke alur CBT; "Lihat Hasil" (jika sudah submit dan hasil sudah keluar); "Lihat Leaderboard" (hanya jika ujian diset Leaderboard dan hasil sudah keluar).
- **Link "Lihat semua"**: Ke halaman Tugas Saya lengkap (list penuh dengan filter/sort sesuai PRD Assignment).
- **State kosong**: "Anda belum memiliki penugasan ujian." jika user tidak termasuk peserta ujian mana pun (atau tidak ada ujian yang memenuhi kriteria tampil).

**Aturan tampil**:
- Hanya ujian yang **user termasuk pesertanya**; hanya ujian dengan status Diterbitkan/Berlangsung/Selesai yang relevan (sesuai PRD Assignment & Manajemen Uji Kompetensi).
- Urutan: mis. batas waktu terdekat dulu; atau belum dikerjakan dulu. Jumlah di dashboard dibatasi (mis. 5–10) agar halaman ringkas.

---

### 3.2 Blok 2: Daftar Jurnal yang Di-Submit (Jurnal Saya)

**Deskripsi**: Bagian dashboard yang menampilkan **daftar jurnal yang di-submit oleh user yang login** (jurnal yang user ini buat/submit). Setiap item menampilkan judul, tanggal submit (atau terakhir diubah), dan **status** (Draft / Menunggu Verifikasi / Diverifikasi / Ditolak / Published). User dapat mengklik untuk ke Detail atau Edit jurnal (sesuai wewenang di WPJurnal).

**Elemen yang Harus Ada**:
- **Judul blok**: "Jurnal Saya" atau "Jurnal yang Saya Submit" atau "Submitan Jurnal".
- **Daftar (kartu atau baris) per jurnal**: **Judul jurnal**, **Tanggal submit** (atau tanggal terakhir diubah), **Status** (Draft / Menunggu Verifikasi / Diverifikasi / Ditolak / Published). Opsional: snippet/abstrak singkat.
- **Aksi per item**: "Lihat" / "Detail" → halaman Detail jurnal; "Edit" (jika status Draft atau sesuai kebijakan WPJurnal); "Submit" (jika Draft dan WPJurnal mendukung submit dari list).
- **Link "Lihat semua"**: Ke halaman lengkap "Jurnal Saya" (list penuh dengan filter status, sort, paginasi) jika halaman tersebut ada di WPJurnal.
- **Tombol "Buat Jurnal"** (opsional): Jika WPJurnal mengizinkan create dari dashboard, tombol ke form Buat/Submit Jurnal.
- **State kosong**: "Anda belum mengirimkan jurnal." atau "Belum ada jurnal yang Anda submit." jika user belum pernah submit jurnal.

**Aturan tampil**:
- Hanya jurnal yang **created_by / submitted_by = user yang login** (atau equivalen di model WPJurnal).
- Status mengikuti definisi di WPJurnal: Draft (belum submit), Menunggu Verifikasi (sudah submit, menunggu verifikator), Diverifikasi, Ditolak, Published.
- Urutan: mis. terbaru dulu (tanggal submit atau updated_at). Jumlah di dashboard dibatasi (mis. 5–10) dengan "Lihat semua".

---

## 4. Layout dan Urutan Dashboard

1. **Header dashboard**: Sapaan singkat (mis. "Selamat datang, [Nama]") atau judul "Dashboard" (opsional).
2. **Blok 1: Daftar Assignment Uji Kompetensi (Tugas Saya)** – di atas atau kiri (prioritas tinggi agar deadline ujian terlihat).
3. **Blok 2: Daftar Jurnal yang Di-Submit (Jurnal Saya)** – di bawah atau kanan.
4. **Opsional**: Widget ringkasan angka (mis. "X ujian belum dikerjakan", "Y jurnal menunggu verifikasi") – fase berikutnya jika dibutuhkan.

Urutan dan tata letak (grid, stacking) dapat disesuaikan dengan design system; yang penting kedua blok jelas dan mudah di-scan.

---

## 5. Antarmuka Pengguna (UI) – Ringkas

### 5.1 Konsistensi
- Dashboard memakai **design system** yang sama dengan modul lain (Tailwind, shadcn/ui, dsb.).
- **Responsif**: Di mobile, blok bisa ditumpuk vertikal; tombol dan link tetap dapat diklik dengan nyaman.
- **Loading**: Saat data assignment atau jurnal dimuat, tampilkan skeleton atau spinner untuk masing-masing blok; jangan block seluruh halaman jika satu sumber lambat.

### 5.2 Navigasi
- Dari dashboard, **Tugas Saya** (list penuh), **CBT**, **Jurnal Saya** (list penuh), **Buat Jurnal**, dan **List/Detail Jurnal (publik)** dapat diakses via menu sidebar atau link di dalam blok. Menu mengikuti wewenang role (PRD RBAC / Auth).

### 5.3 Umpan Balik
- Setelah aksi (mis. kembali dari CBT submit), dashboard dapat menampilkan pesan sukses singkat (toast/alert) jika diperlukan.
- Jika tidak ada data: pesan state kosong per blok seperti disebutkan di §3.1 dan §3.2.

---

## 6. Acceptance Criteria (Ringkas)

**Dashboard secara umum:**  
(1) Halaman Dashboard Widyaprada hanya dapat diakses oleh user dengan role Widyaprada (atau role yang diizinkan melihat dashboard ini).  
(2) Dashboard menampilkan dua blok utama: **Daftar Assignment Uji Kompetensi** dan **Daftar Jurnal yang Di-Submit**.  
(3) Layout responsif; state kosong per blok jelas.

**Blok Assignment (Tugas Saya):**  
(1) Hanya menampilkan ujian yang **user termasuk pesertanya**; aturan status/jadwal sesuai PRD Assignment.  
(2) Setiap item menampilkan: Nama Ujian, Batas Waktu, Status (Belum/Sudah dikerjakan), Hasil (jika ada).  
(3) Aksi "Mulai Ujian", "Lihat Hasil", "Lihat Leaderboard" (jika berlaku) berfungsi dan mengarah ke alur yang benar.  
(4) Link "Lihat semua" ke halaman Tugas Saya lengkap.  
(5) Jika tidak ada penugasan: tampilkan pesan "Anda belum memiliki penugasan ujian."

**Blok Jurnal Saya:**  
(1) Hanya menampilkan jurnal yang **di-submit oleh user yang login**.  
(2) Setiap item menampilkan: Judul, Tanggal submit/ubah, Status (Draft / Menunggu Verifikasi / Diverifikasi / Ditolak / Published).  
(3) Aksi "Lihat", "Edit" (jika diizinkan) berfungsi dan mengarah ke Detail/Edit jurnal sesuai WPJurnal.  
(4) Link "Lihat semua" ke halaman Jurnal Saya lengkap (jika ada); opsional tombol "Buat Jurnal".  
(5) Jika belum ada submitan: tampilkan pesan "Anda belum mengirimkan jurnal." atau setara.

---

## 7. Cakupan Fitur

### 7.1 Termasuk
- **Dashboard Widyaprada**: Halaman utama (atau menu Dashboard) untuk role Widyaprada dengan dua blok: **Daftar Assignment Uji Kompetensi (Tugas Saya)** dan **Daftar Jurnal yang Di-Submit (Jurnal Saya)**.
- **Tugas Saya di dashboard**: Ringkasan penugasan ujian (nama, batas waktu, status, aksi); link ke list penuh dan ke CBT/Hasil/Leaderboard.
- **Jurnal Saya di dashboard**: Ringkasan jurnal yang user submit (judul, tanggal, status); link ke list penuh, Detail, Edit, dan opsional Buat Jurnal.
- Integrasi dengan **PRD Assignment** (data & aturan Tugas Saya) dan **WPJurnal** (data & status jurnal per user).

### 7.2 Tidak Termasuk
- **Alur lengkap Tugas Saya** (list penuh, filter, sort, paginasi) → PRD Assignment.
- **Alur lengkap CBT** (Mulai Ujian, timer, submit, nilai) → PRD CBT.
- **Manajemen Jurnal lengkap** (Create/Submit, Edit, Delete, Verifikasi) → PRD WPJurnal (ongoing).
- **List/Detail Jurnal publik** (jurnal published untuk semua pembaca) → PRD Jurnal.
- **Beranda landing** (slider, panel berita, tautan, panel jurnal) → PRD Beranda (untuk pengunjung/landing, bukan dashboard user).

---

## 8. Persyaratan Produk (Nonteknis)

- **Sumber data**: Assignment dari Manajemen Uji Kompetensi (peserta per ujian); jurnal dari WPJurnal (jurnal per user/creator). Dashboard tidak menyimpan data sendiri; hanya agregasi tampilan.
- **Wewenang**: Hanya user dengan role yang diizinkan (minimal Widyaprada) yang mengakses Dashboard Widyaprada; daftar assignment hanya ujian yang user adalah peserta; daftar jurnal hanya jurnal yang user submit.
- **Konsistensi**: Status dan aksi (Mulai Ujian, Lihat Hasil, Leaderboard, Edit jurnal) harus konsisten dengan aturan di PRD Assignment, PRD CBT, dan WPJurnal.
- **Performa**: Batasi jumlah item per blok di dashboard (mis. 5–10) agar halaman cepat; "Lihat semua" untuk data lengkap.

---

## 9. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: Dashboard Widyaprada dengan Daftar Assignment Uji Kompetensi & Daftar Jurnal yang Di-Submit | - |

---

**Catatan**: Detail teknis (API, schema, endpoint) untuk dashboard dan integrasi dengan modul ujian & jurnal didokumentasikan di SDD. Lihat [PRD_Assignment.md](PRD_Assignment.md), [PRD_Jurnal.md](PRD_Jurnal.md), PRD Manajemen Uji Kompetensi, dan PRD WPJurnal (saat tersedia).
