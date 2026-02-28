# Software Design Document (SDD) – Frontend Aplikasi Widyaprada

**SDD Frontend** mendokumentasikan **desain teknis komponen UI** untuk mengimplementasikan fitur-fitur yang didefinisikan di PRD. Dokumen ini memakai pendekatan **Atomic Design** (Atoms → Molecules → Organisms → Templates → Pages) dengan stack **React** (Next.js), **Tailwind CSS**, dan **shadcn/ui**.

---

## Pendekatan Atomic Design

| Level | Deskripsi | Contoh |
|-------|-----------|--------|
| **Atoms** | Elemen UI terkecil yang tidak bisa dipecah lagi | Button, Input, Label, Badge, Icon |
| **Molecules** | Gabungan beberapa atom untuk satu fungsi | FormField (Label+Input), SearchBar, CardHeader |
| **Organisms** | Komponen kompleks gabungan molecules | LoginForm, DataTable, AssignmentCard |
| **Templates** | Layout halaman (kerangka struktur) | AuthLayout, DashboardLayout |
| **Pages** | Halaman lengkap = Template + Organisms | LoginPage, DashboardPage |

---

## Stack Frontend

| Aspek | Teknologi |
|-------|-----------|
| Framework | Next.js (App Router) |
| UI Library | React |
| Styling | Tailwind CSS |
| Komponen | shadcn/ui (berbasis Radix UI) |
| State | React state, Context, (opsional: React Query/TanStack) |
| Routing | Next.js App Router |
| API | REST + JSON (fetch/axios ke backend Go) |

Lihat [SDD_Tech_Stack.md](../SDD_Tech_Stack.md) untuk detail keputusan teknis.

---

## Daftar SDD Frontend

| SDD Frontend | PRD Terkait | Modul |
|--------------|-------------|-------|
| [SDD_Frontend_Auth_Login](SDD_Frontend_Auth_Login.md) | PRD_Auth_Login | Auth & RBAC |
| [SDD_Frontend_Auth_Registrasi](SDD_Frontend_Auth_Registrasi.md) | PRD_Auth_Registrasi | Auth & RBAC |
| [SDD_Frontend_Auth_Logout](SDD_Frontend_Auth_Logout.md) | PRD_Auth_Logout | Auth & RBAC |
| [SDD_Frontend_Auth_Lupa_Password](SDD_Frontend_Auth_Lupa_Password.md) | PRD_Auth_Lupa_Password | Auth & RBAC |
| [SDD_Frontend_Auth_Manajemen_Pengguna](SDD_Frontend_Auth_Manajemen_Pengguna.md) | PRD_Auth_Manajemen_Pengguna | Auth & RBAC |
| [SDD_Frontend_RBAC](SDD_Frontend_RBAC.md) | PRD_RBAC | Auth & RBAC |
| [SDD_Frontend_Dashboard_Widyaprada](SDD_Frontend_Dashboard_Widyaprada.md) | PRD_Dashboard_Widyaprada | Dashboard |
| [SDD_Frontend_Beranda](SDD_Frontend_Beranda.md) | PRD_Beranda | Landing Page |
| [SDD_Frontend_Berita](SDD_Frontend_Berita.md) | PRD_Berita | Landing Page |
| [SDD_Frontend_Jurnal](SDD_Frontend_Jurnal.md) | PRD_Jurnal | Landing Page |
| [SDD_Frontend_CMS_Landing_Page](SDD_Frontend_CMS_Landing_Page.md) | PRD_CMS_Landing_Page | CMS |
| [SDD_Frontend_Bank_Soal](SDD_Frontend_Bank_Soal.md) | PRD_Bank_Soal | WPUjikom |
| [SDD_Frontend_Paket_Soal](SDD_Frontend_Paket_Soal.md) | PRD_Paket_Soal | WPUjikom |
| [SDD_Frontend_CBT](SDD_Frontend_CBT.md) | PRD_CBT | WPUjikom |
| [SDD_Frontend_Assignment](SDD_Frontend_Assignment.md) | PRD_Assignment | WPUjikom |
| [SDD_Frontend_Manajemen_Uji_Kompetensi](SDD_Frontend_Manajemen_Uji_Kompetensi.md) | PRD_Manajemen_Uji_Kompetensi | WPUjikom |
| [SDD_Frontend_Manajemen_Data_WP](SDD_Frontend_Manajemen_Data_WP.md) | PRD_Manajemen_Data_WP | Manajemen Data WP |

---

## Struktur Folder Komponen (Referensi)

```
frontend/src/
├── components/
│   ├── atoms/          # Button, Input, Label, Badge, Icon, Skeleton, ...
│   ├── molecules/      # FormField, SearchBar, CardHeader, EmptyState, ...
│   ├── organisms/      # LoginForm, Sidebar, DataTable, AssignmentCard, ...
│   └── templates/      # AuthLayout, DashboardLayout, PublicLayout
├── app/
│   ├── auth/           # login, register, forgot-password, reset-password
│   ├── dashboard/
│   ├── beranda/
│   └── ...
├── lib/                # API client, auth, hooks, utils
└── config/
```

---

## Dokumen Terkait

- [SDD Tech Stack](../SDD_Tech_Stack.md)
- [SDD Backend](../backend/README.md)
- [PRD Collection](../prd/PRD%20Collection%20Skeleton.md)
