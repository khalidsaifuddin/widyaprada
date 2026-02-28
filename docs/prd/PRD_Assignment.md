# [PRD] Fitur Assignment (Penugasan Uji Kompetensi)
## Product Requirements Document | WPUjikom

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom (Uji Kompetensi)  
**Fitur**: Assignment – Penugasan Uji Kompetensi ke User, Batas Waktu, Hasil, dan Leaderboard/Privat  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Assignment (Penugasan Uji Kompetensi)
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Assignment memungkinkan **Peserta (Calon/WP)** untuk **mendaftar (apply)** mengikuti uji kompetensi dengan mengisi dan mengupload dokumen persyaratan. **Penugasan ke ujian ditentukan berdasarkan validasi persyaratan** oleh Tim Verval — bukan lagi admin yang assign langsung. Setelah lolos seleksi administrasi, peserta menunggu jadwal Ujikom yang ditetapkan admin, lalu dapat mengerjakan ujian (Mulai Ujikom) melalui CBT ketika jadwal tersedia dan diinformasikan di beranda. Peserta melihat **ujian yang ditugaskan kepadanya** beserta **batas waktu** pengerjaan, mengerjakan melalui CBT, dan **melihat hasil** ketika sudah keluar. Ada **pengaturan visibility hasil** per ujian: **Leaderboard** atau **Privat**. Batas waktu pengerjaan mengacu pada **Jadwal Selesai** ujian.

### 1.3 Role yang Mengakses
- **Peserta (Calon/WP)**: Mendaftar (apply) ke Ujikom dengan memilih jenis ujikom, mengisi dan upload dokumen persyaratan; melihat pengumuman hasil seleksi administrasi di beranda; menunggu jadwal Ujikom; mengerjakan ujian (Mulai Ujikom) via CBT; melihat **hasil sendiri**; jika ujian diset **Leaderboard**, melihat **leaderboard**. Jika **Privat**, hanya melihat hasil sendiri.
- **Admin Uji Kompetensi / Super Admin**: Mengatur peserta (daftar calon), verifikasi & validasi dokumen persyaratan, atur jadwal Ujikom, pengaturan **Tampilkan Leaderboard** (Ya/Tidak) per ujian.
- **Tim Verval (Verifikator)**: Memeriksa persyaratan dokumen; jika ditolak memberikan catatan (tidak ada koreksi — langsung info tidak lulus syarat administrasi).

**Catatan**: Satu pengguna dapat memiliki **lebih dari satu role**. Peserta mengakses fitur Assignment dari **beranda** dan **menu WPUjikom / CBT / Tugas Saya**.

### 1.4 Field/Entitas yang Relevan

**Konsep Assignment (dari sisi produk)**  
- **Apply-first**: Peserta harus **mendaftar (apply)** terlebih dahulu dengan dokumen persyaratan. **Penugasan ke ujian ditentukan oleh validasi persyaratan** — bukan admin assign langsung.  
- **Penugasan** = User (Peserta/Calon WP) lolos seleksi administrasi dan termasuk dalam **daftar peserta** satu ujian. Satu user bisa di-assign ke **banyak ujian**; satu ujian punya **banyak peserta**.  
- **Batas waktu** = **Jadwal Selesai** ujian (tanggal-waktu terakhir peserta boleh mengirim jawaban). Timer pengerjaan per peserta mengikuti **Durasi** ujian sejak klik Mulai (lihat PRD CBT).  
- **Hasil** = Nilai/skor peserta untuk ujian tersebut; tampil setelah proses koreksi (otomatis atau manual).  
- **Pengaturan visibility hasil** = Per ujian, ada opsi **Tampilkan Leaderboard** (Ya/Tidak).  
  - **Ya (Leaderboard)**: Setelah hasil keluar, peserta dapat membuka **Leaderboard** ujian tersebut (ranking berdasarkan nilai; identitas peserta sesuai kebijakan tampilan).  
  - **Tidak (Privat)**: Peserta **hanya** melihat hasil/nilai sendiri; tidak ada akses ke leaderboard atau nilai peserta lain.

**Entitas Ujian (referensi – penambahan field)**  
- Field baru di Manajemen Uji Kompetensi (atau konfigurasi ujian):  
  - **Tampilkan Leaderboard** (Ya / Tidak). Wajib; default boleh **Tidak** (privat).  
