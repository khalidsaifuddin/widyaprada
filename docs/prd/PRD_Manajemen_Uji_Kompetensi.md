# [PRD] Fitur Manajemen Uji Kompetensi
## Product Requirements Document | WPUjikom

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom (Uji Kompetensi)  
**Fitur**: Manajemen Uji Kompetensi (List, Detail, Create, Edit, Delete, Jadwal, Paket Soal, Peserta)  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Manajemen Uji Kompetensi
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Manajemen Uji Kompetensi memungkinkan **Admin Uji Kompetensi** dan **Super Admin** untuk merancang, menjadwalkan, dan mengelola ujian kompetensi bagi Peserta. Pengelola dapat membuat ujian baru (nama, deskripsi, periode, durasi), **mengisi konten ujian** dengan **soal individu** dan/atau **paket soal**, dan mengendalikan status ujian (Draft, Diterbitkan, Berlangsung, Selesai). **Peserta ditentukan berdasarkan validasi persyaratan** — bukan lagi admin memilih langsung; peserta yang lolos seleksi administrasi (validasi dokumen oleh Tim Verval) masuk daftar peserta ujian. Admin mengatur **jadwal Ujikom** untuk peserta yang lolos. Ujian yang sudah diterbitkan dapat diakses peserta melalui modul CBT sesuai jadwal. Hasil ujian dapat dilihat di Manajemen Uji Kompetensi (rekap nilai, daftar peserta) dan oleh peserta di CBT (nilai sendiri).

### 1.3 Role yang Mengakses
- **Admin Uji Kompetensi**: Akses penuh ke Manajemen Uji Kompetensi (CRUD ujian, terbitkan, rekap). Tidak mengubah status verifikasi ujian.
- **Verifikator Uji Kompetensi**: Akses untuk **verifikasi ujian** — melihat list/detail ujian dan melakukan aksi **Verifikasi** (memberi label ujian sudah diverifikasi). Tidak Create/Edit/Delete/Terbitkan ujian.
- **Super Admin**: Akses penuh ke semua modul termasuk Manajemen Uji Kompetensi dan verifikasi.

**Catatan**: Satu pengguna dapat memiliki **lebih dari satu role** (misalnya Peserta sekaligus Admin Uji Kompetensi). **Peserta** tidak mengakses Manajemen Uji Kompetensi; mereka hanya mendaftar (apply), mengerjakan ujian (CBT), dan melihat hasil sendiri.

### 1.4 Field/Entitas Uji Kompetensi (Referensi)

**Entitas: Ujian (Uji Kompetensi)**

| No | Field / Kolom | Wajib | Tipe | Keterangan |
|----|----------------|-------|------|------------|
| 1 | **Kode Ujian** | Ya | Teks (unik) | Identifikasi singkat, misalnya UK-2025-01. |
| 2 | **Nama Ujian** | Ya | Teks | Judul ujian yang tampil ke peserta dan di list. |
| 3 | **Deskripsi** | Opsional | Teks panjang | Informasi tujuan, materi, atau petunjuk umum. |
| 4 | **Periode / Jadwal Mulai** | Ya | Tanggal & Waktu | Kapan ujian dibuka untuk dikerjakan. |
| 5 | **Jadwal Selesai** | Ya | Tanggal & Waktu | Batas akhir pengerjaan; setelah ini ujian tidak bisa diakses peserta. |
| 6 | **Durasi Pengerjaan (menit)** | Ya | Angka | Lama waktu pengerjaan per peserta sejak mulai (timer di CBT). |
| 7 | **Konten Ujian** | Ya | Komposisi | Gabungan **soal individu** (pilih dari Bank Soal) dan/atau **paket soal** (pilih satu atau lebih dari daftar Paket Soal). Satu ujian bisa berisi: hanya soal individu, hanya paket soal, atau keduanya. Pengacakan urutan soal/opsi (per ujian) opsional. |
| 8 | **Peserta** | Ya | Referensi | Daftar peserta yang diizinkan mengerjakan ujian ini. Peserta berasal dari **hasil validasi persyaratan** (apply + verifikasi dokumen) — bukan assign manual. |
| 9 | **Status** | Ya | Pilihan | Draft / Diterbitkan / Berlangsung / Selesai. Draft: masih bisa edit paket dan peserta. Diterbitkan: sudah bisa diakses peserta di jadwal. Berlangsung/Selesai: untuk tampilan dan rekap. |
| 10 | **Status Verifikasi** | Ya | Pilihan | **Belum Diverifikasi** / **Sudah Diverifikasi**. Label untuk membedakan ujian yang sudah diperiksa dan yang belum. Hanya Verifikator Uji Kompetensi (dan Super Admin) yang dapat mengubah status verifikasi. |
| 11 | **Keterangan** | Opsional | Teks | Catatan internal. |
| 12 | **Tanggal Dibuat / Diubah** | Sistem | Timestamp | Otomatis. |

