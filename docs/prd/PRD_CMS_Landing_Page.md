# [PRD] CMS Landing Page
## Product Requirements Document | CMS Slider, CMS Berita, CMS Tautan

**Aplikasi**: Widyaprada  
**Modul**: CMS Landing Page  
**Fitur**: CMS Slider (List, Detail, Create, Edit, Delete), CMS Berita, CMS Tautan (masing-masing full CRUD)  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: CMS Landing Page (CMS Slider, CMS Berita, CMS Tautan)
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
CMS Landing Page memungkinkan Admin Satker dan Super Admin mengelola konten yang tampil di landing page (beranda): **slider/slide** (hero slider di beranda), **berita**, dan **tautan**. Masing-masing punya fitur List, Detail, Create, Edit, Delete sehingga pengelola dapat menambah, mengubah, dan menghapus slide, berita, serta tautan sesuai wewenang (Admin Satker dalam scope satker; Super Admin seluruh sistem).

---

## 2. CMS Slider (Pengelolaan Konten Slider/Slide)

Slider menampilkan beberapa **slide** di beranda (hero). Setiap slide dikelola sebagai satu record (gambar, judul, deskripsi, link, urutan, status).

### 2.1 Field/Kolom Data Slide

| No | Field / Kolom | Wajib | Tipe | Keterangan |
|----|----------------|-------|------|------------|
| 1 | **Gambar/Media** | Ya | Upload/URL | Gambar utama slide (hero). Rasio dan resolusi disarankan konsisten (mis. 1920×800 atau 16:9). |
| 2 | **Judul** | Opsional | Teks | Judul yang tampil di atas gambar (heading). |
| 3 | **Subjudul / Deskripsi** | Opsional | Teks panjang | Teks pendukung atau deskripsi singkat di slide. |
| 4 | **Tautan (URL)** | Opsional | URL | Link saat slide atau tombol CTA diklik. Bisa internal atau eksternal. |
| 5 | **Label tombol CTA** | Opsional | Teks | Teks tombol (mis. "Selengkapnya", "Baca lebih lanjut"). Kosong = tidak tampil tombol. |
| 6 | **Urutan** | Ya | Angka | Urutan tampil slide (1, 2, 3, …). Slide dengan urutan lebih kecil tampil lebih dulu. |
| 7 | **Status** | Ya | Pilihan | Draft / Published. Hanya slide **Published** yang tampil di beranda. |
| 8 | **Tanggal mulai tampil** | Opsional | Tanggal | Jika diisi, slide hanya tampil di beranda pada atau setelah tanggal ini. |
| 9 | **Tanggal selesai tampil** | Opsional | Tanggal | Jika diisi, slide hanya tampil sampai tanggal ini. |
| 10 | **Tanggal Dibuat / Diubah** | Sistem | Timestamp | Otomatis. |

**Catatan**: Rekomendasi UX hero slider: 3–5 slide; gambar dioptimasi untuk performa (LCP). Konten slide dikelola penuh via CMS Slider (Create, Edit, Delete).

### 2.2 CMS Slider – List
- **Judul halaman**: Mis. "CMS Slider" atau "Kelola Slide".
- **Tombol "Tambah Slide"**: Ke form Create.
- **Tabel/List:** Kolom: Gambar (thumbnail), Judul, Urutan, Status, Tanggal mulai/selesai (opsional). Aksi: Detail, Edit, Hapus.
- **Sort/urutkan**: Minimal menurut urutan. Paginasi jika banyak slide.
- **Umpan balik**: Loading, pesan kosong jika belum ada slide, pesan sukses/error setelah Delete.

### 2.3 CMS Slider – Detail
- Menampilkan semua field satu slide (gambar, judul, subjudul, tautan, label CTA, urutan, status, periode tampil, tanggal dibuat/diubah).
- Tombol Edit, Hapus, Kembali ke List.

### 2.4 CMS Slider – Create / Edit
- **Create**: Form dengan field di 2.1; Gambar wajib; Urutan dan Status wajib.
- **Edit**: Form sama, data terisi; bisa ubah urutan dan status.
- **Pratinjau gambar**: Di field URL Gambar, tampilkan **thumbnail pratinjau** gambar sesuai URL yang diisi (dari data existing atau input pengguna). Jika gagal memuat, tampilkan pesan "Gagal memuat gambar". Bermanfaat agar pengelola bisa memastikan gambar yang dipasang sesuai sebelum menyimpan.
- Tombol Simpan, Batal. Validasi: gambar ada; urutan numerik; URL valid jika diisi. Umpan balik: loading saat simpan, pesan sukses/error.

### 2.5 CMS Slider – Delete
- Konfirmasi sebelum hapus (dialog). Setelah konfirmasi: hapus record, kembali ke List atau refresh list; pesan sukses/error. Opsional: wajib isi alasan penghapusan (sesuai kebijakan audit).

---

## 3. CMS Berita

Berita ditampilkan di Panel Berita beranda dan di halaman List/Detail Berita. Hanya berita dengan status **published** yang tampil di beranda dan list publik.

### 3.1 Field/Kolom Data Berita (Referensi Standar CMS Artikel)

