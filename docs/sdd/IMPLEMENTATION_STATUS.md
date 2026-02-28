# SDD Implementation Status

Dokumen ini mencatat status implementasi SDD. Perbarui file ini setiap selesai mengimplementasikan satu SDD.

**Terakhir diperbarui:** 2026-03-01 (Landing Page, CMS, Manajemen Data WP)

---

## Backend

| SDD | Modul | Status | Catatan |
|-----|-------|--------|---------|
| [SDD_Auth_Login](backend/SDD_Auth_Login.md) | Auth & RBAC | ✅ Done | POST /api/v1/auth/login |
| [SDD_Auth_Registrasi](backend/SDD_Auth_Registrasi.md) | Auth & RBAC | ✅ Done | POST /api/v1/auth/register |
| [SDD_Auth_Logout](backend/SDD_Auth_Logout.md) | Auth & RBAC | ✅ Done | POST /api/v1/auth/logout |
| [SDD_Auth_Lupa_Password](backend/SDD_Auth_Lupa_Password.md) | Auth & RBAC | ✅ Done | forgot-password, reset-password |
| [SDD_Auth_Manajemen_Pengguna](backend/SDD_Auth_Manajemen_Pengguna.md) | Auth & RBAC | ✅ Done | CRUD /api/v1/users |
| [SDD_RBAC](backend/SDD_RBAC.md) | Auth & RBAC | ✅ Done | /api/v1/rbac/roles, /api/v1/rbac/permissions |
| [SDD_Dashboard_Widyaprada](docs/sdd/backend/SDD_Dashboard_Widyaprada.md) | Dashboard | ✅ Done | GET /api/v1/dashboard/assignments, /journals |
| [SDD_Bank_Soal](backend/SDD_Bank_Soal.md) | WPUjikom | ✅ Done | CRUD /api/v1/questions, verify, categories |
| [SDD_Paket_Soal](backend/SDD_Paket_Soal.md) | WPUjikom | ✅ Done | CRUD /api/v1/question-packages, verify |
| [SDD_CBT](backend/SDD_CBT.md) | WPUjikom | ✅ Done | /api/v1/cbt: exams, start, questions, answers, submit, history |
| [SDD_Manajemen_Uji_Kompetensi](backend/SDD_Manajemen_Uji_Kompetensi.md) | WPUjikom | ✅ Done | CRUD /api/v1/exams, publish, verify, unverify |
| [SDD_Assignment](backend/SDD_Assignment.md) | WPUjikom | ✅ Done | ujikom/dokumen-persyaratan, apply, status; assignments CRUD, result, leaderboard |
| [SDD_Beranda](docs/sdd/backend/SDD_Beranda.md) | Landing Page | ✅ Done | GET /api/v1/landing/home, /api/v1/beranda/pengumuman |
| [SDD_Berita](docs/sdd/backend/SDD_Berita.md) | Landing Page | ✅ Done | GET /api/v1/berita, /api/v1/berita/:slug |
| [SDD_Jurnal](docs/sdd/backend/SDD_Jurnal.md) | Landing Page | ✅ Done | GET /api/v1/jurnal, /api/v1/jurnal/:id |
| [SDD_CMS_Landing_Page](docs/sdd/backend/SDD_CMS_Landing_Page.md) | CMS | ✅ Done | CRUD /api/v1/cms/slider, /berita, /tautan |
| [SDD_Manajemen_Data_WP](docs/sdd/backend/SDD_Manajemen_Data_WP.md) | Manajemen Data WP | ✅ Done | CRUD /api/v1/wp-data, calon-peserta, verify |

### Ringkasan Backend

- **Done:** 17
- **Not started:** 0

---

## Frontend

