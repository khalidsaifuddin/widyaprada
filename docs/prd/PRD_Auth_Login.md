# [PRD] Fitur Login
## Product Requirements Document | Auth & RBAC

**Aplikasi**: Widyaprada  
**Modul**: Auth & RBAC  
**Fitur**: Login  
**Fokus**: Kebutuhan pengguna, antarmuka (UI), dan pengalaman pengguna (UX)

---

## 1. Informasi Umum

### 1.1 Identitas
- **Nama Fitur**: Login
- **Versi Dokumen**: 1.0
- **Tanggal**: 2025-01-28
- **Status**: Draft
- **Prioritas**: High

### 1.2 Ringkasan untuk Pengguna
Login adalah cara satu pintu bagi pengguna terdaftar untuk masuk ke Aplikasi Widyaprada. Satu halaman login dipakai bersama oleh semua peran (Widyaprada, Admin Satker, Admin Uji Kompetensi, Verifikator Uji Kompetensi, Super Admin). **Satu pengguna dapat memiliki lebih dari satu role** (misalnya Widyaprada sekaligus Admin Uji Kompetensi); setelah login, menu dan wewenang yang tampil adalah **gabungan dari semua role** yang dimiliki pengguna, dan pengguna diarahkan ke halaman pertama yang relevan (tanpa harus memilih “login sebagai apa”). Dengan demikian pengguna bisa langsung mengerjakan tugasnya tanpa bingung.

---

## 2. User Story (Sebagai … Saya ingin … Agar saya bisa …)

Kebutuhan fitur Login dirumuskan per role dalam pola user story berikut, dalam format tabular.

| No | Role | Sebagai … | Saya ingin … | Agar saya bisa … |
|----|------|-----------|---------------|-------------------|
| 1 | Semua | pengguna terdaftar | masuk dengan email atau username dan kata sandi di satu halaman login | mengakses aplikasi tanpa harus memilih "login sebagai apa" dulu |
| 2 | Semua | pengguna terdaftar | melihat pesan yang jelas saat login gagal (kredensial salah, akun nonaktif, atau batasan percobaan) | tahu apa yang harus dilakukan tanpa bingung atau melihat istilah teknis |
| 3 | Semua | pengguna terdaftar | setelah berhasil login langsung diarahkan ke halaman pertama yang sesuai peran saya | langsung mengerjakan tugas tanpa langkah tambahan atau pilih role lagi |
| 4 | Semua | pengguna terdaftar | ada tautan "Lupa kata sandi?" di halaman login | mengatur ulang kata sandi jika lupa tanpa bingung mencari cara |
| 5 | Semua | pengguna terdaftar | tombol Masuk menampilkan keadaan "Memproses…" saat login diproses | yakin bahwa aplikasi sedang bekerja dan tidak hang |
| 6 | Semua | pengguna terdaftar | mengisi kata sandi dengan opsi tampilkan/sembunyikan (ikon mata) | memeriksa ketikan kata sandi jika ragu tanpa harus mengetik ulang |
| 7 | Widyaprada | Widyaprada (PNS jabatan fungsional penjaminan mutu) | setelah login hanya melihat menu WPData, WPJurnal, WPUjikom, dan bagian landing yang untuk saya | fokus ke tugas penjaminan mutu tanpa melihat menu CMS atau Manajemen Pengguna yang tidak saya gunakan |
| 8 | Widyaprada | Widyaprada | setelah login langsung ke dashboard atau ringkasan tugas WP / modul default | langsung bekerja tanpa layar "Pilih role" atau "Pilih modul" |
| 9 | Admin Satker | Admin Satker | setelah login langsung ke dashboard admin atau CMS (Berita/Tautan) dan hanya melihat menu yang relevan dengan admin satker | langsung mengerjakan tugas admin tanpa melihat menu Role & Permission tingkat sistem |
| 10 | Admin Satker | Admin Satker | setelah login mengakses CMS Berita, CMS Tautan, Manajemen Data Widyaprada (di satker/unit kerja/lembaga/instansi saya), dan Manajemen Pengguna (jika diizinkan) serta WPData/WPJurnal/WPUjikom sesuai wewenang | mengelola konten, data Widyaprada di satker saya, dan data satker lainnya tanpa akses yang tidak relevan |
| 11 | Super Admin | Super Admin | setelah login mengakses semua menu termasuk Manajemen Pengguna dan Role & Permission | mengelola pengguna, role, permission, dan konfigurasi sistem secara penuh |
| 12 | Super Admin | Super Admin | setelah login langsung ke dashboard Super Admin atau Manajemen Pengguna / Role & Permission | langsung menjalankan tugas konfigurasi dan pengawasan tanpa langkah tambahan |
| 13 | Semua | pengguna yang sudah login | saat sesi habis diarahkan ke halaman login dengan pesan "Sesi Anda telah berakhir. Silakan masuk kembali." | tahu bahwa saya harus masuk lagi dan tidak mengira aplikasi error |
| 14 | Semua | pengguna yang sudah login | ketika membuka URL halaman login lagi langsung diarahkan ke "rumah" peran saya | tidak terjebak di form login padahal sudah masuk |