| No | Field / Kolom | Wajib | Tipe | Keterangan |
|----|----------------|-------|------|------------|
| 1 | **Judul** | Ya | Teks | Judul berita. |
| 2 | **Slug** | Opsional/Sistem | Teks | Identifier unik untuk URL (biasanya dari judul, auto-generate). |
| 3 | **Konten/Isi** | Ya | Rich text/HTML | Isi lengkap berita. |
| 4 | **Ringkasan/Excerpt** | Opsional | Teks panjang | Snippet untuk panel beranda dan list. |
| 5 | **Gambar/Thumbnail** | Opsional | Upload/URL | Gambar utama berita. |
| 6 | **Tanggal Publikasi** | Opsional | Tanggal-waktu | Kapan berita dipublikasikan; kosong = draft atau jadwal. |
| 7 | **Status** | Ya | Pilihan | Draft / Published. Hanya **Published** tampil di beranda dan list publik. |
| 8 | **Penulis** | Opsional | Teks/Referensi | Nama atau ID penulis. |
| 9 | **Kategori/Tag** | Opsional | Pilihan ganda/teks | Kategori atau tag untuk filter. |
| 10 | **Tanggal Dibuat / Diubah** | Sistem | Timestamp | Otomatis. |

### 3.2 CMS Berita – List
- Judul "CMS Berita" / "Daftar Berita". Tombol "Tambah Berita". Search (judul, isi). Filter (status, kategori). Tabel: Judul, Thumbnail, Tanggal Publikasi, Status, Penulis (opsional). Aksi: Detail, Edit, Hapus. Paginasi, sort.

### 3.3 CMS Berita – Detail
- Tampilan read-only: semua field berita. Tombol Edit, Hapus, Kembali.

### 3.4 CMS Berita – Create / Edit
- Form dengan field di 3.1. Validasi: judul dan konten wajib; slug unik jika dipakai. Simpan → redirect ke Detail atau List; pesan sukses/error.

### 3.5 CMS Berita – Delete
- Konfirmasi; setelah hapus, pesan sukses/error; refresh list atau redirect.

---

## 4. CMS Tautan

Tautan ditampilkan di Panel Tautan beranda. Digunakan untuk link ke dokumen, situs eksternal, atau halaman internal.

### 4.1 Field/Kolom Data Tautan

| No | Field / Kolom | Wajib | Tipe | Keterangan |
|----|----------------|-------|------|------------|
| 1 | **Judul/Label** | Ya | Teks | Teks yang tampil untuk link (mis. "Portal Kemendikbud"). |
| 2 | **URL** | Ya | URL | Alamat tujuan (internal atau eksternal). |
| 3 | **Deskripsi** | Opsional | Teks | Keterangan singkat (untuk admin atau tooltip). |
| 4 | **Urutan** | Opsional | Angka | Urutan tampil di panel (angka kecil = di atas). |
| 5 | **Status** | Ya | Pilihan | Aktif / Nonaktif. Hanya **Aktif** tampil di beranda. |
| 6 | **Buka di tab baru** | Opsional | Boolean | Jika ya, link dibuka di tab baru (disarankan untuk URL eksternal). |
| 7 | **Tanggal Dibuat / Diubah** | Sistem | Timestamp | Otomatis. |

### 4.2 CMS Tautan – List
- Judul "CMS Tautan" / "Daftar Tautan". Tombol "Tambah Tautan". Search (judul). Filter (status). Tabel: Judul, URL, Urutan, Status. Aksi: Detail, Edit, Hapus. Paginasi, sort.

### 4.3 CMS Tautan – Detail
- Tampilan semua field. Tombol Edit, Hapus, Kembali.

### 4.4 CMS Tautan – Create / Edit
- Form field di 4.1. Validasi: judul dan URL wajib; URL format valid. Simpan → pesan sukses/error; redirect.

### 4.5 CMS Tautan – Delete
- Konfirmasi; hapus; pesan sukses/error.

---

## 5. Kebutuhan per Role (Ringkas)

- **Admin Satker:** List, Detail, Create, Edit, Delete untuk Slider, Berita, Tautan **hanya dalam satker/unit kerjanya**. Data slide, berita, dan tautan satker lain tidak terlihat dan tidak bisa diubah oleh Admin Satker. Jika kebijakan produk menetapkan konten landing bersifat global (satu set untuk semua), maka wewenang Admin Satker bisa dibatasi (misalnya hanya view) sesuai keputusan produk.
- **Super Admin:** List, Detail, Create, Edit, Delete **penuh** untuk Slider, Berita, Tautan di seluruh sistem (semua satker).

Keputusan apakah data CMS dibatasi per satker atau global ditetapkan produk; setelah ditetapkan, perilaku wewenang (siapa lihat apa, siapa ubah apa) harus konsisten dan jelas bagi pengguna.

---

## 6. Cakupan Fitur CMS Landing Page

### 6.1 Termasuk
- **CMS Slider**: List, Detail, Create, Edit, Delete untuk slide (gambar, judul, subjudul, tautan, urutan, status, periode tampil). Konten slider beranda dikelola di sini.
- **CMS Berita**: List, Detail, Create, Edit, Delete untuk berita (judul, slug, konten, ringkasan, thumbnail, tanggal publikasi, status, penulis, kategori).
- **CMS Tautan**: List, Detail, Create, Edit, Delete untuk tautan (judul, URL, deskripsi, urutan, status, buka di tab baru).

### 6.2 Tidak Termasuk
- Tampilan beranda (slider, panel berita/tautan/jurnal) → PRD Beranda.
- Halaman List/Detail Berita dan Jurnal untuk publik → PRD Berita, PRD Jurnal.
- WPJurnal (manajemen jurnal, verifikasi) → PRD WPJurnal.

---

## 7. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: CMS Slider, CMS Berita, CMS Tautan (masing-masing List, Detail, Create, Edit, Delete) | - |
| 1.1 | 2025-03-01 | CMS Slider Create/Edit: tambah pratinjau thumbnail gambar sesuai URL | - |

---

**Catatan**: **CMS Slider** mengelola konten **slide** yang ditampilkan di **slider besar beranda** (PRD Beranda). Hanya slide dengan status Published (dan dalam periode tampil jika diisi) yang muncul di beranda.
