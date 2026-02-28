# Software Design Document (SDD) – Aplikasi Widyaprada

**SDD** mendokumentasikan **bagaimana** aplikasi dibangun dari sisi teknis. PRD menjelaskan *apa* yang dibangun (fitur, user story, UI/UX); SDD menjelaskan *cara* implementasinya.

---

## Isi SDD (ruang lingkup)

- **Schema database:** Tabel, kolom, relasi, indeks. Bisa merujuk ke [erd.puml](../erd.puml) dan file di [migrations/](../../migrations/).
- **API:** Endpoint, method, request/response, autentikasi/otorisasi.
- **Konvensi kode:** Struktur folder, penamaan, pola yang dipakai (backend/frontend).
- **Keamanan teknis:** Validasi di backend, penanganan token/sesi, rate limit, dll.
- **Integrasi:** Service eksternal, queue, cache (jika ada).

Persyaratan produk yang menyebut "validasi wewenang", "alasan penghapusan disimpan", "kode unik", dll.: *cara* implementasinya dideskripsikan di sini.

---

## Dokumen terkait

- **Tech stack:** [SDD_Tech_Stack.md](SDD_Tech_Stack.md) – bahasa, framework, database, cache, CSS, UI library.
- **ERD:** [docs/erd.puml](../erd.puml) (diagram entitas dan relasi).
- **Migrations:** [migrations/](../../migrations/) (SQL schema).
- **PRD:** Semua file PRD di [docs/prd/](../prd/); daftar di [PRD Collection Skeleton](../prd/PRD%20Collection%20Skeleton.md).

---

## Status Implementasi

Lihat [IMPLEMENTATION_STATUS.md](../IMPLEMENTATION_STATUS.md) untuk daftar SDD yang sudah dan belum diimplementasi.

---

## Daftar SDD

| SDD | PRD Terkait | Modul |
|-----|-------------|-------|
| [SDD_Auth_Login](SDD_Auth_Login.md) | PRD_Auth_Login | Auth & RBAC |
| [SDD_Auth_Registrasi](SDD_Auth_Registrasi.md) | PRD_Auth_Registrasi | Auth & RBAC |
| [SDD_Auth_Logout](SDD_Auth_Logout.md) | PRD_Auth_Logout | Auth & RBAC |
| [SDD_Auth_Lupa_Password](SDD_Auth_Lupa_Password.md) | PRD_Auth_Lupa_Password | Auth & RBAC |
| [SDD_Auth_Manajemen_Pengguna](SDD_Auth_Manajemen_Pengguna.md) | PRD_Auth_Manajemen_Pengguna | Auth & RBAC |
| [SDD_RBAC](SDD_RBAC.md) | PRD_RBAC | Auth & RBAC |
| [SDD_Dashboard_Widyaprada](SDD_Dashboard_Widyaprada.md) | PRD_Dashboard_Widyaprada | Dashboard |
| [SDD_Bank_Soal](SDD_Bank_Soal.md) | PRD_Bank_Soal | WPUjikom |
| [SDD_Paket_Soal](SDD_Paket_Soal.md) | PRD_Paket_Soal | WPUjikom |
| [SDD_CBT](SDD_CBT.md) | PRD_CBT | WPUjikom |
| [SDD_Manajemen_Uji_Kompetensi](SDD_Manajemen_Uji_Kompetensi.md) | PRD_Manajemen_Uji_Kompetensi | WPUjikom |
| [SDD_Assignment](SDD_Assignment.md) | PRD_Assignment | WPUjikom |
| [SDD_Beranda](SDD_Beranda.md) | PRD_Beranda | Landing Page |
| [SDD_Berita](SDD_Berita.md) | PRD_Berita | Landing Page |
| [SDD_Jurnal](SDD_Jurnal.md) | PRD_Jurnal | Landing Page |
| [SDD_CMS_Landing_Page](SDD_CMS_Landing_Page.md) | PRD_CMS_Landing_Page | CMS |
| [SDD_Manajemen_Data_WP](SDD_Manajemen_Data_WP.md) | PRD_Manajemen_Data_WP | Manajemen Data WP |

## Stack Backend

- **Bahasa**: Go (Golang)
- **Arsitektur**: Clean Architecture (handler → usecase → repository)
- **Database**: PostgreSQL
- **Cache**: Redis (untuk auth, rate limiting)
- **API**: REST + JSON
- **Base path**: /api/v1

