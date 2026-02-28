## SDD – Auth: Logout

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Logout  

Dokumen ini menjelaskan **desain teknis backend** untuk mengimplementasikan PRD `[PRD] Fitur Logout` dengan stack: **Go (Golang)**, Clean Architecture, PostgreSQL, Redis.

---

## 1. Arsitektur & Konteks (Backend)

- **Pattern**: Clean Architecture.
  - `usecase`: `AuthUsecase.Logout(ctx, tokenOrSessionId)`.
  - `delivery/http`: handler `POST /api/v1/auth/logout`.
  - `infrastructure`: invalidasi token/session di Redis (atau blacklist JWT).

---

## 2. Kontrak API

- **Endpoint**: `POST /api/v1/auth/logout`
- **Headers**: `Authorization: Bearer <token>` (wajib).
- **Request**: Body kosong atau opsional.
- **Response Sukses (200)**:

```json
{
  "message": "Anda telah keluar."
}
```

- **Response Gagal (401)**: token invalid atau sudah logout.

---

## 3. Alur Use Case

1. Middleware auth membaca token dari `Authorization` header.
2. Validasi token (JWT verify atau cek session di Redis).
3. Invalidasi: hapus session dari Redis atau tambahkan JWT ke blacklist (jika pakai JWT).
4. Return 200 OK.

---

## 4. Integrasi

- **Login**: Logout harus kompatibel dengan mekanisme session/token yang dipakai Login (simpan session_id atau jti di Redis).
- **Frontend**: Setelah logout berhasil, frontend menghapus token dari storage dan redirect ke `/auth/login?reason=logged_out`.

---

## 5. Non-Fungsional

- Response time target \< 200ms.
- Logging: logout sukses/gagal untuk audit.
