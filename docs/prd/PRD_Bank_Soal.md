# [PRD] Fitur Bank Soal
## Product Requirements Document | WPUjikom

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom (Uji Kompetensi)  
**Fitur**: Bank Soal (List, Detail, Create, Edit, Delete)  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Bank Soal
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Bank Soal memungkinkan pengelola untuk mengelola kumpulan soal uji kompetensi dalam satu tempat. **Hanya Super Admin** yang berhak **menginput dan migrasi soal** (Create, Edit, Delete, import massal). Admin Uji Kompetensi dan Verifikator dapat melihat, memfilter, dan memverifikasi soal — tetapi tidak menambah atau mengubah soal. Pengelola dapat mencari dan memfilter soal berdasarkan kategori kompetensi, tipe soal, dan tingkat kesulitan. Soal mendukung tipe pilihan ganda (PG), benar–salah, dan essay (uraian) agar sesuai dengan kebutuhan asesmen kompetensi jabatan fungsional Widyaprada.

### 1.3 Role yang Mengakses
- **Super Admin**: **Hanya Super Admin** yang berhak **menginput dan migrasi soal** (Create, Edit, Delete, import massal). Akses penuh ke Bank Soal.
- **Admin Uji Kompetensi**: Akses untuk **melihat** list/detail soal, memfilter, menggunakan soal di paket ujian. **Tidak** Create/Edit/Delete soal (hanya Super Admin).
- **Verifikator Uji Kompetensi**: Akses untuk **verifikasi soal** — melihat list/detail soal dan melakukan aksi **Verifikasi** (memberi label soal sudah diverifikasi). Tidak Create/Edit/Delete soal.

**Catatan**: Input dan migrasi soal hanya Super Admin. Admin Uji Kompetensi dan Verifikator dapat melihat dan memverifikasi, tetapi tidak menambah atau mengubah soal. Peserta ujian tidak mengakses Bank Soal; mereka hanya mengerjakan ujian (lihat PRD CBT).

### 1.4 Field/Kolom Soal (Referensi)

| No | Field / Kolom | Wajib | Tipe | Keterangan |
|----|----------------|-------|------|------------|
| 1 | **Kode Soal** | Ya | Teks (unik) | Identifikasi singkat soal, misalnya SOAL-001 atau per kategori. Untuk memudahkan referensi saat menyusun paket. |
| 2 | **Tipe Soal** | Ya | Pilihan | PG (Pilihan Ganda), Benar–Salah, Essay (Uraian). Menentukan tampilan dan cara koreksi. |
| 3 | **Kategori Kompetensi** | Ya | Pilihan/Referensi | Kategori atau dimensi kompetensi Widyaprada (mis. sesuai Permenpanrb 93/2020 atau standar internal). Untuk filter dan analisis hasil. |
| 4 | **Tingkat Kesulitan** | Opsional | Pilihan | Mudah / Sedang / Sukar. Untuk pemilihan soal berimbang saat menyusun paket. |
| 5 | **Teks Soal** | Ya | Teks panjang (bisa rich text) | Pertanyaan atau pernyataan yang ditampilkan ke peserta. Dapat berisi gambar jika kebijakan mendukung. |
| 6 | **Opsi Jawaban** (untuk PG) | Ya (jika PG) | Daftar opsi | Untuk PG: daftar opsi A, B, C, D (atau lebih). Setiap opsi: teks (dan opsional gambar). |
| 7 | **Kunci Jawaban** | Ya | Tergantung tipe | PG: satu opsi yang benar (A/B/C/D). Benar–Salah: Benar atau Salah. Essay: kunci/model jawaban atau rubrik (untuk panduan koreksi). |
| 8 | **Bobot Nilai** | Opsional | Angka | Nilai per soal (default misalnya 1). Untuk perhitungan skor. |
| 9 | **Status** | Ya | Pilihan | Draft / Aktif. Soal Aktif bisa dipakai di paket ujian; Draft masih bisa diedit tanpa memengaruhi ujian yang sudah berjalan. |
| 10 | **Status Verifikasi** | Ya | Pilihan | **Belum Diverifikasi** / **Sudah Diverifikasi**. Label untuk membedakan soal yang sudah diperiksa (diverifikasi) dan yang belum. Hanya role Verifikator Uji Kompetensi (dan Super Admin) yang dapat mengubah status verifikasi. |
| 11 | **Keterangan** | Opsional | Teks | Catatan internal (mis. sumber, tahun, revisi). |
| 12 | **Tanggal Dibuat** | Sistem | Timestamp | Diisi otomatis. |
| 13 | **Tanggal Diubah** | Sistem | Timestamp | Diisi otomatis saat diperbarui. |

