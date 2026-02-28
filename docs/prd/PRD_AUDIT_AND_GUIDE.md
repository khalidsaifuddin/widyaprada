# Audit PRD & Panduan Penggunaan

PRD fokus pada **product requirement** saja. Hal teknis (arsitektur, schema, API) didokumentasikan terpisah di **SDD (Software Design Document)**.

---

## 1. Pemisahan PRD vs SDD

| Dokumen | Isi | Pembaca utama |
|---------|-----|---------------|
| **PRD** | Apa yang dibangun: fitur, user story, acceptance criteria, UI/UX, alur, wewenang per role. Tanpa detail teknis implementasi. | Product, Engineer (untuk memahami "what") |
| **SDD** | Bagaimana dibangun: schema DB, API, konvensi kode, integrasi, keamanan teknis. | Engineer |

**Prinsip:** Siapa pun yang baca PRD harus bisa menjawab "apa yang harus terjadi" dan "apa yang user lihat/lakukan". "Bagaimana di backend/DB" ada di SDD.

---

## 2. Saat membuat PRD baru

Setiap PRD baru wajib memastikan: **software engineer yang baca PRD bisa langsung mulai mengembangkan tanpa harus bertanya lagi.** Tidak ada yang ambigu: fitur apa, siapa akses, alur apa, tampilan apa, pesan apa, dan acceptance criteria terdefinisi sehingga bisa di-test.

Gunakan [PRD_TEMPLATE.md](PRD_TEMPLATE.md); isi hanya bagian yang relevan dengan product. Hal teknis tetap di SDD.

---

## 3. Daftar PRD yang Ada

| PRD | Modul | Ringkasan singkat |
|-----|--------|-------------------|
| [PRD_Auth_Login.md](PRD_Auth_Login.md) | Auth & RBAC | Login satu halaman untuk semua role, redirect per role, pesan error yang ramah. |
| [PRD_Auth_Logout.md](PRD_Auth_Logout.md) | Auth & RBAC | Tombol Keluar, sesi berakhir, redirect ke login + pesan. |
| [PRD_Auth_Lupa_Password.md](PRD_Auth_Lupa_Password.md) | Auth & RBAC | Alur minta reset (email) ke tautan email ke kata sandi baru ke kembali ke login. |
| [PRD_Auth_Manajemen_Pengguna.md](PRD_Auth_Manajemen_Pengguna.md) | Auth & RBAC | CRUD pengguna, satu user bisa lebih dari satu role (multi-select), Create Widyaprada sekaligus buat data WP, delete wajib alasan. |
| [PRD_RBAC.md](PRD_RBAC.md) | Auth & RBAC | Kelola Role & Permission (hanya Super Admin); list/detail/create/edit/delete; assign permission di Edit Role. |
| [PRD_Beranda.md](PRD_Beranda.md) | Landing Page | Slider hero, Panel Berita, Panel Tautan, Panel Jurnal; konten dari CMS. |
| [PRD_Berita.md](PRD_Berita.md) | Landing Page | Halaman publik List Berita & Detail Berita (hanya published). |
| [PRD_Jurnal.md](PRD_Jurnal.md) | Landing Page | Halaman publik List Jurnal & Detail Jurnal (hanya published); sumber data WPJurnal. |
| [PRD_CMS_Landing_Page.md](PRD_CMS_Landing_Page.md) | CMS Landing Page | CMS Slider, CMS Berita, CMS Tautan (masing-masing CRUD); wewenang Admin Satker vs Super Admin. |
| [PRD_Manajemen_Data_WP.md](PRD_Manajemen_Data_WP.md) | WPData | CRUD data Widyaprada (NIP, nama, satker, dll); scope per satker; delete wajib alasan. |
| [PRD_Bank_Soal.md](PRD_Bank_Soal.md) | WPUjikom | Bank Soal: CRUD soal (PG, B-S, Essay), status verifikasi; Admin Uji Kompetensi, Verifikator Uji Kompetensi, Super Admin. |
| [PRD_Paket_Soal.md](PRD_Paket_Soal.md) | WPUjikom | Paket Soal: CRUD paket (playlist soal), status verifikasi; Admin Uji Kompetensi, Verifikator Uji Kompetensi, Super Admin. |
| [PRD_Manajemen_Uji_Kompetensi.md](PRD_Manajemen_Uji_Kompetensi.md) | WPUjikom | Manajemen ujian: jadwal, konten (soal individu + paket soal), status verifikasi, peserta, terbitkan, rekap; Admin & Verifikator Uji Kompetensi, Super Admin. |
| [PRD_CBT.md](PRD_CBT.md) | WPUjikom | CBT: peserta Widyaprada mengerjakan ujian, timer, submit, riwayat/nilai. |

**Template:** [PRD_TEMPLATE.md](PRD_TEMPLATE.md).

---

## 4. Checklist PRD (fokus product)

Untuk setiap PRD, pastikan:

- [ ] Ada ringkasan singkat yang menjelaskan fitur dan nilai untuk pengguna.
- [ ] User story per role jelas (Sebagai … Saya ingin … Agar …).
- [ ] Acceptance criteria per fitur utama bisa di-test tanpa tahu implementasi teknis.
- [ ] UI: elemen yang harus ada per halaman/langkah disebutkan; umpan balik (loading, sukses, error) jelas.
- [ ] Alur (flow) langkah demi langkah bisa diikuti pengguna.
- [ ] Wewenang per role (siapa bisa apa, scope data) eksplisit.
- [ ] In scope / out scope fitur jelas; referensi ke PRD lain jika ada ketergantungan.
- [ ] Tidak ada spesifikasi teknis implementasi (DB, API, kode); itu di SDD.

---

**Changelog**

| Tanggal | Perubahan |
|---------|-----------|
| 2025-02-11 | Dokumen awal: audit PRD fokus product, panduan engineer, pemisahan PRD vs SDD. |
