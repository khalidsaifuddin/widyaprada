## SDD Frontend – Auth: Manajemen Pengguna

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Manajemen Pengguna (List, Detail, Create, Edit, Delete)  
**PRD Terkait**: [PRD_Auth_Manajemen_Pengguna](../../prd/PRD_Auth_Manajemen_Pengguna.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk Manajemen Pengguna dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/auth/manajemen-pengguna` (list), `/auth/manajemen-pengguna/create`, `/auth/manajemen-pengguna/[id]`, `/auth/manajemen-pengguna/[id]/edit`
- **Role**: Admin Satker, Super Admin
- **API**: CRUD users via `/api/v1/users`
- **Layout**: DashboardLayout (dengan Sidebar)

---

## 2. Atomic Design – Komponen

### 2.1 Atoms

| Komponen | Deskripsi |
|----------|-----------|
| `Button` | Tambah, Simpan, Batal, Hapus |
| `Input` | Text, email, password |
| `Label` | Label form |
| `Badge` | Status (Aktif/Nonaktif), Role tags |
| `Checkbox` | Multi-select role |

### 2.2 Molecules

| Komponen | Deskripsi |
|----------|-----------|
| `FormField` | Label + Input + error |
| `SearchBar` | Pencarian nama/email/username |
| `FilterDropdown` | Filter role, satker, status |
| `RoleMultiSelect` | Pilih satu atau lebih role |
| `DeleteConfirmDialog` | Dialog konfirmasi + field wajib alasan penghapusan |

### 2.3 Organisms

| Komponen | Deskripsi |
|----------|-----------|
| `UserListTable` | Tabel: Nama, Email/Username, Role(s), Satker, Status, aksi Detail/Edit/Hapus; search, filter, paginasi |
| `UserDetailCard` | Kartu detail satu pengguna; tombol Edit, Hapus |
| `UserForm` | Form Create/Edit: nama, email, username, password (opsional di Edit), multi-select role, satker; validasi |
| `UserDeleteDialog` | Konfirmasi hapus + textarea alasan penghapusan wajib |

### 2.4 Templates

| Komponen | Deskripsi |
|----------|-----------|
| `DashboardLayout` | Sidebar + content area |

### 2.5 Pages

| Route | Page |
|-------|------|
| `/auth/manajemen-pengguna` | UserListPage (UserListTable) |
| `/auth/manajemen-pengguna/create` | UserCreatePage (UserForm) |
| `/auth/manajemen-pengguna/[id]` | UserDetailPage (UserDetailCard) |
| `/auth/manajemen-pengguna/[id]/edit` | UserEditPage (UserForm) |

---

## 3. State & Validasi

### 3.1 UserListPage
- Search, filter (role, satker, status), sort, paginasi
- Fetch: GET `/api/v1/users?search=&role=&satker=&status=&page=&limit=`

### 3.2 UserForm
- Create: nama, email, username, password (wajib), roles (multi-select, wajib minimal 1), satker
- Edit: sama, password kosong = tidak diubah
- Validasi: email unik, format email, password policy

### 3.3 Delete
- Wajib isi alasan/deskripsi di dalam dialog
- Setelah konfirmasi: DELETE `/api/v1/users/:id` dengan body `{ reason }`

---

## 4. Integrasi API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/users` | List (search, filter, paginate) |
| GET | `/api/v1/users/:id` | Detail |
| POST | `/api/v1/users` | Create |
| PUT | `/api/v1/users/:id` | Edit |
| DELETE | `/api/v1/users/:id` | Delete (body: reason) |

---

## 5. File Lokasi

```
frontend/src/
├── app/auth/manajemen-pengguna/
│   ├── page.tsx
│   ├── create/page.tsx
│   └── [id]/page.tsx, [id]/edit/page.tsx
├── components/organisms/UserListTable.tsx
├── components/organisms/UserDetailCard.tsx
├── components/organisms/UserForm.tsx
└── components/organisms/UserDeleteDialog.tsx
```
