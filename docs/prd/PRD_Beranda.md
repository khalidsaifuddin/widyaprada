# [PRD] Fitur Beranda (Landing Page)
## Product Requirements Document | Landing Page

**Aplikasi**: Widyaprada  
**Modul**: Landing Page  
**Fitur**: Beranda (Slider besar, Panel Berita, Panel Tautan, Panel Jurnal)  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Beranda
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-11
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Beranda adalah halaman utama landing page aplikasi Widyaprada yang **dapat diakses secara publik** tanpa harus login. Pengunjung melihat **slider besar (hero slider)** di bagian atas, lalu **Panel Berita**, **Panel Tautan**, dan **Panel Jurnal**. Di header, pengunjung yang belum login melihat tombol **Login**; setelah login, tombol berubah menjadi **Dashboard**. Untuk **Peserta (Calon/WP)** yang sudah login, beranda juga menampilkan **Panel Pengumuman** yang berisi: (1) **Pengumuman hasil seleksi administrasi** — jika tidak lolos ada informasi; (2) **Info menunggu jadwal Ujikom** (ditentukan admin); (3) **Tombol Mulai Ujikom** ketika jadwal tersedia dan diinformasikan.

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Pengunjung / Semua | pengunjung beranda | melihat slider besar di bagian atas dengan gambar dan pesan utama yang berganti (slide) | mendapat kesan pertama yang jelas dan informasi highlight |
| 2 | Pengunjung / Semua | pengunjung beranda | melihat panel berita berisi berita terbaru yang sudah dipublikasikan | mengakses berita terbaru tanpa masuk ke halaman Berita dulu |
| 3 | Pengunjung / Semua | pengunjung beranda | melihat panel tautan berisi tautan penting | mengakses link terkait (situs lain, dokumen, dll) dengan cepat |
| 4 | Pengunjung / Semua | pengunjung beranda | melihat panel jurnal berisi jurnal yang sudah dipublikasikan | melihat ringkasan jurnal terbaru dari landing |
| 5 | Pengunjung / Semua | pengunjung beranda | mengklik slide, berita, tautan, atau jurnal untuk ke detail/halaman tujuan | menelusuri konten lebih lanjut |
| 5a | Pengunjung / Semua | pengunjung beranda | mengakses beranda dan halaman publik (Berita, Jurnal) tanpa harus login | melihat konten publik kapan saja |
| 5b | Pengunjung / Semua | pengunjung beranda | melihat tombol **Login** di header jika belum login | masuk ke aplikasi jika ingin mengakses fitur terproteksi |
| 5c | Pengunjung / Semua | pengguna yang sudah login | melihat tombol **Dashboard** di header (menggantikan tombol Login) | cepat beralih ke dashboard |
| 6 | Admin Satker / Super Admin | pengelola konten | mengelola konten slider (slide) melalui CMS Slider (List, Detail, Create, Edit, Delete) | mengubah gambar, judul, tautan, dan urutan slide yang tampil di beranda |
| 7 | Peserta | peserta ujian | melihat **pengumuman hasil seleksi administrasi** di beranda (jika tidak lolos ada informasi; jika lolos info menunggu jadwal) | tahu status pendaftaran dan langkah selanjutnya |
| 8 | Peserta | peserta ujian | melihat **info jadwal Ujikom** yang telah ditetapkan admin di beranda | tahu kapan dapat mengerjakan ujian |
| 9 | Peserta | peserta ujian | mengklik **Mulai Ujikom** di beranda ketika jadwal tersedia dan diinformasikan | langsung mengerjakan ujian sesuai jadwal |

**Keterangan**: Konten slider/slide dikelola via **CMS Slider** (lihat PRD CMS Landing Page). Panel Berita, Tautan, dan Jurnal mengambil data dari CMS Berita, CMS Tautan, dan data Jurnal yang sudah published. **Panel Pengumuman** (untuk Peserta yang login) menampilkan pengumuman hasil seleksi administrasi, info jadwal Ujikom, dan tombol Mulai Ujikom jika jadwal tersedia — lihat PRD Assignment.

---

## 3. Antarmuka Pengguna (UI) – Beranda

### 3.1 Slider Besar (Hero Slider)

**Deskripsi**: Area slider/carousel besar di bagian paling atas beranda menampilkan beberapa slide (gambar + teks/CTA) yang berganti otomatis atau dapat diarahkan oleh pengguna.

