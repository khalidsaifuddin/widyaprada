# [PRD] Fitur Berita (List & Detail)
## Product Requirements Document | Landing Page – Halaman Berita

**Aplikasi**: Widyaprada  
**Modul**: Landing Page  
**Fitur**: Berita (List, Detail): halaman tampilan publik  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Berita (List & Detail)
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Fitur Berita menyediakan **halaman List Berita** dan **halaman Detail Berita** untuk pengunjung/pengguna. Keduanya bersifat **tampilan publik**: hanya menampilkan berita yang **sudah dipublikasikan** (status Published) dari CMS Berita. Pengunjung dapat melihat daftar berita, mencari atau memfilter, lalu membuka satu berita untuk membaca isi lengkap. Pengelolaan konten (Create, Edit, Delete) dilakukan di **CMS Berita** (PRD CMS Landing Page).

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Pengunjung / Semua | pembaca berita | melihat daftar berita yang sudah dipublikasikan dalam satu halaman (list) | memilih berita yang ingin dibaca |
| 2 | Pengunjung / Semua | pembaca berita | mencari berita berdasarkan judul atau kata kunci | menemukan berita tertentu dengan cepat |
| 3 | Pengunjung / Semua | pembaca berita | memfilter berita berdasarkan kategori atau tanggal (jika ada) | mempersempit daftar sesuai kebutuhan |
| 4 | Pengunjung / Semua | pembaca berita | mengurutkan daftar (mis. terbaru dulu) | melihat berita paling baru atau paling lama |
| 5 | Pengunjung / Semua | pembaca berita | mengklik satu berita dari list untuk membuka halaman Detail | membaca isi lengkap berita |
| 6 | Pengunjung / Semua | pembaca berita | di halaman Detail melihat judul, tanggal, penulis, isi, dan gambar berita | memahami berita secara lengkap |
| 7 | Pengunjung / Semua | pembaca berita | paginasi atau "load more" jika berita banyak | menjelajah tanpa halaman berat |

**Keterangan**: Konten berita **dikelola** di **CMS Berita** (List, Detail, Create, Edit, Delete). PRD ini hanya mengatur **halaman tampilan publik** List Berita dan Detail Berita.

---

## 3. Sumber Data

- **CMS Berita** (PRD CMS Landing Page): field Judul, Slug, Konten, Ringkasan, Thumbnail, Tanggal Publikasi, Status, Penulis, Kategori.
- **Yang tampil di List & Detail Berita**: Hanya berita dengan **Status = Published**. Urutan default: **Tanggal Publikasi** terbaru dulu (atau tanggal dibuat jika tanggal publikasi kosong).

---

## 4. Antarmuka Pengguna (UI)

### 4.1 Halaman List Berita

**Deskripsi**: Halaman yang menampilkan daftar berita published dalam bentuk card list atau tabel. Pengunjung dapat mencari, memfilter, mengurutkan, dan mengklik satu item untuk ke Detail.

**Elemen yang Harus Ada**:
- **Judul halaman**: Mis. "Berita" atau "Daftar Berita".
- **Kotak pencarian**: Mencari berdasarkan judul atau kata kunci (full-text pada judul/ringkasan/konten sesuai kebijakan).
- **Filter** (opsional): Kategori, rentang tanggal publikasi.
- **Sort**: Pilihan urutan (mis. Terbaru, Terlama); default **Terbaru dulu**.
- **Daftar berita:** Setiap item menampilkan minimal Judul, Tanggal publikasi (atau tanggal dibuat), Ringkasan/snippet (opsional), Thumbnail (opsional). Klik item atau tombol "Baca selengkapnya" mengarah ke Detail Berita (via slug atau ID).
- **Paginasi** atau **Load more**: Jika data banyak agar performa tetap baik.

**Tata Letak dan Keterbacaan**:
- Layout rapi (card atau list); jarak nyaman; thumbnail dan teks tidak terpotong sembarangan.
- Responsif: di mobile tampilan menyesuaikan (satu kolom, card stack).

**Umpan Balik**:
- Saat data dimuat: skeleton atau spinner.
- Jika tidak ada hasil search/filter: pesan "Tidak ada berita yang sesuai."
- Jika belum ada berita published: pesan "Belum ada berita."

---

### 4.2 Halaman Detail Berita

**Deskripsi**: Halaman yang menampilkan **satu berita lengkap** (read-only). Dapat diakses dari List Berita, dari Panel Berita di Beranda, atau via URL langsung (mis. `/berita/{slug}`).

**Elemen yang Harus Ada**:
- **Judul berita**: Heading utama.
- **Meta**: Tanggal publikasi, penulis (jika ada), kategori (jika ada).
- **Gambar/Thumbnail**: Gambar utama berita (jika ada).
- **Konten/Isi**: Isi lengkap berita (rich text/HTML), terbaca dengan nyaman (typography, paragraf).
- **Navigasi**: Tombol/link "Kembali ke Daftar Berita" atau "Berita" agar kembali ke List.

**Tata Letak dan Keterbacaan**:
- Konten terbaca jelas; gambar responsif; tidak ada elemen edit/delete (ini halaman publik).

**Umpan Balik**:
- Saat data dimuat: skeleton atau spinner.
- Jika berita tidak ditemukan (slug salah atau berita di-unpublish): pesan "Berita tidak ditemukan." dan opsi kembali ke List Berita.

---

## 5. Acceptance Criteria (Ringkas)

**List Berita:** (1) Hanya berita dengan status Published yang tampil. (2) Pencarian berdasarkan judul/kata kunci berfungsi. (3) Filter kategori dan tanggal (jika ada) berfungsi. (4) Sort default terbaru dulu; opsi terlama tersedia. (5) Paginasi atau load more jika data banyak. (6) Klik item membuka Detail Berita. (7) State kosong: "Belum ada berita" atau "Tidak ada berita yang sesuai" jika filter/search tidak ada hasil.

**Detail Berita:** (1) Menampilkan judul, meta (tanggal, penulis, kategori), gambar, konten lengkap. (2) Link "Kembali ke Daftar Berita" ke List. (3) Jika berita tidak ditemukan atau tidak published: pesan "Berita tidak ditemukan" dan opsi kembali ke List. (4) Tidak ada elemen edit/delete (halaman publik).

---

## 6. Cakupan Fitur Berita (List & Detail)

### 6.1 Termasuk
- **List Berita:** Halaman daftar berita published; search; filter (kategori, tanggal jika ada); sort (terbaru/terlama); paginasi atau load more; link ke Detail.
- **Detail Berita:** Halaman satu berita lengkap (judul, meta, gambar, konten); link kembali ke List.

### 6.2 Tidak Termasuk
- **CMS Berita** (Create, Edit, Delete, List/Detail untuk admin) → PRD CMS Landing Page.
- Panel Berita di Beranda → PRD Beranda (hanya menampilkan ringkasan; link ke List/Detail Berita).

---

## 7. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: List Berita, Detail Berita (tampilan publik) | - |

---

**Catatan**: Halaman **List Berita** dan **Detail Berita** ini adalah halaman **publik** untuk membaca berita. Pengelolaan konten (tambah, ubah, hapus berita) ada di **CMS Berita** (PRD CMS Landing Page).
