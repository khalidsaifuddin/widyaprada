# [PRD] Fitur Lupa Password
## Product Requirements Document | Auth & RBAC

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Lupa Password  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Lupa Password
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-01-28
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Fitur Lupa Password memungkinkan pengguna terdaftar (Widyaprada, Admin Satker, Super Admin) yang lupa kata sandi untuk mengatur ulang kata sandi secara mandiri. Alur yang sama dipakai untuk semua peran: memasukkan email terdaftar, menerima tautan atau instruksi reset, lalu membuat kata sandi baru. Setelah berhasil, pengguna kembali ke halaman login dan masuk seperti biasa sesuai perannya.

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

Kebutuhan fitur Lupa Password dirumuskan per role dalam pola user story berikut, dalam format tabular.

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Semua | pengguna yang lupa kata sandi | mengakses alur reset kata sandi dari halaman login (tautan "Lupa kata sandi?") | mengatur ulang kata sandi tanpa harus menghubungi admin dulu |
| 2 | Semua | pengguna yang lupa kata sandi | memasukkan email yang terdaftar di satu form yang jelas | meminta tautan atau instruksi reset tanpa bingung field apa yang diisi |
| 3 | Semua | pengguna yang lupa kata sandi | mendapat konfirmasi jelas setelah mengirim permintaan (sukses atau gagal) | tahu langkah berikutnya (cek email atau perbaiki input) tanpa istilah teknis |
| 4 | Semua | pengguna yang lupa kata sandi | mengikuti instruksi yang dikirim (misalnya membuka tautan dari email) lalu memasukkan kata sandi baru dan konfirmasi | menyelesaikan reset tanpa bingung apa yang harus diketik dan di mana |
| 5 | Semua | pengguna yang lupa kata sandi | melihat kata sandi baru dengan opsi tampilkan/sembunyikan (ikon mata) saat mengisi | memeriksa ketikan kata sandi baru jika ragu tanpa mengetik ulang |
| 6 | Semua | pengguna yang lupa kata sandi | setelah berhasil reset diarahkan ke halaman login dengan pesan sukses | tahu bahwa saya bisa langsung masuk dengan kata sandi baru |
| 7 | Semua | pengguna yang lupa kata sandi | jika email tidak terdaftar atau permintaan gagal mendapat pesan yang ramah dan aman | tahu apa yang harus dilakukan (cek email, hubungi admin) tanpa informasi yang membingungkan |
| 8 | Semua | pengguna yang lupa kata sandi | tombol/kirim menampilkan keadaan "Memproses…" atau "Mengirim…" saat permintaan diproses | yakin bahwa aplikasi sedang bekerja dan tidak hang |
| 9 | Widyaprada | Widyaprada yang lupa kata sandi | setelah reset berhasil kembali ke halaman login dan masuk seperti Widyaprada | langsung masuk ke dashboard/area Widyaprada dengan kata sandi baru |
| 10 | Admin Satker | Admin Satker yang lupa kata sandi | setelah reset berhasil kembali ke halaman login dan masuk seperti Admin Satker | langsung masuk ke dashboard/admin satker dengan kata sandi baru |
| 11 | Super Admin | Super Admin yang lupa kata sandi | setelah reset berhasil kembali ke halaman login dan masuk seperti Super Admin | langsung masuk ke area Super Admin dengan kata sandi baru |
| 12 | Semua | pengguna yang lupa kata sandi | ada tautan "Kembali ke login" di tiap langkah alur lupa password | kembali ke login kapan saja jika salah halaman atau ingin membatalkan |

**Keterangan**: *Semua* = berlaku untuk Widyaprada, Admin Satker, dan Super Admin.

---

### 2.1 Kebutuhan Pengguna - Yang Diinginkan
- Mengakses alur **Lupa kata sandi** dari satu tempat yang jelas (tautan di halaman login).
- Hanya memasukkan **email terdaftar** (satu field) untuk meminta reset - tanpa diminta username atau data lain yang mungkin tidak diingat.
- Mendapat **pesan konfirmasi** setelah mengirim permintaan (misalnya "Jika email terdaftar, Anda akan menerima tautan reset di inbox Anda").
- Mengikuti **instruksi sederhana** (buka email → klik tautan atau masukkan kode) lalu mengisi **kata sandi baru** dan **konfirmasi kata sandi**.
- Setelah berhasil, **kembali ke login** dengan pesan sukses dan bisa langsung masuk dengan kata sandi baru.
- Bisa **kembali ke login** kapan saja dari halaman lupa password tanpa tersesat.

### 2.2 Kebutuhan Pengguna - Yang Tidak Diinginkan
- Pesan error yang mengungkap apakah email terdaftar atau tidak (demi keamanan dan privasi).
- Langkah yang terlalu banyak atau istilah teknis (misalnya "token", "expired" tanpa penjelasan ramah).
- Form reset kata sandi tanpa opsi tampilkan/sembunyikan untuk cek ketikan.
- Tidak ada umpan balik saat tombol "Kirim" atau "Reset" diklik (tombol terasa hang).

