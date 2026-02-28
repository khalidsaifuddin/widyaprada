## SDD – Auth: Login

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Login  

Dokumen ini menjelaskan **desain teknis backend** untuk mengimplementasikan PRD `[PRD] Fitur Login` dengan stack:

- Backend: **Go (Golang)** dengan **Clean Architecture**, PostgreSQL, Redis.
- Komunikasi: **REST + JSON**.

---

## 1. Arsitektur & Konteks (Backend)

- **Pattern**: Clean Architecture.
  - `domain`: entity user, session/token, error domain.
  - `usecase`: `AuthUsecase` dengan operasi `Login` dan `GetCurrentUser`.
  - `delivery/http`: handler `POST /api/v1/auth/login`, middleware auth.
  - `infrastructure`:
    - Repository `UserRepository` (PostgreSQL).
    - Service hash password (bcrypt/argon2).
    - Service token/session (JWT + Redis).

---

## 2. Kontrak API

- **Endpoint**: `POST /api/v1/auth/login`
- **Request (JSON)**:

```json
{
  "identifier": "string",
  "password": "string"
}
```

- **Response Sukses (200)**:

```json
{
  "access_token": "jwt-or-session-token",
  "token_type": "bearer",
  "expires_in": 3600,
  "user": {
    "id": "uuid",
    "name": "string",
    "email": "string",
    "username": "string",
    "roles": [{"id": "uuid", "code": "SUPER_ADMIN", "name": "Super Admin"}],
    "default_home_path": "/dashboard"
  }
}
```

- **Response Gagal (400/401)**: pesan aman; tidak membocorkan email terdaftar atau tidak.

---

## 3. Skema Database

- `users`: id, name, email (unique), username (unique), password_hash, is_active, audit fields.
- `roles`: id, code (unique), name.
- `user_roles`: user_id, role_id (composite unique).
- `sessions` (opsional): id, user_id, expires_at, metadata.

---

## 4. Use Case Login

1. Normalisasi input (trim, lower-case email).
2. Ambil user via email atau username.
3. Cek `is_active`.
4. Verifikasi password (bcrypt/argon2).
5. Cek rate limit (Redis).
6. Ambil roles via `user_roles`.
7. Hitung `default_home_path`.
8. Buat token (JWT/session).
9. Return LoginResult.

---

## 5. Keamanan

- HTTPS wajib. Password hash bcrypt/argon2. Error message tidak membocorkan info.
- Rate limiting: Redis `login_attempts:{userOrIp}`, threshold (mis. 5/15 menit).
- JWT: secret kuat, exp pendek (mis. 1 jam).

---

**Detail lengkap**: lihat `docs/SDD_Auth_Login.md`.