**Konten Ujian (komposisi)**
- **Soal individu**: Soal yang dipilih satu per satu dari Bank Soal untuk ujian ini.
- **Paket soal**: Satu atau lebih Paket Soal (entitas terpisah — kumpulan soal seperti playlist) yang dipilih untuk ujian ini; setiap paket menyumbangkan semua soal di dalamnya.
- Total soal ujian = soal individu + semua soal dari setiap paket yang dipilih. Urutan tampil: tetap (sesuai urutan di paket + urutan soal individu) atau acak (pengacakan soal dan/atau opsi jawaban untuk PG) — diatur per ujian.
- Jumlah soal dan total bobot mengikuti komposisi di atas.

**Peserta**
- Daftar peserta per ujian: referensi ke User (calon WP / peserta) yang **lolos seleksi administrasi** (validasi dokumen persyaratan oleh Tim Verval).
- **Jenis Ujikom** (untuk apply): Perpindahan jabatan fungsional (non-WP → WP) — **aktif**; Kenaikan tingkat WP — **disabled** (fase ini).
- Admin mengatur **jadwal Ujikom**; peserta melihat jadwal di beranda dan dapat Mulai Ujikom ketika tersedia.
- Status pengerjaan per peserta: Belum mengerjakan / Sedang / Sudah submit (dan nilai jika sudah dikoreksi).

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Admin Uji Kompetensi / Super Admin | pengelola ujian | melihat daftar ujian dalam satu halaman (kode, nama, jadwal, status) | memantau ujian yang sudah dan akan berjalan |
| 2 | Admin Uji Kompetensi / Super Admin | pengelola ujian | mencari ujian berdasarkan kode atau nama | menemukan ujian tertentu dengan cepat |
| 3 | Admin Uji Kompetensi / Super Admin | pengelola ujian | memfilter daftar berdasarkan status (Draft, Diterbitkan, Berlangsung, Selesai) atau periode | fokus pada kelompok ujian yang relevan |
| 4 | Admin Uji Kompetensi / Super Admin | pengelola ujian | membuka detail satu ujian (jadwal, paket soal, daftar peserta, rekap hasil) | memeriksa konfigurasi dan hasil sebelum/sesudah ujian |
| 5 | Admin Uji Kompetensi / Super Admin | pengelola ujian | menambah ujian baru (Create): nama, jadwal, durasi, lalu mengisi konten (soal individu dan/atau paket soal) dan peserta | menyelenggarakan uji kompetensi sesuai jadwal |
| 6 | Admin Uji Kompetensi / Super Admin | pengelola ujian | mengubah ujian (Edit) yang masih Draft: ubah jadwal, durasi, konten (soal individu/paket soal), peserta | menyesuaikan ujian sebelum diterbitkan |
| 7 | Admin Uji Kompetensi / Super Admin | pengelola ujian | menerbitkan ujian (ubah status ke Diterbitkan) agar peserta bisa mengerjakan di CBT sesuai jadwal | ujian aktif dan terkendali |
| 8 | Admin Uji Kompetensi / Super Admin | pengelola ujian | mengisi konten ujian: tambah soal individu (pilih dari Bank Soal) dan/atau tambah paket soal (pilih dari daftar Paket Soal); opsi pengacakan urutan per ujian | satu ujian bisa memuat gabungan soal individu dan kumpulan soal dari paket |
| 9 | Admin Uji Kompetensi / Super Admin | pengelola ujian | melihat peserta yang lolos validasi dokumen dan mengatur jadwal Ujikom | peserta yang berhak mengerjakan sesuai jadwal (peserta = hasil validasi, bukan assign manual) |
| 10 | Admin Uji Kompetensi / Super Admin | pengelola ujian | menghapus ujian (Delete) dengan konfirmasi dan alasan; hanya ujian Draft atau kebijakan jelas untuk ujian yang sudah berjalan | membersihkan ujian yang batal atau salah input |
| 11 | Admin Uji Kompetensi / Super Admin | pengelola ujian | melihat rekap hasil ujian: daftar peserta, status (sudah/belum submit), nilai | memantau kelulusan dan laporan |
| 12 | Admin Uji Kompetensi / Super Admin | pengelola ujian | paginasi dan loading/umpan balik yang jelas | daftar dan aksi tetap nyaman dipakai |
| 13 | Verifikator Uji Kompetensi / Super Admin | verifikator | melihat daftar dan detail ujian agar bisa memeriksa konfigurasi ujian | memutuskan apakah ujian layak diverifikasi |
| 14 | Verifikator Uji Kompetensi / Super Admin | verifikator | memfilter daftar ujian berdasarkan status verifikasi (Belum / Sudah Diverifikasi) | fokus pada ujian yang belum atau sudah diverifikasi |
| 15 | Verifikator Uji Kompetensi / Super Admin | verifikator | melakukan aksi Verifikasi pada ujian (menandai Sudah Diverifikasi) dan opsional Batalkan Verifikasi | label ujian terverifikasi jelas untuk pelaporan dan audit |

