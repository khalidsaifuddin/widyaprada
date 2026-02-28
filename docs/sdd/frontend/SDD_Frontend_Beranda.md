## SDD Frontend – Beranda (Landing Page)

**Aplikasi**: Widyaprada  
**Modul**: Landing Page  
**Fitur**: Beranda – Slider, Panel Berita, Panel Tautan, Panel Jurnal, Panel Pengumuman  
**PRD Terkait**: [PRD_Beranda](../../prd/PRD_Beranda.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk Beranda dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/` atau `/beranda`
- **Akses**: Publik (pengunjung) + Panel Pengumuman hanya untuk Peserta yang login
- **Layout**: PublicLayout (header/footer publik, tanpa sidebar dashboard)
- **API**: Slider, Berita, Tautan, Jurnal published; Pengumuman peserta

---

## 2. Atomic Design – Komponen

### 2.1 Slider (Hero)

#### Atoms
- `Image` (optimized), `Button` (CTA)

#### Molecules
- `SlideItem` | Satu slide: gambar, judul, subjudul, tautan/CTA |

#### Organisms
- `HeroSlider` | Carousel: navigasi prev/next, indikator titik; auto-play opsional; pause on hover; hanya slide Published, dalam periode tampil |
- Skeleton saat gambar dimuat (minim CLS)

### 2.2 Panel Berita

#### Molecules
- `NewsCard` | Judul, tanggal, snippet, link "Selengkapnya" |

#### Organisms
- `NewsPanel` | Judul "Berita Terbaru" + grid/list NewsCard (3–5 item) + EmptyState "Belum ada berita" |

### 2.3 Panel Tautan

#### Molecules
- `LinkItem` | Judul/label, ikon opsional, link (eksternal: tab baru) |

#### Organisms
- `LinksPanel` | Judul "Tautan Penting" + daftar LinkItem |

### 2.4 Panel Jurnal

#### Molecules
- `JournalCard` | Judul, penulis, tanggal, abstrak singkat, link ke Detail |

#### Organisms
- `JournalPanel` | Judul "Jurnal Terbaru" + grid JournalCard (3–5 item) + EmptyState "Belum ada jurnal" |

### 2.5 Panel Pengumuman (Peserta Login)

#### Organisms
- `AnnouncementPanel` | Pengumuman hasil seleksi administrasi; Info jadwal Ujikom; Tombol "Mulai Ujikom" (jika jadwal tersedia); hanya tampil untuk user dengan role Peserta/Widyaprada |

### 2.6 Urutan Layout

1. HeroSlider – full-width
2. NewsPanel
3. LinksPanel (sejajar atau di bawah NewsPanel)
4. JournalPanel
5. AnnouncementPanel (jika login & role Peserta)

### 2.7 Pages

| Route | Page |
|-------|------|
| `/` atau `/beranda` | BerandaPage = PublicLayout + HeroSlider + NewsPanel + LinksPanel + JournalPanel + AnnouncementPanel (conditional) |

---

## 3. State & Fetch

- Slider: `GET /api/v1/landing/slides` (filter: published, periode)
- Berita: `GET /api/v1/berita?status=published&limit=5`
- Tautan: `GET /api/v1/links?active=true`
- Jurnal: `GET /api/v1/jurnal?status=published&limit=5`
- Pengumuman: `GET /api/v1/assignment/announcement` (auth required)

---

## 4. Responsivitas & Performa

- Gambar slider: lazy load, format WebP/optimized; skeleton untuk LCP
- Grid responsif: 1 kolom mobile, 2–3 kolom desktop
- Auto-play slider: interval 5–7 detik; pause on hover; keyboard accessible

---

## 5. File Lokasi

```
frontend/src/
├── app/page.tsx                     # atau app/beranda/page.tsx
├── components/organisms/HeroSlider.tsx
├── components/organisms/NewsPanel.tsx
├── components/organisms/LinksPanel.tsx
├── components/organisms/JournalPanel.tsx
├── components/organisms/AnnouncementPanel.tsx
└── components/templates/PublicLayout.tsx
```
