## SDD – Auth: Lupa Password

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Lupa Password  

Dokumen ini menjelaskan **desain teknis backend** untuk mengimplementasikan PRD `[PRD] Fitur Lupa Password` dengan stack: **Go (Golang)**, Clean Architecture, PostgreSQL, Redis, email service.

---

## 1. Arsitektur & Konteks (Backend)

- **Pattern**: Clean Architecture.
  - `usecase`: `AuthUsecase.RequestPasswordReset(ctx, email)`, `AuthUsecase.ResetPassword(ctx, token, newPassword)`.
  - `delivery/http`:
    - `POST /api/v1/auth/forgot-password` — request reset.
    - `POST /api/v1/auth/reset-password` — set password baru (dengan token).
  - `infrastructure`:
    - UserRepository, PasswordResetTokenRepository (atau Redis untuk token sementara).
    - EmailService (SMTP atau third-party).

---

## 2. Kontrak API

### 2.1 Request Reset

- **Endpoint**: `POST /api/v1/auth/forgot-password`
- **Request**:

```json
{
  "email": "string"
}
```

- **Response Sukses (200)**: Pesan generik (tidak mengungkap apakah email terdaftar):

```json
{
  "message": "Jika email Anda terdaftar, Anda akan menerima tautan untuk mengatur ulang kata sandi. Periksa juga folder spam."
}
```

### 2.2 Reset Password

- **Endpoint**: `POST /api/v1/auth/reset-password`
- **Request**:

```json
{
  "token": "string",
  "password": "string",
  "password_confirm": "string"
}
```

- **Response Sukses (200)**:

```json
{
  "message": "Kata sandi berhasil diubah. Silakan masuk dengan kata sandi baru."
}
```

- **Response Gagal (400)**: token tidak valid/kedaluwarsa, password tidak memenuhi kebijakan, konfirmasi tidak cocok.

---

## 3. Skema Database

- `password_reset_tokens` (atau Redis):
  - `id`, `user_id`, `token` (unique), `expires_at`, `created_at`.
- Token: random 32+ bytes, hash sebelum simpan; TTL mis. 1 jam.

---

## 4. Alur Use Case

### Request Reset

1. Validasi email (format, required).
2. Cari user by email (jika tidak ada, tetap return sukses — jangan bocorkan).
3. Generate token unik, simpan ke Redis/DB dengan TTL 1 jam.
4. Kirim email dengan link: `{frontend_url}/auth/reset-password?token={token}`.
5. Return pesan generik.

### Reset Password

1. Validasi token (ada, belum expired).
2. Validasi password (panjang, kompleksitas) dan konfirmasi cocok.
3. Hash password baru.
4. Update `users.password_hash`.
5. Invalidasi token (hapus dari Redis/DB).
6. Return sukses.

---

## 5. Keamanan

- Token single-use; invalidasi setelah reset.
- TTL token 1 jam.
- Pesan response tidak mengungkap apakah email terdaftar.
- Rate limiting pada `/forgot-password` (per IP/email) untuk mencegah abuse.
