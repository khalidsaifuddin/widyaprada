# Database Migrations - Aplikasi Widyaprada

Migration PostgreSQL mengikuti konvensi:

- **Schema `ref`**: data referensi (agama, jenis_kelamin, satker, status, role, permission, tipe_soal, kategori_kompetensi, tingkat_kesulitan, dll).
- **Schema `public`**: data utama dan transaksional (user, user_role, widyaprada, berita, slide, tautan, jurnal, soal, paket_soal, ujian, ujian_peserta, ujikom_application, ujikom_application_document, dll).
- Setiap tabel: `id` UUID PK, metadata `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_at`, `deleted_by`.
- Timestamp memakai `TIMESTAMPTZ`.
- FK kolom format: `nama_tabel_id`.

Sesuai PRD terbaru: Auth/RBAC (user–role many-to-many), Manajemen Uji Kompetensi, Bank Soal, Paket Soal, CBT, Assignment (tampilkan_leaderboard), **Apply-first flow** (ujikom_application, ujikom_application_document).

## Urutan eksekusi

1. `001_create_ref_schema.sql` — buat schema ref dan semua tabel referensi (termasuk level_wilayah dan mst_wilayah dengan seed data).
2. `002_create_public_schema.sql` — buat tabel di schema public (user dengan role_id, role_permission, widyaprada, slide, berita, tautan, jurnal).
3. `003_user_role_many_to_many.sql` — ubah user–role ke many-to-many: buat tabel `user_role`, migrasi data, hapus kolom `role_id` dari `user`.
4. `004_ujikom_bank_soal_ujian_cbt.sql` — tambah ref (tipe_soal, kategori_kompetensi, tingkat_kesulitan), status untuk soal/paket/ujian/verifikasi, dan tabel public: soal, soal_opsi, paket_soal, paket_soal_item, ujian, ujian_soal, ujian_paket, ujian_peserta.
5. `005_ujikom_applications.sql` — Apply-first flow: ujikom_application, ujikom_application_document; status ujikom (menunggu_verifikasi, lolos, tidak_lolos).
6. `006_user_ref_adjustments.sql` — user.satker_id nullable (calon peserta); ref.permission tambah group, description; ref.role tambah description.
7. `007_dokumen_persyaratan_ujikom.sql` — ref.dokumen_persyaratan_ujikom: referensi 13 dokumen persyaratan apply uji kompetensi (Non-WP ke WP / Widyaprada Ahli Madya).

## Auto-migration (backend service)

Backend menjalankan auto-migration saat startup ketika `DB_TYPE=postgres`:

1. Membuat tabel `public.schema_migrations` untuk tracking versi
2. Mendeteksi versi database saat ini (berdasarkan migration yang sudah dijalankan)
3. Memindai file migration (`NNN_*.sql`) dan mendeteksi yang belum dijalankan (stale)
4. Menjalankan migration yang tertunda secara berurutan

**Config**: `MIGRATIONS_PATH` (default: `migrations`). Jika backend dijalankan dari `backend/`, set `MIGRATIONS_PATH=../migrations`.

Auto-migration tidak dijalankan untuk SQLite (schema tetap via GORM AutoMigrate).

## Dummy Data (Development)

Saat `ENVIRONMENT=development`, backend menjalankan seed dummy data otomatis untuk **SQLite dan PostgreSQL**. Data meliputi:

- **User login**:
  | Email | Password | Role |
  |-------|----------|------|
  | admin@example.com | Password123! | Super Admin |
  | peserta@example.com | Password123! | Peserta |

- **Berita**: 3 artikel (pengumuman, workshop, pedoman CBT)
- **Bank Soal**: 3 soal pilihan ganda
- **Paket Soal**: 1 paket (Paket Soal Umum Widyaprada)
- **Ujian**: 1 ujian diterbitkan (Uji Kompetensi Widyaprada Dasar 2025), peserta demo sudah terdaftar
- **Slide**: 2 slide hero
- **Jurnal**: 2 artikel jurnal
- **Tautan**: 2 link (Ditjen PAUD Dikdas Dikmen, Pusat Asesmen)

Seed bersifat idempotent: hanya dijalankan jika data belum ada.

## Cara jalankan manual (psql)

```bash
# Dari direktori project, dengan psql:
psql -U <user> -d <database> -f migrations/001_create_ref_schema.sql
psql -U <user> -d <database> -f migrations/002_create_public_schema.sql
psql -U <user> -d <database> -f migrations/003_user_role_many_to_many.sql
psql -U <user> -d <database> -f migrations/004_ujikom_bank_soal_ujian_cbt.sql
psql -U <user> -d <database> -f migrations/005_ujikom_applications.sql
psql -U <user> -d <database> -f migrations/006_user_ref_adjustments.sql
psql -U <user> -d <database> -f migrations/007_dokumen_persyaratan_ujikom.sql
```

Atau gunakan tool migration (golang-migrate, goose, dll) dengan urutan file di atas.