- Peserta dan Jadwal Selesai sudah ada di PRD Manajemen Uji Kompetensi; Assignment memanfaatkan data tersebut.

### 1.5 Lampiran Dokumen Persyaratan (Apply Uji Kompetensi Non-WP ke WP / Widyaprada Ahli Madya)

Berikut **13 dokumen persyaratan** yang wajib dilengkapi calon peserta apply uji kompetensi jabatan fungsional Widyaprada Ahli Madya (dari non-WP ke WP). Sumber kebenaran: `ref.dokumen_persyaratan_ujikom`.

| No | Kode | Dokumen | Tipe Input | Batasan |
|----|------|---------|------------|---------|
| 1 | surat_usul_pimpinan | Surat usul dari pimpinan satuan kerja kepada Sekretaris Direktorat Jenderal PAUD, Pendidikan Dasar, dan Pendidikan Menengah (pelamar di lingkungan Ditjen PAUD Dikdas Dikmen), atau surat usul dari Sekretaris Unit Utama (pelamar dari unit utama lain) | File | - |
| 2 | sk_kenaikan_pangkat_terakhir | Surat keputusan kenaikan pangkat terakhir | File | - |
| 3 | sk_jabatan_terakhir | Surat keputusan jabatan terakhir | File | - |
| 4 | surat_pernyataan_integritas_moralitas | Surat pernyataan pimpinan yang menyatakan calon peserta memiliki integritas dan moralitas yang baik | File | - |
| 5 | surat_keterangan_sehat | Surat keterangan sehat dari pusat pelayanan kesehatan yang berwenang | File | - |
| 6 | fotokopi_ijazah | Fotokopi ijazah pendidikan terakhir | File | - |
| 7 | surat_keterangan_pengalaman_2tahun | Surat keterangan pimpinan yang menyatakan calon peserta memiliki pengalaman dalam pelaksanaan tugas di bidang penjaminan mutu pendidikan paling singkat 2 (dua) tahun | File | - |
| 8 | surat_pernyataan_lowongan | Surat pernyataan pimpinan unit kerja mengenai ketersediaan lowongan kebutuhan pada jenjang jabatan yang akan diduduki. (Calon dari unit kerja lain: lampirkan surat pernyataan dari pimpinan unit kerja tujuan) | File | - |
| 9 | surat_pernyataan_tidak_menuntut | Surat pernyataan tidak menuntut untuk diangkat sebagai pejabat fungsional Widyaprada | File | - |
| 10 | penilaian_skp_2tahun | Penilaian sasaran kinerja pegawai 2 (dua) tahun terakhir (24 buah aktivitas di bulan berbeda) | File | 24 aktivitas di bulan berbeda |
| 11 | portofolio | Portofolio (format terlampir): rincian pelaksanaan tugas dalam 2 (dua) tahun di bidang penjaminan mutu pendidikan, memuat minimal 1 (satu) kegiatan perbulan berisi topik, tahapan pelaksanaan tugas, dan output | Text (per baris) | Min 1 kegiatan/bulan; 2 tahun |
| 12 | essay_inovasi_praktik_baik | Tulisan/essay terkait inovasi dan aksi nyata/praktik baik di bidang penjaminan mutu pendidikan yang sudah dilakukan | Text (essay) | Maksimal 1500 kata |
| 13 | surat_pernyataan_orisinalitas_essay | Surat pernyataan calon peserta perihal orisinalitas tulisan/essay, yang diketahui oleh pimpinan satuan kerja | File | - |