---

## 3. Antarmuka Pengguna (UI)

*Alur Lupa Password dapat terdiri atas beberapa halaman/langkah. Semua role melihat tampilan yang sama.*

### 3.1 Langkah 1 - Meminta reset (dari Login)

**Elemen yang harus ada**:
- **Judul singkat**: Misalnya "Lupa kata sandi?" atau "Atur ulang kata sandi".
- **Penjelasan singkat**: Satu kalimat, misalnya "Masukkan email terdaftar. Kami akan mengirim tautan untuk mengatur ulang kata sandi."
- **Satu field teks**: Label **"Email"** - tempat pengguna memasukkan email yang terdaftar.
- **Tombol utama**: **"Kirim tautan"** atau **"Kirim instruksi"** - satu tombol jelas.
- **Tautan**: **"Kembali ke login"** - kembali ke halaman login.

**Tata letak dan keterbacaan**:
- Form singkat dan tidak penuh; jarak antar elemen nyaman.
- Di layar kecil (ponsel/tablet), form tetap nyaman dipakai dan tombol tidak terpotong.

### 3.2 Umpan balik setelah mengirim permintaan

**Jika permintaan berhasil diterima** (tanpa mengungkap apakah email terdaftar atau tidak, demi keamanan):
- Pesan informatif, misalnya: **"Jika email Anda terdaftar, Anda akan menerima tautan untuk mengatur ulang kata sandi. Periksa juga folder spam."**
- Opsi **"Kembali ke login"** atau tombol serupa.

**Jika ada validasi gagal** (misalnya field kosong):
- Pesan di dekat field atau di atas form, misalnya "Email wajib diisi" atau "Format email tidak valid."

**Sedang memproses**:
- Saat pengguna menekan "Kirim tautan", tombol menampilkan keadaan loading (misalnya "Mengirim…" atau spinner) dan tidak bisa diklik lagi sampai selesai.

### 3.3 Langkah 2 - Mengatur kata sandi baru (setelah klik tautan dari email)

**Elemen yang harus ada**:
- **Judul singkat**: Misalnya "Buat kata sandi baru".
- **Dua field sandi**:
  - **"Kata sandi baru"** - input tersembunyi dengan opsi tampilkan/sembunyikan (ikon mata).
  - **"Konfirmasi kata sandi baru"** - input tersembunyi dengan opsi tampilkan/sembunyikan.
- **Tombol utama**: **"Simpan kata sandi"** atau **"Atur ulang kata sandi"**.
- **Tautan**: **"Kembali ke login"** (opsional, jika tidak mengganggu alur).

**Umpan balik**:
- Jika kata sandi dan konfirmasi tidak sama: pesan jelas, misalnya "Konfirmasi kata sandi tidak sama."
- Jika aturan kata sandi tidak terpenuhi: pesan singkat yang menjelaskan aturan (misalnya minimal 8 karakter, ada huruf dan angka) - tanpa istilah teknis.
- Sedang memproses: tombol menampilkan "Menyimpan…" atau spinner.
- **Jika berhasil**: Pesan sukses, misalnya "Kata sandi berhasil diubah. Silakan masuk dengan kata sandi baru." dan tautan/tombol **"Ke halaman login"**.

**Jika tautan tidak valid atau kedaluwarsa**:
- Pesan ramah, misalnya "Tautan tidak valid atau sudah kedaluwarsa. Silakan minta tautan reset lagi dari halaman Lupa kata sandi."
- Tautan ke halaman **Lupa kata sandi** (langkah 1) atau **Login**.

---

## 4. Pengalaman Pengguna (UX) - Alur Lupa Password

### 4.1 Alur umum (Semua Role)

1. Pengguna di halaman **Login** → mengklik **"Lupa kata sandi?"**.
2. Masuk ke **halaman/langkah meminta reset** → memasukkan **email** terdaftar → menekan **"Kirim tautan"** (atau setara).
3. Melihat umpan balik **"Mengirim…"** pada tombol, lalu pesan konfirmasi (misalnya "Jika email terdaftar, Anda akan menerima tautan…").
4. Pengguna membuka **email** → mengklik **tautan** di email (atau mengikuti instruksi lain yang diberikan).
5. Dibuka **halaman buat kata sandi baru** → mengisi **kata sandi baru** dan **konfirmasi kata sandi** → menekan **"Simpan kata sandi"** (atau setara).
6. Melihat umpan balik "Menyimpan…", lalu **pesan sukses** dan tautan **"Ke halaman login"**.
7. Pengguna mengklik ke login → memasukkan **email/username** dan **kata sandi baru** → masuk seperti biasa sesuai peran (lihat PRD Login).

### 4.2 Perilaku yang diharapkan

