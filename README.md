# Aplikasi Widyaprada

Aplikasi layanan untuk **Widyaprada** — PNS dengan jabatan fungsional penjaminan mutu pendidikan (PAUD, dasar, menengah, masyarakat). Meliputi manajemen data WP, jurnal, uji kompetensi (CBT), CMS landing page, dan fitur pendukung lainnya.

## Tech Stack

| Layer   | Teknologi              |
|---------|------------------------|
| Backend | Go, Clean Architecture |
| Frontend| Next.js, React, Tailwind CSS, shadcn/ui |
| Database| PostgreSQL             |
| Cache   | Redis                  |
| API     | REST + JSON            |

## Struktur Proyek

```
ProjectWidyaprada/
├── backend/        # API Go (Gin, GORM)
├── frontend/       # Aplikasi web Next.js
├── docs/
│   ├── prd/        # Product Requirements Documents
│   ├── sdd/        # Software Design Documents (backend & frontend)
│   └── workflows/  # Alur kerja
└── README.md
```

## Menjalankan Proyek

### Backend

```bash
cd backend
go mod download
go run main.go
```

### Frontend

```bash
cd frontend
npm install
npm run dev
```

## Dokumentasi

- **PRD** — Kebutuhan produk, user story, UI/UX: [`docs/prd/`](docs/prd/)
- **SDD** — Desain teknis backend & frontend: [`docs/sdd/`](docs/sdd/)