**Keterangan**: *Semua* = berlaku untuk semua role. Satu user bisa punya **lebih dari satu role**; setelah login menu dan akses adalah gabungan dari semua role tersebut.

---

### 2.1 Kebutuhan Pengguna - Yang Diinginkan
- Bisa masuk ke aplikasi dengan **satu tempat** (satu halaman login), tanpa harus memilih “login sebagai apa” dulu.
- Memasukkan **identitas** (email atau username) dan **kata sandi** yang sudah diberikan/didaftarkan.
- Mendapat **konfirmasi jelas** jika login gagal (pesan yang mudah dipahami, tanpa istilah teknis).
- Setelah berhasil, **langsung sampai di “rumah” masing-masing** (halaman pertama yang relevan dengan tugasnya), tanpa langkah tambahan.
- Ada **jalan keluar** jika lupa password (link ke fitur Lupa Password).

### 2.2 Kebutuhan Pengguna - Yang Tidak Diinginkan
- Pesan error yang membingungkan atau terlalu teknis.
- Diarahkan ke halaman yang tidak sesuai peran (misalnya Widyaprada masuk ke menu Super Admin).
- Form yang lambat memberi umpan balik (misalnya tombol Login tidak ada indikasi “sedang memproses”).

---

## 3. Antarmuka Pengguna (UI) - Halaman Login

*Halaman login ini sama untuk semua role; tidak ada pilihan “Login sebagai Widyaprada/Admin/Super Admin”. Sistem mengenali peran setelah kredensial divalidasi.*

### 3.1 Elemen yang Harus Ada
- **Judul / branding aplikasi**: Misalnya “Aplikasi Widyaprada” agar pengguna yakin berada di aplikasi yang benar.
- **Satu field teks**: Label **“Email atau Username”** - tempat pengguna memasukkan email atau username yang terdaftar.
- **Satu field sandi**: Label **“Kata sandi”** - input tersembunyi (titik/bintang) dengan opsi **tampilkan/sembunyikan** (ikon mata) agar pengguna bisa cek ketik jika perlu.
- **Tombol utama**: **“Masuk”** atau **“Login”** - satu tombol jelas, warna menonjol.
- **Tautan**: **“Lupa kata sandi?”** - mengarah ke alur Lupa Password (PRD terpisah).

### 3.2 Tata Letak dan Keterbacaan
- Form rapi dan tidak penuh; jarak antar field nyaman.
- Label terlihat jelas, placeholder (jika ada) tidak menggantikan label.
- Tombol “Masuk” mudah terlihat dan mudah diklik (ukuran dan kontras cukup).
- Di layar kecil (ponsel/tablet), form tetap nyaman dipakai dan tombol tidak terpotong.

### 3.3 Pesan dan Umpan Balik di Halaman Login
- **Sedang memproses**: Saat pengguna menekan “Masuk”, tombol menampilkan keadaan loading (misalnya teks “Memproses…” atau spinner) dan tidak bisa diklik lagi sampai selesai - agar pengguna tidak mengira form hang.
- **Field wajib**: Jika email/username atau kata sandi kosong saat submit, beri pesan singkat di dekat field atau di atas form, misalnya “Email atau username wajib diisi” / “Kata sandi wajib diisi”.
- **Login gagal**: Satu pesan umum yang ramah dan aman, misalnya **“Email/username atau kata sandi salah. Silakan coba lagi.”** - tanpa menyebut apakah email terdaftar atau tidak.
- **Akun tidak aktif**: Pesan terpisah yang informatif, misalnya **“Akun ini tidak aktif. Silakan hubungi administrator satker atau Super Admin.”**
- **Terlalu banyak percobaan gagal**: Jika sistem membatasi percobaan login, tampilkan pesan yang jelas, misalnya **“Terlalu banyak percobaan login. Silakan coba lagi setelah [waktu].”** - tanpa detail teknis.
- **Sesi habis** (jika pengguna diarahkan kembali ke login): Pesan seperti **“Sesi Anda telah berakhir. Silakan masuk kembali.”**

Semua pesan error ditampilkan dengan gaya yang konsisten (misalnya warna dan ikon peringatan yang sama) dan mudah dibaca.

---

## 4. Pengalaman Pengguna (UX) - Alur Login