**Catatan implementasi**:
- **Essay**: Koreksi dapat manual (nilai diinput oleh penguji) atau semi-otomatis jika ada rubrik; detail koreksi essay didokumentasikan di SDD atau PRD Manajemen Uji Kompetensi/CBT jika memengaruhi alur peserta.
- **Gambar pada soal/opsi**: In scope jika produk menghendaki; jika fase pertama tanpa gambar, field cukup teks.

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Super Admin | pengelola bank soal | **menginput dan migrasi soal** (Create, Edit, Delete, import massal) | mengisi dan mengelola bank soal — hanya Super Admin yang berhak |
| 2 | Admin Uji Kompetensi / Super Admin | pengelola bank soal | melihat daftar soal dalam satu halaman (kode, tipe, kategori, tingkat kesulitan, status) | dengan cepat menemukan dan memantau soal yang tersedia |
| 3 | Admin Uji Kompetensi / Super Admin | pengelola bank soal | mencari soal berdasarkan kode atau teks soal | menemukan soal tertentu tanpa menggulir panjang |
| 4 | Admin Uji Kompetensi / Super Admin | pengelola bank soal | memfilter daftar berdasarkan tipe soal, kategori kompetensi, tingkat kesulitan, atau status | fokus pada kelompok soal yang relevan saat menyusun ujian |
| 5 | Admin Uji Kompetensi / Super Admin | pengelola bank soal | mengurutkan daftar (kode, tanggal dibuat, kategori) | mengorganisir tampilan sesuai kebutuhan |
| 6 | Admin Uji Kompetensi / Super Admin | pengelola bank soal | membuka detail satu soal dari list | melihat konten lengkap dan kunci jawaban sebelum memakai di paket |
| 7 | Super Admin | pengelola bank soal | menambah soal baru (Create) dengan form yang jelas | mengisi bank soal — hanya Super Admin |
| 8 | Super Admin | pengelola bank soal | mengubah soal yang ada (Edit) dari detail atau list | memperbarui konten atau kunci — hanya Super Admin |
| 9 | Super Admin | pengelola bank soal | menghapus soal (Delete) dengan konfirmasi dan alasan | membersihkan soal — hanya Super Admin |
| 10 | Super Admin | pengelola bank soal | mendapat umpan balik jelas setelah Create/Edit/Delete (sukses atau pesan error) | tahu apakah aksi berhasil |
| 11 | Admin Uji Kompetensi / Super Admin | pengelola bank soal | paginasi pada list jika data banyak | daftar tetap cepat dan nyaman dipakai |
| 12 | Admin Uji Kompetensi / Super Admin | pengelola bank soal | tombol/aksi menampilkan loading saat diproses | yakin bahwa aplikasi sedang bekerja dan tidak hang |
| 13 | Super Admin | pengelola bank soal | saat Delete, wajib mengisi alasan/deskripsi penghapusan dalam dialog konfirmasi | ada jejak dokumentasi — hanya Super Admin yang Delete |
| 13 | Verifikator Uji Kompetensi / Super Admin | verifikator | melihat daftar soal (list) dan detail soal agar bisa memeriksa konten | memutuskan apakah soal layak diverifikasi |
| 14 | Verifikator Uji Kompetensi / Super Admin | verifikator | memfilter daftar soal berdasarkan status verifikasi (Belum / Sudah Diverifikasi) | fokus pada soal yang belum diverifikasi atau yang sudah |
| 15 | Verifikator Uji Kompetensi / Super Admin | verifikator | melakukan aksi Verifikasi pada soal (menandai sebagai Sudah Diverifikasi) dari detail atau list | label soal terverifikasi jelas untuk dipakai di paket/ujian |
| 16 | Verifikator Uji Kompetensi / Super Admin | verifikator | membatalkan verifikasi (kembalikan ke Belum Diverifikasi) jika kebijakan mengizinkan | koreksi jika soal salah diverifikasi |

---

### 2.1 Kebutuhan Pengguna - Yang Diinginkan
- **List**: Daftar soal terstruktur dengan kolom relevan (kode, tipe, kategori, tingkat kesulitan, status, **status verifikasi**), search, filter (termasuk filter **Status Verifikasi**), sort, paginasi.
- **Detail**: Satu halaman untuk melihat soal lengkap (teks, opsi, kunci, bobot, status, **status verifikasi**); untuk Verifikator: tombol **Verifikasi** / **Batalkan Verifikasi**.
- **Create**: Form jelas untuk tambah soal dengan validasi (kode unik, teks wajib, opsi dan kunci wajib untuk PG/B-S).
- **Edit**: Form ubah soal dengan validasi yang sama; soal yang sudah dipakai di ujian yang sedang berjalan dapat dibatasi (hanya metadata yang boleh diubah, atau peringatan).
- **Delete**: Konfirmasi + **wajib alasan penghapusan**; jika soal dipakai di paket ujian yang sudah ada, kebijakan (blok hapus / soft delete / peringatan) harus jelas.
- **Umpan balik**: Loading pada aksi; pesan sukses/error yang ramah.

