## SDD Frontend – Auth: Registrasi

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Registrasi Mandiri Calon Peserta  
**PRD Terkait**: [PRD_Auth_Registrasi](../../prd/PRD_Auth_Registrasi.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk halaman Registrasi dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Route**: `/auth/register` atau `/auth/registrasi`
- **Proteksi**: Link dari halaman login; tidak perlu auth
- **API**: `POST /api/v1/auth/register` (nama, email, nip?)
- **Layout**: AuthLayout

---

## 2. Atomic Design – Komponen

### 2.1 Atoms

| Komponen | Deskripsi |
|----------|-----------|
| `Button` | Tombol Daftar, loading state |
| `Input` | Input teks (nama, email, nip) |
| `Label` | Label form |

### 2.2 Molecules

| Komponen | Deskripsi |
|----------|-----------|
| `FormField` | Label + Input + error |
| `AuthBranding` | Judul aplikasi |

### 2.3 Organisms

| Komponen | Deskripsi |
|----------|-----------|
| `RegistrationForm` | Form: Nama, Email, NIP (opsional), tombol Daftar; link ke Login |
| `RegistrationSuccess` | Pesan konfirmasi: "Registrasi berhasil. Silakan cek email Anda..."; link kembali ke login |

### 2.4 Pages

| Komponen | Deskripsi |
|----------|-----------|
| `RegisterPage` | AuthLayout + RegistrationForm |
| `RegisterSuccessPage` | AuthLayout + RegistrationSuccess (setelah submit sukses) |

---

## 3. State & Validasi

### 3.1 State Lokal

- `nama`, `email`, `nip` (opsional)
- `loading`, `error`
- `success`: boolean (redirect atau tampil RegistrationSuccess)

### 3.2 Validasi Client-side

- Email format valid
- Field wajib: nama, email

### 3.3 Pesan Error

- Email sudah terdaftar: "Email ini sudah terdaftar. Gunakan Lupa Password jika Anda lupa kata sandi."
- Email tidak valid: pesan validasi per field

---

## 4. Integrasi API

| Method | Endpoint | Request | Response |
|--------|----------|---------|----------|
| POST | `/api/v1/auth/register` | `{ name, email, nip? }` | 201 + pesan sukses |

---

## 5. File Lokasi

```
frontend/src/
├── app/auth/register/page.tsx      # RegisterPage
├── components/organisms/RegistrationForm.tsx
```
