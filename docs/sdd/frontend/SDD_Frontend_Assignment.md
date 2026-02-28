## SDD Frontend – Assignment (Penugasan Uji Kompetensi)

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom  
**Fitur**: Assignment – Tugas Saya, Apply Ujikom, Hasil, Leaderboard  
**PRD Terkait**: [PRD_Assignment](../../prd/PRD_Assignment.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk Assignment dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/wpujikom/assignment` (Tugas Saya), `/wpujikom/assignment/apply`, `/wpujikom/assignment/[examId]/hasil`, `/wpujikom/assignment/[examId]/leaderboard`
- **Role**: Peserta (Tugas Saya, Apply, Hasil, Leaderboard); Admin/Verifikator (daftar calon, verifikasi, jadwal) — sebagian di Manajemen Uji Kompetensi
- **Layout**: DashboardLayout
- **API**: `/api/v1/assignment/*`

---

## 2. Atomic Design – Komponen

### 2.1 Tugas Saya (Penugasan)

#### Atoms
- `Button` | Mulai Ujian, Lihat Hasil, Lihat Leaderboard |
- `Badge` | Belum dikerjakan, Sudah dikerjakan |

#### Molecules
- `AssignmentCardItem` | Nama ujian, batas waktu, status, hasil; tombol aksi (sesuai kondisi) |
- `EmptyState` | "Anda belum memiliki penugasan ujian." |

#### Organisms
- `AssignmentList` | Daftar assignment; filter Belum/Sudah dikerjakan; sort batas waktu; paginasi |

### 2.2 Halaman Hasil (per ujian)

#### Organisms
- `ExamResultCard` | Nama ujian, nilai/skor, status Lulus/Tidak; tombol Lihat Leaderboard (jika ujian Leaderboard) |
- Jika Privat: tidak ada tombol leaderboard |

### 2.3 Halaman Leaderboard

#### Atoms
- `Badge` | Highlight baris user login |

#### Molecules
- `LeaderboardRow` | Peringkat, identitas peserta (sesuai kebijakan), nilai |

#### Organisms
- `LeaderboardTable` | Judul ujian; tabel ranking; highlight baris user; paginasi jika > 50 peserta |
- Akses: hanya ujian dengan Tampilkan Leaderboard = Ya dan hasil tersedia; hanya peserta ujian tersebut |

### 2.4 Apply Ujikom (Pendaftaran)

#### Molecules
- `FormField` | Nama, email, jenis ujikom, file upload (13 dokumen persyaratan), portofolio (text), essay |
| `FileUploadField` | Upload file dengan validasi |

#### Organisms
- `ApplyUjikomForm` | Pilih jenis ujikom; isi & upload kelengkapan berkas; submit |
- `ApplyUjikomSuccess` | Konfirmasi "Pendaftaran berhasil. Status: Menunggu verifikasi." |
- `DokumenPersyaratanList` | Daftar 13 dokumen dengan input per item (file/text) |

### 2.5 Admin/Verifikator (dalam Manajemen Uji Kompetensi)

- Daftar calon peserta, verifikasi dokumen, tolak dengan catatan, atur jadwal
- Lihat SDD_Frontend_Manajemen_Uji_Kompetensi

### 2.6 Pages

| Route | Page |
|-------|------|
| `/wpujikom/assignment` | AssignmentListPage (Tugas Saya) |
| `/wpujikom/assignment/apply` | ApplyUjikomPage |
| `/wpujikom/assignment/[examId]/hasil` | ExamResultPage |
| `/wpujikom/assignment/[examId]/leaderboard` | LeaderboardPage (conditional: 403 jika privat) |

---

## 3. State & Validasi

- Apply: validasi file (tipe, ukuran); portofolio minimal 1 kegiatan/bulan; essay maks 1500 kata
- Leaderboard: akses ditolak (403) untuk ujian privat atau non-peserta

---

## 4. Integrasi API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/assignment` | List tugas saya |
| POST | `/api/v1/assignment/apply` | Apply ujikom (dokumen) |
| GET | `/api/v1/assignment/[examId]/result` | Hasil ujian |
| GET | `/api/v1/assignment/[examId]/leaderboard` | Leaderboard (jika Ya) |
| GET | `/api/v1/assignment/announcement` | Pengumuman (beranda) |

---

## 5. File Lokasi

```
frontend/src/
├── app/wpujikom/assignment/
│   ├── page.tsx
│   ├── apply/page.tsx
│   └── [examId]/hasil/page.tsx, [examId]/leaderboard/page.tsx
├── components/organisms/AssignmentList.tsx
├── components/organisms/LeaderboardTable.tsx
├── components/organisms/ApplyUjikomForm.tsx
└── components/molecules/DokumenPersyaratanList.tsx
```
