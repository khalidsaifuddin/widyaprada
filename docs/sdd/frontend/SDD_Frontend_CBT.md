## SDD Frontend – CBT (Computer-Based Test)

**Aplikasi**: Widyaprada  
**Modul**: WPUjikom  
**Fitur**: CBT – Mengerjakan Ujian dan Melihat Hasil  
**PRD Terkait**: [PRD_CBT](../../prd/PRD_CBT.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk CBT dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/wpujikom/cbt` (daftar), `/wpujikom/cbt/[examId]/mulai`, `/wpujikom/cbt/[examId]/kerjakan`, `/wpujikom/cbt/riwayat`
- **Role**: Widyaprada (peserta)
- **Layout**: DashboardLayout; halaman pengerjaan bisa fullscreen/layout khusus
- **API**: `/api/v1/cbt/*`

---

## 2. Atomic Design – Komponen

### 2.1 Daftar Ujian (CBT Beranda)

#### Atoms
- `Button` (Mulai Ujian, Lanjutkan), `Badge` (status)

#### Molecules
- `ExamCard` | Nama ujian, jadwal, durasi, tombol Mulai Ujian / Lanjutkan |
- `EmptyState` | "Tidak ada ujian yang tersedia untuk Anda saat ini." |

#### Organisms
- `CBTExamList` | Daftar kartu ujian tersedia; link "Riwayat Ujian" |

### 2.2 Petunjuk & Mulai Ujian

#### Organisms
- `ExamInstructions` | Ringkasan: nama ujian, jumlah soal, durasi, peringatan "Pastikan koneksi stabil..."; tombol Mulai Ujian, Batal |

### 2.3 Halaman Pengerjaan Soal

#### Atoms
- `Button` (Simpan, Submit Ujian), `RadioGroup`, `Textarea`, `Badge`

#### Molecules
- `TimerDisplay` | Countdown menit:detik; peringatan warna saat sisa 5 menit |
- `QuestionNumberGrid` | Daftar nomor soal dengan indikator terisi/belum; klik untuk pindah |
- `QuestionContent` | Teks soal; PG: opsi A–D; B-S: Benar/Salah; Essay: textarea |
- `SubmitConfirmDialog` | "Anda yakin ingin mengirim jawaban? Setelah submit Anda tidak dapat mengubah jawaban." Batal | Ya, Submit |

#### Organisms
- `CBTWorkspace` | Timer + QuestionNumberGrid + QuestionContent + Simpan + Submit Ujian |
- `CBTQuestionView` | Satu soal per tampilan (atau scroll) |

### 2.4 Halaman Selesai

#### Organisms
- `ExamFinishedMessage` | "Ujian telah berhasil dikirim."; info nilai (jika ada) atau "Nilai akan tampil di Riwayat..."; tombol Kembali ke Daftar, Lihat Riwayat |
- `AutoSubmitMessage` | "Waktu habis. Jawaban Anda telah disimpan." (saat timer 0) |

### 2.5 Riwayat / Hasil Saya

#### Organisms
- `CBTHistoryList` | Daftar ujian yang sudah dikerjakan: nama, tanggal submit, nilai (jika ada), status Lulus/Tidak |
- `CBTResultDetail` | Detail nilai per ujian (opsional) |

### 2.6 Pages

| Route | Page |
|-------|------|
| `/wpujikom/cbt` | CBTListPage (CBTExamList) |
| `/wpujikom/cbt/[examId]/mulai` | CBTInstructionsPage (ExamInstructions) |
| `/wpujikom/cbt/[examId]/kerjakan` | CBTWorkspacePage (CBTWorkspace) |
| `/wpujikom/cbt/selesai` | CBTFinishedPage (ExamFinishedMessage) |
| `/wpujikom/cbt/riwayat` | CBTHistoryPage (CBTHistoryList) |

---

## 3. State & Logika

### 3.1 CBTWorkspace
- `answers`: Record<questionId, value>
- `currentQuestionIndex`: number
- `timeRemaining`: number (detik)
- `submitted`: boolean
- Timer: countdown client-side; sinkronisasi dengan backend jika perlu
- Auto-submit: saat timeRemaining = 0 → kirim jawaban → redirect ke Selesai
- Simpan jawaban: PATCH per soal atau batch (sesuai API)

### 3.2 Validasi
- Satu kali pengerjaan per peserta per ujian (dicek backend)

---

## 4. Integrasi API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/cbt/exams` | Daftar ujian tersedia |
| GET | `/api/v1/cbt/exams/:id` | Detail ujian (petunjuk) |
| POST | `/api/v1/cbt/exams/:id/start` | Mulai ujian (start timer) |
| GET | `/api/v1/cbt/exams/:id/questions` | Daftar soal |
| PATCH | `/api/v1/cbt/exams/:id/answers` | Simpan jawaban (draft) |
| POST | `/api/v1/cbt/exams/:id/submit` | Submit final |
| GET | `/api/v1/cbt/history` | Riwayat ujian + nilai |

---

## 5. Umpan Balik & Aksesibilitas

- Peringatan sisa 5 menit: toast atau perubahan warna timer
- Konfirmasi submit wajib
- Layout responsif desktop/tablet
- Keyboard navigasi antar soal

---

## 6. File Lokasi

```
frontend/src/
├── app/wpujikom/cbt/
│   ├── page.tsx
│   ├── [examId]/mulai/page.tsx
│   ├── [examId]/kerjakan/page.tsx
│   ├── selesai/page.tsx
│   └── riwayat/page.tsx
├── components/organisms/CBTExamList.tsx
├── components/organisms/CBTWorkspace.tsx
├── components/organisms/ExamInstructions.tsx
├── components/molecules/TimerDisplay.tsx
└── components/molecules/QuestionNumberGrid.tsx
```