**Catatan**: Dokumen No 1, 12, 13 berlaku khusus untuk jenis ujikom **perpindahan jabatan** (Non-WP ke WP). Referensi lengkap di database: `ref.dokumen_persyaratan_ujikom` (migration 007).

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Peserta | calon peserta | mendaftar (apply) ke Ujikom dengan memilih **jenis ujikom** (perpindahan jabatan fungsional aktif; kenaikan tingkat WP disabled) | mengajukan diri mengikuti uji kompetensi |
| 2 | Peserta | calon peserta | mengisi dan mengupload **kelengkapan berkas** sesuai lampiran (portofolio diketik per baris) | memenuhi persyaratan administrasi |
| 3 | Peserta | calon peserta | melihat **pengumuman hasil seleksi administrasi** di beranda (jika tidak lolos ada informasi) | tahu status pendaftaran saya |
| 4 | Peserta | calon peserta | melihat info **menunggu jadwal Ujikom** yang akan ditetapkan admin | tahu langkah selanjutnya |
| 5 | Peserta | peserta uji kompetensi | **klik Mulai Ujikom** ketika jadwal tersedia dan diinformasikan di beranda | mengerjakan ujian sesuai jadwal |
| 6 | Peserta | peserta uji kompetensi | melihat di dashboard daftar **penugasan ujian** (ujian yang di-assign ke saya) beserta batas waktu (Jadwal Selesai) dan status (belum/sudah dikerjakan) | tahu ujian mana yang harus saya kerjakan dan kapan batas waktunya |
| 7 | Peserta | peserta uji kompetensi | mengerjakan ujian yang ditugaskan melalui tombol/link ke CBT (Mulai Ujian) sesuai jadwal dan batas waktu | menyelesaikan penugasan sebelum deadline |
| 8 | Peserta | peserta uji kompetensi | melihat **hasil/nilai** saya untuk ujian yang sudah saya kerjakan, ketika hasil sudah keluar | memantau performa dan kelulusan |
| 9 | Peserta | peserta uji kompetensi | jika ujian diset **Leaderboard**, melihat **leaderboard** (ranking peserta berdasarkan nilai) untuk ujian tersebut | membandingkan dengan peserta lain (motivasi/transparansi) |
| 10 | Peserta | peserta uji kompetensi | jika ujian diset **Privat**, hanya melihat hasil saya sendiri dan **tidak** melihat hasil atau ranking peserta lain | privasi hasil ujian terjaga |
| 11 | Peserta | peserta uji kompetensi | membedakan dengan jelas ujian yang belum dikerjakan vs sudah dikerjakan, dan yang sudah ada hasil vs belum | mengatur prioritas dan tidak bingung |
| 12 | Admin / Super Admin | pengelola ujian | melihat daftar **calon peserta** dan melakukan **verifikasi serta validasi dokumen persyaratan** | menyeleksi peserta yang lolos syarat administrasi |
| 13 | Admin / Super Admin | pengelola ujian | menolak pendaftaran dengan **catatan** (tanpa koreksi — langsung info tidak lulus syarat administrasi) | peserta tahu alasan tidak lolos |
| 14 | Admin / Super Admin | pengelola ujian | mengatur **jadwal Ujikom** untuk peserta yang lolos | peserta dapat mengerjakan ujian sesuai jadwal |
| 15 | Admin Uji Kompetensi / Super Admin | pengelola ujian | mengatur **Tampilkan Leaderboard** (Ya/Tidak) per ujian saat Create/Edit ujian | mengendalikan visibility hasil ujian |
| 16 | Peserta | peserta uji kompetensi | tidak melihat ujian yang bukan penugasan saya (hanya ujian yang saya termasuk pesertanya) | daftar penugasan relevan dan aman |

---

### 2.1 Kebutuhan Pengguna - Yang Diinginkan
- **Apply / Pendaftaran Ujikom**: Halaman untuk memilih **jenis ujikom** (Perpindahan jabatan fungsional — aktif; Kenaikan tingkat WP — disabled), mengisi dan upload **kelengkapan berkas** sesuai lampiran (portofolio diketik per baris). Setelah submit, status "Menunggu verifikasi".
- **Dashboard Tugas Saya / Penugasan**: Satu halaman (atau submenu di CBT/WPUjikom) yang menampilkan daftar ujian tempat user termasuk peserta. Setiap item: nama ujian, **batas waktu** (Jadwal Selesai), status (Belum dikerjakan / Sudah dikerjakan), tombol "Mulai Ujian" (jika masih dalam jadwal dan belum submit), dan link "Lihat Hasil" (jika hasil sudah ada).
- **Batas waktu**: Ditampilkan jelas (tanggal & waktu); setelah batas waktu lewat, ujian tidak bisa lagi dikerjakan (konsisten dengan CBT).
- **Hasil**: Setelah hasil keluar, user melihat nilai sendiri (di Tugas Saya atau Riwayat CBT). Tampilan konsisten dengan PRD CBT (Riwayat/Hasil Saya).
- **Leaderboard**: Jika ujian diset **Tampilkan Leaderboard = Ya**, tersedia tombol/link **"Lihat Leaderboard"** (atau serupa). Halaman Leaderboard menampilkan ranking peserta (misalnya peringkat, nama/NIP sesuai kebijakan, nilai). Hanya untuk ujian yang hasilnya sudah tersedia.
- **Privat**: Jika ujian diset **Tampilkan Leaderboard = Tidak**, tidak ada tombol leaderboard; user hanya melihat hasil sendiri. Satu user tidak bisa mengakses hasil atau ranking peserta lain.
- **Admin**: Di Manajemen Uji Kompetensi (Create/Edit ujian), field **Tampilkan Leaderboard** (Ya/Tidak); pengaturan peserta tetap seperti PRD Manajemen Uji Kompetensi.

