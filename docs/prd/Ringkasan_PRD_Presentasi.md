# Ringkasan PRD – Presentasi & Slide Show
## Modul WPUjikom (Uji Kompetensi) | Aplikasi Widyaprada

*Setiap PRD = 2 slide: (1) Definisi, (2) Daftar Fungsi/Fitur*

---

## 1. BANK SOAL

### Slide 1 – Definisi
- **Bank Soal** = tempat mengelola kumpulan soal uji kompetensi dalam satu tempat.
- **Pengguna**: Admin Uji Kompetensi, Verifikator Uji Kompetensi, Super Admin (Widyaprada tidak akses).
- **Fungsi utama**: Menambah, mengubah, melihat detail, menghapus soal; mencari & memfilter; verifikasi soal.
- **Tipe soal**: Pilihan Ganda (PG), Benar–Salah, Essay (uraian).
- **Konteks**: Soal dipakai saat menyusun paket/ujian di Manajemen Uji Kompetensi.

### Slide 2 – Daftar Fungsi/Fitur
- **List**: Daftar soal dengan search, filter (tipe, kategori, tingkat kesulitan, status, status verifikasi), sort, paginasi.
- **Detail**: Lihat soal lengkap (teks, opsi, kunci, bobot, status verifikasi).
- **Create & Edit**: Form tambah/ubah soal (validasi kode unik, field wajib).
- **Delete**: Konfirmasi + **wajib alasan penghapusan**.
- **Verifikasi**: Verifikator menandai soal Belum/Sudah Diverifikasi (aksi Verifikasi & Batalkan Verifikasi).
- **Role**: Admin & Super Admin = CRUD; Verifikator = list, detail, verifikasi saja.

---

## 2. PAKET SOAL

### Slide 1 – Definisi
- **Paket Soal** = kumpulan soal (dari Bank Soal) yang disimpan jadi satu entitas; satu paket berisi banyak soal dan bisa dipakai ulang di banyak ujian.
- **Pengguna**: Admin Uji Kompetensi, Verifikator Uji Kompetensi, Super Admin (Widyaprada tidak akses).
- **Perbedaan**: **Soal** = satu pertanyaan (bahan baku); **Paket Soal** = kumpulan soal; **Ujian** = event dengan jadwal & peserta, kontennya = soal individu + paket soal.
- **Konteks**: Di Manajemen Uji Kompetensi, satu ujian bisa memuat soal individu dan/atau satu atau lebih paket soal.

### Slide 2 – Daftar Fungsi/Fitur
- **List**: Daftar paket dengan search, filter (status, status verifikasi), sort, paginasi (kode, nama, jumlah soal).
- **Detail**: Nama, deskripsi, daftar soal berurutan, status verifikasi.
- **Create & Edit**: Data dasar (kode, nama, deskripsi); tambah/hapus soal dari Bank Soal; atur urutan soal (drag-and-drop/nomor).
- **Delete**: Konfirmasi + **wajib alasan penghapusan**.
- **Verifikasi**: Verifikator menandai paket Belum/Sudah Diverifikasi.
- **Integrasi**: Paket status Aktif bisa dipilih saat menyusun ujian di Manajemen Uji Kompetensi.

---

## 3. MANAJEMEN UJI KOMPETENSI

### Slide 1 – Definisi
- **Manajemen Uji Kompetensi** = merancang, menjadwalkan, dan mengelola ujian kompetensi bagi Widyaprada.
- **Pengguna**: Admin Uji Kompetensi, Verifikator Uji Kompetensi, Super Admin (Widyaprada tidak akses).
- **Fungsi utama**: Buat ujian (nama, jadwal, durasi); isi konten = **soal individu** dari Bank Soal dan/atau **paket soal**; atur peserta; terbitkan; rekap hasil; verifikasi ujian.
- **Status ujian**: Draft → Diterbitkan → Berlangsung → Selesai.
- **Konteks**: Ujian yang Diterbitkan bisa diakses peserta lewat CBT sesuai jadwal.

### Slide 2 – Daftar Fungsi/Fitur
- **List**: Daftar ujian dengan search, filter (status, status verifikasi), sort, paginasi.
- **Detail**: Jadwal, deskripsi, konten ujian (soal individu + paket), daftar peserta, rekap nilai, status verifikasi.
- **Create & Edit**: Data dasar (nama, jadwal, durasi); konten = soal individu + paket soal (bisa campuran); peserta (satker/individu); opsi pengacakan urutan.
- **Terbitkan**: Ubah status Draft → Diterbitkan (ujian bisa dikerjakan di CBT).
- **Delete**: Konfirmasi + alasan; hanya Draft atau sesuai kebijakan.
- **Rekap**: Daftar peserta, status pengerjaan, nilai.
- **Verifikasi**: Verifikator menandai ujian Belum/Sudah Diverifikasi.

