## SDD – Auth: Manajemen Pengguna

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Manajemen Pengguna (List, Detail, Create, Edit, Delete)  

Dokumen ini menjelaskan **desain teknis backend** untuk mengimplementasikan PRD `[PRD] Fitur Manajemen Pengguna` dengan stack: **Go (Golang)**, Clean Architecture, PostgreSQL.

---

## 1. Arsitektur & Konteks (Backend)

- **Pattern**: Clean Architecture.
  - `usecase`: `UserUsecase` — List, Get, Create, Update, Delete.
  - `delivery/http`: REST endpoints di `/api/v1/users`.
  - `infrastructure`: `UserRepository`, `UserRoleRepository`, `WidyapradaDataRepository` (untuk auto-create data WP saat role Widyaprada).

---

## 2. Kontrak API

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/v1/users` | List users (search, filter role/satker/status, pagination) |
| GET | `/api/v1/users/:id` | Detail user |
| POST | `/api/v1/users` | Create user |
| PUT | `/api/v1/users/:id` | Update user |
| DELETE | `/api/v1/users/:id` | Delete user (body: `{"reason":"..."}`) |

**Query params list**: `q`, `role_id`, `satker_id`, `status`, `page`, `page_size`, `sort_by`, `sort_order`.

**Create/Update body**: `name`, `email`, `username`, `password` (create only), `role_ids` (array), `satker_id`, `is_active`, dll.

---

## 3. Skema Database

- `users`: id, name, email, username, password_hash, satker_id, is_active, created_at, updated_at, deleted_at, deleted_reason.
- `user_roles`: user_id, role_id.
- Relasi ke `widyaprada_data` (jika role Widyaprada): auto-create saat Create user dengan role Widyaprada.

---

## 4. Aturan Bisnis

- **Wewenang**:
  - Admin Satker: hanya users dalam satker sendiri.
  - Super Admin: semua users.
- **Create**:
  - Email, username unik.
  - Minimal 1 role.
  - Jika role_ids mengandung Widyaprada → wajib create entitas `widyaprada_data` (profil WP) untuk user.
- **Delete**:
  - Body wajib berisi `reason`.
  - Super Admin tidak boleh hapus diri sendiri.
  - Soft delete disarankan; simpan `deleted_reason`.

---

## 5. RBAC

- Permission: `USER_READ`, `USER_CREATE`, `USER_UPDATE`, `USER_DELETE`.
- Scope: filter by satker_id untuk Admin Satker.