### 2.2 Kebutuhan Pengguna - Yang Tidak Diinginkan
- User melihat ujian yang tidak di-assign kepadanya (bukan peserta) di daftar Tugas Saya.
- User mengerjakan ujian setelah batas waktu (Jadwal Selesai); sistem mencegah akses submit.
- Dalam mode **Privat**, user bisa melihat nilai atau identitas peserta lain (harus diblokir).
- Leaderboard tampil untuk ujian yang diset **Privat** (harus disembunyikan).

---

## 3. Antarmuka Pengguna (UI)

*Assignment mencakup: (1) Dashboard user Widyaprada – Tugas Saya / Penugasan; (2) Tampilan hasil dan Leaderboard; (3) Penambahan pengaturan di Manajemen Uji Kompetensi (Tampilkan Leaderboard).*

### 3.1 Dashboard User – Tugas Saya / Penugasan
- **Judul**: "Tugas Saya" atau "Penugasan Uji Kompetensi" atau "Ujian Saya" (konsisten dengan menu WPUjikom/CBT).
- **Daftar**: Kartu atau tabel per ujian yang di-assign ke user. Kolom/info: **Nama Ujian**, **Batas Waktu** (Jadwal Selesai), **Status** (Belum dikerjakan / Sudah dikerjakan), **Hasil** (nilai jika sudah ada; atau "-" / "Belum keluar").
- **Aksi**: "Mulai Ujian" (jika masih dalam periode dan belum submit) → mengarahkan ke alur CBT; "Lihat Hasil" (jika sudah submit dan hasil sudah keluar); **"Lihat Leaderboard"** (hanya tampil jika ujian diset Leaderboard dan hasil sudah keluar).
- **Pesan kosong**: "Anda belum memiliki penugasan ujian." jika tidak ada ujian yang memenuhi (user tidak termasuk peserta atau tidak ada ujian yang sesuai jadwal/status).
- **Filter/sort** (opsional): Filter "Belum dikerjakan" / "Sudah dikerjakan"; sort berdasarkan batas waktu (terdekat dulu).

### 3.2 Halaman Hasil (per ujian)
- **Konteks**: Dapat diakses dari Tugas Saya atau dari Riwayat CBT.
- **Konten**: Nama ujian, nilai/skor user, status Lulus/Tidak lulus (jika ada standar). Jika ujian **Leaderboard**, tombol **"Lihat Leaderboard"**. Jika **Privat**, tidak ada tombol leaderboard.

### 3.3 Halaman Leaderboard (Detail Eksplisit)

**Akses**
- Hanya untuk ujian dengan **Tampilkan Leaderboard = Ya** dan **hasil ujian sudah tersedia** (nilai peserta sudah dipublikasikan, sesuai alur koreksi di Manajemen Uji Kompetensi/CBT).
- Hanya **peserta ujian tersebut** yang boleh membuka leaderboard ujian itu (user yang tidak termasuk peserta tidak boleh akses).
- Untuk ujian dengan **Tampilkan Leaderboard = Tidak**, halaman Leaderboard **tidak di-link** dan akses langsung (URL) harus ditolak (403 atau redirect).

