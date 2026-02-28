## SDD – CMS Landing Page

**Aplikasi**: Widyaprada  
**Modul**: CMS Landing Page  
**Fitur**: CMS Slider, CMS Berita, CMS Tautan (CRUD)  

Dokumen ini menjelaskan **desain teknis backend** untuk PRD CMS Landing Page dengan stack: Go (Golang), Clean Architecture, PostgreSQL.

---

## 1. Arsitektur dan Konteks

- Usecase: SliderUsecase, BeritaUsecase, TautanUsecase - masing-masing CRUD.
- Delivery: REST di /api/v1/cms/slider, /api/v1/cms/berita, /api/v1/cms/tautan.
- Scope: Admin Satker (satker sendiri), Super Admin (semua).

---

## 2. Kontrak API

### CMS Slider

- GET/POST /api/v1/cms/slider
- GET/PUT/DELETE /api/v1/cms/slider/:id

Body: gambar_url, judul, subjudul, url, label_cta, urutan, status (Draft/Published), tanggal_mulai_tampil, tanggal_selesai_tampil.

### CMS Berita

- GET/POST /api/v1/cms/berita
- GET/PUT/DELETE /api/v1/cms/berita/:id

Body: judul, slug, konten, ringkasan, thumbnail, tanggal_publikasi, status, penulis, kategori.

### CMS Tautan

- GET/POST /api/v1/cms/tautan
- GET/PUT/DELETE /api/v1/cms/tautan/:id

Body: judul, url, deskripsi, urutan, status (Aktif/Nonaktif), buka_di_tab_baru.

---

## 3. Skema Database

- slides: id, image_url, title, subtitle, link_url, cta_label, sort_order, status, date_start, date_end, satker_id, created_at, updated_at.
- articles: id, title, slug, content, excerpt, thumbnail_url, published_at, status, author_id, category_id, satker_id.
- links: id, title, url, description, sort_order, status, open_in_new_tab, satker_id.

---

## 4. Aturan Bisnis

- Admin Satker: filter by satker_id.
- Super Admin: semua data.
- Slug berita unik. Urutan slider/tautan numerik.

---

## 5. RBAC

- Permission: SLIDER_*, BERITA_*, TAUTAN_* (CRUD).
- Scope satker untuk Admin Satker.