---

### 2.1 Kebutuhan Pengguna - Yang Diinginkan
- **List**: Daftar ujian dengan search, filter (status, periode, **status verifikasi**), sort, paginasi.
- **Detail**: Jadwal, deskripsi, konten ujian (ringkasan), daftar peserta, rekap nilai, **status verifikasi**; untuk Verifikator: tombol **Verifikasi** / **Batalkan Verifikasi**.
- **Create/Edit**: Form nama, jadwal, durasi; lalu konfigurasi **konten ujian** (tambah soal individu dari Bank Soal dan/atau pilih Paket Soal; opsi pengacakan urutan). **Peserta** = calon yang lolos validasi dokumen; Admin tidak memilih peserta manual — peserta masuk otomatis berdasarkan hasil verifikasi Tim Verval.
- **Terbitkan**: Aksi ubah status Draft → Diterbitkan dengan konfirmasi; setelah diterbitkan, batasan edit (misalnya tidak mengubah konten ujian yang sudah dipakai).
- **Delete**: Konfirmasi + alasan penghapusan; hanya untuk Draft atau sesuai kebijakan.
- **Rekap**: Di Detail, tampilkan tabel peserta: nama, NIP/satker, status pengerjaan, nilai (jika ada).

### 2.2 Kebutuhan Pengguna - Yang Tidak Diinginkan
- Mengubah paket soal atau peserta setelah ujian sudah ada yang mengerjakan (kecuali kebijakan eksplisit).
- Delete tanpa konfirmasi dan alasan.
- Peserta bisa mengakses ujian yang belum jadwal atau sudah lewat jadwal (dikendalikan di CBT).

---

## 3. Antarmuka Pengguna (UI)

*Manajemen Uji Kompetensi: List ujian, Detail ujian, form Create/Edit, Terbitkan, Delete, aksi Verifikasi. Menu terlihat oleh pengguna dengan role Admin Uji Kompetensi, Verifikator Uji Kompetensi, atau Super Admin (satu user bisa punya lebih dari satu role).*

### 3.1 List Ujian
- **Judul**: "Manajemen Uji Kompetensi" atau "Daftar Ujian".
- **Tombol "Buat Ujian"**: Ke form Create (Admin Uji Kompetensi & Super Admin).
- **Pencarian**: Kode atau nama ujian.
- **Filter**: Status (Draft, Diterbitkan, Berlangsung, Selesai), **Status Verifikasi** (Belum / Sudah Diverifikasi).
- **Tabel**: Kode, Nama, Jadwal Mulai–Selesai, Durasi, Status, **Status Verifikasi**; aksi Detail, Edit (jika Draft; Admin & Super Admin), Hapus; untuk Verifikator: Detail, **Verifikasi** / **Batalkan Verifikasi**.
- **Paginasi** dan umpan balik loading/sukses/error.

