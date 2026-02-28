## SDD – Auth: Registrasi Peserta

**Aplikasi**: Widyaprada  
**Modul**: Auth & WPUjikom  
**Fitur**: Registrasi Mandiri Calon Peserta  

Dokumen ini menjelaskan desain teknis backend untuk PRD `[PRD] Fitur Registrasi Peserta` dengan stack: Go (Golang), Clean Architecture, PostgreSQL.

---

## 1. Arsitektur & Konteks

- Usecase: `RegistrationUsecase` — Register, SendPasswordEmail.
- Delivery: REST `POST /api/v1/auth/register`.
- Infrastructure: UserRepository, EmailService (kirim password ke email).

---

## 2. Kontrak API

- **Endpoint**: `POST /api/v1/auth/register`
- **Request**: nama_lengkap, email (unique), nip (opsional)
- **Response Sukses (201)**: pesan konfirmasi; password dikirim ke email.
- **Response Gagal (400)**: email sudah terdaftar, validasi field.

---

## 3. Skema Database

- Reuse `users`; role default = Peserta/Calon WP.
- Password digenerate sistem (random/strong); hash simpan di users; plaintext dikirim via email.

---

## 4. Aturan Bisnis

- Email unik. Format valid.
- Password digenerate otomatis (bcrypt/argon2 hash); kirim plaintext ke email (one-time).
- Setelah registrasi, user bisa login dengan email + password yang dikirim.

---

## 5. RBAC

- Endpoint publik (tidak perlu auth) untuk registrasi.