### 4.1 Alur Umum (Semua Role)
1. Pengguna membuka aplikasi (atau diarahkan ke halaman login jika belum masuk).
2. Melihat satu form: Email/Username, Kata sandi, tombol Masuk, dan link Lupa kata sandi.
3. Mengisi kedua field dan menekan **Masuk**.
4. Melihat umpan balik “Memproses…” pada tombol.
5. **Jika berhasil**: Halaman beralih ke “halaman pertama” sesuai peran (lihat bagian per role di bawah). Tidak ada langkah pilih role atau konfirmasi tambahan.
6. **Jika gagal**: Pesan error tampil di halaman login; pengguna bisa memperbaiki input dan mencoba lagi.

### 4.2 Perilaku yang Diharapkan
- Pengguna yang sudah login tidak perlu melihat halaman login lagi saat menelusuri aplikasi; jika suatu saat sesi habis, baru diarahkan ke login dengan pesan “Sesi berakhir”.
- Jika pengguna yang sudah login secara sengaja membuka URL halaman login, lebih baik diarahkan ke “halaman pertama” perannya (tidak stuck di form login).

---

## 5. Kebutuhan per Role: Setelah Login

*Halaman login-nya sama; yang berbeda adalah tujuan setelah login dan pengalaman di dalam aplikasi. **Satu pengguna dapat memiliki lebih dari satu role** (misalnya Widyaprada + Admin Uji Kompetensi); dalam hal itu menu dan wewenang setelah login adalah gabungan dari semua role.*

---

### 5.1 Role: Widyaprada

**Siapa**: PNS dengan jabatan fungsional Widyaprada yang bertugas penjaminan mutu pendidikan (PAUD, dasar, menengah, masyarakat). Mereka membutuhkan akses ke data diri, jurnal, dan uji kompetensi.

**Kebutuhan saat baru login**:
- Langsung diarahkan ke **halaman pertama yang relevan** dengan tugas Widyaprada, misalnya:
  - Dashboard Widyaprada, atau
  - Ringkasan tugas (data WP, jurnal, uji kompetensi), atau
  - Langsung ke modul default (misalnya WPData) jika tidak ada dashboard.
- Menu yang terlihat: yang **boleh diakses** menurut role pengguna. Jika pengguna **hanya** Widyaprada: WPData, WPJurnal, WPUjikom, landing. Jika pengguna punya **lebih dari satu role** (mis. Widyaprada + Admin Uji Kompetensi), menu adalah **gabungan** dari semua role (mis. + Bank Soal, Paket Soal, Manajemen Uji Kompetensi).
- Pengguna merasa **“langsung di tempat kerja”** tanpa melihat opsi yang tidak dipakai.

**UI/UX yang diharapkan**:
- Setelah login, tidak ada layar “Pilih role” atau “Pilih modul”; satu langkah langsung ke halaman pertama role Widyaprada.
- Navigasi dan judul halaman konsisten dengan konteks “Widyaprada” (misalnya tidak muncul istilah “Admin Satker” atau “Super Admin” di menu mereka).

---

### 5.2 Role: Admin Satker

**Siapa**: Admin satuan kerja (satker) atau unit kerja/lembaga/instansi yang mengelola konten (berita, tautan), **data Widyaprada di satker/unit kerjanya**, dan mungkin verifikasi di tingkat satker. Mereka tidak mengelola role/permission global, tetapi butuh akses ke CMS, manajemen data Widyaprada dalam satker mereka, dan data/konten terkait satker lainnya.

**Kebutuhan saat baru login**:
- Langsung diarahkan ke **halaman pertama Admin Satker**, misalnya:
  - Dashboard admin, atau
  - CMS (Berita/Tautan) atau daftar konten yang perlu dikelola, atau
  - Manajemen Data Widyaprada (satker), atau
  - Ringkasan tugas admin (berita, tautan, data WP satker, jurnal, verifikasi yang relevan dengan satker).
- Menu yang terlihat: **Landing Page (CMS Berita, CMS Tautan)**, **Manajemen Data Widyaprada** (untuk satker/unit kerja/lembaga/instansi masing-masing), akses ke **WPData / WPJurnal / WPUjikom** sesuai wewenang (misalnya verifikasi), dan **Manajemen Pengguna** jika role ini punya wewenang - **tanpa** menu konfigurasi Role & Permission global dan tanpa akses penuh Super Admin.
- Pengguna merasa **langsung bisa mengerjakan tugas admin**, termasuk mengelola data Widyaprada di satker mereka, tanpa melihat menu yang tidak relevan.

**UI/UX yang diharapkan**:
- Satu langkah dari login ke “rumah” Admin Satker; tidak ada pilihan role.
- Istilah dan tata letak menu konsisten dengan konteks “Admin Satker” (misalnya “CMS Berita”, “Manajemen Data Widyaprada”, “Manajemen Pengguna”, tanpa “Kelola Role/Permission” di level sistem).
- Data Widyaprada yang dikelola terbatas pada satker/unit kerja/lembaga/instansi Admin Satker tersebut.

---

### 5.3 Role: Super Admin