| SDD | Modul | Status | Catatan |
|-----|-------|--------|---------|
| [SDD_Frontend_Auth_Login](frontend/SDD_Frontend_Auth_Login.md) | Auth & RBAC | ✅ Done | /auth/login, POST /api/v1/auth/login |
| [SDD_Frontend_Auth_Registrasi](frontend/SDD_Frontend_Auth_Registrasi.md) | Auth & RBAC | ✅ Done | /auth/register, POST /api/v1/auth/register |
| [SDD_Frontend_Auth_Logout](frontend/SDD_Frontend_Auth_Logout.md) | Auth & RBAC | ✅ Done | POST /api/v1/auth/logout, redirect + toast |
| [SDD_Frontend_Auth_Lupa_Password](frontend/SDD_Frontend_Auth_Lupa_Password.md) | Auth & RBAC | ✅ Done | forgot-password, reset-password |
| [SDD_Frontend_Auth_Manajemen_Pengguna](frontend/SDD_Frontend_Auth_Manajemen_Pengguna.md) | Auth & RBAC | ✅ Done | CRUD /auth/manajemen-pengguna |
| [SDD_Frontend_RBAC](frontend/SDD_Frontend_RBAC.md) | Auth & RBAC | ✅ Done | /auth/role, /auth/permission, CRUD roles & permissions |
| [SDD_Frontend_Dashboard_Widyaprada](frontend/SDD_Frontend_Dashboard_Widyaprada.md) | Dashboard | ✅ Done | /dashboard: AssignmentBlock, JournalBlock, sapaan user |
| [SDD_Frontend_Bank_Soal](frontend/SDD_Frontend_Bank_Soal.md) | WPUjikom | ✅ Done | CRUD v1/questions, verify/unverify, categories |
| [SDD_Frontend_Paket_Soal](frontend/SDD_Frontend_Paket_Soal.md) | WPUjikom | ✅ Done | CRUD v1/question-packages, verify/unverify, pilih soal |
| [SDD_Frontend_CBT](docs/sdd/frontend/SDD_Frontend_CBT.md) | WPUjikom | ✅ Done | /wpujikom/cbt, mulai, kerjakan, selesai, riwayat |
| [SDD_Frontend_Manajemen_Uji_Kompetensi](docs/sdd/frontend/SDD_Frontend_Manajemen_Uji_Kompetensi.md) | WPUjikom | ✅ Done | /wpujikom/uji-kompetensi CRUD, publish, verify |
| [SDD_Frontend_Assignment](docs/sdd/frontend/SDD_Frontend_Assignment.md) | WPUjikom | ✅ Done | Tugas Saya, apply, hasil, leaderboard |
| [SDD_Frontend_Beranda](frontend/SDD_Frontend_Beranda.md) | Landing Page | ✅ Done | /, HeroSlider, NewsPanel, LinksPanel, JournalPanel, AnnouncementPanel |
| [SDD_Frontend_Berita](frontend/SDD_Frontend_Berita.md) | Landing Page | ✅ Done | /berita, /berita/[slug], BeritaList, BeritaDetail |
| [SDD_Frontend_Jurnal](frontend/SDD_Frontend_Jurnal.md) | Landing Page | ✅ Done | /jurnal, /jurnal/[id], JurnalList, JurnalDetail |
| [SDD_Frontend_CMS_Landing_Page](frontend/SDD_Frontend_CMS_Landing_Page.md) | CMS | ✅ Done | /cms/slider, /cms/berita, /cms/tautan CRUD |
| [SDD_Frontend_Manajemen_Data_WP](frontend/SDD_Frontend_Manajemen_Data_WP.md) | Manajemen Data WP | ✅ Done | /wpdata CRUD, list, create, [id], [id]/edit |

### Ringkasan Frontend

- **Done:** 17
- **Not started:** 0

---

## Cara memperbarui

Setelah menyelesaikan implementasi satu SDD:

1. Ubah status dari `❌ Not started` menjadi `✅ Done`
2. Isi kolom **Catatan** dengan endpoint atau ringkasan singkat (opsional)
3. Perbarui **Terakhir diperbarui** di bagian atas
4. Perbarui angka di **Ringkasan**