**Elemen yang Harus Ada**:
- **Slide**: Setiap slide menampilkan **gambar/visual** (wajib) dan opsional **judul**, **subjudul/deskripsi**, serta **tautan/CTA** (tombol atau link). Konten slide dikelola di **CMS Slider** (List, Detail, Create, Edit, Delete).
- **Navigasi**: Indikator posisi (titik atau nomor) dan/atau tombol prev/next agar pengguna tahu ada beberapa slide dan bisa pindah slide.
- **Perilaku**: Auto-play opsional dengan interval wajar (mis. 5–7 detik); pause on hover/focus; aksesibel (keyboard, screen reader).
- **Urutan tampil**: Hanya slide yang **published** dan (jika ada) dalam periode tampil (tanggal mulai–selesai) ditampilkan; urutan sesuai pengaturan di CMS (mis. field urutan).

**Tata Letak dan Keterbacaan**:
- Slider memakai lebar penuh (full-width) atau lebar konten dengan tinggi proporsional (hero size). Gambar responsif agar tidak pecah di mobile.
- Teks di atas gambar harus terbaca (kontras cukup, ukuran font memadai). Rekomendasi: maksimal 3–5 slide agar tidak membebani performa dan perhatian pengguna.
- Di layar kecil (mobile), pertimbangkan satu slide utama atau stack konten agar tetap nyaman dipakai.

**Umpan Balik**:
- Saat gambar slide dimuat: placeholder atau skeleton agar tidak layout shift (CLS). Gambar dioptimasi (format dan ukuran) untuk performa (LCP).

---

### 3.2 Panel Berita

**Deskripsi**: Blok di beranda yang menampilkan beberapa berita terbaru yang **sudah dipublikasikan** (sumber: CMS Berita). Biasanya berisi judul, ringkasan/snippet, tanggal, dan link “Selengkapnya” atau klik ke halaman Detail Berita.

**Elemen yang Harus Ada**:
- **Sumber data**: Hanya berita dengan status **published** dari CMS Berita; bisa dibatasi jumlah (mis. 3–5 terbaru) dan diurutkan berdasarkan tanggal publish.
- **Tampilan**: Judul berita, tanggal (opsional), snippet/ringkasan (opsional), link ke Detail Berita.
- **Judul panel**: Mis. “Berita Terbaru” atau “Berita”.

**Tata Letak**: Rapi, konsisten dengan design system; di bawah slider. Responsif (grid/list menyesuaikan lebar layar).

---

### 3.3 Panel Tautan

**Deskripsi**: Blok yang menampilkan daftar **tautan** (judul + URL) dari CMS Tautan. Pengunjung dapat mengklik untuk membuka link (internal atau eksternal).

**Elemen yang Harus Ada**:
- **Sumber data**: Tautan yang aktif/dipublikasikan dari CMS Tautan; bisa dibatasi jumlah dan diurutkan (mis. urutan atau tanggal).
- **Tampilan**: Judul/label tautan, opsional ikon atau thumbnail; link membuka di tab baru jika eksternal (sesuai kebijakan).
- **Judul panel**: Mis. “Tautan Penting” atau “Tautan”.

**Tata Letak**: Rapi, konsisten; di bawah Panel Berita atau berdampingan sesuai layout. Responsif.

---

### 3.4 Panel Jurnal

**Deskripsi**: Blok yang menampilkan jurnal yang **sudah dipublikasikan** saja. Relasi ke modul WPJurnal (PRD WPJurnal ongoing); untuk beranda hanya menampilkan jurnal yang statusnya published.

**Elemen yang Harus Ada**:
- **Sumber data**: Hanya jurnal dengan status **published**; bisa dibatasi jumlah (mis. 3–5 terbaru) dan diurutkan berdasarkan tanggal publish.
- **Tampilan**: Judul jurnal, penulis/ringkasan (jika ada), tanggal, link ke Detail Jurnal.
- **Judul panel**: Mis. “Jurnal Terbaru” atau “Jurnal”.

**Tata Letak**: Rapi, konsisten; responsif. Jika belum ada data jurnal, panel dapat disembunyikan atau menampilkan pesan “Belum ada jurnal.”

---

## 4. Urutan dan Layout Beranda (Rekomendasi)

