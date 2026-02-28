# [PRD] Fitur CBT (Computer-Based Test)
## Product Requirements Document | WPUjikom

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom (Uji Kompetensi)  
**Fitur**: CBT – Mengerjakan Ujian dan Melihat Hasil  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: CBT (Computer-Based Test) – Ujian Berbasis Komputer
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
CBT memungkinkan **Widyaprada** (peserta uji kompetensi) untuk mengerjakan ujian yang telah disusun oleh Admin Uji Kompetensi melalui Manajemen Uji Kompetensi. Peserta melihat daftar **ujian yang tersedia** (sesuai jadwal dan daftar peserta), memulai ujian, mengerjakan soal dalam waktu yang ditentukan (timer), mengirim jawaban (submit), dan melihat **hasil/nilai** setelah ujian selesai (atau setelah periode koreksi). Pengalaman dirancang agar jelas, adil (waktu dan soal sesuai pengaturan), dan minim kebingungan. Admin Uji Kompetensi dan Super Admin tidak mengerjakan ujian sebagai peserta; mereka memantau hasil melalui Manajemen Uji Kompetensi.

### 1.3 Role yang Mengakses
- **Widyaprada**: Role yang mengakses CBT sebagai **peserta** (mengerjakan ujian dan melihat nilai sendiri). Satu pengguna dapat memiliki **lebih dari satu role** (misalnya Widyaprada sekaligus Admin Uji Kompetensi); akses CBT sebagai peserta tetap mengacu pada adanya role Widyaprada.
- **Admin Uji Kompetensi / Verifikator Uji Kompetensi / Super Admin**: Tidak mengerjakan ujian di CBT sebagai peserta; mereka mengelola atau memverifikasi soal/ujian di modul lain serta melihat rekap hasil di Manajemen Uji Kompetensi.

### 1.4 Konteks dengan Modul Lain
- **Manajemen Uji Kompetensi**: Mendefinisikan ujian (jadwal, durasi, paket soal, daftar peserta). Hanya ujian dengan status **Diterbitkan** dan dalam **periode jadwal** yang tampil di CBT untuk peserta yang terdaftar.
- **Bank Soal**: Sumber soal; peserta tidak melihat Bank Soal, hanya soal yang sudah dibundel dalam ujian.

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Widyaprada | peserta uji kompetensi | melihat daftar ujian yang bisa saya kerjakan (nama, jadwal, durasi) di satu halaman CBT | memilih ujian yang sedang dibuka sesuai jadwal |
| 2 | Widyaprada | peserta uji kompetensi | hanya melihat ujian yang saya termasuk pesertanya dan yang jadwalnya sedang berjalan | tidak bingung dengan ujian yang bukan untuk saya atau sudah lewat |
| 3 | Widyaprada | peserta uji kompetensi | memulai ujian dengan satu klik (Mulai Ujian) dan melihat petunjuk singkat (durasi, jumlah soal, peringatan submit) | siap mengerjakan dengan pemahaman yang jelas |
| 4 | Widyaprada | peserta uji kompetensi | mengerjakan soal satu per satu (atau navigasi antar nomor) dengan timer yang terlihat | mengatur waktu dan tidak kehilangan sisa waktu |
| 5 | Widyaprada | peserta uji kompetensi | menjawab soal PG dengan memilih satu opsi, soal Benar–Salah dengan pilihan Benar/Salah, dan soal Essay dengan mengetik di kotak teks | menjawab sesuai tipe soal |
| 6 | Widyaprada | peserta uji kompetensi | melihat nomor soal dan status jawaban (terisi/belum) agar bisa loncat atau mengecek | tidak ada soal terlewat tanpa sengaja |
| 7 | Widyaprada | peserta uji kompetensi | mengirim jawaban (Submit) dengan konfirmasi jelas; setelah submit tidak bisa mengubah jawaban | yakin bahwa jawaban sudah tercatat dan ujian selesai |
| 8 | Widyaprada | peserta uji kompetensi | mendapat peringatan jika waktu hampir habis (misalnya 5 menit tersisa) | sempat menyimpan atau submit sebelum waktu habis |
| 9 | Widyaprada | peserta uji kompetensi | jika waktu habis sebelum submit, jawaban yang sudah diisi tetap tersimpan dan ujian dianggap selesai (auto-submit) | tidak kehilangan jawaban yang sudah dikerjakan |
| 10 | Widyaprada | peserta uji kompetensi | setelah submit (atau auto-submit) melihat konfirmasi bahwa ujian selesai dan pesan bahwa nilai akan tampil setelah dikoreksi (jika ada soal essay) atau langsung melihat nilai (jika semua otomatis) | tahu langkah berikutnya |
| 11 | Widyaprada | peserta uji kompetensi | melihat daftar ujian yang sudah saya kerjakan beserta nilai (jika sudah tersedia) di halaman CBT atau submenu "Riwayat / Hasil Saya" | memantau hasil uji kompetensi saya |
| 12 | Widyaprada | peserta uji kompetensi | tidak bisa mengerjakan ujian yang sama dua kali (satu peserta satu kali submit per ujian) | keadilan dan konsistensi penilaian |
| 13 | Widyaprada | peserta uji kompetensi | tampilan nyaman di desktop dan tablet (layout responsif) | mengerjakan dari perangkat yang tersedia |