**Konten halaman**
- **Judul**: Nama ujian (dan opsional periode/jadwal).
- **Tabel/daftar ranking** dengan kolom:
  - **Peringkat** (1, 2, 3, …). Jika nilai sama, **peringkat boleh sama** (misalnya dua orang nilai 85 → keduanya peringkat 2; berikutnya peringkat 4) atau dipecah sesuai kebijakan produk (misalnya secondary sort berdasarkan waktu submit). Produk menetapkan satu aturan konsisten.
  - **Identitas peserta**: Sesuai **kebijakan privasi/tampilan** yang ditetapkan produk — misalnya Nama Lengkap, atau Nama + NIP, atau Nama + Satker (tanpa NIP). Opsi anonim (hanya "Peserta 1", "Peserta 2") jika produk menghendaki. Harus konsisten per ujian.
  - **Nilai**: Skor/nilai akhir peserta (format angka atau angka + persen sesuai definisi ujian).
- **Urutan**: Berdasarkan nilai **tinggi ke rendah** (nilai tertinggi = peringkat 1). Secondary sort (jika nilai sama): bebas produk (nama, waktu submit, atau peringkat sama seperti di atas).
- **Highlight**: Baris **user yang sedang login** (peserta saat ini) ditandai (warna/ikon) agar user mudah melihat posisi sendiri.
- **Peserta yang tampil**: Hanya peserta yang **sudah submit** dan **sudah memiliki nilai**. Peserta yang belum mengerjakan atau belum punya nilai **tidak ditampilkan** di leaderboard (atau ditampilkan di bagian terpisah dengan label "Belum ada nilai" — produk memilih satu aturan).
- **Paginasi**: Jika peserta banyak (misalnya > 50), tampilkan dengan **paginasi** atau batas "Top 100" dengan link "Lihat semua" ke halaman berikutnya; agar performa dan UX terjaga.

**Perilaku**
- Leaderboard **read-only** (peserta hanya melihat, tidak bisa mengubah data).
- Admin Uji Kompetensi / Super Admin dapat melihat data ranking yang sama dari sisi Manajemen Uji Kompetensi (rekap nilai peserta sudah ada di PRD Manajemen Uji Kompetensi); tidak wajib halaman terpisah "Leaderboard untuk Admin" selama rekap bisa di-sort berdasarkan nilai.

### 3.4 Manajemen Uji Kompetensi – Pengaturan Leaderboard
- **Lokasi**: Di form Create/Edit Ujian (PRD Manajemen Uji Kompetensi).
- **Field**: **Tampilkan Leaderboard** – pilihan **Ya** / **Tidak** (default: **Tidak**). Keterangan singkat: "Jika Ya, peserta dapat melihat ranking hasil ujian. Jika Tidak, setiap peserta hanya melihat hasil sendiri (mode privat)."
- **Detail Ujian**: Tampilkan juga label "Leaderboard: Ya/Tidak" di blok info agar Admin/Verifikator tahu pengaturan visibility.

---

## 4. Pengalaman Pengguna (UX) - Alur

### 4.0 Alur Pendaftaran (Apply) dan Seleksi Administrasi
1. **Peserta** registrasi akun (password via email) → Login.
2. Peserta pilih **jenis ujikom** (Perpindahan jabatan fungsional — aktif; Kenaikan tingkat WP — disabled).
3. Peserta mengisi dan mengupload **kelengkapan berkas** sesuai lampiran (portofolio diketik per baris).
4. **Tim Verval** memeriksa persyaratan dokumen. Jika **ditolak**: ada catatan; tidak ada koreksi — langsung info tidak lulus syarat administrasi.
5. **Pengumuman hasil seleksi administrasi** di beranda: jika tidak lolos → informasi ditampilkan; jika lolos → lanjut.
6. Info **menunggu jadwal Ujikom** yang akan ditetapkan admin (ditampilkan di beranda).
7. Peserta **klik Mulai Ujikom** ketika jadwal tersedia dan diinformasikan di beranda.

### 4.1 Alur Melihat dan Mengerjakan Penugasan
1. Peserta login → Beranda (pengumuman, jadwal) / WPUjikom → **Tugas Saya** (atau "Penugasan" / "Ujian Saya").
2. Melihat daftar ujian yang di-assign (nama, batas waktu, status).
3. Klik "Mulai Ujian" pada ujian yang belum dikerjakan → dialihkan ke alur CBT (petunjuk → pengerjaan → submit).
4. Setelah submit, status berubah "Sudah dikerjakan"; ketika hasil keluar, nilai tampil dan (jika Leaderboard) tombol "Lihat Leaderboard" tersedia.