1. **Slider besar (hero)** – paling atas, full-width.
2. **Panel Berita** – di bawah slider.
3. **Panel Tautan** – di bawah atau berdampingan dengan Panel Berita (sesuai design).
4. **Panel Jurnal** – di bawah Panel Berita/Tautan.
5. **Panel Pengumuman** (untuk Peserta yang login) – menampilkan: Pengumuman hasil seleksi administrasi; Info jadwal Ujikom; Tombol Mulai Ujikom (jika jadwal tersedia).

Urutan dan tampilan (grid, jumlah item) dapat disesuaikan dengan design system dan kebutuhan produk.

---

## 5. Acceptance Criteria (Ringkas)

**Akses publik & header:**
- Beranda (`/`), Berita (`/berita`, `/berita/[slug]`), dan Jurnal (`/jurnal`, `/jurnal/[id]`) **dapat diakses tanpa login**.
- Header menampilkan tombol **Login** jika pengguna belum login; tombol **Dashboard** jika sudah login.

**Slider:** (1) Hanya slide dengan status Published yang tampil; jika ada tanggal mulai/selesai, hanya dalam periode tersebut. (2) Urutan sesuai pengaturan di CMS. (3) Setiap slide menampilkan gambar (wajib), plus opsional judul, subjudul, tautan/CTA. (4) Ada navigasi (prev/next atau indikator). (5) Auto-play opsional; pause on hover. (6) Responsif; tidak layout shift saat gambar dimuat.

**Panel Berita:** (1) Hanya berita status Published; jumlah dibatasi (misalnya 3–5 terbaru). (2) Tampil judul, tanggal (opsional), snippet (opsional), link ke Detail Berita. (3) Judul panel jelas (misalnya "Berita Terbaru"). (4) Jika belum ada berita published: panel kosong atau pesan "Belum ada berita."

**Panel Tautan:** (1) Hanya tautan yang aktif/dipublikasikan; jumlah dan urutan sesuai CMS. (2) Tampil judul/label; klik membuka URL (eksternal dapat di tab baru sesuai pengaturan). (3) Judul panel jelas (misalnya "Tautan Penting").

**Panel Jurnal:** (1) Hanya jurnal status Published; jumlah dibatasi (misalnya 3–5 terbaru). (2) Tampil judul, penulis/ringkasan (jika ada), tanggal, link ke Detail Jurnal. (3) Jika belum ada jurnal: panel disembunyikan atau pesan "Belum ada jurnal."

---

## 6. Cakupan Fitur Beranda

### 6.1 Termasuk
- **Slider besar:** Tampilan slide di beranda; konten slide (gambar, judul, deskripsi, link, urutan) dikelola via **CMS Slider** (PRD CMS Landing Page).
- **Panel Berita:** Menampilkan berita published dari CMS Berita; link ke List/Detail Berita.
- **Panel Tautan:** Menampilkan tautan dari CMS Tautan; link ke URL yang dikonfigurasi.
- **Panel Jurnal:** Menampilkan jurnal published saja; link ke Detail Jurnal (relasi ke WPJurnal).
- **Panel Pengumuman** (untuk Peserta): Pengumuman hasil seleksi administrasi (jika tidak lolos ada informasi); Info menunggu jadwal Ujikom; Tombol **Mulai Ujikom** ketika jadwal tersedia.

### 6.2 Tidak Termasuk
- CMS Slider, CMS Berita, CMS Tautan (CRUD) → PRD CMS Landing Page.
- Halaman List/Detail Berita, List/Detail Jurnal → PRD Berita, PRD Jurnal.
- Manajemen Jurnal (submit, verifikasi) → PRD WPJurnal.

---

## 7. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-11 | Dokumen awal: Slider besar, Panel Berita, Panel Tautan, Panel Jurnal | - |
| 1.1 | 2025-02-28 | Panel Pengumuman untuk Peserta: hasil seleksi administrasi, info jadwal Ujikom, tombol Mulai Ujikom | - |
| 1.2 | 2025-03-01 | Akses publik beranda tanpa login; tombol Login/Dashboard di header sesuai status auth | - |

---

**Catatan**: Konten **slider/slide** dikelola melalui **CMS Slider** (lihat PRD CMS Landing Page). Beranda hanya menampilkan slide yang published dan memenuhi kriteria tampil (urutan, periode jika ada).