---

### 2.1 Kebutuhan Pengguna - Yang Diinginkan
- **Daftar ujian tersedia**: Hanya ujian yang (1) status Diterbitkan, (2) dalam periode Jadwal Mulai–Selesai, (3) peserta masuk daftar. Tampil nama, jadwal, durasi, tombol "Mulai Ujian".
- **Mulai ujian**: Petunjuk singkat → konfirmasi mulai → timer mulai, soal tampil.
- **Tampilan soal**: Satu soal per layar (atau scroll dengan navigasi nomor); timer selalu terlihat; tombol Simpan jawaban (opsional per soal) dan Submit akhir.
- **Submit**: Konfirmasi ("Yakin kirim jawaban? Tidak bisa diubah.") → Submit → layar selesai + pesan (nilai akan tampil / nilai langsung).
- **Auto-submit**: Saat waktu habis, jawaban tersimpan otomatis dan ujian ditutup.
- **Riwayat / Hasil**: Daftar ujian yang sudah dikerjakan + nilai (jika sudah ada); detail per ujian (nilai, status lulus/tidak jika ada standar).

### 2.2 Kebutuhan Pengguna - Yang Tidak Diinginkan
- Peserta bisa mengerjakan ujian di luar jadwal atau ujian yang tidak termasuk daftar peserta.
- Submit tanpa konfirmasi sehingga risiko salah klik.
- Timer tidak terlihat atau tidak akurat.
- Bisa mengerjakan ujian yang sama lebih dari sekali (kecuali kebijakan ujian ulang yang eksplisit).

---

## 3. Antarmuka Pengguna (UI)

*CBT terdiri atas: halaman Daftar Ujian Tersedia, halaman Petunjuk/Mulai Ujian, halaman Pengerjaan Soal (dengan timer dan navigasi), dialog Konfirmasi Submit, halaman Selesai, dan halaman Riwayat/Hasil. Hanya role Widyaprada yang melihat menu CBT dan halaman ini.*

### 3.1 Daftar Ujian (CBT – Beranda)
- **Judul**: "Uji Kompetensi" atau "CBT" atau "Ujian Tersedia".
- **Daftar kartu/tabel**: Setiap ujian menampilkan Nama Ujian, Jadwal (mulai–selesai), Durasi (menit), tombol **"Mulai Ujian"** (atau "Lanjutkan" jika pernah mulai tapi belum submit).
- **Pesan kosong**: "Tidak ada ujian yang tersedia untuk Anda saat ini." jika tidak ada ujian yang memenuhi syarat.
- **Submenu/link**: "Riwayat Ujian" atau "Hasil Saya" ke halaman riwayat.

### 3.2 Petunjuk & Mulai Ujian
- **Sebelum mulai**: Tampilkan ringkasan (nama ujian, jumlah soal, durasi menit, peringatan "Pastikan koneksi stabil. Setelah Submit jawaban tidak dapat diubah.").
- **Tombol**: "Mulai Ujian" (mulai timer dan tampilkan soal pertama) dan "Batal" (kembali ke daftar).

### 3.3 Halaman Pengerjaan Soal
- **Timer**: Tampil jelas (countdown menit:detik) di bagian atas atau samping; peringatan saat sisa 5 menit (toast atau warna).
- **Soal**: Satu soal per tampilan (atau scroll); teks soal; untuk PG: opsi A/B/C/D (radio); untuk B-S: Benar / Salah; untuk Essay: textarea.
- **Navigasi nomor**: Daftar nomor soal (1–N) dengan indikator terisi/belum; klik untuk pindah soal.
- **Tombol**: "Simpan" (opsional, simpan jawaban tanpa submit), "Submit Ujian" (ke konfirmasi).
- **Auto-submit**: Saat timer 0, tampilkan pesan "Waktu habis. Jawaban Anda telah disimpan." dan redirect ke halaman Selesai.

### 3.4 Konfirmasi Submit
- **Dialog**: "Anda yakin ingin mengirim jawaban? Setelah submit Anda tidak dapat mengubah jawaban." Tombol "Batal" dan "Ya, Submit".
- **Setelah Submit**: Redirect ke halaman Selesai.

### 3.5 Halaman Selesai
- **Pesan**: "Ujian telah berhasil dikirim." / "Terima kasih. Jawaban Anda telah tercatat."
- **Informasi**: "Nilai akan tampil di Riwayat Ujian setelah proses koreksi." (jika ada essay) atau "Nilai Anda: X" (jika koreksi otomatis penuh).
- **Tombol**: "Kembali ke Daftar Ujian" atau "Lihat Riwayat".