### 4.2 Alur Melihat Hasil dan Leaderboard
1. Dari Tugas Saya (atau Riwayat CBT) → pilih ujian yang sudah dikerjakan dan hasil sudah keluar.
2. Melihat nilai sendiri di halaman Hasil.
3. Jika ujian **Leaderboard**: klik "Lihat Leaderboard" → halaman ranking; user bisa melihat posisi sendiri dan peserta lain (sesuai kebijakan tampilan).
4. Jika ujian **Privat**: tidak ada aksi leaderboard; hanya nilai sendiri yang terlihat.

### 4.3 Alur Admin – Verifikasi, Jadwal, dan Penugasan
1. Admin melihat **daftar calon peserta** yang telah apply.
2. **Tim Verval** memeriksa dokumen persyaratan; setuju/tolak. Jika tolak: beri catatan; peserta mendapat info tidak lulus syarat administrasi.
3. Peserta yang lolos masuk daftar peserta ujian; Admin mengatur **jadwal Ujikom**.
4. Admin membuat atau mengedit ujian di Manajemen Uji Kompetensi; mengisi **Tampilkan Leaderboard** (Ya/Tidak).
5. **Penugasan** = peserta yang lolos validasi; Admin tidak assign langsung — assignment ditentukan oleh hasil validasi persyaratan.
6. Setelah ujian diterbitkan dan jadwal tersedia, peserta melihat ujian di Tugas Saya/Beranda dan dapat Mulai Ujikom.

---

## 5. Kebutuhan per Role

### 5.1 Role: Peserta (Calon/WP)
- **Apply**: Pilih jenis ujikom, isi dan upload dokumen persyaratan.
- **Beranda**: Lihat pengumuman hasil seleksi administrasi, info jadwal Ujikom, Mulai Ujikom jika jadwal tersedia.
- **Tugas Saya**: List penugasan (ujian yang saya ikuti), batas waktu, status, Mulai Ujian, Lihat Hasil.
- **Hasil**: Lihat nilai sendiri untuk setiap ujian yang sudah ada hasilnya.
- **Leaderboard**: Hanya jika ujian diset Leaderboard; akses halaman Leaderboard untuk ujian tersebut. Tidak bisa akses leaderboard ujian yang diset Privat.
- **Tidak mengakses**: Manajemen Uji Kompetensi (kecuali punya role Admin/Verifikator); pengaturan Tampilkan Leaderboard.

### 5.2 Role: Admin Uji Kompetensi
- **Daftar calon peserta**: Melihat list calon yang apply.
- **Verifikasi & validasi dokumen**: Memeriksa persyaratan; setuju/tolak dengan catatan jika tolak.
- **Atur jadwal Ujikom**: Menetapkan jadwal untuk peserta yang lolos.
- **Manajemen Uji Kompetensi**: Create/Edit ujian, **Tampilkan Leaderboard** (Ya/Tidak); peserta ditentukan oleh validasi, bukan assign manual. Melihat rekap hasil.

### 5.3 Role: Tim Verval (Verifikator Uji Kompetensi)
- Memeriksa persyaratan dokumen yang diupload calon peserta. Jika ditolak: beri catatan; tidak ada koreksi — langsung info tidak lulus syarat administrasi.
- Bisa melihat info ujian termasuk setting Leaderboard (read-only); verifikasi ujian sesuai PRD Manajemen Uji Kompetensi.

### 5.4 Role: Super Admin
- Akses penuh: mengatur peserta, Tampilkan Leaderboard, dan semua fitur Manajemen Uji Kompetensi; plus akses modul lain.

---

## 6. Cakupan Fitur

