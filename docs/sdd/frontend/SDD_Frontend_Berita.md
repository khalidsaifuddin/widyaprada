## SDD Frontend – Berita (List & Detail Publik)

**Aplikasi**: Widyaprada  
**Modul**: Landing Page  
**Fitur**: Berita (List, Detail) – Tampilan publik  
**PRD Terkait**: [PRD_Berita](../../prd/PRD_Berita.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk halaman List dan Detail Berita publik dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/berita` (list), `/berita/[slug]` (detail)
- **Akses**: Publik (hanya berita status Published)
- **Layout**: PublicLayout
- **API**: `GET /api/v1/berita` (filter: published), `GET /api/v1/berita/:slug`

---

## 2. Atomic Design – Komponen

### 2.1 Halaman List Berita

#### Atoms
- `Input` (search), `Button`, `Image`

#### Molecules
- `SearchBar` | Pencarian judul/kata kunci |
| `NewsCard` | Judul, tanggal publikasi, ringkasan/snippet, thumbnail; link ke detail |
| `EmptyState` | "Tidak ada berita yang sesuai." / "Belum ada berita." |

#### Organisms
- `BeritaList` | Judul halaman; SearchBar; Filter (kategori, tanggal); Sort (Terbaru/Terlama); grid/list NewsCard; paginasi |

### 2.2 Halaman Detail Berita

#### Organisms
- `BeritaDetail` | Judul; meta (tanggal, penulis, kategori); thumbnail; konten (rich text/HTML); responsive |

### 2.3 Pages

| Route | Page |
|-------|------|
| `/berita` | BeritaListPage |
| `/berita/[slug]` | BeritaDetailPage |

---

## 3. State & Fetch

- List: search, filter (kategori, rentang tanggal), sort (default: terbaru), paginasi
- Fetch: `GET /api/v1/berita?status=published&search=&page=&limit=`
- Detail: `GET /api/v1/berita/:slug`

---

## 4. Responsivitas

- Grid: 1 kolom mobile, 2–3 kolom desktop
- Thumbnail responsif; konten readable (max-width, typography)

---

## 5. File Lokasi

```
frontend/src/
├── app/berita/
│   ├── page.tsx
│   └── [slug]/page.tsx
├── components/organisms/BeritaList.tsx
└── components/organisms/BeritaDetail.tsx
```