**Siapa**: Pengguna dengan wewenang penuh untuk mengelola aplikasi: pengguna, role, permission, dan konfigurasi sistem. Mereka butuh akses ke semua modul termasuk Manajemen Pengguna, Role & Permission, dan CMS/landing.

**Kebutuhan saat baru login**:
- Langsung diarahkan ke **halaman pertama Super Admin**, misalnya:
  - Dashboard Super Admin, atau
  - Manajemen Pengguna / Role & Permission, atau
  - Ringkasan sistem (jumlah pengguna, role, dll).
- Menu yang terlihat: **semua** yang ada di aplikasi - Auth (kelola pengguna), **Role & Permission** (Role, Permission, Role–Permission), **Landing Page (CMS)**, **WPData**, **WPJurnal**, **WPUjikom**, sesuai desain menu aplikasi.
- Pengguna merasa **punya kendali penuh** dan tidak ada menu “tersembunyi” yang seharusnya bisa diakses.

**UI/UX yang diharapkan**:
- Satu langkah dari login ke “rumah” Super Admin.
- Navigasi jelas ke Manajemen Pengguna dan Role & Permission agar konfigurasi mudah dijangkau.
- Tidak ada kebingungan antara “Admin Satker” dan “Super Admin”; label dan hierarki menu membedakan dengan jelas.

---

## 6. Ringkasan Perbedaan per Role (Setelah Login)

| Aspek | Widyaprada | Admin Satker | Super Admin |
|-------|------------|--------------|-------------|
| **Halaman pertama setelah login** | Dashboard / ringkasan WP atau modul default (WPData dll) | Dashboard admin / CMS atau daftar konten | Dashboard Super Admin / Manajemen Pengguna atau ringkasan sistem |
| **Menu yang terlihat** | WPData, WPJurnal, WPUjikom, bagian landing yang diizinkan | CMS Berita/Tautan, **Manajemen Data Widyaprada (satker)**, WPData/WPJurnal/WPUjikom sesuai wewenang, Manajemen Pengguna (jika ada) | Semua: Pengguna, Role & Permission, CMS, WPData, WPJurnal, WPUjikom |
| **Yang tidak terlihat** | CMS, Manajemen Pengguna, Role & Permission | Role & Permission global, akses penuh Super Admin | - (semua relevan terlihat) |
| **Pengalaman yang diharapkan** | Langsung ke tugas penjaminan mutu | Langsung ke tugas admin satker | Langsung ke tugas konfigurasi dan pengawasan sistem |

---

## 7. Cakupan Fitur Login

### 7.1 Termasuk dalam Fitur Login
- Satu halaman login untuk semua role (Widyaprada, Admin Satker, Super Admin).
- Input: Email atau Username, Kata sandi; tombol Masuk; link Lupa kata sandi.
- Umpan balik: loading, field wajib, kredensial salah, akun nonaktif, batasan percobaan, sesi habis.
- Redirect ke halaman pertama yang sesuai peran setelah login berhasil.
- Menu dan akses halaman sesuai role (tanpa pilihan “login sebagai” di form).

### 7.2 Tidak Termasuk (Fitur Lain)
- Lupa password → PRD Lupa Password.
- Logout → PRD Logout.
- Manajemen pengguna, role, permission → PRD masing-masing modul.
- Registrasi mandiri calon peserta → PRD Auth Registrasi.

---

## 8. Persyaratan Produk (Nonteknis)

Detail implementasi (sesi, token, validasi backend) didokumentasikan di SDD.

- **Keamanan pesan**: Pesan error tidak boleh mengungkap apakah suatu email/username terdaftar atau tidak; cukup satu pesan umum untuk “kredensial salah”.
- **Ketersediaan:** Halaman login dapat diakses kapan pun pengguna butuh; jika ada pemeliharaan, tampilkan pemberitahuan sederhana.
- **Performa yang terasa**: Waktu dari menekan “Masuk” sampai pindah halaman atau muncul pesan error sebaiknya wajar (tidak sampai puluhan detik) agar pengguna tidak mengira aplikasi error.
- **Akses perangkat:** Form login dapat digunakan dengan nyaman dari komputer, tablet, dan ponsel (layout dan tombol tetap bisa dipakai).

---

## 9. Changelog

| Versi | Tanggal | Perubahan | Penulis |
|-------|---------|-----------|---------|
| 1.0 | 2025-01-28 | Dokumen awal, fokus kebutuhan user & UI/UX, dibagi per role | - |
| 1.1 | 2025-02-11 | Satu user bisa lebih dari satu role; menu setelah login gabungan dari semua role; sebut Verifikator Uji Kompetensi | - |

---

**Catatan**: Dokumen ini hanya mencakup fitur **Login**. Fitur **Lupa Password** dan **Logout** didokumentasikan dalam PRD masing-masing.
