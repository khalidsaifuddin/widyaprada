# [PRD] Fitur Paket Soal
## Product Requirements Document | WPUjikom

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom (Uji Kompetensi)  
**Fitur**: Paket Soal (List, Detail, Create, Edit, Delete)  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Paket Soal
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
**Paket Soal** adalah kumpulan soal (dari Bank Soal) yang disimpan menjadi satu entitas tersendiri — analogi seperti **playlist di pemutar musik**: satu playlist berisi banyak lagu, satu paket soal berisi banyak soal. Paket soal bisa dipakai ulang di banyak ujian tanpa harus memilih soal satu per satu setiap kali. **Admin Uji Kompetensi** dan **Super Admin** dapat membuat paket (nama, deskripsi, daftar soal), mengubah urutan soal dalam paket, menambah atau mengurangi soal, serta menghapus paket. Di **Manajemen Uji Kompetensi**, saat menyusun satu ujian, pengelola bisa mengisi konten ujian dengan **soal individu** (pilih dari Bank Soal satu per satu) dan/atau **satu atau lebih Paket Soal** (pilih paket yang sudah ada). Dengan demikian, satu uji kompetensi bisa memuat gabungan soal individu dan kumpulan soal dari paket.

### 1.3 Perbedaan: Soal vs Paket Soal vs Ujian

| Konsep | Penjelasan |
|--------|------------|
| **Soal** (Bank Soal) | Satu pertanyaan tunggal (PG, Benar–Salah, Essay). Bahan baku. |
| **Paket Soal** | Kumpulan beberapa soal yang disimpan jadi satu “playlist”; bisa dipakai di banyak ujian. |
| **Ujian (Uji Kompetensi)** | Satu event ujian dengan jadwal dan peserta; **konten**-nya = soal individu + paket soal (bisa campuran). |

### 1.4 Role yang Mengakses
- **Admin Uji Kompetensi**: Akses penuh ke Paket Soal (CRUD). Tidak mengubah status verifikasi.
- **Verifikator Uji Kompetensi**: Akses untuk **verifikasi paket soal** — melihat list/detail paket dan melakukan aksi **Verifikasi** (memberi label paket sudah diverifikasi). Tidak Create/Edit/Delete paket.
- **Super Admin**: Akses penuh ke Paket Soal dan semua modul lain, termasuk verifikasi.

**Catatan**: Satu pengguna dapat memiliki **lebih dari satu role** (misalnya Widyaprada sekaligus Verifikator Uji Kompetensi). **Widyaprada** tidak mengakses Paket Soal; mereka hanya mengerjakan ujian (CBT).

### 1.5 Field/Entitas Paket Soal (Referensi)

| No | Field / Kolom | Wajib | Tipe | Keterangan |
|----|----------------|-------|------|------------|
| 1 | **Kode Paket** | Ya | Teks (unik) | Identifikasi singkat, misalnya PAKET-A-2025. |
| 2 | **Nama Paket** | Ya | Teks | Nama paket (mis. "Kompetensi Dasar Widyaprada", "Paket Ujian Ahli Muda"). |
| 3 | **Deskripsi** | Opsional | Teks panjang | Keterangan isi paket atau tujuan pemakaian. |
| 4 | **Daftar Soal** | Ya | Referensi (urut) | Daftar soal yang masuk paket (referensi ke Bank Soal); setiap item punya urutan (order). Minimal 1 soal. |
| 5 | **Status** | Ya | Pilihan | Draft / Aktif. Paket Aktif bisa dipilih saat menyusun ujian; Draft untuk yang masih disusun. |
| 6 | **Status Verifikasi** | Ya | Pilihan | **Belum Diverifikasi** / **Sudah Diverifikasi**. Label untuk membedakan paket yang sudah diperiksa dan yang belum. Hanya Verifikator Uji Kompetensi (dan Super Admin) yang dapat mengubah status verifikasi. |
| 7 | **Keterangan** | Opsional | Teks | Catatan internal. |
| 8 | **Tanggal Dibuat / Diubah** | Sistem | Timestamp | Otomatis. |