---

## 4. ASSIGNMENT (PENUGASAN UJI KOMPETENSI)

### Slide 1 – Definisi
- **Assignment** = fitur agar **Widyaprada** melihat **ujian yang ditugaskan kepadanya**, batas waktu, dan hasil; serta mengerjakan via CBT.
- **Penugasan** = user (Widyaprada) masuk **daftar peserta** ujian; satu user bisa banyak ujian, satu ujian banyak peserta.
- **Batas waktu** = Jadwal Selesai ujian (timer pengerjaan per peserta sesuai CBT).
- **Visibility hasil**: **Leaderboard** (peserta lihat ranking/nilai peserta lain) atau **Privat** (hanya lihat hasil sendiri).
- **Pengaturan**: Admin atur peserta & **Tampilkan Leaderboard** (Ya/Tidak) per ujian di Manajemen Uji Kompetensi.

### Slide 2 – Daftar Fungsi/Fitur
- **Tugas Saya / Penugasan**: Daftar ujian yang di-assign (nama, batas waktu, status belum/sudah dikerjakan); tombol Mulai Ujian (CBT), Lihat Hasil.
- **Hasil**: Nilai sendiri setelah koreksi (di Tugas Saya atau Riwayat CBT).
- **Leaderboard** (jika ujian diset Ya): Halaman ranking (peringkat, identitas, nilai); hanya peserta ujian; highlight baris user login; paginasi jika banyak peserta.
- **Privat** (jika ujian diset Tidak): Hanya hasil sendiri; tidak ada akses ke leaderboard/nilai peserta lain.
- **Admin**: Field **Tampilkan Leaderboard** (Ya/Tidak) di Create/Edit Ujian; mengatur peserta = penugasan.

---

## 5. CBT (COMPUTER-BASED TEST)

### Slide 1 – Definisi
- **CBT** = fitur untuk **Widyaprada** (peserta) **mengerjakan ujian** yang sudah disusun di Manajemen Uji Kompetensi.
- **Akses**: Hanya role Widyaprada; Admin/Verifikator/Super Admin tidak mengerjakan ujian sebagai peserta.
- **Alur**: Lihat daftar ujian tersedia → Mulai ujian → Kerjakan soal (timer) → Submit → Lihat hasil (setelah koreksi).
- **Ujian tampil**: Hanya yang status **Diterbitkan**, dalam **periode jadwal**, dan user termasuk **daftar peserta**.
- **Satu peserta = satu kali pengerjaan** per ujian (no double attempt).

### Slide 2 – Daftar Fungsi/Fitur
- **Daftar ujian tersedia**: Nama, jadwal, durasi; tombol Mulai Ujian (atau Lanjutkan jika belum submit).
- **Petunjuk & Mulai**: Ringkasan (jumlah soal, durasi, peringatan); konfirmasi mulai → timer mulai.
- **Pengerjaan**: Soal PG, Benar–Salah, Essay; timer countdown; navigasi nomor soal (indikator terisi/belum); Simpan (opsional), Submit Ujian.
- **Submit**: Konfirmasi → jawaban tidak bisa diubah; **auto-submit** saat waktu habis.
- **Selesai**: Pesan sukses; info nilai (langsung atau "akan tampil setelah koreksi").
- **Riwayat / Hasil Saya**: Daftar ujian sudah dikerjakan + nilai (jika ada); status Lulus/Tidak lulus jika ada standar.
- **Peringatan**: Waktu hampir habis (mis. 5 menit); layout responsif (desktop & tablet).

---

## Ringkasan Cakupan per Modul

| PRD              | Definisi singkat                          | Siapa yang pakai                    |
|------------------|-------------------------------------------|-------------------------------------|
| Bank Soal        | Kelola kumpulan soal (PG, B-S, Essay)     | Admin, Verifikator, Super Admin     |
| Manajemen Uji Kompetensi | Rancang & jadwalkan ujian, konten + peserta | Admin, Verifikator, Super Admin     |
| Paket Soal       | Kumpulan soal untuk dipakai di banyak ujian | Admin, Verifikator, Super Admin     |
| Assignment       | Penugasan ujian ke user, batas waktu, hasil, Leaderboard/Privat | Widyaprada + Admin/Super Admin      |
| CBT              | Mengerjakan ujian (timer, submit, nilai)   | Widyaprada (peserta)                |

---

*Sumber: PRD_Bank_Soal.md, PRD_Manajemen_Uji_Kompetensi.md, PRD_Paket_Soal.md, PRD_Assignment.md, PRD_CBT.md*
