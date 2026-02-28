# Koleksi PRD – Aplikasi Widyaprada

**PRD (Product Requirements Document)** mendefinisikan **apa** yang dibangun dari sisi produk: fitur, user story, acceptance criteria, UI/UX, wewenang per role. Tidak berisi detail teknis implementasi (schema, API, kode); itu didokumentasikan di **SDD (Software Design Document)** di folder `docs/`.

- **Untuk engineer:** Baca PRD fitur yang dikerjakan, lalu SDD untuk desain teknis.
- **Tujuan PRD:** Software engineer yang baca PRD bisa langsung mulai mengembangkan tanpa harus bertanya lagi; tidak ada yang ambigu atau tidak jelas.

Panduan dan audit PRD: [PRD_AUDIT_AND_GUIDE.md](PRD_AUDIT_AND_GUIDE.md).  
Desain teknis: [docs/sdd/README.md](../sdd/README.md).

---

## Deskripsi Proyek

**Nama:** Aplikasi Widyaprada  

**Deskripsi:** Aplikasi yang merupakan sekumpulan layanan untuk Widyaprada. Widyaprada adalah Pegawai Negeri Sipil (PNS) dengan jabatan fungsional yang bertugas melaksanakan penjaminan mutu pendidikan pada jenjang PAUD, pendidikan dasar, pendidikan menengah, dan pendidikan masyarakat. Mereka bertanggung jawab dalam pemetaan, pendampingan, supervisi, dan pengembangan model mutu untuk memastikan standar pendidikan tercapai.

---

## Auth & RBAC

| Fitur | PRD | Status |
|-------|-----|--------|
| Login | [PRD_Auth_Login.md](PRD_Auth_Login.md) | Ada |
| Registrasi Peserta (Calon WP) | [PRD_Auth_Registrasi.md](PRD_Auth_Registrasi.md) | Ada |
| Lupa Password | [PRD_Auth_Lupa_Password.md](PRD_Auth_Lupa_Password.md) | Ada |
| Logout | [PRD_Auth_Logout.md](PRD_Auth_Logout.md) | Ada |
| Manajemen Pengguna (List, Detail, Create, Edit, Delete) | [PRD_Auth_Manajemen_Pengguna.md](PRD_Auth_Manajemen_Pengguna.md) | Ada |
| Role & Permission (Kelola Role, Kelola Permission) | [PRD_RBAC.md](PRD_RBAC.md) | Ada |

---

## Landing Page

| Fitur | PRD | Status |
|-------|-----|--------|
| Beranda (Slider, Panel Berita, Panel Tautan, Panel Jurnal) | [PRD_Beranda.md](PRD_Beranda.md) | Ada |
| Berita – List & Detail (publik) | [PRD_Berita.md](PRD_Berita.md) | Ada |
| Jurnal – List & Detail (publik) | [PRD_Jurnal.md](PRD_Jurnal.md) | Ada |
| CMS Landing Page (Slider, Berita, Tautan – CRUD) | [PRD_CMS_Landing_Page.md](PRD_CMS_Landing_Page.md) | Ada |

---

## WPData

| Fitur | PRD | Status |
|-------|-----|--------|
| Manajemen Data WP (List, Detail, Create, Edit, Delete) | [PRD_Manajemen_Data_WP.md](PRD_Manajemen_Data_WP.md) | Ada |

---

## WPJurnal

| Fitur | PRD | Status |
|-------|-----|--------|
| Manajemen Jurnal (List, Detail, Create/Submit, Edit, Delete, Verifikasi) | Belum ada | Belum ada PRD |

---

## WPUjikom

**Role WPUjikom:** **Admin Uji Kompetensi** (CRUD Bank Soal, Paket Soal, Manajemen Ujian); **Verifikator Uji Kompetensi** (verifikasi soal, paket soal, dan ujian — memberi label Sudah/Belum Diverifikasi). Widyaprada mengakses CBT sebagai peserta ujian saja. **Satu pengguna dapat memiliki lebih dari satu role** (misalnya Widyaprada sekaligus Admin Uji Kompetensi atau Verifikator Uji Kompetensi). Soal, paket soal, dan ujian memiliki **status verifikasi** (Belum Diverifikasi / Sudah Diverifikasi). Satu ujian dapat memuat **soal individu** dan/atau **paket soal**.

| Fitur | PRD | Status |
|-------|-----|--------|
| Bank Soal | [PRD_Bank_Soal.md](PRD_Bank_Soal.md) | Ada |
| Paket Soal | [PRD_Paket_Soal.md](PRD_Paket_Soal.md) | Ada |
| Manajemen Uji Kompetensi | [PRD_Manajemen_Uji_Kompetensi.md](PRD_Manajemen_Uji_Kompetensi.md) | Ada |
| CBT | [PRD_CBT.md](PRD_CBT.md) | Ada |
| Assignment (Penugasan Uji Kompetensi, Batas Waktu, Hasil, Leaderboard/Privat) | [PRD_Assignment.md](PRD_Assignment.md) | Ada |

---

## Template

PRD baru mengikuti: [PRD_TEMPLATE.md](PRD_TEMPLATE.md). Isi hanya bagian yang relevan dengan **product requirement**; hal teknis didokumentasikan di SDD.
