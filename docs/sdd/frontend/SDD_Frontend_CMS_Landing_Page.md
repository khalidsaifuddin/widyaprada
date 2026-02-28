## SDD Frontend – CMS Landing Page (Slider, Berita, Tautan)

**Aplikasi**: Widyaprada  
**Modul**: CMS Landing Page  
**Fitur**: CMS Slider, CMS Berita, CMS Tautan (masing-masing full CRUD)  
**PRD Terkait**: [PRD_CMS_Landing_Page](../../prd/PRD_CMS_Landing_Page.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk CMS Landing Page dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/cms/slider`, `/cms/berita`, `/cms/tautan` (masing-masing: list, create, [id], [id]/edit)
- **Role**: Admin Satker, Super Admin
- **Layout**: DashboardLayout
- **API**: `/api/v1/cms/slides`, `/api/v1/cms/berita`, `/api/v1/cms/links`

---

## 2. Atomic Design – Komponen

### 2.1 CMS Slider

#### Molecules
- `ImageUpload` | Upload gambar slide |
| `FormField` | Judul, subjudul, URL, label CTA, urutan, status, tanggal mulai/selesai |

#### Organisms
- `SlideListTable` | Tabel: thumbnail, judul, urutan, status; aksi Detail, Edit, Hapus; tombol Tambah Slide |
- `SlideForm` | Create/Edit: gambar, judul, subjudul, URL, label CTA, urutan, status, periode tampil |
- `SlideDetailCard` | Detail lengkap; tombol Edit, Hapus |
- `SlideDeleteDialog` | Konfirmasi (opsional: alasan) |

### 2.2 CMS Berita

#### Organisms
- `BeritaCMSListTable` | Tabel: judul, tanggal, status; search, filter; aksi Detail, Edit, Hapus |
- `BeritaCMSForm` | Create/Edit: judul, slug, konten (rich text), ringkasan, thumbnail, tanggal publikasi, kategori, penulis |
- `BeritaCMSDetailCard` | Detail lengkap berita |
- `RichTextEditor` | WYSIWYG untuk konten berita |

### 2.3 CMS Tautan

#### Organisms
- `LinkListTable` | Tabel: judul, URL, status; aksi Detail, Edit, Hapus |
- `LinkForm` | Create/Edit: judul, URL, status (aktif/nonaktif) |
- `LinkDetailCard` | Detail tautan |

### 2.4 Navigasi

- Menu CMS: submenu Slider, Berita, Tautan (di bawah Landing Page / CMS)

### 2.5 Pages

| Route | Page |
|-------|------|
| `/cms/slider` | SlideListPage |
| `/cms/slider/create`, `/cms/slider/[id]`, `/cms/slider/[id]/edit` | SlideCreatePage, SlideDetailPage, SlideEditPage |
| `/cms/berita` | BeritaCMSListPage |
| `/cms/berita/create`, `/cms/berita/[id]`, `/cms/berita/[id]/edit` | BeritaCMSCreatePage, BeritaCMSDetailPage, BeritaCMSEditPage |
| `/cms/tautan` | LinkListPage |
| `/cms/tautan/create`, `/cms/tautan/[id]`, `/cms/tautan/[id]/edit` | LinkCreatePage, LinkDetailPage, LinkEditPage |

---

## 3. Integrasi API

| Modul | Method | Endpoint |
|-------|--------|----------|
| Slider | GET, POST, GET/:id, PUT/:id, DELETE/:id | `/api/v1/cms/slides` |
| Berita | GET, POST, GET/:id, PUT/:id, DELETE/:id | `/api/v1/cms/berita` |
| Tautan | GET, POST, GET/:id, PUT/:id, DELETE/:id | `/api/v1/cms/links` |

---

## 4. File Lokasi

```
frontend/src/
├── app/cms/
│   ├── slider/page.tsx, create/page.tsx, [id]/page.tsx, [id]/edit/page.tsx
│   ├── berita/page.tsx, create/page.tsx, [id]/page.tsx, [id]/edit/page.tsx
│   └── tautan/page.tsx, create/page.tsx, [id]/page.tsx, [id]/edit/page.tsx
├── components/organisms/SlideListTable.tsx
├── components/organisms/BeritaCMSListTable.tsx
├── components/organisms/LinkListTable.tsx
└── components/molecules/RichTextEditor.tsx
```
