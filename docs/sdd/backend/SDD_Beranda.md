## SDD - Beranda (Landing Page)

**Aplikasi**: Widyaprada  
**Modul**: Landing Page  
**Fitur**: Slider, Panel Berita, Panel Tautan, Panel Jurnal  

Stack: Go (Golang), Clean Architecture, PostgreSQL.

---

## 1. Arsitektur

- Beranda = agregasi data dari CMS Slider, CMS Berita, CMS Tautan, Jurnal (published).
- Usecase: BerandaUsecase - GetSlider, GetBeritaTerbaru, GetTautan, GetJurnalTerbaru.
- Delivery: REST /api/v1/landing/home atau endpoint per panel.

---

## 2. Kontrak API

- GET /api/v1/landing/home - Return: slider, berita, tautan, jurnal.
- GET /api/v1/beranda/pengumuman (auth: Peserta) - Return: hasil seleksi administrasi, info jadwal Ujikom, can_start_ujikom (boolean).
- Atau GET /api/v1/landing/slider, /berita, /tautan, /jurnal dengan query limit.

---

## 3. Sumber Data

- Slider: slides (CMS), status Published.
- Berita: articles (CMS), status Published.
- Tautan: links (CMS), status Aktif.
- Jurnal: journals (WPJurnal), status Published.

---

## 4. RBAC

- Slider, Berita, Tautan, Jurnal: endpoint publik.
- Panel Pengumuman (untuk Peserta yang login): GET /api/v1/beranda/pengumuman — return: hasil seleksi administrasi, info jadwal Ujikom, tombol Mulai Ujikom jika jadwal tersedia. Perlu auth role Peserta.
