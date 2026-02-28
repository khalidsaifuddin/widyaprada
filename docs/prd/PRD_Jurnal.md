# [PRD] Fitur Jurnal (List & Detail)
## Product Requirements Document | Landing Page – Halaman Jurnal

**Aplikasi**: Widyaprada  
**Modul**: Landing Page  
**Fitur**: Jurnal (List, Detail): halaman tampilan publik  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Jurnal (List & Detail)
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Fitur Jurnal menyediakan **halaman List Jurnal** dan **halaman Detail Jurnal** untuk pengunjung/pengguna. Keduanya bersifat **tampilan publik**: hanya menampilkan jurnal yang **sudah dipublikasikan** (status Published). Data jurnal bersumber dari modul **WPJurnal** (PRD WPJurnal masih ongoing). Untuk saat ini, halaman List dan Detail Jurnal di landing page **hanya menampilkan jurnal yang sudah published**; pengelolaan jurnal (submit, verifikasi, edit) ada di modul WPJurnal.

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Pengunjung / Semua | pembaca jurnal | melihat daftar jurnal yang sudah dipublikasikan dalam satu halaman (list) | memilih jurnal yang ingin dibaca |
| 2 | Pengunjung / Semua | pembaca jurnal | mencari jurnal berdasarkan judul, penulis, atau kata kunci | menemukan jurnal tertentu dengan cepat |
| 3 | Pengunjung / Semua | pembaca jurnal | memfilter jurnal (mis. tahun, kategori/jenis jika ada) | mempersempit daftar sesuai kebutuhan |
| 4 | Pengunjung / Semua | pembaca jurnal | mengurutkan daftar (mis. terbaru dulu) | melihat jurnal paling baru atau paling lama |
| 5 | Pengunjung / Semua | pembaca jurnal | mengklik satu jurnal dari list untuk membuka halaman Detail | membaca ringkasan/isi jurnal lengkap |
| 6 | Pengunjung / Semua | pembaca jurnal | di halaman Detail melihat judul, penulis, abstrak, tanggal publish, dan metadata jurnal | memahami jurnal secara lengkap |
| 7 | Pengunjung / Semua | pembaca jurnal | paginasi atau "load more" jika jurnal banyak | menjelajah tanpa halaman berat |

**Keterangan**: Data jurnal **dikelola** di modul **WPJurnal** (Manajemen Jurnal: List, Detail, Create/Submit, Edit, Delete, Verifikasi). PRD ini hanya mengatur **halaman tampilan publik** List Jurnal dan Detail Jurnal di landing page, dengan sumber data jurnal yang **statusnya published saja**.

---

## 3. Sumber Data dan Relasi ke WPJurnal

- **Sumber data**: Modul **WPJurnal** (PRD WPJurnal ongoing). Field lengkap jurnal (judul, penulis, abstrak, tanggal, status, dll) akan mengikuti definisi di PRD WPJurnal.
- **Yang tampil di List & Detail Jurnal (landing)**: Hanya jurnal dengan **status Published**. Jurnal draft, submit, atau menunggu verifikasi **tidak** tampil di halaman List/Detail Jurnal publik ini.
- **Urutan default**: Tanggal publish terbaru dulu (atau tanggal yang setara di model WPJurnal).

---

## 4. Antarmuka Pengguna (UI)

### 4.1 Halaman List Jurnal

**Deskripsi**: Halaman yang menampilkan daftar jurnal published. Pengunjung dapat mencari, memfilter, mengurutkan, dan mengklik satu item untuk ke Detail.

**Elemen yang Harus Ada**:
- **Judul halaman**: Mis. "Jurnal" atau "Daftar Jurnal".
- **Kotak pencarian**: Mencari berdasarkan judul, penulis, atau kata kunci (sesuai field yang tersedia di WPJurnal).
- **Filter** (opsional): Tahun, kategori/jenis jurnal (jika ada di model WPJurnal).
- **Sort**: Pilihan urutan (mis. Terbaru, Terlama); default **Terbaru dulu**.
- **Daftar jurnal:** Setiap item menampilkan minimal Judul, Penulis (atau ringkasan penulis), Tanggal/Tahun, Abstrak/snippet (opsional). Klik item atau "Baca selengkapnya" mengarah ke Detail Jurnal.
- **Paginasi** atau **Load more**: Jika data banyak.

