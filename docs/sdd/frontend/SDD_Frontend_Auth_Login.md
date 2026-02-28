## SDD Frontend – Auth: Login

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Login  
**PRD Terkait**: [PRD_Auth_Login](../../prd/PRD_Auth_Login.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk halaman Login dengan pendekatan **Atomic Design**, stack React (Next.js), Tailwind CSS, shadcn/ui.

---

## 1. Arsitektur & Konteks

- **Route**: `/auth/login`
- **Proteksi**: Jika user sudah login, redirect ke `default_home_path` dari response login
- **API**: `POST /api/v1/auth/login` (identifier, password)
- **Layout**: AuthLayout (tanpa sidebar, full-height form center)

---

## 2. Atomic Design – Komponen

### 2.1 Atoms

| Komponen | Deskripsi | Props/Use |
|----------|-----------|-----------|
| `Button` | Tombol dengan variant primary, loading state | `onClick`, `disabled`, `loading`, `children` |
| `Input` | Input teks (text, password) | `type`, `value`, `onChange`, `placeholder`, `error` |
| `Label` | Label form | `htmlFor`, `children` |
| `Icon` (Eye/EyeOff) | Toggle tampilkan/sembunyikan password | `onClick`, `show` |

### 2.2 Molecules

| Komponen | Deskripsi | Atoms digunakan |
|----------|-----------|-----------------|
| `FormField` | Label + Input + pesan error | Label, Input |
| `PasswordField` | Input password + ikon show/hide | Input, Icon (Eye) |
| `AuthBranding` | Judul/branding aplikasi (Logo + "Aplikasi Widyaprada") | - |
| `LoginButton` | Tombol Masuk dengan loading state | Button |

### 2.3 Organisms

| Komponen | Deskripsi | Molecules digunakan |
|----------|-----------|---------------------|
| `LoginForm` | Form lengkap: Email/Username, Password, Masuk, Lupa kata sandi | FormField, PasswordField, LoginButton, Link |

### 2.4 Templates

| Komponen | Deskripsi |
|----------|-----------|
| `AuthLayout` | Layout full-height, center card; tanpa sidebar; background netral |

### 2.5 Pages

| Komponen | Deskripsi |
|----------|-----------|
| `LoginPage` | AuthLayout + LoginForm + AuthBranding |

---

## 3. Struktur Halaman & State

### 3.1 State Lokal (LoginForm)

- `identifier`: string (email atau username)
- `password`: string
- `showPassword`: boolean
- `loading`: boolean (saat submit)
- `error`: string | null (pesan error dari API atau validasi)

### 3.2 Validasi Client-side

- Field wajib: identifier dan password tidak boleh kosong
- Pesan: "Email atau username wajib diisi", "Kata sandi wajib diisi"

### 3.3 Alur Submit

1. Validasi client-side
2. Set loading = true, error = null
3. `POST /api/v1/auth/login` dengan `{ identifier, password }`
4. Sukses: simpan token, redirect ke `user.default_home_path` (atau `/dashboard`)
5. Gagal: set error dari response (aman, tidak bocorkan email terdaftar); loading = false

### 3.4 Umpan Balik UI

- Loading: tombol "Memproses…" atau spinner, disabled
- Error: alert/banner di atas form dengan pesan (kredensial salah, akun nonaktif, terlalu banyak percobaan)
- Sesi habis (query param): tampilkan toast "Sesi Anda telah berakhir. Silakan masuk kembali."

---

## 4. Responsivitas & Aksesibilitas

- Layout responsif: form lebar penuh di mobile, max-width card di desktop
- Label terhubung dengan input (htmlFor, id)
- Error message terhubung dengan aria-describedby
- Focus trap dalam form
- Keyboard navigasi (Tab, Enter)

---

## 5. Integrasi API

| Method | Endpoint | Request | Response |
|--------|----------|---------|----------|
| POST | `/api/v1/auth/login` | `{ identifier, password }` | `{ access_token, user: { default_home_path, ... } }` |

---

## 6. File & Lokasi Komponen

```
frontend/src/
├── app/auth/login/page.tsx         # LoginPage
├── components/
│   ├── atoms/Button.tsx, Input.tsx
│   ├── molecules/FormField.tsx, PasswordField.tsx
│   ├── organisms/LoginForm.tsx
│   └── templates/AuthLayout.tsx
└── lib/auth.ts                     # login(), getUserProfile()
```
