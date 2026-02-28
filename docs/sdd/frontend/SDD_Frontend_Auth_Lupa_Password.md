## SDD Frontend – Auth: Lupa Password

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Lupa Password  
**PRD Terkait**: [PRD_Auth_Lupa_Password](../../prd/PRD_Auth_Lupa_Password.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk alur Lupa Password dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/auth/forgot-password`, `/auth/reset-password?token=...`
- **Link**: Dari halaman login "Lupa kata sandi?"
- **API**: `POST /api/v1/auth/forgot-password`, `POST /api/v1/auth/reset-password`
- **Layout**: AuthLayout

---

## 2. Atomic Design – Komponen

### 2.1 Langkah 1 – Meminta Reset (Forgot Password)

#### Atoms
- `Input`, `Label`, `Button`

#### Molecules
- `FormField`

#### Organisms
- `ForgotPasswordForm`: Field email, tombol "Kirim", link "Kembali ke login"
- `ForgotPasswordSuccess`: Pesan "Jika email terdaftar, Anda akan menerima tautan reset di inbox Anda"

### 2.2 Langkah 2 – Reset Kata Sandi (Dari Tautan Email)

#### Organisms
- `ResetPasswordForm`: Field password baru, konfirmasi password, tombol "Reset", opsi show/hide password
- `ResetPasswordSuccess`: "Kata sandi berhasil direset. Silakan login dengan kata sandi baru."; link ke login

### 2.3 Pages

| Route | Page |
|-------|------|
| `/auth/forgot-password` | ForgotPasswordPage |
| `/auth/reset-password` | ResetPasswordPage (baca token dari query) |

---

## 3. State & Validasi

### 3.1 ForgotPasswordForm
- `email`, `loading`, `error`, `success`

### 3.2 ResetPasswordForm
- `password`, `confirmPassword`, `showPassword`, `loading`, `error`
- Validasi: password minimal 8 karakter; password === confirmPassword

---

## 4. Integrasi API

| Method | Endpoint | Request |
|--------|----------|---------|
| POST | `/api/v1/auth/forgot-password` | `{ email }` |
| POST | `/api/v1/auth/reset-password` | `{ token, password, confirm_password }` |

---

## 5. File Lokasi

```
frontend/src/
├── app/auth/forgot-password/page.tsx
├── app/auth/reset-password/page.tsx
├── components/organisms/ForgotPasswordForm.tsx
└── components/organisms/ResetPasswordForm.tsx
```