### 3.6 Riwayat / Hasil Saya
- **Daftar**: Ujian yang sudah dikerjakan: nama ujian, tanggal submit, nilai (jika sudah ada), status (Lulus/Tidak lulus jika ada standar).
- **Detail** (opsional): Per ujian bisa dibuka untuk melihat nilai per kategori atau feedback (sesuai kebijakan produk).

---

## 4. Pengalaman Pengguna (UX) - Alur

### 4.1 Alur Mengerjakan Ujian
1. Widyaprada login → menu WPUjikom → CBT (atau "Ujian Tersedia").
2. Melihat daftar ujian yang tersedia → klik "Mulai Ujian" pada satu ujian.
3. Membaca petunjuk → klik "Mulai Ujian" → timer mulai, soal pertama tampil.
4. Menjawab soal, berpindah nomor jika perlu; bisa "Simpan" sementara.
5. Klik "Submit Ujian" → konfirmasi → "Ya, Submit" → halaman Selesai.
6. (Alternatif) Timer habis → auto-submit → halaman Selesai.

### 4.2 Alur Melihat Hasil
1. Dari menu CBT → "Riwayat Ujian" / "Hasil Saya".
2. Melihat daftar ujian yang sudah dikerjakan dan nilai (jika tersedia).

### 4.3 Aturan Waktu dan Sesi
- Satu peserta hanya bisa satu **sesi pengerjaan** per ujian (mulai sekali, submit sekali). Jika keluar di tengah (refresh/close), kebijakan: lanjutkan dengan sisa waktu atau dianggap submit (ditentukan produk; disarankan lanjutkan dengan sisa waktu tersimpan).
- Timer per peserta dihitung sejak peserta klik "Mulai Ujian"; tidak bergantung pada peserta lain.

---

## 5. Kebutuhan per Role

### 5.1 Role: Widyaprada
- **CBT**: Melihat daftar ujian tersedia (hanya yang jadwal + daftar peserta), mulai ujian, kerjakan soal, submit, lihat riwayat dan nilai sendiri.
- **Tidak mengakses**: Bank Soal, Manajemen Uji Kompetensi, rekap nilai peserta lain.

### 5.2 Role: Admin Uji Kompetensi
- **Tidak mengerjakan ujian** di CBT. Mengelola Bank Soal dan Manajemen Uji Kompetensi; melihat rekap hasil di Manajemen Uji Kompetensi.

### 5.3 Role: Super Admin
- Sama seperti Admin Uji Kompetensi untuk konteks CBT; plus akses penuh ke modul lain.

---

## 6. Cakupan Fitur

### 6.1 Termasuk
- Daftar ujian tersedia (filter: jadwal + peserta; status Diterbitkan).
- Mulai ujian (petunjuk + konfirmasi).
- Tampilan soal (PG, Benar–Salah, Essay) dengan timer dan navigasi nomor.
- Simpan jawaban (opsional per soal atau otomatis).
- Submit dengan konfirmasi; auto-submit saat waktu habis.
- Halaman Selesai dengan pesan dan informasi nilai (jika sudah ada).
- Riwayat ujian yang sudah dikerjakan + nilai (jika tersedia).
- Satu kali pengerjaan per peserta per ujian (no double attempt).
- Peringatan waktu hampir habis (misalnya 5 menit).

### 6.2 Tidak Termasuk
- Bank Soal, Manajemen Uji Kompetensi → PRD masing-masing.
- Proctoring (kamera, deteksi kecurangan) → fase berikutnya atau PRD terpisah.
- Sertifikat/export nilai peserta → fase berikutnya.
- Ujian ulang (retake) → kebijakan produk; jika diizinkan, atur di Manajemen Uji Kompetensi dan tampilkan di CBT.

---

## 7. Persyaratan Produk (Nonteknis)

- **Wewenang**: Hanya Widyaprada yang melihat dan mengerjakan ujian di CBT. Daftar peserta per ujian mengacu ke data dari Manajemen Uji Kompetensi.
- **Waktu**: Timer akurat (server-side atau sinkronisasi); auto-submit dan penyimpanan jawaban andal saat waktu habis atau submit.
- **Integritas**: Jawaban disimpan per soal (draft) dan final saat submit; tidak bisa mengedit setelah submit.
- **Ketersediaan**: Halaman CBT dapat diakses dalam periode ujian; di luar periode daftar ujian kosong atau hanya riwayat.
- **Akses perangkat**: Layout responsif untuk desktop dan tablet.

---

## 8. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: CBT untuk peserta Widyaprada; daftar ujian, pengerjaan, submit, riwayat/nilai | - |

---

**Catatan**: Integrasi teknis dengan Manajemen Uji Kompetensi (jadwal, paket soal, peserta, penyimpanan jawaban, perhitungan nilai) didokumentasikan di SDD. Soal essay yang memerlukan koreksi manual: nilai diinput dari sisi Manajemen Uji Kompetensi (atau modul koreksi); setelah nilai di-set, peserta melihat di Riwayat CBT.