**Catatan**: Saat paket dipakai di ujian, pengaturan pengacakan urutan soal (dan opsi PG) bisa diatur **per ujian** di Manajemen Uji Kompetensi; paket sendiri hanya menyimpan “urutan default” atau daftar soal (urutan acak opsional diterapkan saat ujian dibuat).

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Admin Uji Kompetensi / Super Admin | pengelola paket soal | melihat daftar paket soal dalam satu halaman (kode, nama, jumlah soal, status) | memantau paket yang ada dan memilih paket saat buat ujian |
| 2 | Admin Uji Kompetensi / Super Admin | pengelola paket soal | mencari paket berdasarkan kode atau nama | menemukan paket tertentu dengan cepat |
| 3 | Admin Uji Kompetensi / Super Admin | pengelola paket soal | memfilter daftar berdasarkan status (Draft / Aktif) | fokus pada paket yang siap dipakai |
| 4 | Admin Uji Kompetensi / Super Admin | pengelola paket soal | membuka detail satu paket (nama, deskripsi, daftar soal berurutan) | memeriksa isi paket sebelum dipakai di ujian atau sebelum edit |
| 5 | Admin Uji Kompetensi / Super Admin | pengelola paket soal | menambah paket baru (Create): nama, kode, deskripsi, lalu menambah soal dari Bank Soal ke dalam paket | membuat “playlist” soal yang bisa dipakai ulang |
| 6 | Admin Uji Kompetensi / Super Admin | pengelola paket soal | mengatur urutan soal dalam paket (drag-and-drop atau nomor urut) | urutan tampil sesuai kebutuhan |
| 7 | Admin Uji Kompetensi / Super Admin | pengelola paket soal | menambah atau menghapus soal dari paket (Edit) tanpa mengubah Bank Soal | paket tetap relevan dan bisa dipakai di banyak ujian |
| 8 | Admin Uji Kompetensi / Super Admin | pengelola paket soal | menghapus paket (Delete) dengan konfirmasi dan alasan penghapusan | membersihkan paket yang tidak dipakai |
| 9 | Admin Uji Kompetensi / Super Admin | pengelola paket soal | mendapat umpan balik jelas setelah Create/Edit/Delete | tahu aksi berhasil atau gagal |
| 10 | Admin Uji Kompetensi / Super Admin | pengelola paket soal | paket yang status Aktif bisa dipilih saat menyusun ujian di Manajemen Uji Kompetensi | satu paket dipakai di banyak ujian tanpa susun ulang |
| 11 | Verifikator Uji Kompetensi / Super Admin | verifikator | melihat daftar dan detail paket soal agar bisa memeriksa isi paket | memutuskan apakah paket layak diverifikasi |
| 12 | Verifikator Uji Kompetensi / Super Admin | verifikator | memfilter daftar paket berdasarkan status verifikasi (Belum / Sudah Diverifikasi) | fokus pada paket yang belum atau sudah diverifikasi |
| 13 | Verifikator Uji Kompetensi / Super Admin | verifikator | melakukan aksi Verifikasi pada paket (menandai Sudah Diverifikasi) dan opsional Batalkan Verifikasi | label paket terverifikasi jelas untuk referensi ujian |

---

### 2.1 Kebutuhan Pengguna - Yang Diinginkan
- **List**: Daftar paket dengan kolom relevan (kode, nama, jumlah soal, status, **status verifikasi**), search, filter (status, **status verifikasi**), sort, paginasi.
- **Detail**: Nama, deskripsi, daftar soal dalam urutan, **status verifikasi**; untuk Verifikator: tombol **Verifikasi** / **Batalkan Verifikasi**.
- **Create**: Form nama, kode, deskripsi; lalu tambah soal dari Bank Soal (pilih satu/satu atau filter lalu pilih banyak), atur urutan.
- **Edit**: Ubah nama/deskripsi; tambah/hapus soal dari paket; ubah urutan soal.
- **Delete**: Konfirmasi + **wajib alasan penghapusan**. Jika paket sedang dipakai di ujian (Draft), kebijakan jelas (peringatan / blok hapus).
- **Umpan balik**: Loading pada aksi; pesan sukses/error yang ramah.