- Dari halaman login, **satu klik** ke "Lupa kata sandi?" langsung masuk ke langkah meminta reset; tidak ada langkah perantara yang membingungkan.
- Di setiap langkah alur lupa password, pengguna bisa **kembali ke login** tanpa tersesat.
- Setelah kata sandi berhasil diubah, **satu aksi** (klik "Ke halaman login") membawa pengguna ke form login; tidak perlu menutup browser atau mencari URL.
- Pengalaman setelah masuk kembali **sama untuk semua peran** (Widyaprada, Admin Satker, Super Admin): mengikuti PRD Login (redirect ke "rumah" masing-masing).

---

## 5. Kebutuhan per Role: Setelah Berhasil Reset

*Alur Lupa Password itu sendiri sama untuk semua role. Yang berbeda hanya setelah pengguna kembali login dengan kata sandi baru - mengikuti perilaku per role di PRD Login.*

### 5.1 Role: Widyaprada

**Setelah reset berhasil**:
- Kembali ke halaman login, masuk dengan email/username dan kata sandi baru.
- Setelah login berhasil, diarahkan ke **halaman pertama Widyaprada** (dashboard atau modul default) dan hanya melihat menu WPData, WPJurnal, WPUjikom, serta bagian landing yang diizinkan - sama seperti di PRD Login.

### 5.2 Role: Admin Satker

**Setelah reset berhasil**:
- Kembali ke halaman login, masuk dengan email/username dan kata sandi baru.
- Setelah login berhasil, diarahkan ke **halaman pertama Admin Satker** (dashboard admin, CMS, atau Manajemen Data Widyaprada satker) dan melihat menu yang relevan dengan admin satker - sama seperti di PRD Login.

### 5.3 Role: Super Admin

**Setelah reset berhasil**:
- Kembali ke halaman login, masuk dengan email/username dan kata sandi baru.
- Setelah login berhasil, diarahkan ke **halaman pertama Super Admin** dan mengakses semua menu - sama seperti di PRD Login.

---

## 6. Ringkasan Perbedaan per Role

| Aspek | Widyaprada | Admin Satker | Super Admin |
|-------|------------|--------------|-------------|
| **Alur Lupa Password** | Sama: minta reset → terima tautan → buat kata sandi baru → kembali ke login | Sama | Sama |
| **Setelah reset & login lagi** | Masuk ke dashboard/area Widyaprada (WPData, WPJurnal, WPUjikom, landing) | Masuk ke dashboard/admin satker (CMS, Manajemen Data WP satker, dll) | Masuk ke area Super Admin (semua menu) |
| **Jika butuh bantuan** | Hubungi Admin Satker atau Super Admin instansi/satker | Hubungi Super Admin | Hubungi tim teknis atau kelola sendiri |

---

## 7. Cakupan Fitur Lupa Password

### 7.1 Termasuk dalam fitur ini
- Tautan **"Lupa kata sandi?"** di halaman login yang mengarah ke alur reset.
- Halaman/langkah **meminta reset** dengan satu field email dan tombol kirim.
- Pesan konfirmasi setelah kirim (ramah dan aman, tanpa mengungkap ada/tidaknya email terdaftar).
- Halaman **buat kata sandi baru** (setelah pengguna mengikuti tautan/instruksi dari email) dengan field kata sandi baru dan konfirmasi, serta opsi tampilkan/sembunyikan.
- Pesan sukses dan **arah ke halaman login** setelah kata sandi berhasil diubah.
- Tautan **"Kembali ke login"** di langkah yang relevan.
- Umpan balik loading pada tombol dan pesan error/validasi yang jelas.

### 7.2 Tidak termasuk (fitur lain)
- Login → PRD Login.
- Logout → PRD Logout.
- Manajemen pengguna (reset kata sandi oleh admin) → PRD Manajemen Pengguna.
- Registrasi atau aktivasi akun → PRD terpisah jika ada.

---

## 8. Persyaratan Produk (Nonteknis)

Detail implementasi (pengiriman email, token reset, validasi) didokumentasikan di SDD.

- **Pesan setelah kirim:** Pesan setelah "Kirim tautan" tidak boleh mengungkap apakah email terdaftar atau tidak. Cukup satu pesan umum (misalnya "Jika email Anda terdaftar, Anda akan menerima tautan untuk mengatur ulang kata sandi. Periksa juga folder spam.").
- **Ketersediaan:** Halaman lupa password dapat diakses kapan pun pengguna butuh.
- **Performa:** Waktu dari menekan "Kirim" atau "Simpan kata sandi" sampai muncul pesan konfirmasi/sukses harus wajar agar pengguna tidak mengira aplikasi error.
- **Akses perangkat:** Form dapat digunakan dengan nyaman dari komputer, tablet, dan ponsel.

---

## 9. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-01-28 | Dokumen awal, pola sama dengan PRD Login, fokus user & UI/UX, dibagi per role | - |

---

**Catatan**: Dokumen ini hanya mencakup fitur **Lupa Password**. Fitur **Login** dan **Logout** didokumentasikan dalam PRD masing-masing.
