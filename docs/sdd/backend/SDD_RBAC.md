## SDD – Auth & RBAC: Role, Permission, Kelola Role

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Manajemen Role, Permission, Kelola Role (Role–Permission)  

Dokumen ini menjelaskan **desain teknis backend** untuk mengimplementasikan PRD `[PRD] Fitur RBAC` dengan stack: **Go (Golang)**, Clean Architecture, PostgreSQL.

---

## 1. Arsitektur & Konteks (Backend)

- **Pattern**: Clean Architecture.
  - `usecase`: `RoleUsecase`, `PermissionUsecase`.
  - `delivery/http`: REST di `/api/v1/rbac/roles`, `/api/v1/rbac/permissions`.
  - `infrastructure`: `RoleRepository`, `PermissionRepository`, `RolePermissionRepository`.

---

## 2. Kontrak API

### Role

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/rbac/roles` | List roles (q, page, page_size) |
| GET | `/api/v1/rbac/roles/:id` | Detail role + permissions |
| POST | `/api/v1/rbac/roles` | Create role |
| PUT | `/api/v1/rbac/roles/:id` | Update role + state permission_ids |
| DELETE | `/api/v1/rbac/roles/:id` | Delete role (body: reason) |

### Permission

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/rbac/permissions` | List permissions (q, group filter) |
| GET | `/api/v1/rbac/permissions/:id` | Detail permission |
| POST | `/api/v1/rbac/permissions` | Create permission |
| PUT | `/api/v1/rbac/permissions/:id` | Update permission |
| DELETE | `/api/v1/rbac/permissions/:id` | Delete permission (body: reason) |

---

## 3. Skema Database

- `roles`: id, code (unique), name, description, created_at, updated_at, deleted_at, deleted_reason.
- `permissions`: id, code (unique), name, group, description, created_at, updated_at, deleted_at, deleted_reason.
- `role_permissions`: role_id, permission_id (composite PK).

---

## 4. Aturan Bisnis

- **Delete role**: Cek `user_roles` — jika ada user memakai role → error `ErrRoleInUse` ("Role ini masih digunakan oleh N pengguna...").
- **Delete permission**: Cek `role_permissions` — jika masih di-assign → error `ErrPermissionInUse`.
- **Update role permissions**: State final — kirim list permission_ids; backend sync (insert baru, delete yang hilang).

---

## 5. RBAC

- Hanya **Super Admin** yang boleh akses endpoint RBAC.
- Permission: `ROLE_READ`, `ROLE_CREATE`, `ROLE_UPDATE`, `ROLE_DELETE`; `PERMISSION_READ`, `PERMISSION_CREATE`, dll.

---

**Detail lengkap**: lihat `docs/SDD_RBAC.md`.