### 3.2 Detail Ujian
- **Judul**: Nama ujian atau "Detail Ujian".
- **Blok info**: Kode, Nama, Deskripsi, Jadwal Mulai, Jadwal Selesai, Durasi, Status, **Status Verifikasi**.
- **Blok Konten Ujian**: Ringkasan komposisi — soal individu + paket soal yang dipakai; total jumlah soal dan bobot. Link ke Bank Soal / Paket Soal jika perlu.
- **Blok Peserta**: Tabel peserta (nama, NIP, satker, status pengerjaan, nilai). Opsi export (out of scope atau fase berikutnya).
- **Blok Rekap**: Jumlah peserta, sudah submit, rata-rata nilai (jika ada).
- **Tombol**: Edit (jika Draft), Terbitkan (jika Draft), Hapus, Kembali ke List (Admin & Super Admin). Untuk **Verifikator** (dan Super Admin): **Verifikasi** / **Batalkan Verifikasi**, Kembali ke List.

### 3.3 Create / Edit Ujian (Form)
- **Step atau satu halaman**: (1) Data dasar: Kode, Nama, Deskripsi, Jadwal Mulai, Jadwal Selesai, Durasi. (2) **Konten ujian**: Tambah **soal individu** (pilih dari Bank Soal) dan/atau tambah **paket soal**; opsi pengacakan urutan soal/opsi per ujian. (3) Peserta: Daftar peserta = **calon yang lolos validasi dokumen**; Admin mengatur **jadwal Ujikom**; peserta muncul berdasarkan hasil verifikasi Tim Verval (tidak pilih manual).
- **Tombol**: Simpan (Draft), Simpan & Terbitkan (opsional), Batal.
- **Validasi**: Jadwal Selesai > Jadwal Mulai; Durasi > 0; minimal 1 soal (dari soal individu dan/atau dari paket). Peserta minimal 1 (berasal dari hasil validasi).

### 3.4 Terbitkan Ujian
- Dari Detail (status Draft): tombol "Terbitkan" → konfirmasi ("Ujian akan dapat diakses peserta sesuai jadwal. Lanjutkan?") → Ya → status berubah Diterbitkan; pesan sukses.

### 3.5 Delete Ujian
- Dialog konfirmasi + **field wajib Alasan penghapusan**. Hanya untuk Draft atau sesuai kebijakan (ujian yang sudah berjalan: blok atau soft delete).

---

## 4. Pengalaman Pengguna (UX) - Alur

### 4.1 Alur Buat Ujian
1. Klik "Buat Ujian" → isi data dasar (nama, jadwal, durasi).
2. Konfigurasi **konten ujian**: tambah soal individu (dari Bank Soal) dan/atau pilih Paket Soal; atur opsi pengacakan jika perlu.
3. **Peserta** = calon yang lolos validasi dokumen (Tim Verval); Admin mengatur **jadwal Ujikom** untuk peserta yang lolos.
4. Simpan sebagai Draft → muncul di List.
5. Buka Detail → "Terbitkan" jika siap; setelah itu peserta yang lolos validasi bisa mengerjakan di CBT sesuai jadwal (lihat pengumuman di beranda).

### 4.2 Alur Lihat Hasil
1. Buka Detail ujian yang status Berlangsung/Selesai.
2. Lihat blok Rekap dan tabel Peserta (status submit, nilai).

### 4.3 Alur Edit
1. Hanya ujian status Draft yang bisa diedit penuh (konten ujian & peserta). Setelah Terbitkan, hanya metadata tertentu yang boleh diubah (kebijakan produk).

---

## 5. Kebutuhan per Role

### 5.1 Role: Admin Uji Kompetensi
- **Manajemen Uji Kompetensi**: List, Detail, Create, Edit, Terbitkan, Delete (untuk Draft), lihat Rekap. Tidak mengubah status verifikasi ujian.
- **Menu**: WPUjikom → Bank Soal, Paket Soal, Manajemen Uji Kompetensi (menu gabungan jika user punya banyak role).

