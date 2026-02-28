# [PRD] Fitur Registrasi Peserta (Calon WP)
## Product Requirements Document | Auth & Ujikom

**Aplikasi**: Widyaprada  
**Modul**: Auth & WPUjikom  
**Fitur**: Registrasi Mandiri Calon Peserta Uji Kompetensi  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Registrasi Peserta (Calon WP)
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-02-28
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Registrasi memungkinkan **calon peserta** (calon Widyaprada/non-WP yang mengikuti perpindahan jabatan fungsional ke WP) untuk mendaftar akun secara mandiri. Setelah mendaftar, **password dikirim melalui email aktif** ke alamat email yang didaftarkan. Alur ini diperlukan karena jenis Ujikom yang diakomodir saat ini adalah **perpindahan jabatan fungsional dari non-WP ke WP**, sehingga belum ada data peserta awal di sistem — peserta harus mendaftar terlebih dahulu.

### 1.3 Konteks
- Jenis Ujikom yang diakomodir saat ini: **Perpindahan jabatan fungsional (non-WP → WP)**.
- Kenaikan tingkat WP: **nonaktif/disabled** untuk fase ini.
- Peserta yang mendaftar adalah calon (belum memiliki data WP di sistem).

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Calon Peserta | calon WP | mendaftar akun melalui form registrasi dengan email aktif | menerima password melalui email dan login ke aplikasi |
| 2 | Calon Peserta | calon WP | mengisi data dasar (nama, email, NIP, dll) saat registrasi | profil awal terisi untuk proses selanjutnya |
| 3 | Calon Peserta | calon WP | mendapat konfirmasi bahwa registrasi berhasil | tahu langkah selanjutnya (cek email untuk password) |
| 4 | Calon Peserta | calon WP | mendapat password melalui email ke alamat yang didaftarkan | bisa login ke aplikasi setelah menerima email |
| 5 | Semua | pengunjung | melihat link "Daftar" atau "Registrasi" di halaman login | bisa mengakses form registrasi jika belum punya akun |

---

## 3. Antarmuka Pengguna (UI)

### 3.1 Form Registrasi
- **Lokasi**: Halaman terpisah atau modal; link dari halaman login ("Belum punya akun? Daftar di sini").
- **Field wajib** (sesuai kebijakan): Nama lengkap, Email (unik, aktif), NIP (opsional sesuai kebijakan).
- **Tombol**: "Daftar" / "Registrasi".
- **Validasi**: Email format valid, email belum terdaftar, field wajib terisi.

### 3.2 Setelah Submit Sukses
- Pesan: "Registrasi berhasil. Silakan cek email Anda untuk mendapatkan password. Gunakan email dan password tersebut untuk login."
- Link kembali ke halaman login.

### 3.3 Pesan Error
- Email sudah terdaftar: "Email ini sudah terdaftar. Gunakan Lupa Password jika Anda lupa kata sandi."
- Email tidak valid: pesan validasi per field.

---

## 4. Pengalaman Pengguna (UX) - Alur

### 4.1 Alur Registrasi
1. Pengguna membuka halaman login → klik "Daftar" / "Registrasi".
2. Form registrasi tampil; pengguna mengisi nama, email (aktif), dan field lain yang wajib.
3. Klik "Daftar" → sistem memproses.
4. **Sukses**: Pesan konfirmasi + instruksi cek email. Password dikirim ke email yang didaftarkan.
5. Pengguna membuka email, mendapat password, lalu login dengan email + password.

---

## 5. Cakupan Fitur

### 5.1 Termasuk
- Form registrasi mandiri calon peserta.
- Pengiriman password ke email setelah registrasi berhasil.
- Validasi email unik dan format.
- Link dari halaman login ke registrasi.

### 5.2 Tidak Termasuk
- Registrasi oleh admin → PRD Manajemen Pengguna.
- Aktivasi email (klik link) — password langsung dikirim; bisa ditambah di fase berikutnya.
- Verifikasi dokumen — PRD Pendaftaran Ujikom / Manajemen Data WP.

---

## 6. Persyaratan Produk (Nonteknis)

- **Email aktif**: Sistem mengirim password ke email yang didaftarkan; email harus valid dan dapat diterima.
- **Keamanan**: Password digenerate oleh sistem (atau user set saat registrasi — kebijakan produk); jika digenerate, harus kuat dan acak.
- **Unik**: Email tidak boleh duplikat dengan pengguna yang sudah ada.

---

## 7. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-02-28 | Dokumen awal: Registrasi mandiri calon peserta, password via email | - |

---

**Catatan**: Setelah registrasi dan login, peserta mengikuti alur **Pendaftaran Ujikom** (pilih jenis ujikom, upload dokumen persyaratan). Lihat PRD_Assignment.md, PRD_Manajemen_Uji_Kompetensi.md.
