# Frontend – Aplikasi Widyaprada

Skeleton frontend Next.js (App Router) mengacu pada contoh dashboard-sr.  
Digunakan untuk modul: Auth & RBAC, Landing, WPData, WPJurnal, WPUjikom (Bank Soal, Paket Soal, Uji Kompetensi, CBT, Assignment).

## Setup

```bash
cd frontend
npm install
cp .env.example .env   # sesuaikan NEXT_PUBLIC_API_BASE_URL jika perlu
npm run dev
```

Buka http://localhost:3000. Halaman login: `/auth/login`.

## Struktur (yang dipertahankan)

- **App Router**: `src/app/layout.tsx`, `page.tsx`, `auth/login`, `dashboard`, `wpdata`, `wpjurnal`, `wpujikom/*`
- **Auth**: middleware (cookie `auth_token`), AuthWrapper, lib/auth (IndexedDB + cookie), halaman login
- **Config**: `src/config/index.ts`, `navigation.json` (menu sidebar)
- **Lib**: `api.ts`, `auth.ts`, `utils.ts`, `navigation.ts`, `security.ts`
- **Komponen (Atomic Design)**: `atoms/`, `molecules/`, `organisms/` (Sidebar, AuthWrapper, Card, ConfirmDialog, ErrorDialog), `templates/`

## Yang tidak dipakai (dibuang dari contoh)

- Feature flags, under-construction, test-auth
- Progress wilayah/sekolah/nasional, rekap, form kesehatan siswa
- Data pokok (sekolah, guru, siswa)
- Quick links, FloatingBubbles, VersionDisplay
- Chart (PieChart, GaugeChart, BarChart, ProvinceProgressTabs)
- Captcha dan slider di halaman login (skeleton login sederhana)

## Environment

- `NEXT_PUBLIC_API_BASE_URL` – URL API backend (default: http://localhost:8080/api)

Logo: ganti `public/logo-widyaprada.svg` dengan aset resmi jika perlu.
