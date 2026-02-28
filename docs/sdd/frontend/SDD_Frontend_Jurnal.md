## SDD Frontend – Jurnal (List & Detail Publik)

**Aplikasi**: Widyaprada  
**Modul**: Landing Page  
**Fitur**: Jurnal (List, Detail) – Tampilan publik  
**PRD Terkait**: [PRD_Jurnal](../../prd/PRD_Jurnal.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk halaman List dan Detail Jurnal publik dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/jurnal` (list), `/jurnal/[slug]` atau `/jurnal/[id]` (detail)
- **Akses**: Publik (hanya jurnal status Published)
- **Layout**: PublicLayout
- **API**: `GET /api/v1/jurnal` (filter: published), `GET /api/v1/jurnal/:id` atau `/:slug`

---

## 2. Atomic Design – Komponen

### 2.1 Halaman List Jurnal

#### Atoms
- `Input`, `Button`, `Image`

#### Molecules
- `SearchBar` | Pencarian judul, penulis, kata kunci |
| `JournalCard` | Judul, penulis, tanggal/tahun, abstrak/snippet; link ke detail |
| `EmptyState` | "Tidak ada jurnal yang sesuai." / "Belum ada jurnal." |

#### Organisms
- `JurnalList` | Judul halaman; SearchBar; Filter (tahun, kategori); Sort (Terbaru/Terlama); grid/list JournalCard; paginasi atau load more |

### 2.2 Halaman Detail Jurnal

#### Organisms
- `JurnalDetail` | Judul; meta (penulis, tanggal publikasi, DOI/ISSN); abstrak; konten lengkap (read-only) |

### 2.3 Pages

| Route | Page |
|-------|------|
| `/jurnal` | JurnalListPage |
| `/jurnal/[slug]` atau `/jurnal/[id]` | JurnalDetailPage |

---

## 3. State & Fetch

- List: search, filter (tahun, kategori), sort (default: terbaru), paginasi
- Fetch: `GET /api/v1/jurnal?status=published&search=&page=&limit=`
- Detail: `GET /api/v1/jurnal/:id` atau `/:slug`

---

## 4. File Lokasi

```
frontend/src/
├── app/jurnal/
│   ├── page.tsx
│   └── [slug]/page.tsx
├── components/organisms/JurnalList.tsx
└── components/organisms/JurnalDetail.tsx
```