### 5.2 Role: Verifikator Uji Kompetensi
- **Manajemen Uji Kompetensi**: List, Detail, aksi **Verifikasi** / **Batalkan Verifikasi**, lihat Rekap. Tidak Create/Edit/Delete/Terbitkan.
- **Menu**: WPUjikom → Bank Soal, Paket Soal, Manajemen Uji Kompetensi (untuk verifikasi).

### 5.3 Role: Super Admin
- Sama dengan Admin Uji Kompetensi; plus verifikasi dan semua modul lain (RBAC, Manajemen Pengguna, CMS, dll).

### 5.4 Role: Peserta (Calon/WP)
- Tidak mengakses Manajemen Uji Kompetensi (kecuali jika user juga punya role Admin atau Verifikator). Sebagai peserta: apply dulu, lalu CBT (mengerjakan ujian, lihat nilai sendiri).

---

## 6. Cakupan Fitur

### 6.1 Termasuk
- List ujian dengan search, filter (status, **status verifikasi**), sort, paginasi.
- Detail ujian: info, konten ujian, peserta, rekap hasil, **status verifikasi**.
- Create dan Edit ujian (Admin & Super Admin); Verifikasi/Batalkan Verifikasi (Verifikator & Super Admin).
- Aksi Terbitkan (Draft → Diterbitkan) (Admin & Super Admin).
- Delete dengan konfirmasi dan alasan (untuk Draft atau sesuai kebijakan).
- **Status Verifikasi**: label Belum Diverifikasi / Sudah Diverifikasi; aksi verifikasi hanya untuk Verifikator Uji Kompetensi dan Super Admin.
- Tampilan rekap: daftar peserta, status pengerjaan, nilai.

### 6.2 Tidak Termasuk
- Bank Soal (soal individu) → PRD Bank Soal.
- Paket Soal (entitas playlist soal) → PRD Paket Soal.
- CBT (mengerjakan ujian, timer, submit, nilai peserta) → PRD CBT.
- Koreksi essay manual (bisa di Manajemen Uji Kompetensi atau modul terpisah; integrasi dengan nilai peserta di SDD).
- Export nilai/ sertifikat (fase berikutnya atau PRD terpisah).

---

## 7. Persyaratan Produk (Nonteknis)

- **Jadwal**: Jadwal Selesai harus setelah Jadwal Mulai; Durasi konsisten dengan pengalaman CBT (timer).
- **Status**: Transisi status (Draft → Diterbitkan → Berlangsung → Selesai) dapat otomatis berdasarkan waktu atau manual; aturan jelas (misalnya Berlangsung saat sekarang antara Mulai–Selesai).
- **Delete**: Alasan penghapusan wajib; kebijakan untuk ujian yang sudah diterbitkan/berjalan ditetapkan produk.
- **Integrasi**: Konten ujian memakai soal dari Bank Soal dan/atau Paket Soal; peserta = User (calon WP) yang lolos validasi dokumen persyaratan oleh Tim Verval. Detail integrasi di SDD.

---

## 8. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: Manajemen Uji Kompetensi untuk Admin Uji Kompetensi & Super Admin | - |
| 1.1 | 2025-02-11 | Konten ujian = soal individu + paket soal (bisa campuran); referensi PRD Paket Soal | - |
| 1.2 | 2025-02-11 | Status Verifikasi ujian; role Verifikator Uji Kompetensi; satu user bisa lebih dari satu role | - |
| 1.3 | 2025-02-28 | Apply-first: peserta ditentukan validasi dokumen Tim Verval; Admin tidak assign manual; jenis ujikom (perpindahan aktif, kenaikan disabled); jadwal Ujikom | - |

---

**Catatan**: Role **Admin Uji Kompetensi** dan **Verifikator Uji Kompetensi** diberi permission sesuai PRD. Satu pengguna dapat memiliki **lebih dari satu role** (misalnya Widyaprada + Admin Uji Kompetensi). Satu ujian dapat memuat **soal individu** dan/atau **paket soal**. Lihat PRD_RBAC.md, PRD_Bank_Soal.md, PRD_Paket_Soal.md, PRD_CBT.md, PRD_Auth_Manajemen_Pengguna.md.
