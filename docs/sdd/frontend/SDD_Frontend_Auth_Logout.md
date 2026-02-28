## SDD Frontend – Auth: Logout

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Logout  
**PRD Terkait**: [PRD_Auth_Logout](../../prd/PRD_Auth_Logout.md)

Dokumen ini menjelaskan **desain teknis frontend** untuk fitur Logout dengan pendekatan **Atomic Design**.

---

## 1. Arsitektur & Konteks

- **Lokasi UI**: Tombol/tautan "Keluar" di Sidebar (organisms) atau header profil
- **API**: `POST /api/v1/auth/logout` (dengan token di header)
- **Post-logout**: Redirect ke `/auth/login?message=logout_success`

---

## 2. Atomic Design – Komponen

### 2.1 Atoms

| Komponen | Deskripsi |
|----------|-----------|
| `Button` / `Link` | Tombol "Keluar" dengan ikon (ArrowRightEndOnRectangle) |

### 2.2 Organisms

| Komponen | Deskripsi |
|----------|-----------|
| `Sidebar` | Sudah ada; tombol Keluar di bagian bawah, trigger logout |
| `UserMenu` (opsional) | Dropdown profil dengan opsi Keluar |

### 2.3 Perilaku

1. User klik "Keluar"
2. Opsional: dialog konfirmasi "Yakin keluar?" (sesuai PRD: langsung keluar tanpa konfirmasi)
3. Panggil `logout()` → clear token, POST /api/v1/auth/logout
4. Redirect ke `/auth/login?message=logout_success`
5. Halaman login tampilkan toast: "Anda telah keluar."
6. Tombol "Keluar" tampilkan loading saat proses

---

## 3. State

- `loading`: boolean saat logout diproses
- Tidak perlu state global; cukup di Sidebar/UserMenu

---

## 4. Integrasi API

| Method | Endpoint | Request |
|--------|----------|---------|
| POST | `/api/v1/auth/logout` | Header: Authorization Bearer {token} |

---

## 5. File Lokasi

```
frontend/src/
├── lib/auth.ts                     # logout()
├── components/organisms/Sidebar.tsx # Tombol Keluar
```