### 2.2 Kebutuhan Pengguna - Yang Tidak Diinginkan
- List tanpa search/filter.
- Delete tanpa konfirmasi dan tanpa alasan.
- Soal di Bank Soal terhapus otomatis saat dihapus dari paket (hapus dari paket hanya memutus relasi, tidak menghapus soal).

---

## 3. Antarmuka Pengguna (UI)

*Paket Soal: List paket, Detail paket, form Create/Edit, Delete, aksi Verifikasi. Menu terlihat oleh pengguna dengan role Admin Uji Kompetensi, Verifikator Uji Kompetensi, atau Super Admin (satu user bisa punya lebih dari satu role).*

### 3.1 List Paket Soal
- **Judul**: "Paket Soal" atau "Daftar Paket Soal".
- **Tombol "Tambah Paket"**: Ke form Create (hanya Admin Uji Kompetensi & Super Admin).
- **Pencarian**: Kode atau nama paket.
- **Filter**: Status (Draft / Aktif), **Status Verifikasi** (Belum / Sudah Diverifikasi).
- **Tabel**: Kode, Nama, Jumlah Soal, Status, **Status Verifikasi**; aksi Detail, Edit, Hapus (Edit/Hapus untuk Admin & Super Admin); untuk Verifikator: Detail, **Verifikasi** / **Batalkan Verifikasi**.
- **Paginasi** dan umpan balik loading/sukses/error.

### 3.2 Detail Paket Soal
- **Judul**: Nama paket atau "Detail Paket Soal".
- **Info**: Kode, Nama, Deskripsi, Status, **Status Verifikasi**, Jumlah soal, tanggal dibuat/diubah.
- **Daftar soal**: Tabel atau list berurutan — nomor urut, kode soal, tipe, kategori (dari Bank Soal). Link ke detail soal jika perlu.
- **Tombol**: Edit, Hapus, Kembali ke List (Admin & Super Admin). Untuk **Verifikator** (dan Super Admin): **Verifikasi** / **Batalkan Verifikasi**, Kembali ke List.

### 3.3 Create / Edit Paket Soal
- **Data dasar**: Kode Paket, Nama Paket, Deskripsi, Status.
- **Daftar soal**: Area untuk menambah soal dari Bank Soal (pencarian/filter Bank Soal, checkbox atau "Tambah ke paket"); daftar soal yang sudah masuk paket dengan urutan (drag-and-drop atau field nomor urut). Opsi hapus soal dari paket (hapus dari daftar paket saja).
- **Tombol**: Simpan, Batal.
- **Validasi**: Kode unik; minimal 1 soal dalam paket.

### 3.4 Delete Paket Soal
- **Pemicu**: Tombol Hapus di Detail atau List.
- **Dialog**: Konfirmasi + **field wajib Alasan penghapusan**. Jika paket dipakai di ujian (Draft), tampilkan peringatan; kebijakan blok atau izinkan dengan peringatan.
- **Tombol**: Batal, Ya Hapus (setelah alasan diisi).

---

## 4. Pengalaman Pengguna (UX) - Alur

### 4.1 Alur Buat Paket Soal
1. Klik "Tambah Paket" → isi kode, nama, deskripsi.
2. Tambah soal dari Bank Soal (cari/filter, pilih, tambah ke paket); atur urutan jika perlu.
3. Simpan (Draft atau Aktif) → muncul di List; paket bisa dipilih di Manajemen Uji Kompetensi saat menyusun ujian.

### 4.2 Alur Edit Paket
1. Dari List atau Detail, klik Edit.
2. Ubah data dasar dan/atau daftar soal (tambah, hapus dari paket, ubah urutan) → Simpan.

