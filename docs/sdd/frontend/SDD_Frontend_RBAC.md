## SDD Frontend – RBAC (Manajemen Role & Permission)

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Manajemen Role, Manajemen Permission, Kelola Role (Role–Permission)  
**PRD Terkait**: [PRD_RBAC](../../prd/PRD_RBAC.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk fitur RBAC dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/auth/role`, `/auth/permission`, `/auth/role/[id]`, `/auth/permission/[id]`
- **Role**: Hanya Super Admin
- **Layout**: DashboardLayout
- **API**: `/api/v1/roles`, `/api/v1/permissions`

---

## 2. Atomic Design – Komponen

### 2.1 Manajemen Role

#### Atoms
- `Button`, `Input`, `Label`, `Badge`

#### Molecules
- `FormField`, `SearchBar`, `FilterDropdown`, `DeleteConfirmDialog`

#### Organisms
- `RoleListTable`: Kolom kode, nama, jumlah permission; search, filter; aksi Detail, Edit, Hapus
- `RoleDetailCard`: Detail role + daftar permission yang di-assign; tombol Edit, Hapus
- `RoleForm`: Form Create/Edit (kode, nama)
- `RolePermissionManager`: Di Detail/Edit role; checkbox/list untuk pilih/batalkan permission per role
- `RoleDeleteDialog`: Konfirmasi + alasan penghapusan wajib

### 2.2 Manajemen Permission

#### Organisms
- `PermissionListTable`: Kolom kode, nama, modul/group; search; aksi Detail, Edit, Hapus
- `PermissionDetailCard`: Detail permission
- `PermissionForm`: Form Create/Edit (kode, nama, deskripsi, modul)
- `PermissionDeleteDialog`: Konfirmasi + alasan

### 2.3 Navigasi

- Menu: "Kelola Role" dan "Kelola Permission" di area yang sama dengan Manajemen Pengguna (submenu Auth)

### 2.4 Pages

| Route | Page |
|-------|------|
| `/auth/role` | RoleListPage |
| `/auth/role/create` | RoleCreatePage |
| `/auth/role/[id]` | RoleDetailPage (dengan RolePermissionManager) |
| `/auth/role/[id]/edit` | RoleEditPage |
| `/auth/permission` | PermissionListPage |
| `/auth/permission/create` | PermissionCreatePage |
| `/auth/permission/[id]` | PermissionDetailPage |
| `/auth/permission/[id]/edit` | PermissionEditPage |

---

## 3. Integrasi API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/roles` | List role |
| GET | `/api/v1/roles/:id` | Detail + permissions |
| POST | `/api/v1/roles` | Create |
| PUT | `/api/v1/roles/:id` | Edit (termasuk assign permission) |
| DELETE | `/api/v1/roles/:id` | Delete (+ reason) |
| GET | `/api/v1/permissions` | List permission |
| GET | `/api/v1/permissions/:id` | Detail |
| POST | `/api/v1/permissions` | Create |
| PUT | `/api/v1/permissions/:id` | Edit |
| DELETE | `/api/v1/permissions/:id` | Delete (+ reason) |

---

## 4. File Lokasi

```
frontend/src/
├── app/auth/role/
│   ├── page.tsx, create/page.tsx
│   └── [id]/page.tsx, [id]/edit/page.tsx
├── app/auth/permission/
│   ├── page.tsx, create/page.tsx
│   └── [id]/page.tsx, [id]/edit/page.tsx
├── components/organisms/RoleListTable.tsx
├── components/organisms/RoleDetailCard.tsx
├── components/organisms/RolePermissionManager.tsx
├── components/organisms/PermissionListTable.tsx
└── components/organisms/PermissionForm.tsx
```