### 2.2 Kebutuhan Pengguna - Yang Tidak Diinginkan
- List tanpa search/filter sehingga sulit menemukan soal saat bank besar.
- Form Create/Edit tanpa validasi jelas (field mana yang salah).
- Delete tanpa konfirmasi dan tanpa alasan penghapusan.
- Pesan error yang teknis atau membingungkan.

---

## 3. Antarmuka Pengguna (UI)

*Bank Soal terdiri atas: halaman List, halaman Detail, form Create, form Edit, aksi Delete, dan aksi Verifikasi/Batalkan Verifikasi. Semua mengikuti design system aplikasi. Menu Bank Soal terlihat oleh pengguna yang memiliki role Admin Uji Kompetensi, Verifikator Uji Kompetensi, atau Super Admin (satu user bisa punya lebih dari satu role).*

### 3.1 List Soal
- **Judul**: "Bank Soal" atau "Daftar Soal".
- **Tombol "Tambah Soal"**: Ke form Create.
- **Kotak pencarian**: Kode atau teks soal.
- **Filter**: Tipe Soal, Kategori Kompetensi, Tingkat Kesulitan, Status (Draft/Aktif), **Status Verifikasi** (Belum Diverifikasi / Sudah Diverifikasi).
- **Tabel**: Kolom minimal Kode, Tipe, Kategori, Tingkat Kesulitan, Status, **Status Verifikasi**; kolom aksi Detail, Edit, Hapus (Edit/Hapus **hanya untuk Super Admin**); untuk Verifikator: aksi Detail, **Verifikasi** / **Batalkan Verifikasi**.
- **Paginasi** jika data banyak.
- **Umpan balik**: Loading saat data dimuat; pesan kosong jika tidak ada hasil; pesan sukses/error setelah Delete.

### 3.2 Detail Soal
- **Judul**: "Detail Soal" atau kode soal.
- **Tampilan**: Teks soal, tipe, kategori, tingkat kesulitan; untuk PG: daftar opsi dan penanda kunci; untuk Essay: kunci/rubrik (jika ada). Bobot, status, **status verifikasi**, keterangan, tanggal dibuat/diubah.
- **Tombol**: Edit, Hapus, Kembali ke List (**hanya Super Admin**). Untuk **Verifikator Uji Kompetensi** (dan Super Admin): **Verifikasi** / **Batalkan Verifikasi** (sesuai status saat ini), Kembali ke List.

### 3.3 Create Soal (Form)
- **Judul**: "Tambah Soal".
- **Field**: Kode Soal, Tipe Soal, Kategori Kompetensi, Tingkat Kesulitan, Teks Soal; jika PG: Opsi A/B/C/D (atau dinamis) + pilihan Kunci; jika Benar–Salah: pilihan Kunci; jika Essay: Kunci/Rubrik (opsional). Bobot Nilai, Status, Keterangan.
- **Tombol**: Simpan, Batal.
- **Validasi**: Pesan per field (kode sudah dipakai, teks wajib, kunci wajib untuk PG/B-S).

### 3.4 Edit Soal (Form)
- Sama seperti Create dengan data terisi. Jika soal sudah dipakai di ujian, tampilkan peringatan atau batasi field yang boleh diubah (kebijakan produk).

### 3.5 Delete Soal
- **Pemicu**: Tombol Hapus di Detail atau List.
- **Dialog konfirmasi**: Pesan jelas + **field wajib Alasan/deskripsi penghapusan**. Tombol Batal dan Ya, Hapus (Hapus aktif setelah alasan diisi). Loading saat proses hapus.

---

## 4. Pengalaman Pengguna (UX) - Alur

### 4.1 Alur List
1. Admin Uji Kompetensi atau Super Admin membuka menu WPUjikom → Bank Soal.
2. Melihat daftar soal; dapat mencari, memfilter, mengurutkan.
3. Klik "Tambah Soal" → form Create; atau klik baris/Detail → Detail; atau Edit/Hapus dari list.

### 4.2 Alur Detail
1. Dari List, klik satu soal → halaman Detail.
2. Melihat konten lengkap; klik Edit untuk ubah atau Hapus untuk konfirmasi Delete.

### 4.3 Alur Create / Edit
1. Isi form; Simpan → loading → sukses: redirect ke List atau Detail; gagal: pesan error per field/umum.

