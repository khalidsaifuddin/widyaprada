## SDD – Assignment (Penugasan Uji Kompetensi)

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom  
**Fitur**: Apply Pendaftaran, Tugas Saya, Batas Waktu, Hasil, Leaderboard/Privat  

Dokumen ini menjelaskan **desain teknis backend** untuk mengimplementasikan PRD `[PRD] Fitur Assignment` dengan stack: **Go (Golang)**, Clean Architecture, PostgreSQL.

---

## 1. Arsitektur & Konteks (Backend)

- **Apply-first**: Peserta mendaftar (apply) dengan dokumen persyaratan; penugasan ditentukan oleh validasi Tim Verval — bukan admin assign langsung.
- **Penugasan** = user lolos validasi dokumen dan termasuk peserta ujian. Data dari `exam_participants` (peserta yang lolos validasi).
- **Assignment Usecase**: Apply (pilih jenis ujikom, upload dokumen), list ujian where user = peserta; include batas waktu (jadwal_selesai), status pengerjaan, nilai, pengaturan leaderboard.

---

## 2. Kontrak API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/ujikom/dokumen-persyaratan` | Daftar dokumen persyaratan sesuai jenis ujikom (query: jenis_ujikom). Sumber: `ref.dokumen_persyaratan_ujikom`. |
| POST | `/api/v1/ujikom/apply` | Apply pendaftaran Ujikom (jenis ujikom, upload dokumen) |
| GET | `/api/v1/ujikom/apply/status` | Status pendaftaran user (menunggu verifikasi / lolos / tidak lolos) |
| GET | `/api/v1/assignments` | Tugas Saya (ujian yang user ikuti) — filter: status, sort batas_waktu |
| GET | `/api/v1/assignments/:examId/result` | Hasil ujian user untuk exam tersebut |
| GET | `/api/v1/assignments/:examId/leaderboard` | Leaderboard (hanya jika tampilkan_leaderboard=Ya & hasil ada) |

**Response list assignment**: exam_id, exam_name, deadline (jadwal_selesai), status (belum_dikerjakan|sudah_dikerjakan), score, can_view_leaderboard.

---

## 3. Skema Database

- `ujikom_application`: id, user_id, jenis_ujikom (perpindahan_jabatan/kenaikan_tingkat), status_id, catatan_tolak, created_at.
- `ujikom_application_document`: id, ujikom_application_id, document_type (kode dari ref), file_path, portofolio_text (untuk portofolio/essay), created_at.
- `ref.dokumen_persyaratan_ujikom`: referensi 13 dokumen persyaratan (lihat migration 007). Kolom: id, kode, nama, urutan, tipe_input (file|text_portofolio|text_essay), batasan, deskripsi, untuk_jenis_ujikom.
- `document_type` di `ujikom_application_document` harus sesuai `kode` di `ref.dokumen_persyaratan_ujikom`. API dapat mengembalikan daftar dokumen wajib dari ref sesuai `jenis_ujikom` dan `untuk_jenis_ujikom`.
- Memanfaatkan `exams`, `exam_participants`, `exam_attempts`, `exams.tampilkan_leaderboard`.
- Peserta yang lolos validasi → insert ke `exam_participants`.

### 3.1 Referensi Dokumen Persyaratan (`ref.dokumen_persyaratan_ujikom`)

| Kode | Tipe Input | Batasan | Untuk Jenis Ujikom |
|------|------------|---------|--------------------|
| surat_usul_pimpinan | file | - | perpindahan_jabatan |
| sk_kenaikan_pangkat_terakhir | file | - | - |
| sk_jabatan_terakhir | file | - | - |
| surat_pernyataan_integritas_moralitas | file | - | - |
| surat_keterangan_sehat | file | - | - |
| fotokopi_ijazah | file | - | - |
| surat_keterangan_pengalaman_2tahun | file | - | - |
| surat_pernyataan_lowongan | file | - | - |
| surat_pernyataan_tidak_menuntut | file | - | - |
| penilaian_skp_2tahun | file | 24 aktivitas di bulan berbeda | - |
| portofolio | text_portofolio | Min 1 kegiatan/bulan; 2 tahun | - |
| essay_inovasi_praktik_baik | text_essay | Maks 1500 kata | perpindahan_jabatan |
| surat_pernyataan_orisinalitas_essay | file | - | perpindahan_jabatan |

---

## 4. Aturan Bisnis

- **Leaderboard**: Hanya jika `exams.tampilkan_leaderboard = true` dan nilai peserta sudah ada. Akses hanya untuk peserta ujian tersebut.
- **Privat**: Jika `tampilkan_leaderboard = false`, endpoint leaderboard return 403.
- **Leaderboard content**: Peringkat, identitas (sesuai kebijakan), nilai. Urutan nilai tinggi ke rendah.

---

## 5. RBAC

- Role **Widyaprada** (peserta). Hanya melihat assignment sendiri dan leaderboard ujian yang diikuti (jika diizinkan).