**Tata Letak dan Keterbacaan**:
- Layout rapi (card atau list); responsif (mobile-friendly).

**Umpan Balik**:
- Saat data dimuat: skeleton atau spinner.
- Jika tidak ada hasil search/filter: pesan "Tidak ada jurnal yang sesuai."
- Jika belum ada jurnal published: pesan "Belum ada jurnal."

---

### 4.2 Halaman Detail Jurnal

**Deskripsi**: Halaman yang menampilkan **satu jurnal lengkap** (read-only). Dapat diakses dari List Jurnal, dari Panel Jurnal di Beranda, atau via URL langsung.

**Elemen yang Harus Ada**:
- **Judul jurnal**: Heading utama.
- **Meta**: Penulis, tanggal/tahun publikasi, identifikasi (DOI/ISSN jika ada).
- **Abstrak**: Ringkasan jurnal.
- **Konten/isi lengkap** (jika ada): Isi full text atau link ke dokumen (sesuai model WPJurnal).
- **Navigasi**: Tombol/link "Kembali ke Daftar Jurnal" atau "Jurnal".

**Tata Letak dan Keterbacaan**:
- Konten terbaca jelas; tidak ada elemen edit/verifikasi (ini halaman publik).

**Umpan Balik**:
- Saat data dimuat: skeleton atau spinner.
- Jika jurnal tidak ditemukan atau belum published: pesan "Jurnal tidak ditemukan." dan opsi kembali ke List Jurnal.

---

## 5. Acceptance Criteria (Ringkas)

**List Jurnal:** (1) Hanya jurnal dengan status Published yang tampil. (2) Pencarian berdasarkan judul, penulis, atau kata kunci berfungsi. (3) Filter tahun dan kategori (jika ada) berfungsi. (4) Sort default terbaru dulu; opsi terlama tersedia. (5) Paginasi atau load more jika data banyak. (6) Klik item membuka Detail Jurnal. (7) State kosong: "Belum ada jurnal" atau "Tidak ada jurnal yang sesuai" jika filter/search tidak ada hasil.

**Detail Jurnal:** (1) Menampilkan judul, meta (penulis, tanggal/tahun, DOI/ISSN jika ada), abstrak, konten/link. (2) Link "Kembali ke Daftar Jurnal" ke List. (3) Jika jurnal tidak ditemukan atau tidak published: pesan "Jurnal tidak ditemukan" dan opsi kembali ke List. (4) Tidak ada elemen edit/verifikasi (halaman publik).

**Sumber data:** Field jurnal (judul, penulis, abstrak, tanggal, status, dll) mengikuti definisi di modul WPJurnal. PRD WPJurnal belum final; untuk implementasi awal cukup tampilkan field yang sudah disepakati (judul, penulis, abstrak, tanggal publish). Detail struktur lengkap ada di PRD WPJurnal atau SDD.

---

## 6. Cakupan Fitur Jurnal (List & Detail)

### 6.1 Termasuk
- **List Jurnal:** Halaman daftar jurnal published saja; search; filter (tahun, kategori jika ada); sort (terbaru/terlama); paginasi atau load more; link ke Detail.
- **Detail Jurnal:** Halaman satu jurnal lengkap (judul, penulis, abstrak, tanggal, konten/link); link kembali ke List.

### 6.2 Tidak Termasuk
- **Manajemen Jurnal** (List, Detail, Create/Submit, Edit, Delete, Verifikasi) → PRD WPJurnal.
- Panel Jurnal di Beranda → PRD Beranda (hanya menampilkan ringkasan; link ke List/Detail Jurnal).

---

## 7. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: List Jurnal, Detail Jurnal (tampilan publik; hanya published). Relasi ke WPJurnal (ongoing) | - |

---

**Catatan**: Halaman **List Jurnal** dan **Detail Jurnal** ini adalah halaman **publik** untuk membaca jurnal. Hanya jurnal dengan **status published** yang ditampilkan. Pengelolaan dan verifikasi jurnal ada di modul **WPJurnal** (PRD WPJurnal). Field detail jurnal (judul, penulis, abstrak, dll) akan mengikuti model data yang ditetapkan di PRD WPJurnal.