### 4.4 Alur Delete
1. Klik Hapus → dialog dengan alasan wajib → isi alasan → Ya, Hapus → loading → sukses/error.

---

## 5. Kebutuhan per Role

### 5.1 Role: Admin Uji Kompetensi
- **Siapa**: Pengelola uji kompetensi (manajemen ujian, paket soal). **Tidak** menginput/migrasi soal — hanya Super Admin.
- **Bank Soal**: List, Detail (hanya melihat). **Tidak** Create, Edit, Delete soal. Dapat memverifikasi (jika juga punya role Verifikator).
- **Menu**: Melihat menu WPUjikom → Bank Soal (read-only untuk soal), Paket Soal, Manajemen Uji Kompetensi.

### 5.2 Role: Verifikator Uji Kompetensi
- **Siapa**: Pemberi label verifikasi pada soal, paket soal, dan ujian. Hanya aksi verifikasi (dan melihat list/detail), tidak Create/Edit/Delete.
- **Bank Soal**: List, Detail, aksi **Verifikasi** / **Batalkan Verifikasi**. Dapat memfilter berdasarkan status verifikasi.
- **Menu**: Melihat menu WPUjikom → Bank Soal, Paket Soal, Manajemen Uji Kompetensi (untuk verifikasi saja).

### 5.3 Role: Super Admin
- **Bank Soal**: **Hanya Super Admin** yang berhak **menginput dan migrasi soal** (Create, Edit, Delete, import massal). List, Detail, Create, Edit, Delete, Verifikasi/Batalkan Verifikasi (penuh); plus akses ke semua modul lain.

### 5.4 Role: Widyaprada
- **Tidak mengakses** Bank Soal (kecuali jika user juga punya role Admin Uji Kompetensi atau Verifikator). Sebagai peserta, hanya mengakses CBT sesuai PRD CBT.

---

## 6. Cakupan Fitur

### 6.1 Termasuk
- List soal dengan search, filter (tipe, kategori, tingkat kesulitan, status, **status verifikasi**), sort, paginasi.
- Detail soal (tampilan lengkap termasuk status verifikasi).
- **Create, Edit, Delete soal — hanya Super Admin** (input dan migrasi soal). Verifikasi/Batalkan Verifikasi (Verifikator Uji Kompetensi & Super Admin).
- Delete dengan konfirmasi dan **wajib alasan penghapusan** (Super Admin).
- **Status Verifikasi**: label Belum Diverifikasi / Sudah Diverifikasi; aksi verifikasi hanya untuk Verifikator Uji Kompetensi dan Super Admin.

### 6.2 Tidak Termasuk
- Paket Soal (kumpulan soal seperti playlist) → PRD Paket Soal.
- Manajemen Uji Kompetensi (jadwal ujian, konten: soal individu + paket soal, peserta) → PRD Manajemen Uji Kompetensi.
- CBT (mengerjakan ujian) → PRD CBT.
- Import/export soal massal (dapat ditambah di PRD terpisah atau fase berikutnya).
- Manajemen Pengguna, RBAC → PRD masing-masing.

---

## 7. Persyaratan Produk (Nonteknis)

- **Unik**: Kode Soal unik di sistem.
- **Delete**: Alasan penghapusan wajib diisi; jika soal sudah dipakai di paket ujian, kebijakan (blok / soft delete / peringatan) ditetapkan produk dan didokumentasikan.
- **Umpan balik**: Loading dan pesan sukses/error pada setiap aksi.
- **Akses**: Bank Soal dapat diakses oleh Admin Uji Kompetensi (read-only untuk soal), Verifikator (verifikasi), Super Admin (penuh). **Hanya Super Admin** yang berhak menginput dan migrasi soal (Create, Edit, Delete).

---

## 8. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: Bank Soal untuk Admin Uji Kompetensi & Super Admin; role Admin Uji Kompetensi | - |
| 1.1 | 2025-02-11 | Status Verifikasi (Belum/Sudah Diverifikasi); role Verifikator Uji Kompetensi; satu user bisa lebih dari satu role | - |
| 1.2 | 2025-02-28 | Hanya Super Admin yang berhak menginput dan migrasi soal (Create, Edit, Delete) | - |

---

**Catatan**: Role **Admin Uji Kompetensi** dan **Verifikator Uji Kompetensi** perlu ditambahkan di sistem RBAC oleh Super Admin. Satu pengguna dapat memiliki **lebih dari satu role** (misalnya Widyaprada + Admin Uji Kompetensi). Lihat PRD_RBAC.md, PRD_Manajemen_Uji_Kompetensi.md, PRD_Auth_Manajemen_Pengguna.md.