### 4.3 Alur Delete
1. Klik Hapus → isi alasan di dialog → Ya Hapus → sukses/error.

---

## 5. Kebutuhan per Role

### 5.1 Role: Admin Uji Kompetensi
- **Paket Soal**: List, Detail, Create, Edit, Delete. Tidak mengubah status verifikasi.
- **Menu**: WPUjikom → Bank Soal, Paket Soal, Manajemen Uji Kompetensi (jika user punya role ini; menu gabungan jika punya banyak role).

### 5.2 Role: Verifikator Uji Kompetensi
- **Paket Soal**: List, Detail, aksi **Verifikasi** / **Batalkan Verifikasi**. Tidak Create/Edit/Delete.
- **Menu**: WPUjikom → Bank Soal, Paket Soal, Manajemen Uji Kompetensi (untuk verifikasi).

### 5.3 Role: Super Admin
- Paket Soal: List, Detail, Create, Edit, Delete, Verifikasi/Batalkan Verifikasi; plus semua modul lain.

### 5.4 Role: Widyaprada
- Tidak mengakses Paket Soal (kecuali jika user juga punya role Admin atau Verifikator). Sebagai peserta hanya CBT (mengerjakan ujian).

---

## 6. Cakupan Fitur

### 6.1 Termasuk
- List paket soal dengan search, filter (status, **status verifikasi**), sort, paginasi.
- Detail paket (info + daftar soal + **status verifikasi**).
- Create dan Edit paket (Admin & Super Admin); Verifikasi/Batalkan Verifikasi (Verifikator & Super Admin).
- Delete dengan konfirmasi dan alasan penghapusan.
- **Status Verifikasi**: label Belum Diverifikasi / Sudah Diverifikasi; aksi verifikasi hanya untuk Verifikator Uji Kompetensi dan Super Admin.
- Paket dengan status Aktif dapat dipilih di Manajemen Uji Kompetensi saat menyusun ujian.

### 6.2 Tidak Termasuk
- Bank Soal (soal individu) → PRD Bank Soal.
- Manajemen Uji Kompetensi (ujian, jadwal, peserta) → PRD Manajemen Uji Kompetensi; di sana paket dipakai sebagai salah satu sumber konten ujian.
- CBT → PRD CBT.
- Import/export paket (fase berikutnya jika diperlukan).

---

## 7. Persyaratan Produk (Nonteknis)

- **Unik**: Kode Paket unik di sistem.
- **Relasi**: Menghapus soal dari paket hanya memutus relasi; soal di Bank Soal tetap ada. Menghapus soal dari Bank Soal dapat memengaruhi paket yang memakai soal tersebut (kebijakan: tampilkan peringatan atau hapus dari paket; ditetapkan produk/SDD).
- **Delete paket**: Alasan penghapusan wajib. Jika paket sudah dipakai di ujian, kebijakan (blok hapus / peringatan / cascade) ditetapkan produk.
- **Akses**: Paket Soal dapat diakses oleh pengguna dengan role Admin Uji Kompetensi, Verifikator Uji Kompetensi, atau Super Admin. Satu user dapat memiliki lebih dari satu role.

---

## 8. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: Paket Soal sebagai entitas terpisah (playlist soal); dipakai di Manajemen Uji Kompetensi | - |
| 1.1 | 2025-02-11 | Status Verifikasi; role Verifikator Uji Kompetensi; satu user bisa lebih dari satu role | - |

---

**Catatan**: Di **Manajemen Uji Kompetensi**, satu ujian dapat memuat **soal individu** dan/atau **paket soal**. Role **Verifikator Uji Kompetensi** hanya melakukan verifikasi (label Sudah/Belum Diverifikasi). Satu pengguna dapat memiliki lebih dari satu role. Lihat PRD_Manajemen_Uji_Kompetensi.md, PRD_Auth_Manajemen_Pengguna.md.