### 6.1 Termasuk
- **Dashboard Tugas Saya**: Daftar ujian yang di-assign ke user (peserta), dengan batas waktu (Jadwal Selesai), status (belum/sudah dikerjakan), link Mulai Ujian (CBT), dan Lihat Hasil.
- **Batas waktu**: Tampil jelas; pengerjaan hanya dalam periode ujian (sesuai CBT); setelah Jadwal Selesai tidak bisa submit.
- **Tampilan hasil**: User melihat hasil/nilai sendiri ketika sudah keluar (di Tugas Saya atau Riwayat CBT).
- **Pengaturan Tampilkan Leaderboard** (Ya/Tidak) per ujian di Manajemen Uji Kompetensi.
- **Mode Leaderboard**: Jika Ya, halaman Leaderboard untuk ujian tersebut (ranking berdasarkan nilai); akses hanya untuk ujian dengan setting Ya dan hasil sudah ada. Detail lengkap (akses, konten, peringkat, identitas, paginasi) di **§3.3 Halaman Leaderboard**.
- **Mode Privat**: Jika Tidak, tidak ada leaderboard; peserta hanya melihat hasil sendiri; tidak ada akses ke hasil peserta lain.
- Integrasi dengan **Manajemen Uji Kompetensi** (peserta, jadwal) dan **CBT** (pengerjaan, submit, nilai).

### 6.2 Tidak Termasuk
- Bank Soal, Paket Soal → PRD masing-masing.
- Alur pengerjaan soal (timer, submit) → PRD CBT.
- CRUD ujian dan verifikasi ujian → PRD Manajemen Uji Kompetensi.
- Batas waktu **per peserta** yang berbeda dari Jadwal Selesai ujian (misalnya deadline per orang); fase berikutnya jika dibutuhkan.
- Export sertifikat/leaderboard (fase berikutnya atau PRD terpisah).

---

## 7. Persyaratan Produk (Nonteknis)

- **Apply-first**: Penugasan ditentukan oleh validasi persyaratan; Admin tidak assign langsung. Peserta apply dulu, Tim Verval validasi, lalu peserta yang lolos masuk daftar ujian.
- **Penugasan**: Sumber kebenaran daftar peserta = hasil validasi + Manajemen Uji Kompetensi. Satu user bisa di-assign ke banyak ujian; batas waktu = Jadwal Selesai ujian.
- **Visibility**: Leaderboard hanya untuk ujian dengan **Tampilkan Leaderboard = Ya**. Mode Privat (Tidak) harus ditegakkan: tidak menampilkan dan tidak mengizinkan akses ke nilai/identitas peserta lain.
- **Leaderboard**: Hanya peserta ujian yang boleh akses leaderboard ujian tersebut. Isi leaderboard: peserta yang sudah submit dan sudah punya nilai; urutan nilai tinggi ke rendah; identitas sesuai kebijakan privasi produk; akses direct URL untuk ujian privat harus ditolak (403/redirect). Aturan peringkat nilai sama (peringkat sama vs dipecah) ditetapkan produk dan konsisten.
- **Konsistensi**: Tugas Saya hanya menampilkan ujian yang status Diterbitkan dan user termasuk peserta; aturan jadwal dan satu kali pengerjaan per ujian mengikuti PRD CBT.
- **Integrasi**: Assignment memanfaatkan data ujian (jadwal, peserta, nilai) dan CBT (pengerjaan, hasil); detail teknis di SDD.

---

## 8. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: Assignment (penugasan uji kompetensi), batas waktu, hasil, Leaderboard vs Privat | - |
| 1.1 | 2025-02-11 | Leaderboard dijelaskan secara detail dan eksplisit: akses, konten (peringkat, identitas, nilai), aturan nilai sama, paginasi, peserta yang tampil; persyaratan produk §7 | - |
| 1.2 | 2025-02-28 | Apply-first flow: peserta apply dulu dengan dokumen persyaratan; penugasan ditentukan validasi Tim Verval; pengumuman di beranda; jenis ujikom (perpindahan aktif, kenaikan disabled) | - |
| 1.3 | 2025-02-28 | Lampiran dokumen persyaratan lengkap: 13 dokumen (surat usul, SK, pernyataan, SKP, portofolio, essay, dll) untuk apply Non-WP ke WP / Widyaprada Ahli Madya; referensi `ref.dokumen_persyaratan_ujikom` | - |

---

**Catatan**: Integrasi teknis dengan Manajemen Uji Kompetensi (peserta, field Tampilkan Leaderboard) dan CBT (daftar ujian, nilai) didokumentasikan di SDD. Lihat PRD_Manajemen_Uji_Kompetensi.md, PRD_CBT.md, PRD_Bank_Soal.md, PRD_Paket_Soal.md.
